# Homebrew Tap Strategy

## Recommendation
Use a separate tap repository (`homebrew-lfx`) instead of a folder inside the main repo. This keeps Homebrew metadata isolated, enables clean PR-based updates, and avoids coupling release automation to the main repo layout.

## Local Dev Install
For local testing without a tap, install the repo formula directly:
```bash
./scripts/lfx-dev
```
This script builds a local tarball and installs a temporary formula for the current machine.
It creates a local tap (default `lfx/local`) and installs `lfx` from it.

## Tap Repo Structure
```
homebrew-lfx/
  Formula/
    lfx.rb
  README.md
```

## Formula Template (tap repo)
Create `Formula/lfx.rb` with the following structure:
```ruby
class Lfx < Formula
  desc "Fast, pretty-by-default control-plane CLI for lf"
  homepage "https://github.com/brandon-gottshall/lfx"
  version "{{VERSION}}"

  on_macos do
    on_arm do
      url "https://github.com/brandon-gottshall/lfx/releases/download/v#{version}/lfx-darwin-arm64.tar.gz"
      sha256 "{{DARWIN_ARM64_SHA256}}"
    end
    on_intel do
      url "https://github.com/brandon-gottshall/lfx/releases/download/v#{version}/lfx-darwin-amd64.tar.gz"
      sha256 "{{DARWIN_AMD64_SHA256}}"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/brandon-gottshall/lfx/releases/download/v#{version}/lfx-linux-arm64.tar.gz"
      sha256 "{{LINUX_ARM64_SHA256}}"
    end
    on_intel do
      url "https://github.com/brandon-gottshall/lfx/releases/download/v#{version}/lfx-linux-amd64.tar.gz"
      sha256 "{{LINUX_AMD64_SHA256}}"
    end
  end

  def install
    bin.install "lfx"
  end

  test do
    system "#{bin}/lfx", "--help"
  end
end
```

## Automated Update Approach
- On release tag in the main repo:
  - Use `dist/checksums.txt` and `dist/artifacts.txt` from the release workflow.
  - A separate workflow (or a small script) updates the tap formula with new version and checksums.
  - Open a PR against `homebrew-lfx` with the updated `Formula/lfx.rb`.

Minimal automation option:
- A GitHub Action in this repo that runs on tag, checks out the tap repo, updates `Formula/lfx.rb`, and opens a PR using a bot token.

Rationale: keeps formula updates reviewable and avoids direct pushes.
