package proc

import (
	"strconv"
	"strings"
)

// Default diskstats file path
const diskstatsPath = "/proc/diskstats"

// DiskStat single disk stat
// https://www.kernel.org/doc/Documentation/iostats.txt
type DiskStat struct {
	Major        int    `json:"major"`          // major number
	Minor        int    `json:"minor"`          // minor number
	Name         string `json:"name"`           // device name
	ReadIOs      uint64 `json:"read_ios"`       // reads completed successfully
	ReadMerged   uint64 `json:"read_merged"`    // reads merged
	ReadSectors  uint64 `json:"read_sectors"`   // sectors read
	ReadTime     uint64 `json:"read_time"`      // milliseconds spent reading
	WriteIOs     uint64 `json:"write_ios"`      // writes completed
	WriteMerged  uint64 `json:"write_merged"`   // writes merged
	WriteSectors uint64 `json:"write_sectors"`  // sectors written
	WriteTime    uint64 `json:"write_time"`     // milliseconds spent writing
	IOProgress   uint64 `json:"io_progress"`    // I/Os currently in progress
	IOTime       uint64 `json:"io_time"`        // milliseconds spent doing I/Os
	IOTimeWeight uint64 `json:"id_time_weight"` // weighted of milliseconds spent doing I/Os
}

// DiskStats all disk stats
type DiskStats struct {
	Disks []DiskStat `json:"disks"`
}

// GetDiskStats read disk stats info from path file
func GetDiskStats(path string) (*DiskStats, error) {
	lines, err := readFile(path, diskstatsPath)
	if err != nil {
		return nil, err
	}

	var diskStats = DiskStats{}

	for _, line := range lines {
		disk := DiskStat{}
		fields := strings.Fields(line)
		if len(fields) <= 0 {
			continue
		}
		disk.Major, _ = strconv.Atoi(fields[0])
		disk.Minor, _ = strconv.Atoi(fields[1])
		disk.Name = fields[2]
		disk.ReadIOs, _ = strconv.ParseUint(fields[3], 10, 64)
		disk.ReadMerged, _ = strconv.ParseUint(fields[4], 10, 64)
		disk.ReadSectors, _ = strconv.ParseUint(fields[5], 10, 64)
		disk.ReadTime, _ = strconv.ParseUint(fields[6], 10, 64)
		disk.WriteIOs, _ = strconv.ParseUint(fields[7], 10, 64)
		disk.WriteMerged, _ = strconv.ParseUint(fields[8], 10, 64)
		disk.WriteSectors, _ = strconv.ParseUint(fields[9], 10, 64)
		disk.WriteTime, _ = strconv.ParseUint(fields[10], 10, 64)
		disk.IOProgress, _ = strconv.ParseUint(fields[11], 10, 64)
		disk.IOTime, _ = strconv.ParseUint(fields[12], 10, 64)
		disk.IOTimeWeight, _ = strconv.ParseUint(fields[13], 10, 64)

		diskStats.Disks = append(diskStats.Disks, disk)
	}
	return &diskStats, nil
}
