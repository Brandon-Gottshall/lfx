#!/usr/bin/env bash
set -euo pipefail

REPO="brandon-gottshall/lfx"

create_issue() {
  local title="$1"
  local body_file="$2"
  gh issue create --repo "$REPO" --title "$title" --body-file "$body_file"
}

cat <<'BODY' > /tmp/lfx_issue_1.md
Implement `lfx theme list` by reading `registry/themes/<theme-name>/colors` and add a minimal `default-dark` theme.

Acceptance:
- Listing is deterministic and ordered.
- Seed theme `default-dark` is visible.
BODY

cat <<'BODY' > /tmp/lfx_issue_2.md
Load `~/.config/lfx/config.toml` on startup and warn on missing or invalid config before falling back to defaults.

Acceptance:
- Missing config emits a warning and continues.
- Invalid config emits a warning and continues with defaults.
BODY

cat <<'BODY' > /tmp/lfx_issue_3.md
Implement `lfx doctor` to check registry presence and `~/.config/lf` existence with clear exit codes.

Acceptance:
- Reports missing registry or lf config with actionable guidance.
BODY

cat <<'BODY' > /tmp/lfx_issue_4.md
Implement `lfx theme set <name>` to copy colors into `~/.config/lf/colors` and update a scoped, idempotent `lfrc` block.

Acceptance:
- Re-running does not duplicate lines.
- Switching themes updates deterministically.
BODY

cat <<'BODY' > /tmp/lfx_issue_5.md
Implement `lfx plugin install <name>` to install `registry/plugins/<plugin-name>.lfrc` into `~/.config/lf/plugins` and update the same `lfrc` block.

Acceptance:
- Install is idempotent and scoped.
BODY

cat <<'BODY' > /tmp/lfx_issue_6.md
Implement `lfx icons set <pack>` to read `registry/icons/<icon-pack>.json`, download assets, verify checksum, and install into `~/.config/lf/icons`.

Acceptance:
- Network only on explicit command.
- Checksum mismatch aborts without changes.
BODY

cat <<'BODY' > /tmp/lfx_issue_7.md
Add `lfx dev:tests` to run the MVP test suite in a development environment.

Acceptance:
- Returns non-zero on failures.
- Covers MVP acceptance tests in `docs/MVP_PLAN.md`.
BODY

cat <<'BODY' > /tmp/lfx_issue_8.md
Finalize release artifacts, checksums, and tap update PR flow.

Acceptance:
- Release produces binaries + checksums + artifact URLs.
- Homebrew tap update PR is generated.
BODY

create_issue "mvp-1: seed registry + theme list" /tmp/lfx_issue_1.md
create_issue "mvp-2: config loader warning" /tmp/lfx_issue_2.md
create_issue "mvp-3: doctor baseline" /tmp/lfx_issue_3.md
create_issue "mvp-4: theme apply + idempotent lfrc" /tmp/lfx_issue_4.md
create_issue "mvp-5: plugin install (hotkeys-hud)" /tmp/lfx_issue_5.md
create_issue "mvp-6: icons set (metadata download + checksum)" /tmp/lfx_issue_6.md
create_issue "mvp-7: dev test harness" /tmp/lfx_issue_7.md
create_issue "mvp-8: release + Homebrew tap automation" /tmp/lfx_issue_8.md
