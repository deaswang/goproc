package proc

import (
	"errors"
	"strings"
)

// Default version path
const versionPath = "/proc/version"

// Version the version struct
type Version struct {
	LinuxVersion string `json:"linux_version"`
	BuildUser    string `json:"build_user"`
	GccVersion   string `json:"gcc_version"`
	BuildTime    string `json:"build_time"`
}

// GetVersion read the version file
func GetVersion(path string) (*Version, error) {
	lines, err := readFile(path, versionPath)
	if err != nil {
		return nil, err
	}
	if len(lines) < 1 {
		return nil, errors.New("Empty version file")
	}
	var version = Version{}
	linux := strings.TrimSpace(lines[0][strings.Index(lines[0], "Linux version")+13 : strings.Index(lines[0], "(gcc (GCC) ")])
	gcc := strings.TrimSpace(lines[0][strings.Index(lines[0], "(gcc (GCC) ")+10 : strings.Index(lines[0], "#1")])
	time := strings.TrimSpace(lines[0][strings.Index(lines[0], "#1")+2:])
	version.LinuxVersion = linux[:strings.Index(linux, " ")]
	version.BuildUser = linux[strings.Index(linux, "(")+1 : strings.Index(linux, ")")]
	version.GccVersion = gcc[:len(gcc)-1]
	version.BuildTime = time

	return &version, nil
}
