package proc

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

const procPath = "/proc"

// BaseProcess simple base process info
type BaseProcess struct {
	ID      int    `json:"id"`
	Cmdline string `json:"cmdline"`
	State   string `json:"state"`
}

// Process detail process info
type Process struct {
	ID      int                          `json:"id"`
	Cmdline string                       `json:"cmdline"`
	Environ string                       `json:"environ"`
	Fdinfo  []ProcessFdInfo              `json:"fdinfo"`
	Io      map[string]uint64            `json:"io"`
	Limits  map[string]map[string]string `json:"limits"`
	Stat    *ProcessStat                 `json:"stat"`
	Statm   *ProcessStatm                `json:"statm"`
	Status  map[string]string            `json:"status"`
}

// GetProcesses get all process base info
func GetProcesses() (map[int]BaseProcess, error) {
	files, err := os.ReadDir(procPath)
	if err != nil {
		return nil, err
	}
	var ret = make(map[int]BaseProcess)
	for _, fi := range files {
		if fi.IsDir() {
			pid, err := strconv.Atoi(fi.Name())
			if err == nil {
				p := BaseProcess{}
				p.ID = pid
				p.Cmdline, _ = GetProcessCmdline(pid)
				stat, err := GetProcessStat(pid)
				if err == nil {
					p.State = stat.State
				}
				ret[pid] = p
			}
		}
	}
	return ret, nil
}

// GetProcess single process detail info
func GetProcess(pid int) (*Process, error) {
	p, err := os.Stat(filepath.Join(procPath, strconv.Itoa(pid)))
	if err != nil || !p.IsDir() {
		return nil, errors.New("process is not exist")
	}

	var process = Process{ID: pid}
	if process.Cmdline, err = GetProcessCmdline(pid); err != nil {
		process.Cmdline = ""
	}

	if process.Environ, err = GetProcessEnviron(pid); err != nil {
		process.Environ = ""
	}

	if process.Fdinfo, err = GetProcessFdInfo(pid); err != nil {
		process.Fdinfo = nil
	}

	if process.Io, err = GetProcessIO(pid); err != nil {
		process.Io = nil
	}

	if process.Limits, err = GetProcessLimits(pid); err != nil {
		process.Limits = nil
	}

	if process.Stat, err = GetProcessStat(pid); err != nil {
		process.Stat = nil
	}

	if process.Statm, err = GetProcessStatm(pid); err != nil {
		process.Statm = nil
	}

	if process.Status, err = GetProcessStatus(pid); err != nil {
		process.Status = nil
	}

	return &process, nil
}
