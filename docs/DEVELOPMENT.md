# Development

## Dev Install (Homebrew)
Local repo build:
```bash
./scripts/ifx-dev
```
This installs a dev build using a local tap (`ifx/local`). Cleanup:
```bash
brew uninstall lfx && brew untap ifx/local
```

Remote repo build:
```bash
IFX_GH_OWNER=yourname IFX_GH_REPO=ifx ./scripts/ifx-dev remote
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
- `./scripts/ifx-build-tarball`

If all pass, the post-commit hook runs:
- `./scripts/ifx-dev` (using the tarball from `artifacts/ifx/build/latest.env`)

Logs are written to:
- `artifacts/ifx/pre-commit/<timestamp>_<sha>.log`
- `artifacts/ifx/post-commit/<timestamp>_<sha>.log`
To disable the hook:
```bash
git config --unset core.hooksPath
```
