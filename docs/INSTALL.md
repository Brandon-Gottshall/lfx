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

## Development
See `docs/DEVELOPMENT.md` for local dev install and commit-hook workflows.

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
