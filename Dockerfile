FROM golang:1.10-alpine as builder
RUN apk update && apk add --no-cache git make

ARG PROJECT
ARG APP_VERSION
ARG GITHASH
ARG BUILDSTAMP

WORKDIR /go/src/github.com/eschudt/${PROJECT}
COPY . /go/src/github.com/eschudt/${PROJECT}/

RUN make build-local

FROM alpine:3.7
RUN apk update && apk add --no-cache ca-certificates
COPY --from=builder /app /app

EXPOSE 8080

ENTRYPOINT ["/app"]
