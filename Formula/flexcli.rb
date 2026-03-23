class Flexcli < Formula
  desc "Management CLI for FlexCoach AI fitness platform"
  homepage "https://github.com/f1dot4/homebrew-flexcli"
  url "https://github.com/f1dot4/homebrew-flexcli/archive/refs/tags/v0.1.11.tar.gz"
  sha256 "ef7f7cc64072265c11b2e12f81729b4932bf4ac531b015c3517f11f4258c9ad5"
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
