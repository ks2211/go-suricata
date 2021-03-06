package client

import "context"

const (
	addHostBit    string = "add-hostbit"
	removeHostBit string = "remove-hostbit"
	listHostBit   string = "list-hostbit"
)

// AddHostBitRequest holds the ip, bit name and expiration for adding a host bit
// Not used in v3
type AddHostBitRequest struct {
	IPAddr  string `json:"ipaddress"`
	BitName string `json:"hostbit"`
	Expire  int    `json:"expire"`
}

// RemoveHostBitRequest holds the ip, bit name for removing a host bit
// Not used in v3
type RemoveHostBitRequest struct {
	IPAddr  string `json:"ip"`
	BitName string `json:"hostbit"`
}

// ListHostBitRequest holds the host for list hostbit command
// Not used in v3
type ListHostBitRequest struct {
	IPAddr string `json:"ipaddress"`
}

// ListHostBitsResponse holds response from list-hostbits <ip>
type ListHostBitsResponse struct {
	Count    int `json:"count"`
	Hostbits []struct {
		Expire int    `json:"expire"`
		Name   string `json:"name"`
	} `json:"hostbits"`
}

// AddOrRemoveHostBitResponse holds response from adding or removing a hostbit <ip> <bitname>
type AddOrRemoveHostBitResponse string

// String is a helper method to turn go type into struct
func (a AddOrRemoveHostBitResponse) String() string {
	return string(a)
}

// AddHostBitCommand adds host bit
// Not implemented in v3
func (s *Socket) AddHostBitCommand(ctx context.Context, req AddHostBitRequest) (string, error) {
	addHostbitResp := new(AddOrRemoveHostBitResponse)
	err := s.DoCommand(ctx, addHostBit, req, addHostbitResp)
	return addHostbitResp.String(), err
}

// RemoveHostBitCommand does "remove-hostbit"
// Not implemented in v3
func (s *Socket) RemoveHostBitCommand(ctx context.Context, req RemoveHostBitRequest) (string, error) {
	removeHostBitResp := new(AddOrRemoveHostBitResponse)
	err := s.DoCommand(ctx, removeHostBit, req, removeHostBitResp)
	return removeHostBitResp.String(), err
}

// ListHostBitCommand does "list-hostbit"
// Not implemented in v3
func (s *Socket) ListHostBitCommand(ctx context.Context, req ListHostBitRequest) (ListHostBitsResponse, error) {
	hostbitsResp := ListHostBitsResponse{}
	err := s.DoCommand(ctx, listHostBit, req, &hostbitsResp)
	return hostbitsResp, err
}
