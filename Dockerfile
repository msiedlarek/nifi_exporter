FROM golang:1.14-alpine AS builder
WORKDIR /go/src/github.com/msiedlarek/nifi_exporter
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/nifi_exporter

FROM scratch
COPY --from=builder /go/bin/nifi_exporter /nifi_exporter
ENTRYPOINT ["/nifi_exporter"]
CMD ["/config/config.yml"]
