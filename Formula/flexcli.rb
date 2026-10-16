class Flexcli < Formula
  desc "Management CLI for FlexCoach AI fitness platform"
  homepage "https://github.com/f1dot4/homebrew-flexcli"
  url "https://github.com/f1dot4/homebrew-flexcli/archive/refs/tags/v0.3.0.tar.gz"
  sha256 "532f4de72d604d58a114341e832332d87c1810acd82480547e003f77162d50f8"
  license "MIT"

  depends_on "go" => :build

  def install
    # Build the binary from the cmd/flexcli directory
    # std_go_args handles common flags for brew-built Go apps
    system "go", "build", *std_go_args(output: bin/"flexcli"), "./cmd/flexcli"
  end

  test do
    # Simple check to ensure the binary runs and shows help
    output = shell_output("#{bin}/flexcli help")
    assert_match "FlexCLI - FlexCoach Command Line Interface", output
  end
end
