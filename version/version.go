package version

import (
	"os"
	"runtime"
	"strconv"
	"text/template"
	"time"
)

var (
	Version   string = "0.0.0"
	GitCommit string = ""
	GitState  string = ""
	BuildDate string = ""
)

var versionTemplate = `Version:     {{.Version}}
Go version:  {{.GoVersion}}
Git commit:  {{.GitCommit}}{{if eq .GitState "dirty"}}
Git State:   {{.GitState}}{{end}}
Built:       {{.BuildDate}}
OS/Arch:     {{.Os}}/{{.Arch}}`

type VersionInfo struct {
	Version   string
	GoVersion string
	GitCommit string
	GitState  string
	BuildDate string
	Os        string
	Arch      string
}

func New() *VersionInfo {
	i, err := strconv.ParseInt(BuildDate, 10, 64)
	if err != nil {
		panic(err)
	}

	tu := time.Unix(i, 0)

	return &VersionInfo{
		Version: Version,
		GoVersion: runtime.Version(),
		GitCommit: GitCommit,
		GitState: GitState,
		BuildDate: tu.String(),
		Os: runtime.GOOS,
		Arch: runtime.GOARCH,
	}
}

func (i *VersionInfo) ShowVersion() {
	tmpl, err := template.New("version").Parse(versionTemplate)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, i); err != nil {
		panic(err)
	}
}
