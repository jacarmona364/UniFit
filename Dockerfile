FROM golang:1.25 AS builder

ENV CGO_ENABLED=0

RUN go install github.com/go-task/task/v3/cmd/task@latest && \
    go install gotest.tools/gotestsum@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM gcr.io/distroless/static-debian12:debug

ENV PATH=$PATH:/usr/local/go/bin:/usr/local/bin

ENV GOCACHE=/go-cache
ENV GOMODCACHE=/go-cache/mod

SHELL ["/busybox/sh", "-c"]
RUN mkdir -p /go-cache/mod && chmod -R 777 /go-cache

COPY --from=builder /usr/local/go /usr/local/go

COPY --from=builder /go/bin/task /usr/local/bin/task
COPY --from=builder /go/bin/gotestsum /usr/local/bin/gotestsum

WORKDIR /app/test
COPY --from=builder --chown=nonroot:nonroot /app /app/test

USER nonroot:nonroot

ENV CGO_ENABLED=0

ENTRYPOINT ["task", "test"]
