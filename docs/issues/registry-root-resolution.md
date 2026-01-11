# Issue: Stable Registry Root Resolution for Installed Binaries

## Summary
`lfx theme list` currently uses `os.Getwd()` as the registry root. This works when running from the repo, but installed binaries need a stable registry path or discovery strategy.

## Problem
When `lfx` is installed outside the repo, `os.Getwd()` will not contain `registry/`, so list commands will fail or show empty output even though the registry should be available.

## Proposed Direction
- Decide on a stable registry root location for installed binaries.
- Options to evaluate:
  - Embed registry assets at build time and expose them via an internal path.
  - Install registry data to a known share dir and locate via runtime paths.
  - Allow a config override (config.toml) with a default fallback.

## Acceptance
- `lfx theme list` works when run outside the repo.
- Registry root resolution is deterministic and documented.
- No GitHub ownership or URLs are baked into runtime logic.

## Notes
This issue is out of scope for MVP Issue #1 but should be addressed before releases.
