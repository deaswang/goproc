package proc

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GetProcessIO get io info for pid
func GetProcessIO(pid int) (map[string]uint64, error) {
	path := filepath.Join(procPath, strconv.Itoa(pid), "io")
	b, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")

	var pio = make(map[string]uint64)

	for _, line := range lines {

		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			continue
		}
		k := fields[0]
		v, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}
		pio[k] = v
	}
	return pio, nil
}
