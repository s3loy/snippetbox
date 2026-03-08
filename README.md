# snippetbox

[English](./README.md) | [中文](./README.zh-CN.md)

## English

`snippetbox` is an implementation of the project from *Let's Go (2022)*. It is used to practice core Go web development topics such as routing, middleware, template rendering, form handling, session management, and persistence.

### Minimal Requirements

- Go (version defined in `go.mod`)
- MySQL (default DSN: `web:pass@/snippetbox?parseTime=true`)
- Local TLS files: `tls/cert.pem` and `tls/key.pem`

### Generate TLS Certificates

At first `tls/cert.pem` and `tls/key.pem` do not exist, generate them with Go's built-in helper:

PowerShell:

```powershell
Push-Location tls
go run "$(go env GOROOT)\src\crypto\tls\generate_cert.go" --rsa-bits=2048 --host=localhost
Pop-Location
```

Bash:

```bash
pushd tls
go run "$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits=2048 --host=localhost
popd
```

### Project Tree (Simplified)

```text
.
|-- cmd/
|   `-- web/                 # HTTP server entry and handlers
|-- internal/
|   |-- assert/              # test assertions
|   |-- models/              # data access layer
|   `-- validator/           # form validation helpers
|-- tls/                     # local TLS cert and key
|-- ui/
|   |-- html/                # templates
|   `-- static/              # css/js/img assets
|-- go.mod
`-- README.md
```


### Run

```bash
go run ./cmd/web
```

Optional flags:

- `-addr` (default `:4000`)
- `-dsn` (default `web:pass@/snippetbox?parseTime=true`)

Then open: `https://localhost:4000`

### Test

```bash
go test ./...
```

