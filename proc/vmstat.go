package proc

import (
	"strconv"
	"strings"
)

// Default vmstat file path
const vmstatPath = "/proc/vmstat"

// GetVmStat read the vmstat file
func GetVmStat(path string) (map[string]uint64, error) {
	lines, err := readFile(path, vmstatPath)
	if err != nil {
		return nil, err
	}

	vmstatMap := make(map[string]uint64)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		vmstatMap[fields[0]], _ = strconv.ParseUint(fields[1], 10, 64)
	}

	return vmstatMap, nil
}
