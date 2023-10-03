FROM golang:1.21 as base 

ENV PATH="${GOPATH}/bin:${PATH}"
RUN go install golang.org/x/tools/gopls@v0.10.0
RUN go install github.com/tpng/gopkgs@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.3.3
RUN go install github.com/go-delve/delve/cmd/dlv@latest
CMD [ "tail", "-f", "/dev/null" ]
