# snippetbox

[English](./README.md) | [中文](./README.zh-CN.md)

## 中文

`snippetbox` 是书籍 *Let's Go (2022)* 的项目实现，用于练习 Go Web 开发中的基础能力，包括路由、中间件、模板渲染、表单处理、会话管理与数据持久化。

### 最小运行要求

- Go（版本以 `go.mod` 为准）
- MySQL（默认 DSN：`web:pass@/snippetbox?parseTime=true`）
- 本地 TLS 证书：`tls/cert.pem` 与 `tls/key.pem`

### TLS 证书生成

 `tls/cert.pem` 与 `tls/key.pem` 默认不存在，可使用 Go 内置脚本生成：

PowerShell：

```powershell
Push-Location tls
go run "$(go env GOROOT)\src\crypto\tls\generate_cert.go" --rsa-bits=2048 --host=localhost
Pop-Location
```

Bash：

```bash
pushd tls
go run "$(go env GOROOT)/src/crypto/tls/generate_cert.go" --rsa-bits=2048 --host=localhost
popd
```

### 项目结构（简版）

```text
.
|-- cmd/
|   `-- web/                 # HTTP 服务入口与处理逻辑
|-- internal/
|   |-- assert/              # 测试断言
|   |-- models/              # 数据访问层
|   `-- validator/           # 表单校验辅助
|-- tls/                     # 本地 TLS 证书与私钥
|-- ui/
|   |-- html/                # 模板文件
|   `-- static/              # css/js/img 静态资源
|-- go.mod
`-- README.md
```

### 启动

```bash
go run ./cmd/web
```

可选参数：

- `-addr`（默认 `:4000`）
- `-dsn`（默认 `web:pass@/snippetbox?parseTime=true`）

启动后访问：`https://localhost:4000`

### 测试

```bash
go test ./...
```
