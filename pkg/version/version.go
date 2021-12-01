package version

import (
	"fmt"
	"runtime"
)

// 包含的版本信息
type Info struct {
	GitTag       string `json:"gitTag"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info Info) String() string {
	return info.GitTag
}

// 获取版本信息   gitTag、gitCommit、gitTreeState 等变量的值是通过 -ldflags -X importpath.name=value 在编译时传到程序中的，需要在makefile中添加
func Get() Info {
	return Info{
		GitTag:       gitTag,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
