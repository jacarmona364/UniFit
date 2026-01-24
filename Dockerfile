FROM debian:stable-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends wget ca-certificates git && \
    rm -rf /var/lib/apt/lists/* && \
    useradd -m -s /bin/bash test

RUN wget -q https://go.dev/dl/go1.25.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.25.5.linux-amd64.tar.gz && \
    rm go1.25.5.linux-amd64.tar.gz

RUN sh -c "$(wget -O - https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

USER test
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/home/test/go
ENV PATH=$PATH:$GOPATH/bin

RUN go install gotest.tools/gotestsum@latest

WORKDIR /app/test

COPY --chown=test:test go.mod go.sum* ./

RUN go mod download

COPY --chown=test:test Taskfile.yml ./

ENTRYPOINT ["task", "test"]
