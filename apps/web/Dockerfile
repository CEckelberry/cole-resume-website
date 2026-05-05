# syntax=docker/dockerfile:1.7
#
# Multi-stage build for the SvelteKit web service.
# Build context: repository root (so pnpm-workspace.yaml resolves correctly).
# Run from repo root: `docker build -f apps/web/Dockerfile .`

# --- Stage 1: build ---
# pnpm 11 requires Node 22+, so the build image is pinned to node:22-alpine
# even though SvelteKit's runtime works on Node 20. The runtime stage uses
# nodejs22-debian12 to match.
FROM node:22-alpine AS build
WORKDIR /workspace

# pnpm via corepack — pinned through packageManager field in root package.json.
RUN corepack enable

# Copy workspace manifests first so pnpm install layer caches well.
COPY pnpm-workspace.yaml pnpm-lock.yaml package.json ./
COPY apps/web/package.json apps/web/package.json

RUN --mount=type=cache,id=pnpm,target=/root/.local/share/pnpm/store \
  pnpm install --frozen-lockfile --filter web...

# Copy the rest of the web app and any future workspace deps.
COPY apps/web apps/web

# Build SvelteKit app (adapter-node emits apps/web/build), then use
# `pnpm deploy` to assemble a self-contained directory with only prod deps —
# resolves any workspace links into a flat node_modules tree the runtime
# stage can copy without pnpm available.
RUN --mount=type=cache,id=pnpm,target=/root/.local/share/pnpm/store \
  pnpm --filter web build && \
  pnpm --filter web deploy --prod --legacy /workspace/deploy

# --- Stage 2: runtime ---
FROM gcr.io/distroless/nodejs22-debian12 AS runtime
WORKDIR /app

ENV NODE_ENV=production
ENV PORT=3000
EXPOSE 3000

COPY --from=build /workspace/apps/web/build /app/build
COPY --from=build /workspace/deploy/package.json /app/package.json
COPY --from=build /workspace/deploy/node_modules /app/node_modules

CMD ["build/index.js"]
