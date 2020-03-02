package client

import (
	"encoding/json"
)

// AddHostBitRequest holds the ip, bit name and expiration for adding a host bit
// Not used in v3
type AddHostBitRequest struct {
	IPAddr  string `json:"ipaddress"`
	BitName string `json:"hostbit"`
	Expire  int    `json:"expire"`
}

// RemoveHostBit holds the ip, bit name for removing a host bit
// Not used in v3
type RemoveHostBit struct {
	IPAddr  string `json:"ip"`
	BitName string `json:"hostbit"`
}

// ListHostBitRequest holds the host for list hostbit command
// Not used in v3
type ListHostBitRequest struct {
	IPAddr string `json:"ipaddress"`
}

// AddHostBitCommand adds host bit
// Not implemented in v3
func (s *socket) AddHostBitCommand(ipAddr, bitName string, expire int) (string, error) {
	// create and marshal the "add-hostbit" socket message with
	response, err := s.DoCommand(addHostBitCommand, AddHostBitRequest{
		ipAddr,
		bitName,
		expire,
	})
	if err != nil {
		return "", err
	}
	// unmarshal into the iface stat struct
	var addHostbit AddOrRemoveHostBitResponse
	if err := json.Unmarshal(response.Message, &addHostbit); err != nil {
		return "", err
	}
	return addHostbit.String(), nil
}

// RemoveHostBitCommand does "remove-hostbit"
// Not implemented in v3
func (s *socket) RemoveHostBitCommand(ipAddr, bitName string) (string, error) {
	// create and marshal the "remove-hostbit" socket message with
	response, err := s.DoCommand(removeHostBitCommand, RemoveHostBit{
		ipAddr,
		bitName,
	})
	if err != nil {
		return "", err
	}
	// unmarshal into the iface stat struct
	var removeHostBit AddOrRemoveHostBitResponse
	if err := json.Unmarshal(response.Message, &removeHostBit); err != nil {
		return "", err
	}
	return removeHostBit.String(), nil
}

// ListHostBitCommand does "list-hostbit"
// Not implemented in v3
func (s *socket) ListHostBitCommand(ipAddr string) (*ListHostBitsResponse, error) {
	// create and marshal the "memcap-show" socket message with
	response, err := s.DoCommand(listHostBitCommand, ListHostBitRequest{
		ipAddr,
	})
	if err != nil {
		return nil, err
	}
	// unmarshal into the iface stat struct
	var hostbits *ListHostBitsResponse
	if err := json.Unmarshal(response.Message, &hostbits); err != nil {
		return nil, err
	}
	return hostbits, nil
}
