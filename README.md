# x360-progress

The Xbox 360 progress REST service (hive root `LiveServiceUrl.progress`): the profile data the dashboard recovers after a sign-in - title history, achievements, avatar awards. Empty replies for now: a new profile has none.

The console reaches it through the LiveHive: `LiveServiceUrl.progress=http://progress.xboxlive.com` (x360-liveinfo) and OpenPak DNS
resolving `progress.xboxlive.com`. Requests carry `Authorization: XBL2.0 x=<token>` from [`x360-xsts`](../x360-xsts); the token is
verified with the same signing key (`x360-common/xstoken`) and may read only its own xuid.

Part of OpenPak's Xbox 360 stack: one repository and one image per service around the shared Go library
[`x360-common`](../x360-common) (`replace` + committed `vendor/`; after changing it run `go mod vendor` here and commit).

## Configuration

`X360_PROGRESS_ADDR` (default `:10041`), `X360_PROGRESS_XSTS_KEY` (the token signing key file, read-only). See `deploy/`.

## Build and test

`go test -mod=vendor ./...`. Tag `vX.Y.Z` and CI publishes `ghcr.io/openpak/x360-progress:latest` (amd64 + arm64).
The response shapes come from the XAM analysis in the Wharf notes (xrl-schema-opcodes.md section 12). Licensed AGPL-3.0, see LICENSE.
