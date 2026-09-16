# Build report

Baseline: `5e4d6912bae5f20a6224fa84ecd0e34b7838118e`

## Native build

- `go vet ./...`: PASS.
- `go build -trimpath -o /tmp/novel-studio-jiujiu ./cmd/novel-studio`: PASS.
- Native binary reported version `v0.3.1-0.20260913201948-5e4d6912bae5` and the exact baseline commit.
- CI doctor smoke with a deliberately missing config returned exit 1 as expected; platform, version, workspace, dashboard, and source-build checks passed.

## Dashboard

- Python dashboard: PASS, 53 tests.
- Node dashboard clients: PASS, 24 tests.

## Container

- Existing Colima backend was started because Docker Desktop was not installed and no daemon was initially running.
- `docker build --build-arg GOPROXY=https://proxy.golang.org,direct -t novel-studio:jiujiu-phase0 .`: PASS.
- Local image ID: `sha256:a9b0b804ad98718c5ba2c770349e349d0527a75df861bdaef5e9cd041a634890`.
- Architecture: linux/arm64.
- Embedded Dashboard `/api/health`: PASS with `{"ok": true}`.
- Smoke container was removed after the check. The image was not pushed.

Observation: the container reports `novel-studio dev`, `commit: unknown`, `built: unknown` because the fixed Dockerfile does not inject build-version ldflags. This does not change the source baseline but weakens runtime provenance and should be tracked as a non-blocking gap.

Result: **PASS** for native build, Dashboard, Docker build, and runtime health.
