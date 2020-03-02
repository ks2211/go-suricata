package client

import (
	"encoding/json"
)

// IFaceStatRequest holds the interface name for the iface-stat command
type IFaceStatRequest struct {
	IFace string `json:"iface"`
}

// IFaceListCommand gets the list of interfaces available
func (s *socket) IFaceListCommand() (*IFaceListResponse, error) {
	// create and marshal the "iface-list" socket message with no args
	response, err := s.DoCommand(ifaceListCommand, nil)
	if err != nil {
		return nil, err
	}
	// unmarshal into the iface list struct
	var ifaceList *IFaceListResponse
	if err := json.Unmarshal(response.Message, &ifaceList); err != nil {
		return nil, err
	}

	return ifaceList, nil
}

// IFaceStatCommand gets information about a specific interface
func (s *socket) IFaceStatCommand(iface string) (*IFaceStatResponse, error) {
	// create and marshal the "iface-stat" socket message with the ifaceStatRequest arg
	response, err := s.DoCommand(ifaceStatCommand, IFaceStatRequest{
		iface,
	})
	if err != nil {
		return nil, err
	}
	// unmarshal into the iface stat struct
	var ifaceInfo *IFaceStatResponse
	if err := json.Unmarshal(response.Message, &ifaceInfo); err != nil {
		return nil, err
	}

	return ifaceInfo, nil
}
