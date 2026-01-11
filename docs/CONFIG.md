# Configuration

## Related Docs
- `docs/PRD.md`
- `docs/CLI.md`
- `docs/INVARIANTS.md`

## Location
- Primary config: `~/.config/lfx/config.toml`
- Cache (future): `~/.cache/lfx/`
- State (future): `~/.local/share/lfx/`

## Format
- TOML
- User-editable and source of truth
- No runtime state or cache data

## Schema (v1)
```toml
config_version = 1

[ui]
# Theme name from registry/themes/<theme-name>/colors
# Icons name from registry/icons/<icon-pack>.json
# Font is a system-available default for macOS and Linux
# Use "monospace" to defer to OS defaults (e.g., Menlo or DejaVu Sans Mono)

theme = "default-dark"
icons = "nerd-fonts"
font  = "monospace"

[extensions]
# Explicit enablement per extension
hotkeys-hud = true

[paths]
# Optional overrides (defaults are internal)
themes = "~/.config/lfx/themes"
icons  = "~/.config/lfx/icons"
extensions = "~/.config/lfx/extensions"
```

## Rules
- Presence does not imply activation; use explicit booleans.
- Invalid config must fail loudly with actionable errors.
- No silent rewrites; keep user edits intact.
- Defaults are part of the public API and change only with intent.
- The CLI loads config on startup; missing config falls back to defaults.
- If the config is missing or invalid, the CLI should emit a warning before continuing with defaults.
