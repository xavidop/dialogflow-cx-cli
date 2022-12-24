FROM alpine:latest

COPY dialogflow-cx-cli_*.apk /tmp/
RUN apk add --no-cache --allow-untrusted /tmp/dialogflow-cx-cli_*.apk