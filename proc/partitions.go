package proc

import (
	"strconv"
	"strings"
)

// Default partitions file path
const partitionsPath = "/proc/partitions"

// Partition define one line partition
type Partition struct {
	Major  int    `json:"major"`
	Minor  int    `json:"minor"`
	Blocks uint64 `json:"blocks"`
	Name   string `json:"name"`
}

// Partitions define partitions file
type Partitions struct {
	Partitions []Partition `json:"partitions"`
}

// GetPartitions read the partitions file
func GetPartitions(path string) (*Partitions, error) {
	lines, err := readFile(path, partitionsPath)
	if err != nil {
		return nil, err
	}

	var partitions = Partitions{}

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		partition := Partition{}
		if partition.Major, err = strconv.Atoi(fields[0]); err != nil {
			continue
		}
		if partition.Minor, err = strconv.Atoi(fields[1]); err != nil {
			continue
		}
		if partition.Blocks, err = strconv.ParseUint(fields[2], 10, 64); err != nil {
			continue
		}
		partition.Name = fields[3]
		partitions.Partitions = append(partitions.Partitions, partition)
	}

	return &partitions, nil
}
