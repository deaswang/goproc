package proc

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// GetProcessCmdline get cmdline info for pid
func GetProcessCmdline(pid int) (string, error) {
	path := filepath.Join(procPath, strconv.Itoa(pid), "cmdline")

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}
