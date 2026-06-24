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
