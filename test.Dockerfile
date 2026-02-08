FROM golang:bookworm

RUN go install github.com/go-task/task/v3/cmd/task@latest && \
    go install gotest.tools/gotestsum@latest

RUN useradd -u 1001 -m tester
USER tester

ENV GOCACHE=/tmp/gocache
ENV GOMODCACHE=/tmp/gomodcache

RUN mkdir -p /tmp/gocache /tmp/gomodcache && \
    chmod -R 777 /tmp/gocache /tmp/gomodcache

WORKDIR /app/test

ENTRYPOINT ["task", "test"] 
