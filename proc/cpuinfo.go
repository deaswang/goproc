package proc

import (
	"strconv"
	"strings"
)

// Default cpuinfo file path
const cpuinfoPath = "/proc/cpuinfo"

// Processor ...
// single processor data struct
type Processor struct {
	ID         int      `json:"id"`
	VendorID   string   `json:"vendor_id"`
	Family     int      `json:"cpu_family"`
	Model      int      `json:"model"`
	ModelName  string   `json:"model_name"`
	MHz        float64  `json:"cpu_mhz"`
	CacheSize  int64    `json:"cache_size"`
	PhysicalID int      `json:"physical_id"`
	CoreID     int      `json:"core_id"`
	Cores      int      `json:"cpu_cores"`
	Fpu        bool     `json:"fpu"`
	FpuExcep   bool     `json:"fpu_exception"`
	Wp         bool     `json:"wp"`
	Flags      []string `json:"flags"`
}

// CPUInfo store all data for cpu info
type CPUInfo struct {
	Processors []Processor `json:"processors"`
}

// GetCPUInfo read cpu info from cpuinfo file path
func GetCPUInfo(path string) (*CPUInfo, error) {
	lines, err := readFile(path, cpuinfoPath)
	if err != nil {
		return nil, err
	}

	var cpuinfo = CPUInfo{}
	var processor = Processor{ID: -1}

	for _, line := range lines {
		var key string
		var value string

		splitStr := strings.Split(line, ":")
		if len(splitStr) > 1 {
			value = strings.TrimSpace(splitStr[1])
		}
		key = strings.TrimSpace(splitStr[0])

		switch key {
		case "processor":
			if processor.ID < 0 {
				processor.ID, _ = strconv.Atoi(value)
			} else {
				cpuinfo.Processors = append(cpuinfo.Processors, processor)
				id, _ := strconv.Atoi(value)
				processor = Processor{ID: id}
			}
		case "vendor_id":
			processor.VendorID = value
		case "cpu family":
			processor.Family, _ = strconv.Atoi(value)
		case "model":
			processor.Model, _ = strconv.Atoi(value)
		case "model name":
			processor.ModelName = value
		case "cpu MHz":
			processor.MHz, _ = strconv.ParseFloat(value, 64)
		case "cache size":
			processor.CacheSize = parseSizeByte(value)
		case "physical id":
			processor.PhysicalID, _ = strconv.Atoi(value)
		case "core id":
			processor.CoreID, _ = strconv.Atoi(value)
		case "cpu cores":
			processor.Cores, _ = strconv.Atoi(value)
		case "fpu":
			processor.Fpu = parseBool(value)
		case "fpu_exception":
			processor.FpuExcep = parseBool(value)
		case "wp":
			processor.Wp = parseBool(value)
		case "flags":
			processor.Flags = strings.Fields(value)
		}
	}
	cpuinfo.Processors = append(cpuinfo.Processors, processor)
	return &cpuinfo, nil
}

