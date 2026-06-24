<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/log-datadog</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/log-datadog"><img src="https://pkg.go.dev/badge/github.com/togo-framework/log-datadog.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/log-datadog
```

<!-- /togo-header -->

<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">log-datadog</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

**Datadog** log shipping for togo. Forwards your app's `slog` logs to
[Datadog Logs](https://www.datadoghq.com) via the HTTP intake API, in addition to
the existing local/stdout output — so structured logs show up in Datadog with the
service tag.

```bash
togo install togo-framework/log-datadog
```

Install alongside `togo-framework/log`. Blank-importing the plugin registers it.

## Env

| Var | Required | Description |
|---|---|---|
| `DD_API_KEY` | yes | Datadog API key. When unset the plugin is a no-op. |
| `DD_SITE` | no | Datadog site (default `datadoghq.com`; EU = `datadoghq.eu`). |
| `DD_SERVICE` | no | Service name tag (default `togo`). |

## How it works

On boot (after the `log` plugin) it wraps the kernel logger so every record is
**also** POSTed asynchronously to `https://http-intake.logs.<site>/api/v2/logs`.
Shipping is non-blocking and never stalls request handling.

MIT © togo-framework

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
