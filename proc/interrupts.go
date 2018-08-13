package proc

import (
	"errors"
	"strings"
)

// Default interrupts file path
const interruptsPath = "/proc/interrupts"

// Interrupt one line interrupt
type Interrupt struct {
	Name   string            `json:"name"`
	Counts map[string]uint64 `json:"counts"`
	Description string       `json:"description"`
}

// Interrupts all interrupts
type Interrupts struct {
	Entry  []Interrupt  `json:"entry"`
}

// GetInterrupts get interrupts info from path
func GetInterrupts(path string) (*Interrupts, error) {
	lines, err := readFile(path, interruptsPath)
	if err != nil {
		return nil, err
	}

	var interrupts = Interrupts{}
	var nameCPU []string

	for i, line := range lines {
		if i == 0 {
			nameCPU = strings.Fields(line)
			if len(nameCPU) <= 0 {
				return nil, errors.New("can`t get cpu")
			}
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		interrupt := Interrupt{}
		interrupt.Counts = make(map[string]uint64, len(nameCPU))
		interrupt.Name = strings.TrimSuffix(fields[0], ":")
		j := 1
		for ; j < len(nameCPU)+1 && j < len(fields); j++ {
			interrupt.Counts[nameCPU[j-1]] = parseUint(fields[j])
		}
		interrupt.Description = strings.Join(fields[j:], " ")
		interrupts.Entry = append(interrupts.Entry, interrupt)
	}

	return &interrupts, nil
}