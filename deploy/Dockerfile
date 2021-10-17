FROM golang:1.16-alpine AS builder
RUN apk update &&\
    apk --no-cache add make git build-base linux-headers

WORKDIR /opt/fa
COPY . .

RUN make clean && make build

FROM alpine:latest

WORKDIR /opt/fa

COPY deploy/entrypoint.sh /
COPY deploy/glibc-2.33-r0.apk /tmp/
COPY --from=builder /opt/fa/bin/ /opt/fa/bin/
COPY --from=builder /opt/fa/conf/ /opt/fa/conf/

RUN apk update && apk --no-cache add tini tzdata ca-certificates && \
   touch repo.list && apk add --repositories-file=repo.list --allow-untrusted --no-network --no-cache /tmp/glibc-2.33-r0.apk &&\
   mkdir -p /opt/fa/log && \
   chmod +x /entrypoint.sh /opt/fa/bin/*

ENV PATH /opt/fa/bin:$PATH

ENV GIN_MODE=release

ENTRYPOINT ["/sbin/tini", "--", "/entrypoint.sh"]