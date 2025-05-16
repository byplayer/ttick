# development tips

## How to find attribute name to set ldflags

```bash
go build cmd/ttick/ttick.go

go tool nm ttick
```

Then you can find target string in output.
