# SSO service

gRPC authorization (SSO) service.

## Run

From the module root (`sso/`):

```bash
go run ./cmd/sso/
```

With an explicit config file:

```bash
go run ./cmd/sso/ -config config/local.yaml
```

From `cmd/sso/`:

```bash
go run . -config ../../config/local.yaml
```

## Config

Priority: `-config` flag → `CONFIG_PATH` env → `config/local.yaml`.

**PowerShell:**

```powershell
$env:CONFIG_PATH = "config/test.yaml"
go run ./cmd/sso/
```

Config files: `config/local.yaml`, `config/test.yaml`, `config/prod.yaml`.
