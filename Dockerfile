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
    nano \
    && rm -rf /var/lib/apt/lists/*

RUN go install github.com/spf13/cobra-cli@v1.3.0

RUN go install golang.org/x/tools/cmd/goimports@v0.14.0 \
    && go install golang.org/x/lint/golint@v0.0.0-20210508222113-6edffad5e616 \
    && go install github.com/go-delve/delve/cmd/dlv@v1.23.0

WORKDIR /workspace

COPY . /workspace/

CMD [ "tail", "-f", "/dev/null" ]

