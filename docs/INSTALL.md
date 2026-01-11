# Installation

## Related Docs
- `docs/SECURITY.md`
- `docs/CONFIG.md`
- `docs/HOMEBREW_TAP.md`

## Homebrew (Primary)
```bash
brew tap brandon-gottshall/lfx
brew install lfx
```

Update:
```bash
brew upgrade lfx
```

## Homebrew (Dev Install)
Local repo build (one command):
```bash
./scripts/lfx-dev
```
This creates a temporary local tap (default `lfx/local`) and installs from it.
To remove: `brew uninstall lfx && brew untap lfx/local`.

Remote repo build (one command):
```bash
LFX_GH_OWNER=yourname LFX_GH_REPO=lfx ./scripts/lfx-dev remote
```
This installs from the GitHub tarball for the given ref.

## Curl Fallback
Use this when Homebrew is not available.

```bash
curl -fsSL https://github.com/brandon-gottshall/lfx/releases/latest/download/lfx-$(uname -s)-$(uname -m).tar.gz \
  | tar -xz
sudo mv lfx /usr/local/bin/lfx
```

## Verify
```bash
lfx --help
```

## Notes
- lfx writes into `~/.config/lf` when installing or applying extensions.
- No network access occurs unless the invoked command requires it (e.g., icon pack download).
- Ownership may change later; retapping with the new owner is sufficient.
