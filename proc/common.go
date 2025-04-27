package proc

import (
	"os"
	"strconv"
	"strings"
)

// readFile read lines from path, if path is not exist read defaultPath
func readFile(path, defaultPath string) (lines []string, err error) {
	f, err := os.Stat(path)
	if err != nil || f.IsDir() {
		path = defaultPath
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines = strings.Split(string(b), "\n")

	return lines, nil
}

// parseSizeByte parse size value with Byte
func parseSizeByte(value string) uint64 {
	space := strings.IndexAny(value, " \t")
	if space < 0 {
		space = len(value)
	}
	ret, err := strconv.ParseUint(value[:space], 10, 64)
	if err != nil {
		return 0
	}
	if strings.HasSuffix(value, "GB") || strings.HasSuffix(value, "gB") {
		ret *= 1024 * 1024 * 1024
	}
	if strings.HasSuffix(value, "MB") || strings.HasSuffix(value, "mB") {
		ret *= 1024 * 1024
	}
	if strings.HasSuffix(value, "KB") || strings.HasSuffix(value, "kB") {
		ret *= 1024
	}
	return ret
}

// parseBool parse bool value, add yes/y check
func parseBool(value string) bool {
	if s, err := strconv.ParseBool(value); err == nil {
		return s
	}
	switch value {
	case "yes", "YES", "y", "Y":
		return true
	case "no", "NO", "n", "N":
		return false
	}
	return false
}
