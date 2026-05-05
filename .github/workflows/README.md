# CI / CD workflows

## `ci.yml`

Runs on every push to `main` and every pull request. Two parallel jobs:

| Job | Steps |
|---|---|
| `web` | pnpm install → `pnpm lint` → `pnpm check` (svelte-check) → `pnpm test` (vitest) |
| `api` | `go vet` → `go test -race` → `go build` |

Both jobs use repo-only inputs — no secrets needed.

## Future `deploy.yml`

When the GCP project comes online, a `deploy` job will be added that runs
only on push to `main` (gated on `ci.yml` succeeding). It will:

1. Authenticate to Google Cloud via Workload Identity Federation
2. Build both Docker images and push to Artifact Registry, tagged with the commit SHA
3. Run database migrations against Cloud SQL
4. Deploy both Cloud Run services and direct traffic to the new revisions
5. Smoke test `/` and `/api/health`

### Secrets the deploy job will expect

Set these in repo settings → Secrets and variables → Actions:

| Name | Purpose |
|---|---|
| `WIF_PROVIDER` | Full Workload Identity provider resource path, e.g. `projects/123/locations/global/workloadIdentityPools/portfolio/providers/github` |
| `WIF_SERVICE_ACCOUNT` | Service account email the WIF token impersonates, e.g. `gha-deployer@my-proj.iam.gserviceaccount.com` |
| `GCP_PROJECT_ID` | The Google Cloud project ID for the portfolio |
| `GCP_REGION` | Cloud Run region (e.g. `us-west1`) |
| `ARTIFACT_REGISTRY` | Repo path, e.g. `us-west1-docker.pkg.dev/my-proj/portfolio` |

The Terraform module under `infra/terraform/` provisions the WIF pool +
provider and outputs both `WIF_PROVIDER` and `WIF_SERVICE_ACCOUNT` so the
operator can paste them straight into repo secrets.
