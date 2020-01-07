package v4

import (
	"encoding/json"
)

// Constants
const (
	addHostBitCommand    string = "add-hostbit"
	removeHostBitCommand string = "remove-hostbit"
	listHostBitCommand   string = "list-hostbit"
)

// AddHostBitRequest holds the ip, bit name and expiration for adding a host bit
type AddHostBitRequest struct {
	IPAddr  string `json:"ipaddress"`
	BitName string `json:"hostbit"`
	Expire  int    `json:"expire"`
}

// RemoveHostBit holds the ip, bit name for removing a host bit
type RemoveHostBit struct {
	IPAddr  string `json:"ip"`
	BitName string `json:"hostbit"`
}

// ListHostBitRequest holds the host for list hostbit command
type ListHostBitRequest struct {
	IPAddr string `json:"ipaddress"`
}

// AddHostBitCommand adds host bit
func (s *SocketV4) AddHostBitCommand(ipAddr, bitName string, expire int) (string, error) {
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
func (s *SocketV4) RemoveHostBitCommand(ipAddr, bitName string) (string, error) {
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
func (s *SocketV4) ListHostBitCommand(ipAddr string) (ListHostBitsResponse, error) {
	// create and marshal the "memcap-show" socket message with
	response, err := s.DoCommand(listHostBitCommand, ListHostBitRequest{
		ipAddr,
	})
	if err != nil {
		return ListHostBitsResponse{}, err
	}
	// unmarshal into the iface stat struct
	var hostbits ListHostBitsResponse
	if err := json.Unmarshal(response.Message, &hostbits); err != nil {
		return ListHostBitsResponse{}, err
	}
	return hostbits, nil
}
