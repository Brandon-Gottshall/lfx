# Installation

## Related Docs
- `docs/SECURITY.md`
- `docs/CONFIG.md`
- `docs/HOMEBREW_TAP.md`

## Homebrew (Primary)
```bash
brew tap brandon-gottshall/ifx
brew install ifx
```

Update:
```bash
brew upgrade ifx
```

## Homebrew (Dev Install)
Local repo build (one command):
```bash
./scripts/ifx-dev
```
This creates a temporary local tap (default `ifx/local`) and installs from it.
To remove: `brew uninstall lfx && brew untap ifx/local`.

Remote repo build (one command):
```bash
IFX_GH_OWNER=yourname IFX_GH_REPO=ifx ./scripts/ifx-dev remote
```
This installs from the GitHub tarball for the given ref.

## Curl Fallback
Use this when Homebrew is not available.

```bash
curl -fsSL https://github.com/brandon-gottshall/ifx/releases/latest/download/ifx-$(uname -s)-$(uname -m).tar.gz \
  | tar -xz
sudo mv ifx /usr/local/bin/ifx
```

## Verify
```bash
ifx --help
```

## Notes
- ifx writes into `~/.config/lf` when installing or applying extensions.
- No network access occurs unless the invoked command requires it (e.g., icon pack download).
- Ownership may change later; retapping with the new owner is sufficient.
