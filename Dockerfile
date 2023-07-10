#build stage
FROM golang:1.20.5-alpine AS builder
RUN apk add --no-cache git upx

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

RUN go build \
    -ldflags="-s -w" \
    -o app -v .

RUN upx app

#final stage
FROM alpine:latest
LABEL Name=jpsalertcrud Version=0.0.1
RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app .

ENTRYPOINT ["./app"]

EXPOSE 8080
