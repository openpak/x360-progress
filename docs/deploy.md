# Deploy: x360-progress

Part of OpenPak's Xbox 360 stack: one repository and one image per service around the shared Go library
[`x360-common`](../../x360-common) (`replace` + committed `vendor/`; after changing it run `go mod vendor` here and commit).

## Configuration

`X360_PROGRESS_ADDR` (default `:10041`), `X360_PROGRESS_XSTS_KEY` (the token signing key file, read-only). See `deploy/`.

## Build and test

`go test -mod=vendor ./...`. Tag `vX.Y.Z` and CI publishes `ghcr.io/openpak/x360-progress:latest` (amd64 + arm64).

## TLS listener

XAM sends Authorization only over HTTPS. Set `X360_SERVICE_TLS_ADDR`, `X360_SERVICE_TLS_CERT`, and
`X360_SERVICE_TLS_KEY` to enable the separate TLS listener (TLS 1.2, RSA AES-CBC). Use an exact-host
leaf signed by the x360-xuacs CA, so service 12 can validate it; Traefik must pass TLS through.
The XUID in console REST paths is decimal; hexadecimal diagnostic paths are also accepted.

## Image and release

- Image: `ghcr.io/openpak/x360-progress`, built from the `Dockerfile` (`EXPOSE 10041/tcp`, runs as a non-root user).
- Release: `.github/workflows/release.yml` publishes `:<version>`, `:<major>.<minor>` and `:latest` on tags `v*.*.*` (amd64 and arm64). Deployed hosts run podman-auto-update against `:latest`, so pushing a tag is the deploy.
- Quadlet: `deploy/openpak-x360-progress.container` (plain HTTP on `:10041`, TLS listener on `:10441`).
- Traefik: `deploy/traefik-tls.yml` passes TCP for `progress.xboxlive.com` through to the TLS listener on `:10441`.
