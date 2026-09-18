#!/bin/bash
set -e

# Support conditional operations
gh issue create \
  --title "feat(mock): Support conditional operations (only-if-key-absent)" \
  --body "## Description
We need to support conditional operations to ensure atomicity for initialization flows.

## Scope
Extend \`Operation\` with a condition type (e.g., \`only-if-key-absent\`) and validate it in \`Apply\` and \`Simulate\`.

## Acceptance Criteria
- [ ] Add condition enum/string to \`Operation\`
- [ ] Implement validation logic for \`only-if-key-absent\`
- [ ] Add test cases

## Out of Scope
Complex conditions (e.g., comparing current values)."

# Add --watch mode
gh issue create \
  --title "feat(cmd): Add \`--watch\` mode to hot-reload config" \
  --body "## Description
Developers currently have to restart the RPC server to load new config changes. A watch mode would streamline this.

## Scope
Watch the config file for modifications and safely reload the in-memory state engine without dropping the RPC server.

## Acceptance Criteria
- [ ] Add \`--watch\` flag to CLI
- [ ] Use a file watcher (e.g. fsnotify) to reload config on change
- [ ] Safely overwrite \`Engine\` state

## Out of Scope
Diffing state changes; just fully replace the state."

# Add getNetwork method
gh issue create \
  --title "feat(rpc): Add \`getNetwork\` RPC method" \
  --body "## Description
Clients often call \`getNetwork\` to verify the RPC endpoint. We should mock this to return a local testing network passphrase.

## Scope
Add a \`getNetwork\` handler to the JSON-RPC server.

## Acceptance Criteria
- [ ] Add \`getNetwork\` method handler
- [ ] Return a hardcoded mock passphrase/network details
- [ ] Add test cases

## Out of Scope
Dynamic network configurations."

# Publish Go SDK client
gh issue create \
  --title "feat(client): Publish a Go SDK client for the RPC server" \
  --body "## Description
To make testing easier for Go projects, we should provide a typed Go client for our mocked RPC server.

## Scope
Create a \`pkg/client\` package that connects to our local JSON-RPC server and exposes methods.

## Acceptance Criteria
- [ ] Implement HTTP client wrapper
- [ ] Add typed request/response structs
- [ ] Add integration tests

## Out of Scope
Client generation from OpenAPI/JSON schema."

echo "Issues created successfully."
