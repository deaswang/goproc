package proc

import (
	"strconv"
	"strings"
)

// Default mounts file path
const mountsPath = "/proc/mounts"

// Mount define one line mount
type Mount struct {
	Device     string   `json:"device"`
	MountPoint string   `json:"mount_point"`
	FSType     string   `json:"fstype"`
	Options    []string `json:"options"`
	Dump       int      `json:"dump"`
	Pass       int      `json:"pass"`
}

// Mounts define mounts file
type Mounts struct {
	Mounts  []Mount   `json:"mounts"`
}

// GetMounts read the mounts file
func GetMounts(path string) (*Mounts, error) {
	lines, err := readFile(path, mountsPath)
	if err != nil {
		return nil, err
	}

	var mounts = Mounts{}

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}
		mount := Mount{}
		mount.Device = fields[0]
		mount.MountPoint = fields[1]
		mount.FSType = fields[2]
		mount.Options = strings.Split(fields[3], ",")
		mount.Dump, _ = strconv.Atoi(fields[4])
		mount.Pass, _ = strconv.Atoi(fields[5])
		
		mounts.Mounts = append(mounts.Mounts, mount)
	}

	return &mounts, nil
}
