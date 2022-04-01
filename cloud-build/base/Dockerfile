FROM plexinc/pms-docker:1.25.8.5663-e071c3d62

ARG GO_VERSION=1.18
ARG GO_ARCH=amd64

RUN apt-get update && apt-get install -y build-essential

WORKDIR /usr/local

RUN curl -O -L "https://go.dev/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz"
RUN tar -xf go${GO_VERSION}.linux-${GO_ARCH}.tar.gz && rm go${GO_VERSION}.linux-${GO_ARCH}.tar.gz

ENV GOPATH="$HOME:/go"
ENV PATH="$PATH:/usr/local/go/bin:$GOPATH/bin"