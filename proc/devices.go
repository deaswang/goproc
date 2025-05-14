package proc

import (
	"strconv"
	"strings"
)

// Default devices file path
const devicesPath = "/proc/devices"

// CharacterDevice single character device
type CharacterDevice struct {
	Major_number int    `json:"major_number"`
	Name         string `json:"name"`
}

// BlockDevice single block device
type BlockDevice struct {
	Major_number int    `json:"major_number"`
	Name         string `json:"name"`
}

// Devices all devices
type Devices struct {
	CharacterDevices []CharacterDevice `json:"character_devices"`
	BlockDevices     []BlockDevice     `json:"block_devices"`
}

// GetDevices read devices file
func GetDevices(path string) (*Devices, error) {
	lines, err := readFile(path, devicesPath)
	if err != nil {
		return nil, err
	}
	var devices = Devices{}
	// 0: unknow devices 1: character devices 2: block devices
	device_type := 0

	for _, line := range lines {
		if strings.Compare(line, "Character devices:") == 0 {
			device_type = 1
			continue
		}
		if strings.Compare(line, "Block devices:") == 0 {
			device_type = 2
			continue
		}
		switch device_type {
		case 1:
			fields := strings.Fields(line)
			if len(fields) < 2 {
				continue
			}
			var character_device = CharacterDevice{}
			character_device.Major_number, err = strconv.Atoi(fields[0])
			if err != nil {
				continue
			}
			character_device.Name = fields[1]
			devices.CharacterDevices = append(devices.CharacterDevices, character_device)
		case 2:
			fields := strings.Fields(line)
			if len(fields) < 2 {
				continue
			}
			var block_device = BlockDevice{}
			block_device.Major_number, err = strconv.Atoi(fields[0])
			if err != nil {
				continue
			}
			block_device.Name = fields[1]
			devices.BlockDevices = append(devices.BlockDevices, block_device)
		}
	}

	return &devices, nil
}
