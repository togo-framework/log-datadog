// Package datadog ships togo's slog logs to Datadog Logs via the HTTP intake API,
// in addition to the app's existing log output. Install alongside togo-framework/log;
// blank-import registers it.
//
// Env: DD_API_KEY (required — no-op when empty), DD_SITE (default datadoghq.com),
// DD_SERVICE (default togo).
package datadog

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("log-datadog", togo.PriorityService, func(k *togo.Kernel) error {
		key := os.Getenv("DD_API_KEY")
		if key == "" {
			return nil // unconfigured → no-op
		}
		site := envOr("DD_SITE", "datadoghq.com")
		dd := &handler{
			url:     "https://http-intake.logs." + site + "/api/v2/logs",
			key:     key,
			service: envOr("DD_SERVICE", "togo"),
			source:  "go",
			level:   slog.LevelInfo,
			client:  &http.Client{Timeout: 10 * time.Second},
		}
		// Keep the existing handler + also ship to Datadog.
		k.Log = slog.New(tee{[]slog.Handler{k.Log.Handler(), dd}})
		return nil
	})
}

func envOr(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

// handler is an slog.Handler that POSTs records to the Datadog log intake.
type handler struct {
	url, key, service, source string
	level                     slog.Level
	client                    *http.Client
	attrs                     []slog.Attr
}

func (h *handler) Enabled(_ context.Context, l slog.Level) bool { return l >= h.level }

func (h *handler) Handle(_ context.Context, r slog.Record) error {
	m := map[string]any{
		"ddsource":  h.source,
		"service":   h.service,
		"message":   r.Message,
		"status":    r.Level.String(),
		"timestamp": r.Time.UnixMilli(),
	}
	for _, a := range h.attrs {
		m[a.Key] = a.Value.Any()
	}
	r.Attrs(func(a slog.Attr) bool { m[a.Key] = a.Value.Any(); return true })
	body, _ := json.Marshal([]map[string]any{m})
	// Ship asynchronously so logging never stalls request handling.
	go func() {
		req, err := http.NewRequest(http.MethodPost, h.url, bytes.NewReader(body))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("DD-API-KEY", h.key)
		if resp, err := h.client.Do(req); err == nil {
			resp.Body.Close()
		}
	}()
	return nil
}

func (h *handler) WithAttrs(as []slog.Attr) slog.Handler {
	n := *h
	n.attrs = append(append([]slog.Attr{}, h.attrs...), as...)
	return &n
}
func (h *handler) WithGroup(string) slog.Handler { return h }

// tee fans each record out to multiple handlers (the original + Datadog).
type tee struct{ hs []slog.Handler }

func (t tee) Enabled(ctx context.Context, l slog.Level) bool {
	for _, h := range t.hs {
		if h.Enabled(ctx, l) {
			return true
		}
	}
	return false
}
func (t tee) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range t.hs {
		if h.Enabled(ctx, r.Level) {
			_ = h.Handle(ctx, r.Clone())
		}
	}
	return nil
}
func (t tee) WithAttrs(as []slog.Attr) slog.Handler {
	n := make([]slog.Handler, len(t.hs))
	for i, h := range t.hs {
		n[i] = h.WithAttrs(as)
	}
	return tee{n}
}
func (t tee) WithGroup(name string) slog.Handler {
	n := make([]slog.Handler, len(t.hs))
	for i, h := range t.hs {
		n[i] = h.WithGroup(name)
	}
	return tee{n}
}
