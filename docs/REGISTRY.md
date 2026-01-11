# Registry Format

## Overview
The registry is filesystem-first and human-readable. It lives in-repo under `registry/` and is the source of truth for what lfx can install. The MVP uses conventions instead of manifest files.

## Related Docs
- `docs/PRD.md`
- `docs/CLI.md`
- `docs/CONTRIBUTING.md`
- `docs/INVARIANTS.md`

## Layout (MVP Contract)
- Themes: `registry/themes/<theme-name>/colors`
- Plugins: `registry/plugins/<plugin-name>.lfrc`
- Icons: `registry/icons/<icon-pack>.json` (metadata only, no assets)

## Sample Content
- A minimal theme is included at `registry/themes/default-dark/colors` for smoke testing.

## Naming Conventions
- Directory and file names are `kebab-case`.
- Theme directory name is the theme identifier used by the CLI.
- Plugin file name (without extension) is the plugin identifier used by the CLI.
- Icon pack file name (without extension) is the pack identifier used by the CLI.

## Icon Pack Metadata (JSON)
Icon packs are metadata-only and fetched on demand. Keep the schema minimal and stable.

Example:
```json
{
  "name": "nvim-devicons",
  "url": "<upstream_tarball_url>",
  "checksum": "sha256:<hex>",
  "unpack_dir": "<top_level_dir>"
}
```

## Conventions
- Themes are vendored; no network required to list or apply them.
- Icon packs are fetched only when explicitly set or installed.
- Plugins ship as file-based snippets; any `lfrc` edits must be idempotent and scoped.
- Avoid embedding ownership or hosting assumptions in the registry content.
