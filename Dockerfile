FROM golang:latest AS builder

ENV CGO_ENABLED=0

WORKDIR /app

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY Taskfile.yml ./
RUN task setup

RUN mkdir -p /tmp/cache-structure/mod && \
    chmod -R 777 /tmp/cache-structure

FROM gcr.io/distroless/static-debian12

ENV PATH=$PATH:/usr/local/go/bin:/usr/local/bin
ENV GOCACHE=/go-cache
ENV GOMODCACHE=/go-cache/mod

COPY --from=builder /usr/local/go /usr/local/go
COPY --from=builder /go/bin/task /usr/local/bin/task
COPY --from=builder /go/bin/gotestsum /usr/local/bin/gotestsum

COPY --from=builder /tmp/cache-structure /go-cache

WORKDIR /app/test

USER nonroot

ENV CGO_ENABLED=0

ENTRYPOINT ["task", "test"]
