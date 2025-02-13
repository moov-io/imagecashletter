FROM golang:1.24 as builder
WORKDIR /go/src/github.com/moov-io/imagecashletter
RUN apt-get update && apt-get install make gcc g++
COPY . .
RUN make build-server

FROM debian:stable-slim
LABEL maintainer="Moov <oss@moov.io>"
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /go/src/github.com/moov-io/imagecashletter/bin/server /bin/server
# USER moov

ENV HTTP_PORT=8083
ENV HEALTH_PORT=9093

EXPOSE ${HTTP_PORT}/tcp
EXPOSE ${HEALTH_PORT}/tcp

ENTRYPOINT ["/bin/server"]
