# Agent Work Rules (Operational)

This document is the mechanical contract for agent-assisted changes in this repo.

## Correctness Oracle
Before claiming “done”, run:
```bash
make build
make test
make lint
```
If any command fails, the change is not complete.

## Default Scope Constraints
- Smallest diff that achieves the objective.
- No refactors unless required for correctness.
- No dependency changes unless explicitly requested.
- No file moves/renames unless explicitly requested.
- Touch only the files required by the task.

## Boundaries
- `cmd/lfx/`: CLI entrypoint + argument handling.
- `internal/`: implementation packages (registry, install, config, paths, doctor, UI).
- `registry/`: data-only catalog of themes/plugins/icons.
- `docs/`: contracts and planning.

Changes that cross boundaries must be explicitly justified in the PR.

## Safety Rules
- Do not write into `~/.config/lf` except as install/apply targets.
- Do not upgrade the Go toolchain without re-evaluating CVEs and updating docs.
- Keep config edits idempotent and scoped.

## Plan Readiness Checklist
Implementation plans must list:
- Oracle commands (`make build`, `make test`, `make lint`).
- New invariants and where they are enforced.
- Rollback strategy if behavior regresses.
