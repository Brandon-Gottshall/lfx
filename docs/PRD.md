# Product Requirements Document (PRD)

## Overview
lfx is a fast, pretty-by-default external control-plane CLI for lf. It manages themes, icons, and plugins by writing files into `~/.config/lf` while keeping `lf` itself untouched.

## Related Docs
- `docs/CONFIG.md`
- `docs/REGISTRY.md`
- `docs/CLI.md`
- `docs/INVARIANTS.md`
- `docs/SECURITY.md`
- `docs/process/agent-work-rules.md`
- `docs/CHECKPOINTS.md`

## Goals
- Provide a simple CLI to list, install, and apply themes, icons, and plugins.
- Keep installs reproducible and idempotent, especially for `~/.config/lf/lfrc` edits.
- Offer a curated theme set vendored in-repo for fast, offline usage.
- Treat icon packs as external assets referenced via metadata; download only on demand.
- Ship a minimal but useful plugin example (hotkeys HUD) to validate the flow.

## MVP Scope
- Commands: `list`, `set/apply`, `install`, `uninstall` (per extension type).
- Theme selection: curated community themes in `registry/themes/<theme-name>/colors`.
- Icon selection: metadata-only entries in `registry/icons/<icon-pack>.json`; CLI fetches and installs assets when set.
- Plugins: hotkeys HUD as a simple `registry/plugins/hotkeys-hud.lfrc` snippet installable via CLI.
- Idempotent edits to `~/.config/lf/lfrc` with no duplicate lines.
- Config file: `~/.config/lfx/config.toml` (TOML, user-editable source of truth).

## Non-Goals
- No lf core modifications.
- No automated merging of ruler templates.
- No interactive TUI picker (optional later).
- No network calls unless explicitly required by a command (e.g., `lfx icons set <pack>`).

## Target Users
- lf users who want a managed, consistent extension setup.
- Contributors maintaining curated themes and plugin snippets.
- Automation/agent workflows that need deterministic, inspectable changes.

## User Stories
- As a user, I can list available themes and apply one in one command.
- As a user, I can install a hotkeys HUD plugin without manually editing `lfrc`.
- As a contributor, I can add a theme via a PR using the filesystem conventions.

## Success Criteria
- Applying the same theme twice does not duplicate config lines.
- Installing the hotkeys HUD produces a predictable `.lfrc` section.
- Themes are usable without any network access after clone.
- Icon packs are only downloaded when explicitly selected.

## Example Commands
- `lfx themes list`
- `lfx themes set catppuccin`
- `lfx icons list`
- `lfx icons set nvim-devicons`
- `lfx plugins install hotkeys-hud`
- `lfx plugins uninstall hotkeys-hud`
