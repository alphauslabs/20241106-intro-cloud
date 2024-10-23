FROM golang:1.23.2
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /simwork
ENTRYPOINT ["/simwork"]
CMD ["-multi"]