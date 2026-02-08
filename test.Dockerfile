FROM golang:bookworm

RUN go install github.com/go-task/task/v3/cmd/task@latest

RUN useradd -u 1001 -m tester
USER tester
