FROM golang:1.10-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

WORKDIR /go/src/git.dhbw.chd.cx/savood/backend
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN GOOS=linux go build -o /go/bin/backend -ldflags="-s -w" cmd/savood-server/main.go



FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /go/bin/backend /usr/local/bin/


EXPOSE 8080
CMD ["backend", "--port", "8080"]