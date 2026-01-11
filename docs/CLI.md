# CLI Spec (MVP)

## Related Docs
- `docs/PRD.md`
- `docs/REGISTRY.md`
- `docs/INVARIANTS.md`
- `docs/CONFIG.md`

## Principles
- No interactive UI.
- Fast startup and minimal IO.
- Idempotent config edits.
- `~/.config/lf` is a deployment target only.

## Commands

### lfx dev:tests
Run the MVP test suite in a development environment.

Example:
```bash
lfx dev:tests
```

### lfx theme list
List vendored themes from `registry/themes/<theme-name>/colors`.

Example:
```bash
lfx theme list
```

### lfx theme set <theme-name>
Apply a theme by copying its `colors` file into `~/.config/lf/colors/` and updating the `lfrc` block.

Example:
```bash
lfx theme set catppuccin
```

### lfx icons list
List icon packs from `registry/icons/<icon-pack>.json`.

Example:
```bash
lfx icons list
```

### lfx icons set <icon-pack>
Download and apply the icon pack described by the metadata file. Validate checksum and update `lfrc` block.

Example:
```bash
lfx icons set nvim-devicons
```

### lfx plugin list
List plugins from `registry/plugins/<plugin-name>.lfrc`.

Example:
```bash
lfx plugin list
```

### lfx plugin install <plugin-name>
Install plugin snippet into `~/.config/lf/plugins/` and update `lfrc` block.

Example:
```bash
lfx plugin install hotkeys-hud
```

### lfx doctor
Verify target layout and registry presence. Recommended but optional.

Example:
```bash
lfx doctor
```

## Config Block Convention
Use a stable, scoped block in `~/.config/lf/lfrc`, e.g.:
```text
# lfx:begin
# lfx:managed
# lfx:end
```
The block must be rewritten deterministically.
