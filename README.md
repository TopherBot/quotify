# Quotify

**Quotify** is a tiny command‑line tool written in Go that prints a random inspirational quote.

## Features
- One‑line command: `quotify`
- No external dependencies (uses only the Go standard library)
- Unit test & CI pipeline with linting, testing and binary build

## Installation
```bash
# Using go install (requires Go 1.22+)
go install github.com/yourname/quotify@latest
```

Or download the pre‑built binary from the GitHub Releases page.

## Usage
```bash
quotify
```
You should see something like:
```
"The only limit to our realization of tomorrow is our doubts of today." – Franklin D. Roosevelt
```

## Development
```bash
# Clone the repo
git clone https://github.com/yourname/quotify.git
cd quotify

# Run tests
go test ./...

# Build locally
go build -o quotify .
```

## CI
The repository ships with a GitHub Actions workflow that runs on every push and PR:
- **Lint** (`go vet` & `staticcheck`)
- **Unit Tests** with coverage
- **Build** the binary for Linux, macOS and Windows

Badge (replace with your actual workflow badge):

![CI](https://github.com/yourname/quotify/actions/workflows/ci.yml/badge.svg)

## License
MIT – see the `LICENSE` file.
