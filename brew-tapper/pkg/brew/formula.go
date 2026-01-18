package brew

type Formula struct {
	Name, Version, DarwinSha256, DarwinArm64Sha256, LinuxSha256, LinuxArm64Sha256 string
}

func (f *Formula) NotSpecified() bool {
	return len(f.Name) == 0 || len(f.Version) == 0
}
