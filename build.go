package build

import (
	"encoding/json"
	"runtime"
)

// define replaceable variables
var (
	version     = ""
	revision    = ""
	branch      = ""
	buildNumber = ""
	url         = ""
	engine      = runtime.Version()
)

// define struct with info about build
type Build struct {
	Version     string `json:"version"`
	Revision    string `json:"revision"`
	Branch      string `json:"branch"`
	BuildNumber string `json:"build_number"`
	URL         string `json:"url"`
	Engine      string `json:"engine"`
}

func (b Build) String() string {
	data, _ := json.Marshal(b)
	return string(data)
}

// Retrieve current build info func
//goland:noinspection GoUnusedExportedFunction
func GetBuildInfo() Build {
	return Build{
		Version:     version,
		Revision:    revision,
		Branch:      branch,
		BuildNumber: buildNumber,
		URL:         url,
		Engine:      engine,
	}
}
