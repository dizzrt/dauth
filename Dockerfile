FROM --platform=${BUILDPLATFORM} golang:1.25 AS builder

ARG GOPROXY=https://goproxy.cn,direct
ARG TARGETOS
ARG TARGETARCH

COPY ./go.mod ./go.sum /
RUN go mod download

WORKDIR /src
COPY . .

RUN GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" go build -o ./bin/app ./cmd/main.go

FROM --platform=${TARGETPLATFORM} debian:stable-slim

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
    && apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    libterm-readline-perl-perl \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /app
COPY --from=builder /src/bin ./
COPY --from=builder /src/config ./config

EXPOSE 8081 50051

ENTRYPOINT [ "./app" ]
