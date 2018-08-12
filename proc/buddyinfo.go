package proc

import (
	"strconv"
	"strings"
)

// Default buddyinfo file path
const buddyinfoPath = "/proc/buddyinfo"

// Buddy single buddy info
type Buddy struct {
	Node   int    `json:"node"`
	Zone   string `json:"zone"`
	Pages  []int  `json:"pages"`
}

// BuddyInfo buddy info
type BuddyInfo struct {
	Buddys  []Buddy  `json:"buddys"`
}

// GetBuddyInfo read buddyinfo file info from path
func GetBuddyInfo(path string) (*BuddyInfo, error) {
	lines, err := readFile(path, buddyinfoPath)
	if err != nil {
		return nil, err
	}

	var buddyinfo = BuddyInfo{}

	for _, line := range lines {
		buddy := Buddy{}
		fields := strings.Fields(line)
		if len(fields) <= 0 {
			continue
		}
		i := 0
		if fields[i] != "Node" {
			continue
		}
		i++
		buddy.Node, err = strconv.Atoi(strings.TrimRight(fields[i], ","))
		if err != nil {
			continue
		}
		i++
		if fields[i] != "zone" {
			continue
		}
		i++
		buddy.Zone = fields[i]
		for i++; i<len(fields); i++ {
			page, err := strconv.Atoi(fields[i])
			if err != nil {
				continue
			}
			buddy.Pages = append(buddy.Pages, page)
		}

		buddyinfo.Buddys = append(buddyinfo.Buddys, buddy)
	}
	return &buddyinfo, nil
}


