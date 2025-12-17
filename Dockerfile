FROM golang:1.22-bookworm

ENV DEBIAN_FRONTEND=noninteractive
ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    python3-venv \
    git \
    curl \
    ca-certificates \
    build-essential \
    bash \
    && rm -rf /var/lib/apt/lists/*

RUN go install github.com/spf13/cobra-cli@latest

RUN go install golang.org/x/tools/cmd/goimports@latest \
    && go install golang.org/x/lint/golint@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

RUN pip3 install --no-cache-dir \
    pip \
    virtualenv

WORKDIR /workspace

CMD [ "bash" ]

