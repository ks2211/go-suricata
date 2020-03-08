package client

const (
	ifaceList string = "iface-list"
	ifaceStat string = "iface-stat"
)

// IFaceStatRequest holds the interface name for the iface-stat command
type IFaceStatRequest struct {
	IFace string `json:"iface"`
}

// IFaceListResponse is resopnse from "iface-list"
type IFaceListResponse struct {
	Count  int      `json:"count"`
	Ifaces []string `json:"ifaces"`
}

// IFaceStatResponse is response from "iface-stat <iface>"
type IFaceStatResponse struct {
	Drop             int `json:"drop"`
	InvalidChecksums int `json:"invalid-checksums"`
	Pkts             int `json:"pkts"`
}

// IFaceListCommand gets the list of interfaces available
func (s *Socket) IFaceListCommand() (*IFaceListResponse, error) {
	ifaceListResp := &IFaceListResponse{}
	err := s.DoCommand(ifaceList, nil, ifaceListResp)
	return ifaceListResp, err
}

// IFaceStatCommand gets information about a specific interface
func (s *Socket) IFaceStatCommand(ifaceStatRequest IFaceStatRequest) (*IFaceStatResponse, error) {
	ifaceInfo := &IFaceStatResponse{}
	err := s.DoCommand(ifaceStat, ifaceStatRequest, ifaceInfo)
	return ifaceInfo, err
}
