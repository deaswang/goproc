package proc

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// GetProcessLimits get limits info for pid
func GetProcessLimits(pid int) (map[string]map[string]string, error) {
	path := filepath.Join(procPath, strconv.Itoa(pid), "limits")
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	if len(lines) < 2 {
		return nil, errors.New("File empty")
	}
	indexLimit := strings.Index(lines[0], "Limit")
	indexSoft := strings.Index(lines[0], "Soft Limit")
	indexHard := strings.Index(lines[0], "Hard Limit")
	indexUnits := strings.Index(lines[0], "Units")
	lines = lines[1:]
	var pls = make(map[string]map[string]string)

	for _, line := range lines {
		if len(line) < indexUnits {
			continue
		}
		pl := make(map[string]string)
		nameLimit := strings.TrimSpace(line[indexLimit:indexSoft])
		pl["Soft Limit"] = strings.TrimSpace(line[indexSoft:indexHard])
		pl["Hard Limit"] = strings.TrimSpace(line[indexHard:indexUnits])
		pl["Units"] = strings.TrimSpace(line[indexUnits:])
		pls[nameLimit] = pl
	}
	return pls, nil
}
