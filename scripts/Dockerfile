FROM --platform=$BUILDPLATFORM golang:1.25 AS builder
ARG GOPROXY=https://goproxy.cn,direct
ARG TARGETOS
ARG TARGETARCH

ARG ENV
ARG VERSION
ARG GIT_COMMIT

ADD go.mod go.sum /
RUN go mod download

WORKDIR /src
COPY . .

RUN GOOS="$TARGETOS" GOARCH="$TARGETARCH" go build \
    --tags "${ENV},docker" \
    -ldflags="-w -s \
    -X github.com/dizzrt/dauth/internal/conf.Version=${VERSION}" \
    -o ./bin/dauth ./main.go

FROM debian:stable-slim
ARG ENV

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list.d/debian.sources \
    && apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates  \
    libterm-readline-perl-perl \
    netbase \
    && rm -rf /var/lib/apt/lists/ \
    && apt-get autoremove -y && apt-get autoclean -y

WORKDIR /app
COPY --from=builder /src/bin/* ./
COPY --from=builder /src/configs ./configs

EXPOSE 8081 50051

ENV ENV=${ENV}

# HEALTHCHECK --interval=5m --timeout=60s CMD curl -f http://localhost:8080/healthz || exit 1

ENTRYPOINT [ "/app/dauth", "run" ]
