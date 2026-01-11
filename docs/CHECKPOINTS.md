# MVP Checkpoint Order

This order prioritizes fast development cycles and low-merge-risk changes first, while deferring higher‑risk or cross‑cutting work until core foundations are stable.

## 1) Registry Seed + Theme Listing
- Add minimal theme content and list output.
- Lowest risk; isolated to registry read logic.

## 2) Config Loader Warning
- Read `~/.config/lfx/config.toml` and warn on missing/invalid.
- Minimal surface area and clear failure modes.

## 3) Doctor Baseline
- `lfx doctor` checks registry presence and `~/.config/lf` existence.
- Independent from install/apply logic.

## 4) Theme Apply + Idempotent `lfrc`
- Implement `theme set` with deterministic file write + scoped config block.
- This is the core install primitive for other extension types.

## 5) Plugin Install (hotkeys-hud)
- Reuse `lfrc` block rules; install snippet into target.
- Builds directly on theme/apply machinery.

## 6) Icons Set (metadata download + checksum)
- Introduces network + archive handling; keep after idempotent edits are proven.

## 7) Dev Test Harness
- `lfx dev:tests` to validate MVP behaviors.
- Hook in once core commands exist.

## 8) Release + Homebrew Automation
- Binaries, checksums, release assets, tap updates.
- Best done after CLI behavior stabilizes.
