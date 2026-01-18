package brew

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/sirupsen/logrus"
)

var (
	extensions = []string{
		".zip",
		".tar",
		".tar.gz",
		".tgz",
		".tar.bz2",
		".tbz2",
		".tar.xz",
		".txz",
		".tar.lz4",
		".tlz4",
		".tar.sz",
		".tsz",
		".rar",
		".bz2",
		".gz",
		".lz4",
		".sz",
		".xz",
	}
	goarch64bit = []string{
		"amd64",
		"arm64",
		"arm64be",
		"ppc64",
		"ppc64le",
		"mips64",
		"mips64le",
		"s390x",
		"sparc64",
	}
	goos = []string{
		"android",
		"darwin",
		"dragonfly",
		"freebsd",
		"linux",
		"nacl",
		"netbsd",
		"openbsd",
		"plan9",
		"solaris",
		"windows",
		"zos",
	}
)

func (f *Formula) Guess(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("not found any binary archive in %q", path)
	}
	for _, file := range files {
		if file.IsDir() || !isSupportedArchive(file.Name()) {
			continue
		}
		guessOs, guessArch, guessName, guessVersion, err := guess(filepath.Base(file.Name()))
		if err != nil {
			logrus.Debugln(err)
			continue
		}
		if len(f.Name) == 0 {
			f.Name = guessName
		}
		if len(f.Version) == 0 {
			f.Version = guessVersion
		}
		if guessOs == "linux" {
			if guessArch == "arm64" {
				if len(f.LinuxArm64Sha256) == 0 {
					if f.LinuxArm64Sha256, err = hash(filepath.Join(path, file.Name())); err != nil {
						return err
					}
				}
			} else { // amd64 or unspecified
				if len(f.LinuxSha256) == 0 {
					if f.LinuxSha256, err = hash(filepath.Join(path, file.Name())); err != nil {
						return err
					}
				}
			}
		}
		if guessOs == "darwin" {
			if guessArch == "arm64" {
				if len(f.DarwinArm64Sha256) == 0 {
					if f.DarwinArm64Sha256, err = hash(filepath.Join(path, file.Name())); err != nil {
						return err
					}
				}
			} else { // amd64 or unspecified
				if len(f.DarwinSha256) == 0 {
					if f.DarwinSha256, err = hash(filepath.Join(path, file.Name())); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func isSupportedArchive(source string) bool {
	for _, suffix := range extensions {
		if strings.HasSuffix(source, suffix) {
			return true
		}
	}
	return false
}

func hash(file string) (checksum string, err error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	hasher := sha256.New()
	if _, err = io.Copy(hasher, f); err != nil {
		return
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func guess(file string) (os, arch, name, version string, err error) {
	var fileName = truncateExtension(file)
	chunks := strings.Split(fileName, "-")
	if len(chunks) == 1 {
		if chunks = strings.Split(file, "_"); len(chunks) == 1 {
			err = fmt.Errorf("%s does not contains any dash or underline, consider it not a binary archive", file)
			return
		}
	}
	for _, chunk := range chunks {
		if _, err := semver.NewVersion(chunk); err == nil {
			version = chunk
			continue
		}
		if containsIgnoreCase(goarch64bit, chunk) {
			arch = chunk
			continue
		}
		if containsIgnoreCase(goos, chunk) {
			os = chunk
			continue
		}
		name = chunk
	}
	return
}

func containsIgnoreCase(ss []string, s string) bool {
	s = strings.ToLower(s)
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

func truncateExtension(file string) string {
	var ext = filepath.Ext(file)
	return file[0 : len(file)-len(ext)]
}
