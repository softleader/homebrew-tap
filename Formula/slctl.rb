class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "4.0.0"
  
  if OS.mac?
    if Hardware::CPU.arm?
      url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-arm64-#{version}.tgz"
      sha256 ""
    else
      url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-amd64-#{version}.tgz"
      sha256 ""
    end
  elsif OS.linux?
    url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-amd64-#{version}.tgz"
    sha256 ""
  end

  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end
