# chatex-server

## Requirements
 - go > 1.12

To run the server locally run ` go run cmd/main.go `, the server will start at PORT 8080 by default
## NOTE
 - Vendoring is enabled, but the ` /vendor ` directory is not being commited to git, to avoid large size commits
 - On installing a new dependency or when pulling the code, run 
   - ` go mod vendor ` (to avoid Inconsistent vendoring)
   - ` go mod download ` (not tried)