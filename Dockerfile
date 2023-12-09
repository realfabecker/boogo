# syntax=docker/dockerfile:1.3-labs
FROM golang:1.18 as base

# specify a non root user
RUN groupadd --gid 1000 gopher \
  && useradd --uid 1000 --gid 1000 -m gopher \
  && mkdir p /home/gopher/go \
  && chown -R 1000:1000 /home/gopher/go

# change user for install
USER gopher

# declare working directory
WORKDIR /home/gopher/go
ENV GOPATH="/home/gopher/go"
ENV PATH="${GOPATH}/bin:${PATH}"

FROM base as dev
# declare project working directory
WORKDIR /home/gopher/go/src/github.com/falconzord/magacmd
# installing development dependencies
RUN <<EOF
  go install golang.org/x/tools/gopls@v0.10.0
  go install github.com/tpng/gopkgs@latest
  go install github.com/ramya-rao-a/go-outline@latest
  go install honnef.co/go/tools/cmd/staticcheck@v0.3.3
  go install github.com/go-delve/delve/cmd/dlv@latest
EOF
