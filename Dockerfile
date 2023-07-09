FROM golang:alpine AS build
LABEL MAINTAINER = 'Juan Pablo serna'

RUN apk add --update git
WORKDIR /go/src/github.com/JuanPabloSerna/jps-alert-crud
COPY . .
RUN TAG=$(git describe --tags --abbrev=0) \
    && LDFLAGS=$(echo "-s -w -X main.version="$TAG) \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/jps-alert-crud -ldflags "$LDFLAGS" cmd/jps-alert-crud/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/jps-alert-crud /go/bin/jps-alert-crud
ENTRYPOINT ["/go/bin/jps-alert-crud"]