# Contributing

## Related Docs
- `docs/REGISTRY.md`
- `docs/PRD.md`
- `docs/INVARIANTS.md`
- `docs/SECURITY.md`

## Repo Model
- Single mono-repo for CLI, registry metadata, and vendored themes.
- All changes flow through PRs; no direct commits to main.

## What Goes Where
- Themes: add under `registry/themes/<name>/` and vendor the theme file.
- Plugins: add under `registry/plugins/<name>/` with a file-based snippet.
- Icons: add metadata under `registry/icons/<name>/` only; link to upstream assets.

## Contribution Workflow
1) Fork and create a topic branch.
2) Add or update the registry entry.
3) Include a short README or manifest notes if behavior is non-obvious.
4) Open a PR with a concise description and example usage.

## Tooling
- Use asdf for dependencies; `.tool-versions` is the source of truth.
- Ensure the pinned Go version is installed before running builds or tests.

## Theme Contributions
- Provide a single sourceable file (e.g., `theme.lf`).
- Ensure naming is `kebab-case` and matches manifest `name`.

## Plugin Contributions
- Keep snippets minimal and idempotent.
- Any `lfrc` edits must be scoped with a named block (e.g., `# lfx:begin hotkeys-hud`).

## Icon Pack Contributions
- Do not vendor assets.
- Provide stable URLs, checksums, and a clear unpack layout.

## PR Checklist
- Registry entry matches the filesystem conventions in `docs/REGISTRY.md`.
- Example command in PR description (e.g., `lfx themes set <name>`).
- No runtime modification of `lf` beyond install targets.
