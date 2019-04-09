FROM golang:1.11-alpine3.9 as builder

RUN apk add --update git

COPY . .

ENV GO111MODULE="on"
RUN unset GOPATH && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /bin/semver

FROM alpine:3.9
COPY --from=builder /bin/semver /bin/ 
RUN apk add --update ca-certificates

ENTRYPOINT /bin/semver

