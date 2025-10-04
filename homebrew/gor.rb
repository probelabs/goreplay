class Gor < Formula
  desc "Real-time HTTP traffic replay tool written in Go"
  homepage "https://goreplay.org"
  url "https://github.com/grofers/goreplay/archive/refs/tags/v2.0.0.tar.gz"
  sha256 "PLACEHOLDER_SHA256"
  license "LGPL-3.0"
  head "https://github.com/grofers/goreplay.git", branch: "master"

  depends_on "go" => :build

  def install
    # Sync vendor directory to ensure consistency
    system "go", "mod", "vendor"
    
    # Set build flags
    ldflags = "-X main.VERSION=#{version}"
    
    # Build the binary
    system "go", "build",
           "-mod=vendor",
           "-buildvcs=false",
           "--tags", "ngo",
           "-ldflags", ldflags,
           "-o", bin/"gor",
           "./cmd/gor/"
  end

  def caveats
    <<~EOS
      GoReplay (gor) requires raw socket access for packet capture.
      You may need to run it with sudo:
        sudo gor --input-raw :8080 --output-http http://staging.example.com
    EOS
  end

  test do
    # Test that the binary runs and shows help
    output = shell_output("#{bin}/gor 2>&1", 2)
    assert_match "Gor is a simple http traffic replication tool", output
    assert_match "Current Version:", output
  end
end
