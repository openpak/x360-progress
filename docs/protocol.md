# Protocol: x360-progress

The Xbox 360 progress REST service (hive root `LiveServiceUrl.progress`): the profile data the dashboard recovers after a sign-in - title history, achievements, avatar awards. Empty replies for now: a new profile has none.

The console reaches it through the LiveHive: `LiveServiceUrl.progress=https://progress.xboxlive.com` (x360-liveinfo) and OpenPak DNS
resolving `progress.xboxlive.com`. Requests carry `Authorization: XBL2.0 x=<token>` from [`x360-xsts`](../../x360-xsts); the token is
verified with the same signing key (`x360-common/xstoken`) and may read only its own xuid.
