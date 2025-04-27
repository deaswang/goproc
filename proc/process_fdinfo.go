package proc

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ProcessFdInfo the process fdinfo struct
type ProcessFdInfo struct {
	ID string `json:"id"`
	// Name  string `json:"name"`
	Pos   uint64 `json:"pos"`
	Flags string `json:"flags"`
	MntID int    `json:"mnt_id"`
}

// GetProcessFdInfo get process fdinfo file list
func GetProcessFdInfo(pid int) ([]ProcessFdInfo, error) {
	// fdPath := filepath.Join(procPath, strconv.Itoa(pid), "fd")
	fdInfoPath := filepath.Join(procPath, strconv.Itoa(pid), "fdinfo")
	files, err := os.ReadDir(fdInfoPath)
	if err != nil {
		return nil, err
	}
	var fdinfos = make([]ProcessFdInfo, 0)
	for _, fi := range files {
		b, err := os.ReadFile(filepath.Join(fdInfoPath, fi.Name()))
		if err != nil {
			continue
		}
		lines := strings.Split(string(b), "\n")
		fdinfo := ProcessFdInfo{}
		fdinfo.ID = fi.Name()
		for _, line := range lines {
			fields := strings.Split(line, ":")
			if len(fields) < 2 {
				continue
			}
			if fields[0] == "pos" {
				fdinfo.Pos, _ = strconv.ParseUint(fields[1], 10, 64)
			} else if fields[0] == "flags" {
				fdinfo.Flags = strings.TrimSpace(fields[1])
			} else if fields[0] == "mnt_id" {
				fdinfo.MntID, _ = strconv.Atoi(fields[1])
			}
		}
		fdinfos = append(fdinfos, fdinfo)
	}
	return fdinfos, nil
}
