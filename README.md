## dora-proxy

Transparent proxy for selected Dora API endpoints with special handling for `/api/v1/slot/head`.

Endpoints proxied:

- `POST /api/v1/validator` → upstream `/api/v1/validator`
- `GET  /api/v1/epoch/latest` → upstream `/api/v1/epoch/latest`
- `GET  /api/v1/slot/{slotOrHash}` → upstream `/api/v1/slot/{slotOrHash}`
  - Special case: when `{slotOrHash}` is `head`, the proxy resolves the current head beacon block root via the consensus REST API and forwards the request as `/api/v1/slot/0x...`.

Environment variables:

- `PROXY_LISTEN_ADDR` (default `:8081`) — listen address
- `PROXY_UPSTREAM_BASE_URL` (default `http://localhost:8080`) — Dora upstream base; `/api` is appended automatically
- `PROXY_CONSENSUS_API_URL` (default `http://localhost:5052`) — Beacon node REST base (used to resolve head)

Run:

```bash
PROXY_LISTEN_ADDR=:8088 PROXY_UPSTREAM_BASE_URL=https://light-beacon.fusionist.io PROXY_CONSENSUS_API_URL=http://your-beacon-node:5052 go run .
```



### Docker

Build the image:

```bash
docker build -t dora-proxy:latest .
```

Run the container (maps port 8081 and sets upstreams):

```bash
docker run --name dora-proxy -d --restart=unless-stopped -p 8081:8081 \
  -e PROXY_UPSTREAM_BASE_URL=https://light-beacon.fusionist.io \
  -e PROXY_CONSENSUS_API_URL=http://your-beacon-node:5052 \
  dora-proxy:latest
```

Notes:
- `PROXY_LISTEN_ADDR` defaults to `:8081` in the container; override with `-e PROXY_LISTEN_ADDR=...` if needed.
