FROM golang:1.14.1-alpine as builder
RUN apk add --no-cache ca-certificates git
ENV PROJECT /github.com/tommy-sho/k8s-grpc-health-check/gateway
WORKDIR /go/src/$PROJECT

ENV GO111MODULE on
COPY . .
RUN go install .

FROM alpine as release
RUN apk add --no-cache ca-certificates git
COPY --from=builder /go/bin/gateway /gateway
RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
ENTRYPOINT ["/gateway"]