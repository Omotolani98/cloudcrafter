class Cloudcrafter < Formula
    desc "A CLI tool for managing multi-cloud infrastructure provisioning"
    homepage "https://github.com/Omotolani98/cloudcrafter"
    url "https://github.com/Omotolani98/cloudcrafter/releases/download/v1.0.7/cloudcrafter-darwin-amd64.tar.gz" # Update to your release URL
    sha256 "0019dfc4b32d63c1392aa264aed2253c1e0c2fb09216f8e2cc269bbfb8bb49b5" # Replace with the SHA256 hash of the tarball
    version "1.0.7"
  
    def install
      bin.install "cloudcrafter"
    end
  
    test do
      system "#{bin}/cloudcrafter", "--version"
    end
end