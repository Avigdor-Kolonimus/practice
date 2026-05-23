# Contracts (gRPC Auth / SSO)

Protobuf definitions and generated Go stubs for the Auth gRPC API: register, login, admin check, and logout.

This directory is a **standalone Go module** published from the parent repository via Git tags.

## Layout

```
proto/sso/sso.proto   # Auth service API
gen/go/sso/           # Generated Go code (protoc)
go.mod                # Module definition
Taskfile.yaml         # Code generation tasks
```

The service implementation lives in the sibling [`sso/`](../sso/) module at the repository root.

## Module path

```
github.com/Avigdor-Kolonimus/practice/grpc_authorization_service/contracts
```

Import generated code:

```go
import ssov1 "github.com/Avigdor-Kolonimus/practice/grpc_authorization_service/contracts/gen/go/sso"
```

Pin a version with `go get` after pushing a release tag (see below).

## Generate Go code

From this directory:

```bash
task generate
# or: task gen
```

Or run `protoc` directly:

```bash
protoc -I proto proto/sso/*.proto \
  --go_out=./gen/go/ --go_opt=paths=source_relative \
  --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
```

Requires [protoc](https://grpc.io/docs/protoc-installation/), `protoc-gen-go`, and `protoc-gen-go-grpc`:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Commit both `proto/` and `gen/go/` changes before tagging a release.

## GitHub tags (Go module releases)

Consumers install this module with `go get` using **Git tags on the parent repository**. Tags must use the `contracts/` prefix because this module lives in a subdirectory.

### Tag format

| Tag | `go get` |
|-----|----------|
| `contracts/v0.1.0` | `go get github.com/Avigdor-Kolonimus/practice/grpc_authorization_service/contracts@v0.1.0` |
| `contracts/v1.0.0` | `go get github.com/Avigdor-Kolonimus/practice/grpc_authorization_service/contracts@v1.0.0` |

Use [semantic versioning](https://go.dev/doc/modules/version-numbers) (`vMAJOR.MINOR.PATCH`). Breaking changes in protos or generated code require a **major** version bump.

### Create and push a tag

Run from the **repository root** (parent of `contracts/`):

```bash
git tag -a contracts/v0.1.0 -m "contracts v0.1.0: initial Auth gRPC API"
git push origin contracts/v0.1.0
```

Or in GitHub: **Releases → Draft a new release** → tag name `contracts/v0.1.0`.

### Checklist before tagging

1. Generated code matches protos (`task generate` in `contracts/`).
2. `go.mod` and `go.sum` are tidy (`go mod tidy`).
3. Tag name is `contracts/vX.Y.Z` (prefix required for nested modules).

### Verify

```bash
go get github.com/Avigdor-Kolonimus/practice/grpc_authorization_service/contracts@v0.1.0
```

If it fails, check the remote tag (`git ls-remote --tags origin`) and repository access (public, or `GOPRIVATE` / credentials for private repos).

## Auth API

| RPC | Description |
|-----|-------------|
| `Register` | Register a new user |
| `Login` | Log in and receive an auth token |
| `IsAdmin` | Check if a user is an admin |
| `Logout` | Log out and invalidate the auth token |

See [`proto/sso/sso.proto`](proto/sso/sso.proto).
