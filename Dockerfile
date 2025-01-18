FROM golang:1.23 as base
RUN groupadd --gid 1000 gopher \
    && useradd --uid 1000 --gid 1000 -m gopher \
    && chown -R 1000:1000 /home/gopher
USER gopher
ENV GOPATH="/home/gopher"
ENV PATH="${GOPATH}/bin:${PATH}"

FROM base as dev
COPY dev.bash .
RUN bash dev.bash
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
