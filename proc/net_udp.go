package proc

import (
	"strconv"
	"strings"
)

// Default net udp path
const netUDPPath = "/proc/net/udp"

// NetUDP the net udp info
type NetUDP struct {
	NetSocket
	Drops uint64 `json:"drops"`
}

// GetNetUDP get net udp info
func GetNetUDP(path string) ([]NetUDP, error) {

	lines, err := readFile(path, netUDPPath)
	if err != nil {
		return nil, err
	}

	udp := make([]NetUDP, len(lines) - 2)

	for i, line := range lines[1:] {

		f := strings.Fields(line)

		if len(f) < 13 {
			continue
		}

		s, err := parseNetSocket(f)

		if err != nil {
			return nil, err
		}

		e := NetUDP{
			NetSocket: *s,
		}

		if e.Drops, err = strconv.ParseUint(f[12], 10, 64); err != nil {
			return nil, err
		}
		udp[i] = e
	}

	return udp, nil
}
