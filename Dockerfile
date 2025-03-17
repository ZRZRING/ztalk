FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ztalk .

FROM debian:bullseye-slim

COPY ./wait-for-it.sh ./wait-for-it.sh
COPY ./templates ./templates
COPY ./static ./static
COPY ./config ./config

COPY --from=builder /build/ztalk /

RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
    chmod 755 /wait-for-it.sh

EXPOSE 8080

# ENTRYPOINT ["/ztalk", "config/release.yaml"]