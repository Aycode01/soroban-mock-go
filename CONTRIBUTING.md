# Contributing

Thank you for your interest in contributing to `soroban-mock-go`!

## Local Development

You will need Go 1.22+.

### Building

```bash
go build ./...
```

### Testing

Run all unit tests with the race detector enabled:

```bash
go test -race ./...
```

### Linting

This project uses `golangci-lint`. Please ensure your code passes before submitting a pull request.

```bash
golangci-lint run
```

## Pull Request Process

1. Fork the repo and create your branch from `main`.
2. Write tests for any new features or bug fixes.
3. Ensure the CI checks (tests and linting) pass.
4. Follow the **Conventional Commits** format for your commit messages (e.g. `feat(rpc): add getNetwork endpoint`).
5. Open a Pull Request!
