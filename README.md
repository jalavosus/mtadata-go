# mtadata

Work-in-progress data parser and queryer for MTA GTFS data.

# Building

- `mkdir -p ./bin`
- `go build -o ./bin/api ./cmd/api`

# Running

### No local TLS

To just run the thing:

- (build the thing)
- `./bin/api`

### Local TLS (on the gRPC side)

To run the API with local tls enabled for the gRPC server:

- `chmod +x ./scripts/gentls.sh`
- `./scripts/gentls.sh`
- (build the thing)
- `./bin/api -t tlscerts/server.cert -k tlscerts/server.key`