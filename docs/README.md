# log-datadog — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package datadog ships togo's slog logs to Datadog Logs via the HTTP intake API,
in addition to the app's existing log output. Install alongside togo-framework/log;
blank-import registers it.

Env: DD_API_KEY (required — no-op when empty), DD_SITE (default datadoghq.com),
DD_SERVICE (default togo).

## Install

```bash
togo install togo-framework/log-datadog
```

Set `LOG_DRIVER=datadog`.

## Configuration

Environment variables read by this plugin (extracted from the source):

| Env var | Notes |
|---|---|
| `DD_API_KEY` | _see provider docs_ |
| `G` | _see provider docs_ |

## Usage

```go
// Structured logs/errors are forwarded to the configured sink automatically
// once this driver is installed and its env is set.
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/log-datadog
- README: ../README.md
