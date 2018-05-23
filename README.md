# Savood Backend in GO

## Build

```bash
go build -o ./bin/savood-server cmd/savood-server/main.go
```

```bash
GOOS=linux go build -o ./bin/savood-server -ldflags="-s -w" cmd/savood-server/main.go
```
