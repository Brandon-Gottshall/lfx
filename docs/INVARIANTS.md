# Invariants

## Hard Constraints
- `lf` is not modified; only files under `~/.config/lf` are written as install targets.
- The CLI does not live inside `~/.config/lf` and does not treat it as a workspace.
- No network access unless explicitly required by the invoked command.
- Config edits to `~/.config/lf/lfrc` are idempotent with no duplicate lines.
- All installs are inspectable on disk (no hidden state or opaque caches).
- Registry entries are filesystem-first and human-readable.
- Do not bake GitHub ownership or URLs into code paths; keep them in docs and CI only.

## Configuration Safety
- `lfrc` modifications must be scoped to a clearly marked block (e.g., `# lfx:begin` / `# lfx:end`).
- CLI must never overwrite user-owned files outside its declared targets.
- Uninstall must fully remove only what lfx installed.

## Failure Modes and Required Behavior
- Missing permissions in `~/.config/lf`: abort with a clear error; do not partially install.
- Network failure when fetching icon packs: abort; leave existing assets intact.
- Checksum or signature mismatch: abort; do not install corrupted assets.
- Partial writes (e.g., interrupted install): detect on next run and repair or rollback.

## Recovery Expectations
- `lfx doctor` (or equivalent) should report incomplete installs and suggest fixes.
- Re-running `install` or `set` should converge to the same final state.
