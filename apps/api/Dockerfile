# syntax=docker/dockerfile:1.7
#
# Multi-stage build for the Go API service.
# Build context: apps/api (the Go service is self-contained, no workspace needed).
# Run from apps/api: `docker build -f Dockerfile .`

# --- Stage 1: build ---
FROM golang:1.26-alpine AS build
WORKDIR /src

# Cache the module download separately from the source compile.
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod download

COPY . .

ARG BUILD_SHA=dev
ARG VERSION=dev
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 GOOS=linux \
  go build \
    -trimpath \
    -ldflags="-s -w -X main.buildSHA=${BUILD_SHA} -X main.version=${VERSION}" \
    -o /server ./cmd/server

# --- Stage 2: runtime ---
FROM gcr.io/distroless/static-debian12 AS runtime
WORKDIR /app

ENV PORT=8080
ENV ENVIRONMENT=production
EXPOSE 8080

COPY --from=build /server /app/server

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
