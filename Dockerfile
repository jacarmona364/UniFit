FROM golang:latest AS builder

ENV CGO_ENABLED=0

RUN go install github.com/go-task/task/v3/cmd/task@latest

COPY Taskfile.yml .

RUN task install

