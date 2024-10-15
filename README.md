Refers to ioapi bug https://github.com/oapi-codegen/oapi-codegen/issues/1804

To reproduce:

```bash
$ go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml endpoints.yaml
$ go run ./server.go
```

it will throw:

```
Starting
Error loading swagger spec:
failed to resolve "foo_bar" in fragment in URI: "#/components/schemas/foo_bar": map key "foo_bar" not foundexit status 1

```