package proc

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Default uptime path
const uptimePath = "/proc/uptime"

// Uptime the uptime struct
type Uptime struct {
	Total string `json:"total"`
	Idle  string `json:"idle"`
}

// GetUptime read the uptime file
func GetUptime(path string) (*Uptime, error) {
	lines, err := readFile(path, uptimePath)
	if err != nil {
		return nil, err
	}
	if len(lines) < 1 {
		return nil, errors.New("Empty uptime file")
	}
	var uptime = Uptime{}
	fields := strings.Fields(lines[0])
	if len(fields) != 2 {
		return nil, errors.New("Wrong format uptime")
	}
	total, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return nil, errors.New("Wrong format uptime")
	}
	idle, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return nil, errors.New("Wrong format uptime")
	}
	uptime.Total = (time.Duration(total) * time.Second).String()
	uptime.Idle = (time.Duration(idle) * time.Second).String()
	return &uptime, nil
}
