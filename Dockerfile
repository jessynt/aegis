FROM alpine

RUN \
  apk add \
    --no-cache \
    --no-progress \
    tzdata \
    ca-certificates \
  && rm -rf /var/cache/apk/*

COPY bin/aegis* /usr/local/bin/