FROM golang:1.24-alpine

RUN apk add --no-cache ca-certificates

WORKDIR /go/src/github.com/pierre-emmanuelJ/iptv-proxy

# Optimize build caching
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Build properly with modules enabled
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o iptv-proxy .

FROM alpine:3
COPY --from=0  /go/src/github.com/pierre-emmanuelJ/iptv-proxy/iptv-proxy /
ENTRYPOINT ["/iptv-proxy"]
