package proc

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ProcessStatm the statm file info
// http://man7.org/linux/man-pages/man5/proc.5.html
type ProcessStatm struct {
	Size     uint64 `json:"size"`
	Resident uint64 `json:"resident"`
	Share    uint64 `json:"share"`
	Text     uint64 `json:"text"`
	Lib      uint64 `json:"lib"`
	Data     uint64 `json:"data"`
	Dirty    uint64 `json:"dirty"`
}

// GetProcessStatm get statm info
func GetProcessStatm(pid int) (*ProcessStatm, error) {
	path := filepath.Join(procPath, strconv.Itoa(pid), "statm")

	b, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}
	fields := strings.Fields(string(b))

	var statm = ProcessStatm{}

	for i, f := range fields {

		switch i {
		case 0:
			if statm.Size, err = strconv.ParseUint(f, 10, 64); err != nil {
				return nil, err
			}
		case 1:
			if statm.Resident, err = strconv.ParseUint(f, 10, 64); err != nil {
				return nil, err
			}
		case 2:
			if statm.Share, err = strconv.ParseUint(f, 10, 64); err != nil {
				return nil, err
			}
		case 3:
			if statm.Text, err = strconv.ParseUint(f, 10, 64); err != nil {
				return nil, err
			}
		case 4:
			if statm.Lib, err = strconv.ParseUint(f, 10, 64); err != nil {
				return nil, err
			}
		case 5:
			if statm.Data, err = strconv.ParseUint(f, 10, 64); err != nil {
				return nil, err
			}
		case 6:
			if statm.Dirty, err = strconv.ParseUint(f, 10, 64); err != nil {
				return nil, err
			}
		}
	}
	return &statm, nil
}
