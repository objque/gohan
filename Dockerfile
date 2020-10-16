FROM golang:1-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

ARG RELEASE=unset
ARG COMMIT=unset
ARG BUILD_TIME=unset
ENV PROJECT=github.com/objque/gohan

WORKDIR /go/src/github.com/objque/gohan
COPY migrations /var/gohan/migrations
COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN go build -v -a \
    -gcflags "all=-trimpath=${WORKDIR}" \
    -ldflags "-w -s \
       -X ${PROJECT}/internal/version.Release=${RELEASE} \
       -X ${PROJECT}/internal/version.Commit=${COMMIT} \
       -X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
    -o /usr/local/bin/gohan ./cmd/gohan/...

FROM alpine:latest

RUN addgroup -S gohan && adduser -S gohan -G gohan
USER gohan
WORKDIR /home/gohan

COPY --from=builder --chown=gohan:gohan /var/gohan/migrations /var/gohan/migrations
COPY --from=builder --chown=gohan:gohan /usr/local/bin/gohan /usr/local/bin/gohan

ENTRYPOINT ["/usr/local/bin/gohan"]
CMD []