# Agent-Assisted Development Contract

This repo is designed for agent-first development: explicit contracts, deterministic tooling, and low-noise diffs. Agents must optimize for correctness, reproducibility, and minimal surface area changes.

## Core Principles
- Make changes predictable and reviewable.
- Favor explicit configuration over hidden state.
- Keep outputs deterministic and idempotent.
- Respect the toolchain pin and security boundaries.

## Authority Boundaries
- `docs/` defines the contract (CLI behavior, registry layout, config schema).
- `AGENTS.md` defines operational invariants and security boundaries.
- CI is the enforcement mechanism; failing CI means the change is not done.
- `docs/DEVELOPMENT.md` defines the local dev and hook workflow.

## Expectations
- Keep changes scoped to the task.
- Avoid repo-wide formatting unless instructed.
- Update docs when behavior changes.
