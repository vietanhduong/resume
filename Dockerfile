FROM golang:1.16-buster AS build
ENV GO111MODULE=on
WORKDIR /src
COPY . .
RUN go build -o /out/resume

FROM debian:buster

COPY --from=build /out/resume /
COPY public /public

RUN apt-get update && \
    apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["/resume"]