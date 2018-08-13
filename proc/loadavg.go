package proc

import (
	"errors"
	"strconv"
	"strings"
)

// Default loadavg file path
const loadavgPath = "/proc/loadavg"

// LoadAvg the loadavg info
type LoadAvg struct {
	Load1Min       float64 `json:"load1min"`
	Load5Min       float64 `json:"load5min"`
	Load15Min      float64 `json:"load15min"`
	ProcessRunning int     `json:"process_running"`
	ProcessTotal   int     `json:"process_total"`
	LastPID        int     `json:"last_pid"`
}

// GetLoadAvg get the loadavg info from path
func GetLoadAvg(path string) (*LoadAvg, error) {
	lines, err := readFile(path, loadavgPath)
	if err != nil {
		return nil, err
	}

	if len(lines) <= 0 {
		return nil, errors.New("empty file")
	}
	fields := strings.Fields(lines[0])
	if len(fields) < 5 {
		return nil, errors.New("file format error")
	}
	process := strings.Split(fields[3], "/")
	if len(process) != 2 {
		return nil, errors.New("file format error")
	}
	var loadavg = LoadAvg{}
	if loadavg.Load1Min, err = strconv.ParseFloat(fields[0], 64); err != nil {
		return nil, err
	}
	if loadavg.Load5Min, err = strconv.ParseFloat(fields[1], 64); err != nil {
		return nil, err
	}
	if loadavg.Load15Min, err = strconv.ParseFloat(fields[2], 64); err != nil {
		return nil, err
	}
	if loadavg.ProcessRunning, err = strconv.Atoi(process[0]); err != nil {
		return nil, err
	}
	if loadavg.ProcessTotal, err = strconv.Atoi(process[1]); err != nil {
		return nil, err
	}
	if loadavg.LastPID, err = strconv.Atoi(fields[4]); err != nil {
		return nil, err
	}

	return &loadavg, nil
}
