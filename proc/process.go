package proc

import (
	"io/ioutil"
	"strconv"
)

const procPath = "/proc"

// BaseProcess
type BaseProcess struct {
	ID   int   `json:"id"`
}

// Process
type Process struct {

}

// GetProcesses get all process base info
func GetProcesses() (map[int]BaseProcess, error) {
	files, err := ioutil.ReadDir(procPath)
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
				ret[pid] = p
			}
		}
	}
	return ret, nil
}

// GetProcess single process detail info
func GetProcess(pid int) (*Process, error) {
	return nil, nil
}

