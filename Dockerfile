FROM golang:1.25.6 AS builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ADD . /src
RUN cd /src && go build -o goapp -ldflags "-w -s" main.go

FROM golang:1.25.6-alpine

WORKDIR /src
COPY --from=builder /src/goapp /src/
COPY --from=builder /src/.env.example /src/.env.example


CMD ["/src/goapp"]

