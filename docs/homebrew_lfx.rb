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
