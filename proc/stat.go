package proc

import (
	"strconv"
	"strings"
	"time"
)

// Default stat path
const statPath = "/proc/stat"

// Stat the stat struct
type Stat struct {
	CPUStatAll   CPUStat   `json:"cpu_all"`
	CPUStats     []CPUStat `json:"cpus"`
	Intr         []uint64  `json:"intr"`
	Ctxt         uint64    `json:"ctxt"`
	BootTime     time.Time `json:"btime"`
	Processes    uint64    `json:"processes"`
	ProcsRunning uint64    `json:"procs_running"`
	ProcsBlocked uint64    `json:"procs_blocked"`
	Softirq      []uint64  `json:"softirq"`
}

// CPUStat the stat of CPU
type CPUStat struct {
	ID        string `json:"id"`
	User      uint64 `json:"user"`
	Nice      uint64 `json:"nice"`
	System    uint64 `json:"system"`
	Idle      uint64 `json:"idle"`
	IOWait    uint64 `json:"iowait"`
	IRQ       uint64 `json:"irq"`
	SoftIRQ   uint64 `json:"softirq"`
	Steal     uint64 `json:"steal"`
	Guest     uint64 `json:"guest"`
	GuestNice uint64 `json:"guest_nice"`
}

// GetStat read the stat file
func GetStat(path string) (*Stat, error) {
	lines, err := readFile(path, statPath)
	if err != nil {
		return nil, err
	}

	var stat = Stat{}
	for i, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		if fields[0][:3] == "cpu" {
			cpuStat := CPUStat{}
			cpuStat.ID = fields[0]
			cpuStat.User, _ = strconv.ParseUint(fields[1], 10, 64)
			cpuStat.Nice, _ = strconv.ParseUint(fields[2], 10, 64)
			cpuStat.System, _ = strconv.ParseUint(fields[3], 10, 64)
			cpuStat.Idle, _ = strconv.ParseUint(fields[4], 10, 64)
			cpuStat.IOWait, _ = strconv.ParseUint(fields[5], 10, 64)
			cpuStat.IRQ, _ = strconv.ParseUint(fields[6], 10, 64)
			cpuStat.SoftIRQ, _ = strconv.ParseUint(fields[7], 10, 64)
			cpuStat.Steal, _ = strconv.ParseUint(fields[8], 10, 64)
			cpuStat.Guest, _ = strconv.ParseUint(fields[9], 10, 64)
			cpuStat.GuestNice, _ = strconv.ParseUint(fields[10], 10, 64)
			if i == 0 {
				stat.CPUStatAll = cpuStat
			} else {
				stat.CPUStats = append(stat.CPUStats, cpuStat)
			}
		} else if fields[0] == "intr" {
			for j := 1; j < len(fields); j++ {
				intr, err := strconv.ParseUint(fields[j], 10, 64)
				if err != nil {
					continue
				}
				stat.Intr = append(stat.Intr, intr)
			}
		} else if fields[0] == "ctxt" {
			stat.Ctxt, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "btime" {
			seconds, _ := strconv.ParseInt(fields[1], 10, 64)
			stat.BootTime = time.Unix(seconds, 0)
		} else if fields[0] == "processes" {
			stat.Processes, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "procs_running" {
			stat.ProcsRunning, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "procs_blocked" {
			stat.ProcsBlocked, _ = strconv.ParseUint(fields[1], 10, 64)
		} else if fields[0] == "softirq" {
			for j := 1; j < len(fields); j++ {
				softirq, err := strconv.ParseUint(fields[j], 10, 64)
				if err != nil {
					continue
				}
				stat.Softirq = append(stat.Softirq, softirq)
			}
		}
	}
	return &stat, nil
}
