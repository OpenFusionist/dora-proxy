## dora-proxy

A transparent proxy for selected Dora API endpoints, enhanced with additional fields for compatibility with the Beacon Explorer API.

### Why this exists

Beacon Explorer is costly to operate (depends on Bigtable and other heavy components) and its API differs from Dora. This proxy uses Dora as the data source while returning responses compatible with the Beacon Explorer API.


### Endpoints

- POST `/api/v1/validator` → 上游 `/api/v1/validator`
  - What it does：
    - `status` mapping: `active_ongoing → active_online`; `status:withdrawal_done+is_slashed=true → slashed`; `status:withdrawal_done+is_slashed=false → exited`.
    - add `lastattestationslot` (from consensus API).

- GET `/api/v1/epoch/latest` → upstream `/api/v1/epoch/latest`
  - What it does: transparent pass-through, no transformation.

- GET `/api/v1/slot/{slotOrHash}` → upstream `/api/v1/slot/{slotOrHash}`
  - What it does:
    - Supports `{slotOrHash}=head`: resolves the current head block root via consensus REST, then forwards to upstream.
    - Enrich with the following fields:
      - Eth1: `eth1data_depositcount`, `eth1data_depositroot`, `eth1data_blockhash`
      - Execution payload: `exec_logs_bloom`, `exec_parent_hash`,`exec_random`,`exec_receipts_root`,`exec_state_root`,`exec_timestamp`
      - Sync aggregate: `syncaggregate_bits`, `syncaggregate_signature`
      - Randao reveal: `randaoreveal`
      - Signature: `signature`

### Config & run

- `PROXY_LISTEN_ADDR` (default `:8081`) — listen address
- `PROXY_UPSTREAM_BASE_URL` (default `http://localhost:8080`) — Dora upstream base
- `PROXY_CONSENSUS_API_URL` (default `http://localhost:5052`) — Beacon node

Run:

```bash
PROXY_LISTEN_ADDR=:8088 PROXY_UPSTREAM_BASE_URL=https://light-beacon.fusionist.io PROXY_CONSENSUS_API_URL=http://your-beacon-node:5052 go run .
```



### Docker

Build the image:

```bash
docker build -t dora-proxy:latest .
docker run --name dora-proxy -d --restart=unless-stopped -p 8081:8081 \
  -e PROXY_UPSTREAM_BASE_URL=https://light-beacon.fusionist.io \
  -e PROXY_CONSENSUS_API_URL=http://your-beacon-node:5052 \
  dora-proxy:latest
```

Notes: In the container, `PROXY_LISTEN_ADDR` defaults to `:8081`; override with `-e PROXY_LISTEN_ADDR=...` if needed.
