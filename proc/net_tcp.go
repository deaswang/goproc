package proc

import (
	"strconv"
	"strings"
)

const netTCPPath = "/proc/net/tcp"

// NetTCP the net/tcp file info
type NetTCP struct {
	NetSocket
	RetransmitTimeout       uint64 `json:"retransmit_timeout"`
	PredictedTick           uint64 `json:"predicted_tick"`
	AckQuick                uint8  `json:"ack_quick"`
	AckPingpong             bool   `json:"ack_pingpong"`
	SendingCongestionWindow uint64 `json:"sending_congestion_window"`
	SlowStartSizeThreshold  int64  `json:"slow_start_size_threshold"`
}

// GetNetTCP get net tcp info
func GetNetTCP(path string) ([]NetTCP, error) {

	lines, err := readFile(path, netTCPPath)
	if err != nil {
		return nil, err
	}

	tcp := make([]NetTCP, len(lines) - 2)

	for i, line := range lines[1:] {

		f := strings.Fields(line)

		if len(f) < 17 {
			continue
		}

		s, err := parseNetSocket(f)

		if err != nil {
			return nil, err
		}

		var n int64
		e := NetTCP{
			NetSocket: *s,
		}

		if e.RetransmitTimeout, err = strconv.ParseUint(f[12], 10, 64); err != nil {
			return nil, err
		}

		if e.PredictedTick, err = strconv.ParseUint(f[13], 10, 64); err != nil {
			return nil, err
		}

		if n, err = strconv.ParseInt(f[14], 10, 8); err != nil {
			return nil, err
		}
		e.AckQuick = uint8(n >> 1)
		e.AckPingpong = ((n & 1) == 1)

		if e.SendingCongestionWindow, err = strconv.ParseUint(f[15], 10, 64); err != nil {
			return nil, err
		}

		if e.SlowStartSizeThreshold, err = strconv.ParseInt(f[16], 10, 32); err != nil {
			return nil, err
		}

		tcp[i] = e
	}

	return tcp, nil
}
