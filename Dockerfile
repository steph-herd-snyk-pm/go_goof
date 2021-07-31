FROM golang:1.15.4-alpine

ENV SNYK_VERSION=v1.664.0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN apk add --no-cache --update curl bash gcc

RUN cd /usr/local/bin && \
    curl -Lo ./snyk "https://github.com/snyk/snyk/releases/download/${SNYK_VERSION}/snyk-alpine" && \
    chmod -R +x ./snyk

CMD snyk test --severity-threshold=high