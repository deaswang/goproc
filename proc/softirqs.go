package proc

import (
	"errors"
	"strconv"
	"strings"
)

// Default softirqs file path
const softirqsPath = "/proc/softirqs"

// GetSoftirqs get softirqs info from path
func GetSoftirqs(path string) (map[string]map[string]uint64, error) {
	lines, err := readFile(path, softirqsPath)
	if err != nil {
		return nil, err
	}

	softirqs := make(map[string]map[string]uint64)
	var nameCPU []string
	var numCPU int

	for i, line := range lines {
		if i == 0 {
			nameCPU = strings.Fields(line)
			numCPU = len(nameCPU)
			if numCPU <= 0 {
				return nil, errors.New("can`t get cpu")
			}
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		name := strings.TrimSuffix(fields[0], ":")
		softirqs[name] = make(map[string]uint64, numCPU)
		j := 1
		for ; j < numCPU+1 && j < len(fields); j++ {
			softirqs[name][nameCPU[j-1]], err = strconv.ParseUint(fields[j], 10, 64)
			if err != nil {
				return nil, err
			}
		}
	}

	return softirqs, nil
}
