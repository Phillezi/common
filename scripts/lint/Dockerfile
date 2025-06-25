FROM --platform=$BUILDPLATFORM golang:alpine AS builder

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

COPY . .

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=cache,target="/root/.cache/go-build" \
    CGO_ENABLED=0 GOCACHE=/root/.cache/go-build GOOS=$TARGETOS GOARCH=$TARGETARCH go build -mod=readonly -ldflags "-w -s" -o ./bin/glint .

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache \
#    gcompat \
    ca-certificates

COPY --from=builder /app/bin/* /usr/bin/

ENTRYPOINT [ "/usr/bin/glint" ]
