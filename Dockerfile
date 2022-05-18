FROM golang:1.18 as builder
WORKDIR /go/src/github.com/moov-io/imagecashletter
RUN apt-get update && apt-get install make gcc g++
COPY . .
RUN make build-server

FROM debian:stable-slim
LABEL maintainer="Moov <support@moov.io>"
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /go/src/github.com/moov-io/imagecashletter/bin/server /bin/server
# USER moov

EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
