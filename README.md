# Savood Backend in GO

## Code Generation

Ensure basepath to be "v2" and scheme to be "http" not "https" because SSL is terminated in k8s cluster 

```bash
swagger generate server -f ../api-definition/swagger.yml
```


## Build

```bash
go build -o ./bin/savood-server cmd/savood-server/main.go
```

```bash
GOOS=linux go build -o ./bin/savood-server -ldflags="-s -w" cmd/savood-server/main.go
```
