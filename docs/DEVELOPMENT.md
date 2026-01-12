# Development

## Dev Install (Homebrew)
Local repo build:
```bash
./scripts/lfx-dev
```
This installs a dev build using a local tap (`lfx/local`). Cleanup:
```bash
brew uninstall lfx && brew untap lfx/local
```

Remote repo build:
```bash
LFX_GH_OWNER=yourname LFX_GH_REPO=lfx ./scripts/lfx-dev remote
```

## Commit Hook (Build + Install)
Enable the hook for development:
```bash
./scripts/install-hooks.sh
```
Hooks are committed in the repo but are opt-in until you run the installer.
On each commit, the hook runs:
- `make build`
- `make test`
- `make lint`
- `./scripts/lfx-build-tarball`

If all pass, the post-commit hook runs:
- `./scripts/lfx-dev` (using the tarball from `artifacts/lfx/build/latest.env`)

Logs are written to:
- `artifacts/lfx/pre-commit/<timestamp>_<sha>.log`
- `artifacts/lfx/post-commit/<timestamp>_<sha>.log`
To disable the hook:
```bash
git config --unset core.hooksPath
```
