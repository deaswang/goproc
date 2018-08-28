package proc

import (
	"errors"
	"strconv"
	"strings"
)

// Default lock file path
const locksPath = "/proc/locks"

// Lock one line lock
type Lock struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	AccessOther string `json:"access_other"`
	Access      string `json:"access"`
	Pid         int    `json:"pid"`
	DeviceMajor int64  `json:"device_major"`
	DeviceMinor int64  `json:"device_minor"`
	DeviceInode int64  `json:"device_inode"`
	Start       int64  `json:"start"`
	End         int64  `json:"end"`
}

// Locks all locks
type Locks struct {
	Ls []Lock `json:"ls"`
}

// GetLocks get locks info from path
func GetLocks(path string) (*Locks, error) {
	lines, err := readFile(path, locksPath)
	if err != nil {
		return nil, err
	}

	var locks = Locks{}

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 8 {
			continue
		}
		lock := Lock{}
		if lock.ID, err = strconv.Atoi(strings.TrimSuffix(fields[0], ":")); err != nil {
			return nil, err
		}
		lock.Type = fields[1]
		lock.AccessOther = fields[2]
		lock.Access = fields[3]
		if lock.Pid, err = strconv.Atoi(fields[4]); err != nil {
			return nil, err
		}
		devices := strings.Split(fields[5], ":")
		if len(devices) < 3 {
			return nil, errors.New("file device format error")
		}
		if lock.DeviceMajor, err = strconv.ParseInt(devices[0], 16, 64); err != nil {
			return nil, err
		}
		if lock.DeviceMinor, err = strconv.ParseInt(devices[1], 16, 64); err != nil {
			return nil, err
		}
		if lock.DeviceInode, err = strconv.ParseInt(devices[2], 10, 64); err != nil {
			return nil, err
		}
		if lock.Start, err = strconv.ParseInt(fields[6], 10, 64); err != nil {
			return nil, err
		}
		if lock.End, err = strconv.ParseInt(fields[7], 10, 64); err != nil {
			if fields[7] == "EOF" {
				lock.End = -1
			} else {
				return nil, err
			}
		}
		locks.Ls = append(locks.Ls, lock)
	}

	return &locks, nil
}
