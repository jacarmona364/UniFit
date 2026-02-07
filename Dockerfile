FROM golang:bookworm

RUN go install github.com/go-task/task/v3/cmd/task@latest

RUN useradd -u 1001 -m -s /bin/bash usuario_tests

RUN mkdir -p /tmp/gocache /tmp/gomodcache && \
    chown -R 1001:1001 /tmp/gocache /tmp/gomodcache && \
    chmod -R 777 /tmp/gocache /tmp/gomodcache

ENV GOCACHE=/tmp/gocache
ENV GOMODCACHE=/tmp/gomodcache

WORKDIR /app/test

USER usuario_tests

ENTRYPOINT ["task", "test"]
