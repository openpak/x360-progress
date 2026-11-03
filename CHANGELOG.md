# Changelog

- 2026-11-03: v0.9.4: No secrets in logs: the `request` line logs the path only, not the query string. Test added.

- 2026-11-03: v0.9.3: Every request logs one stderr line, `http <method> <path> -> <status> <duration>ms` (x360-common v0.5.1 `httplog`): no query, headers, cookies, tokens or bodies; `/healthz` skipped.

- 2026-10-19: v0.9.2: serve HTTPS with XAM-compatible TLS 1.2 RSA suites; XAM omits Authorization
  for HTTP LiveHive URLs. Accept decimal XUID paths sent by the dashboard.

- 2026-10-18: v0.9.0 first release: empty-list replies for the resources the dashboard recovers after a sign-in,
  behind the XSTS token.
