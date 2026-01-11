# MVP Issues (Planned)

This document defines the MVP issue set and the conservative execution order. Create GitHub issues from this list before implementation.

## Related Docs
- `docs/CHECKPOINTS.md`
- `docs/MVP_PLAN.md`
- `docs/PRD.md`
- `docs/CLI.md`
- `docs/REGISTRY.md`
- `docs/CONFIG.md`
- `docs/INVARIANTS.md`
- `docs/SECURITY.md`
- `docs/HOMEBREW_TAP.md`
- `docs/process/agent-work-rules.md`

## Order (Conservative)
1) Registry seed + theme list — Issue #1
2) Config loader warning — Issue #2
3) Doctor baseline — Issue #3
4) Theme apply + idempotent lfrc — Issue #4
5) Plugin install (hotkeys-hud) — Issue #5
6) Icons set (metadata download + checksum) — Issue #6
7) Dev test harness (lfx dev:tests) — Issue #7
8) Release + Homebrew tap automation — Issue #8

## GitHub Issue Links
- #1 https://github.com/Brandon-Gottshall/lfx/issues/1
- #2 https://github.com/Brandon-Gottshall/lfx/issues/2
- #3 https://github.com/Brandon-Gottshall/lfx/issues/3
- #4 https://github.com/Brandon-Gottshall/lfx/issues/4
- #5 https://github.com/Brandon-Gottshall/lfx/issues/5
- #6 https://github.com/Brandon-Gottshall/lfx/issues/6
- #7 https://github.com/Brandon-Gottshall/lfx/issues/7
- #8 https://github.com/Brandon-Gottshall/lfx/issues/8

## Issue Creation Script
- `scripts/create_mvp_issues.sh` creates the MVP issue set from this document.

## Issue Set

## Recommended Models (Per Issue)
- GPT 5.1-codex-mini (low-high reasoning)
- Claude 4.5 Haiku
- GPT 5.2-codex (Low-xhigh reasoning)
- Claude 4.5 Sonnet
- Claude 4.5 Opus

### 1. Seed registry + theme list
Create the minimal registry data and implement `lfx theme list` from `registry/themes/<theme-name>/colors`.

Acceptance:
- Listing is deterministic and ordered.
- Seed theme `default-dark` is visible.

### 2. Config loader warning
Load `~/.config/lfx/config.toml` on startup and warn on missing/invalid config before falling back to defaults.

Acceptance:
- Missing config emits a warning and continues.
- Invalid config emits a warning and continues with defaults.

### 3. Doctor baseline
Implement `lfx doctor` to check registry presence and `~/.config/lf` existence with clear exit codes.

Acceptance:
- Reports missing registry or lf config with actionable guidance.

### 4. Theme apply + idempotent lfrc
Implement `lfx theme set <name>` to copy colors into `~/.config/lf/colors` and update a scoped, idempotent `lfrc` block.

Acceptance:
- Re-running does not duplicate lines.
- Switching themes updates deterministically.

### 5. Plugin install (hotkeys-hud)
Implement `lfx plugin install <name>` to install `registry/plugins/<plugin-name>.lfrc` into `~/.config/lf/plugins` and update the same `lfrc` block.

Acceptance:
- Install is idempotent and scoped.

### 6. Icons set (metadata download + checksum)
Implement `lfx icons set <pack>` to read `registry/icons/<icon-pack>.json`, download assets, verify checksum, and install into `~/.config/lf/icons`.

Acceptance:
- Network only on explicit command.
- Checksum mismatch aborts without changes.

### 7. Dev test harness
Add `lfx dev:tests` to run the MVP test suite in a development environment.

Acceptance:
- Returns non-zero on failures.
- Covers MVP acceptance tests in `docs/MVP_PLAN.md`.

### 8. Release + Homebrew tap automation
Finalize release artifacts, checksums, and tap update PR flow.

Acceptance:
- Release produces binaries + checksums + artifact URLs.
- Homebrew tap update PR is generated.
