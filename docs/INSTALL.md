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

## Development
See `docs/DEVELOPMENT.md` for local dev install and commit-hook workflows.

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
