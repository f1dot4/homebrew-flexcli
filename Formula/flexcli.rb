class Flexcli < Formula
  desc "Management CLI for FlexCoach AI fitness platform"
  homepage "https://github.com/f1dot4/homebrew-flexcli"
  url "https://github.com/f1dot4/homebrew-flexcli/archive/refs/tags/v0.1.7.tar.gz"
  sha256 "0b9f0aa4d2b6c29d7f38b43a0d5bbe9b9775ce5133dd9828ff93eee753b884d6"
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
