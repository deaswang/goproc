package proc

import (
	"strconv"
	"strings"
)

// Default cgroups file path
const cgroupsPath = "/proc/cgroups"

// CGroup single control group
type CGroup struct {
	SubsysName string `json:"subsys_name"`
	Hierarchy  int    `json:"hierarchy"`
	NumCgroups int    `json:"num_cgroups"`
	Enabled    bool   `json:"enabled"`
}

// CGroups control groups
type CGroups struct {
	CGroups []CGroup `json:"cgroups"`
}

// GetCGroups read cgroups file info from path
func GetCGroups(path string) (*CGroups, error) {
	lines, err := readFile(path, cgroupsPath)
	if err != nil {
		return nil, err
	}
	var cgroups = CGroups{}
	for _, line := range lines {
		if strings.Index(line, "#") == 0 {
			continue
		}
		cgroup := CGroup{}
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		i := 0
		cgroup.SubsysName = fields[i]
		i++
		cgroup.Hierarchy, err = strconv.Atoi(fields[i])
		if err != nil {
			continue
		}
		i++
		cgroup.NumCgroups, err = strconv.Atoi(fields[i])
		if err != nil {
			continue
		}
		i++
		cgroup.Enabled, err = strconv.ParseBool(fields[i])
		if err != nil {
			continue
		}
		cgroups.CGroups = append(cgroups.CGroups, cgroup)
	}

	return &cgroups, nil
}
