# ifx

ifx is an external control-plane CLI for lf that manages themes, icons, and plugins by writing files into `~/.config/lf`.

## Quick Start
```bash
brew tap brandon-gottshall/ifx
brew install ifx

lfx theme list
```

## Tooling
- Dependencies are managed with asdf via `.tool-versions`.
- Install Go 1.25.5 (pinned) before building or running tests.

## Configuration
- Config file: `~/.config/ifx/config.toml` (TOML, user-editable source of truth).
- See `docs/CONFIG.md` for schema and defaults.

## Docs Index
- `docs/PRD.md`
- `docs/MVP_PLAN.md`
- `docs/CHECKPOINTS.md`
- `docs/CLI.md`
- `docs/REGISTRY.md`
- `docs/CONFIG.md`
- `docs/INVARIANTS.md`
- `docs/SECURITY.md`
- `docs/INSTALL.md`
- `docs/DEVELOPMENT.md`
- `docs/CONTRIBUTING.md`
- `docs/GH_INIT.md`
- `docs/ISSUES_MVP.md`
- `docs/process/agent-work-rules.md`
- `docs/process/agent-assisted-development-contract.md`

## Security
Go 1.25.5 is the latest patch release of the Go 1.25 series at time of pinning. Known toolchain-level CVEs prior to 1.25.5 (notably in `crypto/x509`) are patched in this version. As of the pin date, there are no known unpatched CVEs in the Go standard library affecting typical CLI workloads.

Go’s security guarantees do not extend to third-party modules, application logic, or handling of untrusted input beyond stdlib contracts. Security posture is conditional: CLI tools that do not parse untrusted certificates or operate as network services are outside the known CVE impact envelope, while tools that do process untrusted certs, HTTP traffic, or structured attacker-controlled input must rely on correct application-level handling.

Pinning to Go 1.25.5 mitigates known toolchain vulnerabilities as of the pin date, but does not eliminate dependency- or application-level risk.

## Status
This repository is in early setup; see `docs/` for project plans and design docs.
