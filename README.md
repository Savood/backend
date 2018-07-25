# Savood Backend in GO

A detailed description can be found in the presentation repository, wich links to important code.


## Code Generation

Ensure basepath to be "v2" and scheme to be "http" not "https" because SSL is terminated in k8s cluster 

```bash
swagger generate server -P models.Principal -f ../api-definition/swagger.yml
```


## Build

```bash
go build -o ./bin/savood-server cmd/savood-server/main.go
```

```bash
GOOS=linux go build -o ./bin/savood-server -ldflags="-s -w" cmd/savood-server/main.go
```

### Docker-Build

```bash
docker build --tag savood:test . 
CONTAINER_IMAGE=savood CI_BUILD_REF=test docker-compose up 
```
