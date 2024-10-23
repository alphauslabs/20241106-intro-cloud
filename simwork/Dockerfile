FROM golang:1.23.2-bookworm
COPY . /go/src/github.com/alphauslabs/simwork/
WORKDIR /go/src/github.com/alphauslabs/simwork/
RUN CGO_ENABLED=0 GOOS=linux go build -v -trimpath -installsuffix cgo -o simwork .

FROM debian:stable-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app/
COPY --from=0 /go/src/github.com/alphauslabs/simwork/simwork .
ENTRYPOINT ["/app/simwork"]
CMD ["-multi"]
