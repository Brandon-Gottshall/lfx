# MVP Implementation Plan (lfx)

## Related Docs
- `docs/PRD.md`
- `docs/CLI.md`
- `docs/REGISTRY.md`
- `docs/INVARIANTS.md`
- `docs/SECURITY.md`
- `docs/CONFIG.md`
- `docs/process/agent-work-rules.md`
- `docs/CHECKPOINTS.md`
- `docs/ISSUES_MVP.md`

## Phase 0: Repo Wiring
Goal: establish buildable Go CLI, docs, and CI for fast iteration.
- Create Go module and `cmd/lfx` entry.
- Wire lipgloss-based output helpers.
- Add Makefile targets for build/test/lint.
- Add agent-work rules to keep changes scoped and reviewable.

Acceptance Tests:
- `go build ./cmd/lfx` completes in <2s on a warm cache.
- `lfx --help` renders styled output without errors.

## Phase 1: Registry + Listing
Goal: parse filesystem-first registry without manifests.
- Implement registry scan for:
  - `registry/themes/<theme-name>/colors`
  - `registry/plugins/<plugin-name>.lfrc`
  - `registry/icons/<icon-pack>.json`
- Implement `lfx theme list`, `lfx icons list`, `lfx plugin list`.
- Add a minimal seed theme under `registry/themes/default-dark/colors` for smoke tests.

Acceptance Tests:
- Listing is deterministic and ordered.
- Empty registry returns a clean message and exit 0.

## Phase 2: Apply/Install Logic
Goal: install/apply with reversible edits and idempotency.
- Theme set: copy colors file into `~/.config/lf/colors`, update `lfrc` block.
- Icons set: download pack, verify checksum, install into `~/.config/lf/icons`, update `lfrc` block.
- Plugin install: install `.lfrc` snippet and update `lfrc` block.
- Add backup or predictable overwrite policy for `lfrc`.
- Load config from `~/.config/lfx/config.toml`; warn on missing or invalid config before falling back to defaults.
- Add `lfx dev:tests` to run the MVP suite in a development environment.

Acceptance Tests:
- “lfx dev:tests runs the MVP suite in a development environment”
- “fresh machine -> brew install -> set theme/icons -> install hotkeys-hud -> open lf and see it applied”
- “rerun same commands -> no duplicate lfrc lines”
- “switch theme -> colors file changes deterministically”
- “lfx doctor reports missing registry or lf config when absent”
- “missing or invalid config prints a warning and continues with defaults”

## Phase 3: Doctor + Hardening
Goal: verify state and surface problems quickly.
- Implement `lfx doctor` to check registry presence and `~/.config/lf` existence.
- Add clear error messages and exit codes.

Acceptance Tests:
- `lfx doctor` reports missing `~/.config/lf` with actionable guidance.

## Phase 4: Release & Homebrew
Goal: ship installable artifacts and brew tap.
- GitHub Actions: build/test/lint on macOS + Linux.
- Release job: build binaries, checksums, attach to release.
- Homebrew tap formula updated on release.

Acceptance Tests:
- Tagging a release produces binaries + checksums.
- `brew tap <owner>/lfx && brew install lfx` succeeds on macOS.
