# Issue 4 Plan: Theme Apply + Idempotent `lfrc`

This plan is intentionally exhaustive to prevent unplanned yak shaving. It is based on a ruthless discovery pass against current code and docs.

## Discovery Snapshot (Read-Only Facts)

### Current Code State
- CLI commands implemented: `lfx theme list`, `lfx doctor` (`cmd/lfx/main.go`).
- Install/apply logic is stubbed (`internal/install/install.go`).
- Registry list exists for themes (`internal/registry/registry.go`).
- Paths helpers: `LfConfigDir` and `LfxConfigDir` (`internal/paths/paths.go`).
- Theme registry entries: `registry/themes/<name>/colors` (`registry/themes/default-dark/colors`, `registry/themes/warp-phenomenon/colors`).

### Current Docs (Behavioral Requirements)
- Issue #4: Apply theme + idempotent lfrc (`docs/ISSUES_MVP.md`).
- CLI spec: `lfx theme set <name>` copies colors into `~/.config/lf/colors` and updates `lfrc` block (`docs/CLI.md`).
- Invariants: idempotent `lfrc` edits, scoped block (`docs/INVARIANTS.md`).
- Checkpoints: Issue 4 is the install/apply primitive (`docs/CHECKPOINTS.md`).
- MVP plan: “copy colors file into `~/.config/lf/colors`, update lfrc block, add backup/predictable overwrite” (`docs/MVP_PLAN.md`).
- Contributing: block may be named for plugins (`docs/CONTRIBUTING.md`).

### Known Decisions (From User)
- Theme destination is a file: `~/.config/lf/colors`.

### Pending Decisions (Must Be Locked Before Code)
- Exact `lfrc` block markers and layout.
  - Proposed: `# lfx:begin` / `# lfx:managed` / `# lfx:end` (matches `docs/CLI.md`).
  - Need to decide whether to include a single managed block for all installs or a per-feature block (e.g., `# lfx:begin theme`).

## Goals & Non-Goals

### Goals (Issue #4)
- Implement `lfx theme set <name>`:
  - Copy `registry/themes/<name>/colors` -> `~/.config/lf/colors`.
  - Update a scoped, idempotent `~/.config/lf/lfrc` block.
- Re-running is idempotent; switching themes is deterministic.
- Clear errors on missing registry or file/permission issues.

### Non-Goals (Deferred)
- Icons set, plugins install (Issues #5/#6).
- Registry root resolution beyond current `os.Getwd()` logic.
- Config-driven theme defaults or persistence beyond `lfrc` block.

## Ruthless Step-by-Step Plan

### 0) Lock Decisions (Before Coding)
1) Confirm `lfrc` block markers (standardize across docs and code):
   - Option A (default): `# lfx:begin`, `# lfx:managed`, `# lfx:end` (from `docs/CLI.md`).
   - Option B: `# lfx:begin theme` / `# lfx:end theme` if you want scoped per-feature blocks.
2) Confirm whether to write a backup of `lfrc`:
   - MVP plan suggests “backup or predictable overwrite policy”.
   - If backup is required now, specify naming (e.g., `lfrc.lfx.bak`).

### 1) Documentation Alignment (Only If Needed)
If Step 0 changes any conventions:
1) Update `docs/CLI.md` block example to match chosen markers.
2) Update `docs/INVARIANTS.md` block naming.
3) Update `docs/CONTRIBUTING.md` for block naming rule.
4) Update `docs/MVP_PLAN.md` if backup policy is decided.

### 2) Implement Theme Apply Command Wiring
1) Add CLI command `lfx theme set <name>`:
   - Validate args.
   - Resolve repo root (`os.Getwd()` for now).
   - Construct registry theme colors path: `registry/themes/<name>/colors`.
   - Resolve target path: `paths.LfConfigDir()/colors`.
   - Delegate to install logic.
2) Update `printHelp()` to include `theme set`.

### 3) Implement Install Logic in `internal/install`
Build a minimal, reusable install helper that:
1) Ensures `~/.config/lf` exists or errors with actionable message.
2) Copies the colors file to `~/.config/lf/colors`:
   - Use `os.ReadFile` + `os.WriteFile`.
   - Preserve line endings; do not mutate theme file content.
3) Updates `~/.config/lf/lfrc` with a scoped managed block:
   - Read file if exists; treat missing as empty.
   - Replace existing block if found; else append new block.
   - Must be deterministic: same input yields same file.
   - No duplicate lines when re-running.
4) Block contents for theme set (minimum viable):
   - Comment line referencing the colors file path (lf loads it automatically).
   - Block structure example (subject to decision):
     ```
     # lfx:begin
     # lfx:managed
     # colors file at ~/.config/lf/colors
     # lfx:end
     ```
5) If backup policy is adopted, create/overwrite `lfrc.lfx.bak` before writing.

### 4) Testing (Must Be Meaningful)
Create tests to ensure behavior, not just coverage.
1) Add unit tests in `internal/install`:
   - `TestWriteLfrcBlock_Idempotent`: run twice, file is identical.
   - `TestWriteLfrcBlock_ReplacesExistingBlock`: old block replaced; unrelated content preserved.
   - `TestApplyTheme_CopiesColors`: target colors file matches source.
2) Add CLI-level smoke test:
   - `cmd/lfx/main_test.go`: run `go run ./cmd/lfx theme set default-dark` with temp `XDG_CONFIG_HOME` to avoid touching real `~/.config/lf`.
   - Assert `colors` and `lfrc` exist under temp.
3) Update tests to use temp directories and avoid global state.

### 5) Error Handling & UX
1) Ensure missing theme returns a clear error:
   - "theme not found: <name>" and exit non-zero.
2) Ensure missing registry returns a clear error:
   - re-use existing doctor guidance or new errors.
3) Ensure permission errors are surfaced verbatim.

### 6) Verification Commands
Run locally:
- `make build`
- `make test`
- `make lint`

### 7) Final Documentation Touches
If behavior differs from docs after implementation, update:
- `docs/CLI.md`
- `docs/INVARIANTS.md`
- `docs/MVP_PLAN.md`

## Risk Checklist
- Avoid writing anywhere outside `~/.config/lf`.
- Ensure `lfrc` edits are scoped and idempotent.
- Avoid embedding repo paths or ownership in file content.

## Exit Criteria (Issue #4)
- `lfx theme set <name>` applies colors file and lfrc block.
- Re-running does not duplicate lfrc entries.
- Switching themes updates deterministically.
- Tests cover install behavior and CLI-level flow.
