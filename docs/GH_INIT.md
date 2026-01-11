# GitHub Initialization Checklist

## 1) Local Git Init
- `git init`
- `git add .`
- `git commit -m "chore: initial commit"`
- `git branch -M main`

## 2) Create GitHub Repo
- Repo name (placeholder): `lfx`
- Description (placeholder): "Fast, pretty-by-default control-plane CLI for lf themes, icons, and plugins."
- Visibility: Public
- License: MIT

## 3) Branch Strategy
- Protected branch: `main`
- Feature branches: `feature/<short-name>` or `fix/<short-name>`
- Merge via PRs only

## 4) Commit Style
- Use Conventional Commits: `feat:`, `fix:`, `docs:`, `chore:`, `refactor:`, `test:`
- Keep subject lines under 72 chars

## 5) Issue & PR Templates (Minimal)
- Add `.github/ISSUE_TEMPLATE/bug_report.md`
- Add `.github/ISSUE_TEMPLATE/feature_request.md`
- Add `.github/PULL_REQUEST_TEMPLATE.md`

Suggested PR template fields:
- Summary
- Testing
- Related issues

## 6) GitHub Actions CI
- Workflows under `.github/workflows/`
- Required jobs:
  - Build: macOS + Linux
  - Test: unit/integration
  - Lint/format: `gofmt` + lint tool (pick one)

## 7) Release Flow
- Tag release: `git tag -a vX.Y.Z -m "vX.Y.Z"`
- Push tag: `git push origin vX.Y.Z`
- CI builds artifacts for macOS and Linux
- Attach binaries and checksums to GitHub Release
- Update Homebrew tap formula and bump version

## 8) Supply Chain Basics
- Publish SHA256 checksums alongside artifacts
- Keep release notes minimal but explicit
- Optional: signed tags for releases
