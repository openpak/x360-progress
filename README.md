# Xbox 360 Progress Server

OpenPak's Xbox 360 Progress Server.

| | |
| --- | --- |
| Network | Xbox 360 |
| Role | REST service for the profile progress the dashboard recovers after a sign-in: title history, achievements, avatar awards |
| Used by | Xbox 360 console, through the LiveHive `LiveServiceUrl.progress` set by [`x360-liveinfo`](https://github.com/openpak/x360-liveinfo) |
| Depends on | [`x360-xsts`](https://github.com/openpak/x360-xsts) (token signing key), [`x360-common`](https://github.com/openpak/x360-common) (shared Go library) |
| Image | `ghcr.io/openpak/x360-progress` |
| Licence | AGPL-3.0-only |

## What it serves

Every request needs `Authorization: XBL2.0 x=<token>` and may read only the xuid in its own token.

| Endpoint | What |
| --- | --- |
| `GET /users/xuid(<xuid>)/history/titles` | Title history |
| `GET /users/xuid(<xuid>)/titleachievements`, `.../achievements` | Achievements |
| `GET /users/xuid(<xuid>)/avatar/titleawards`, `.../avatar/awards` | Avatar awards |

## Build

```sh
go test -mod=vendor ./...
```

## Docs

- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)
- [Notes](docs/notes.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
