class Slctl < Formula
  desc "Slctl is a command line interface for running commands against SoftLeader Services"
  homepage "https://github.com/softleader/slctl"
  version "3.8.1"
  
  if OS.mac?
    if Hardware::CPU.arm?
      url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-arm64-#{version}.tgz"
      sha256 "0c9b99f2bf3d526ae7197f37675e6ed1fe059c29ee93f2ca51b4fb41503fb5fd"
    else
      url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-darwin-amd64-#{version}.tgz"
      sha256 "0c9b99f2bf3d526ae7197f37675e6ed1fe059c29ee93f2ca51b4fb41503fb5fd"
    end
  elsif OS.linux?
    if Hardware::CPU.arm?
      url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-arm64-#{version}.tgz"
      sha256 "5eb622c25a37c1a7a373f4036f51a7363663032e4436f2e33bb92cd7febb4628"
    else
      url "https://github.com/softleader/slctl/releases/download/#{version}/slctl-linux-amd64-#{version}.tgz"
      sha256 "5eb622c25a37c1a7a373f4036f51a7363663032e4436f2e33bb92cd7febb4628"
    end
  end

  def install
    bin.install "slctl"
  end

  def caveats; <<~EOS
    To begin working with slctl, run the 'slctl init' command.
  EOS
  end
end
