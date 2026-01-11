# Repository Guidelines

lfx is an open-source (MIT & Public) project owned by Moons Out Labs, a legally declared DBA for Moons Out Media LLC (Ohio). Keep this document accurate as the repo evolves.

## Project Identity
- Name: lfx (lf control‑plane / extension CLI).
- Owner: Moons Out Labs (DBA for Moons Out Media LLC, Ohio).
- License: MIT; see `LICENSE` for the full text.

## Project Goals & Boundaries
- Purpose: an external control-plane CLI and registry for lf extensions (themes, icons, plugins, HUDs).
- lf is not modified; `~/.config/lf` is a deployment target only and never a workspace.
- Extensions are file-based and sourceable; distribution is via GitHub (mono-repo + PRs).
- Design must be stable, inspectable, and agent-friendly.
- Non-goals: no runtime plugin API inside lf, no UI rendering inside the CLI, no hidden state or magic.
- Config is stored at `~/.config/lfx/config.toml` (TOML, user-editable source of truth).

## Architecture Summary
- Control plane: `lfx` CLI (Go) manages registry, install/apply logic, and writes files into `~/.config/lf`.
- Runtime: lf loads the installed files; no runtime API changes or embedded logic.

## Project Structure & Module Organization
This repository is currently a minimal scaffold for the lf control‑plane CLI. As code lands, keep it organized by responsibility:
- `cmd/lfx/`: CLI entry point and command wiring (fast startup).
- `internal/`: registry resolution, install/apply logic, config, and tooling.
- `internal/config/`: config loader for `~/.config/lfx/config.toml`.
- `assets/`: bundled templates or default registry metadata.
- `tests/`: unit/integration tests and fixtures.
- `docs/`: design notes and registry format docs.
- `docs/process/`: agent-first development rules and contracts.
- `registry/`: filesystem-first registry (themes/plugins/icons).

If directories differ, update this section to match the actual tree.

## Build, Test, and Development Commands
No build system is committed yet. When added, document exact commands here with examples, e.g.:
- `make build`: compile the `lfx` binary.
- `make test`: run unit tests.
- `./lfx --help`: verify CLI help output.

## Tooling
- Dependencies are managed via asdf and `.tool-versions`. Keep Go version pinned and updated there.

## Coding Style & Naming Conventions
- Language: Go CLI with pretty output (lipgloss) and fast startup.
- Indentation: 2 spaces for Markdown/config, Go defaults for code.
- Naming: use `snake_case` for files, `kebab-case` for CLI subcommands, and `UpperCamelCase` for types.
- Keep registry schemas and lockfiles in lowercase keys.
- Code standards: small functions, explicit errors, deterministic behavior, no hidden state.
- Use `gofmt` and keep CLI output stable for automation.

## Toolchain Pinning
- Go is pinned to 1.25.5 and must be enforced via `go.mod` and `.tool-versions`.
- `go.mod` must use `go 1.25` with `toolchain go1.25.5`.
- Do not upgrade the Go version casually; reassess security boundaries first.

## Documentation-First Workflow
- When in doubt, document first. Treat docs as the primary source of truth.
- Before making planning or sequencing decisions, consult the relevant docs and update them if needed.
- Prefer adding internal links to keep the knowledge graph tight and discoverable.

## Issue Model Recommendations
- Default list for MVP issues:
  - GPT 5.1-codex-mini (low-high reasoning)
  - Claude 4.5 Haiku
  - GPT 5.2-codex (Low-xhigh reasoning)
  - Claude 4.5 Sonnet
  - Claude 4.5 Opus

## Go Toolchain Security Boundaries
Go 1.25.5 is the latest patch release of the Go 1.25 series at time of pinning. Known toolchain-level CVEs prior to 1.25.5 (notably in `crypto/x509`) are patched in this version. As of the pin date, there are no known unpatched CVEs in the Go standard library affecting typical CLI workloads.

Go’s security guarantees do not extend to third-party modules, application logic, or handling of untrusted input beyond stdlib contracts. Security posture is conditional: CLI tools that do not parse untrusted certificates or operate as network services are outside the known CVE impact envelope, while tools that do process untrusted certs, HTTP traffic, or structured attacker-controlled input must rely on correct application-level handling.

Pinning to Go 1.25.5 mitigates known toolchain vulnerabilities as of the pin date, but does not eliminate dependency- or application-level risk.
- Do not upgrade Go without re-evaluating CVEs and recording the rationale in docs.
- Use `govulncheck` for dependency analysis when adding or updating modules.
- Treat “Go is memory-safe” as a partial invariant, not a blanket guarantee.

## Testing Guidelines
- Prefer fast unit tests for registry parsing, dependency resolution, and install planning.
- Integration tests should use a temp target (e.g., `/tmp/lfx-test/`) instead of `~/.config/lf`.
- Name tests after behavior, e.g., `test_resolve_pin_exact_version`.

## Commit & Pull Request Guidelines
- No commit convention is established yet. Use Conventional Commits (`feat:`, `fix:`, `docs:`) to build a consistent history.
- PRs should include: a short problem statement, a summary of changes, and testing performed. Add screenshots only for CLI UX changes (e.g., help output).

## Release Expectations
- Versioned tags are required for releases.
- Homebrew install must be first-class (formula or tap), with a curl fallback documented in `docs/INSTALL.md`.

## Adding Extensions via PR
- Themes: add `registry/themes/<theme-name>/colors`; keep names `kebab-case`.
- Plugins: add `registry/plugins/<plugin-name>.lfrc` with a scoped block.
- Icons: add `registry/icons/<icon-pack>.json` only; link to upstream assets and include checksums.

## Configuration & Security
- Do not modify runtime lf configs directly; installs must target `~/.config/lf` only as an output.
- Registry is filesystem-first under `registry/`; avoid hidden caches.
- Maintain idempotent `lfrc` edits with no duplicate lines.
- Ensure any network access is opt-in (`lfx icons set <pack>`), and document checksums in lockfiles.
- Favor reproducible installs, explicit registries, and minimal defaults.
