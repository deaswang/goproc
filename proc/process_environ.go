package proc

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// GetProcessEnviron get process environ info
func GetProcessEnviron(pid int) (string, error) {
	path := filepath.Join(procPath, strconv.Itoa(pid), "environ")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}
