package proc

import (
	"errors"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var (
	ipv4RegExp = regexp.MustCompile("^[0-9a-fA-F]{8}:[0-9a-fA-F]{4}$")
	ipv6RegExp = regexp.MustCompile("^[0-9a-fA-F]{32}:[0-9a-fA-F]{4}$")
)

// NetSocket store netsock info
type NetSocket struct {
	LocalAddress         string `json:"local_address"`
	RemoteAddress        string `json:"remote_address"`
	Status               uint8  `json:"st"`
	TxQueue              uint64 `json:"tx_queue"`
	RxQueue              uint64 `json:"rx_queue"`
	UID                  uint32 `json:"uid"`
	Inode                uint64 `json:"inode"`
	SocketReferenceCount uint64 `json:"ref"`
}

// parseNetSocket parse netsocket
func parseNetSocket(f []string) (*NetSocket, error) {

	if len(f) < 11 {
		return nil, errors.New("can't parse net socket line: " + strings.Join(f, " "))
	}

	if !strings.Contains(f[4], ":") {
		return nil, errors.New("can't parse tx/rx queues: " + f[4])
	}

	q := strings.Split(f[4], ":")

	socket := &NetSocket{}

	var s uint64
	var u uint64
	var err error

	if socket.LocalAddress, err = NetIPDecoder(f[1]); err != nil {
		return nil, err
	}

	if socket.RemoteAddress, err = NetIPDecoder(f[2]); err != nil {
		return nil, err
	}

	if s, err = strconv.ParseUint(f[3], 16, 8); err != nil {
		return nil, err
	}

	if socket.TxQueue, err = strconv.ParseUint(q[0], 16, 64); err != nil {
		return nil, err
	}

	if socket.RxQueue, err = strconv.ParseUint(q[1], 16, 64); err != nil {
		return nil, err
	}

	if u, err = strconv.ParseUint(f[7], 10, 32); err != nil {
		return nil, err
	}

	if socket.Inode, err = strconv.ParseUint(f[9], 10, 64); err != nil {
		return nil, err
	}

	if socket.SocketReferenceCount, err = strconv.ParseUint(f[10], 10, 64); err != nil {
		return nil, err
	}

	socket.Status = uint8(s)
	socket.UID = uint32(u)

	return socket, nil
}

// NetIPDecoder decode an IP address with port from a given hex string
func NetIPDecoder(s string) (string, error) {
	var h []byte
	if ipv4RegExp.MatchString(s) {
		h = make([]byte, 4)
	} else if ipv6RegExp.MatchString(s) {
		h = make([]byte, 16)
	} else {
		return "", errors.New("wrong ipv4 address: " + s)
	}

	fields := strings.Split(s, ":")
	hf := fields[0]
	pf := fields[1]

	for i := 0; i < len(h); i++ {
		n, _ := strconv.ParseUint(hf[2*i:2*i+2], 16, 8)
		h[i] = byte(n)
	}

	host := net.IP(h).String()

	n, _ := strconv.ParseUint(pf, 16, 64)
	port := strconv.FormatUint(n, 10)
	ret := host + ":" + port

	return ret, nil
}
