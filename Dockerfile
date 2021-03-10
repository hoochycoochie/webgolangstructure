FROM golang:1.13.4 AS builder
ENV DATA_DIRECTORY /go/src/hoochycoochie/webgolangstructure
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENBALED=0
COPY . .
RUN go build -ldflags="-X hoochycoochie/webgolangstructure/internal/config.Version=$APP_VERSION" hoochycoochie/webgolangstructure/cmd/server
FROM alpine:3.10
ENV DATA_DIRECTORY /go/src/hoochycoochie/webgolangstructure/
RUN apk add --update --no-cache \
    ca-certificates
COPY --from=builder ${DATA_DIRECTORY}server /webgolangstructure
ENTRYPOINT [ "/webgolangstructure" ]