package v3

import (
	"encoding/json"

	"github.com/ks2211/go-suricata/client"
)

// Constants
const (
	ifaceListCommand string = "iface-list"
	ifaceStatCommand string = "iface-stat"
)

// IFaceStatRequest holds the interface name for the iface-stat command
type IFaceStatRequest struct {
	IFace string `json:"iface"`
}

// IFaceListCommand gets the list of interfaces available
func (s *SocketV3) IFaceListCommand() (client.IFaceListResponse, error) {
	// create and marshal the "iface-list" socket message with no args
	response, err := s.DoCommand(ifaceListCommand, nil)
	if err != nil {
		return client.IFaceListResponse{}, err
	}
	// unmarshal into the iface list struct
	var ifaceList client.IFaceListResponse
	if err := json.Unmarshal(response.Message, &ifaceList); err != nil {
		return client.IFaceListResponse{}, err
	}

	return ifaceList, nil
}

// IFaceStatCommand gets information about a specific interface
func (s *SocketV3) IFaceStatCommand(iface string) (client.IFaceStatResponse, error) {
	// create and marshal the "iface-stat" socket message with the ifaceStatRequest arg
	response, err := s.DoCommand(ifaceStatCommand, IFaceStatRequest{
		iface,
	})
	if err != nil {
		return client.IFaceStatResponse{}, err
	}
	// unmarshal into the iface stat struct
	var ifaceInfo client.IFaceStatResponse
	if err := json.Unmarshal(response.Message, &ifaceInfo); err != nil {
		return client.IFaceStatResponse{}, err
	}

	return ifaceInfo, nil
}
