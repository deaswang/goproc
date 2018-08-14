package proc

import (
	"strconv"
	"strings"
)

// Default misc file path
const miscPath = "/proc/misc"

// GetMisc read the misc file
func GetMisc(path string) (map[int]string, error) {
	lines, err := readFile(path, miscPath)
	if err != nil {
		return nil, err
	}

	miscMap := make(map[int]string)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		deviceNum, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}
		miscMap[deviceNum] = fields[1]
	}

	return miscMap, nil
}
