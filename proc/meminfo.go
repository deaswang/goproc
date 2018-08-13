package proc

import (
	// "reflect"
	"strings"
)

// Default meminfo file path
const meminfoPath = "/proc/meminfo"

// GetMemInfo read the meminfo file
func GetMemInfo(path string) (map[string]uint64, error) {
	lines, err := readFile(path, meminfoPath)
	if err != nil {
		return nil, err
	}

	meminfoMap := make(map[string]uint64)

	for _, line := range lines {
		fields := strings.SplitN(line, ":", 2)
		if len(fields) < 2 {
			continue
		}
		meminfoMap[fields[0]] = parseSizeByte(strings.TrimSpace(fields[1]))
	}

	return meminfoMap, nil
}
