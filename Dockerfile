FROM alpine:latest

COPY cxcli_*.apk /tmp/
RUN apk add --no-cache --allow-untrusted /tmp/cxcli_*.apk