package proc

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// GetProcessStatus get status info for pid
func GetProcessStatus(pid int) (map[string]string, error) {
	path := filepath.Join(procPath, strconv.Itoa(pid), "status")
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	var status = make(map[string]string)
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			continue
		}
		status[fields[0]] = strings.TrimSpace(fields[1])
	}
	return status, nil
}
