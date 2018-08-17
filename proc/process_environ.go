package proc

import (
	"io/ioutil"
	"strconv"
	"strings"
	"path/filepath"
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
