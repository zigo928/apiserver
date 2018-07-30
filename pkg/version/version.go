package version

type Info struct {
	GitTag string `json:"gitTag"`
	GitCommit string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	Compiler string `json:"compiler"`
	Platform string `json:"platform"`
}

func (info Info) String() string {
	return info.GitTag
}

func Get() Info {
	return Info{
		GitTag: gitTag,
		GitCommit: gitCommit,
		GitTreeState: gitTreeState,
		BuildDate: buildDate,
		GoVersion: goVersion,
		Compiler: compiler,
	}
}

