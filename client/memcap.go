package client

import (
	"context"
	"errors"
	"strings"
)

const (
	memcapSet  string = "memcap-set"
	memcapShow string = "memcap-show"
	memcapList string = "memcap-list"
)

// MemCapSetRequest is the json object needed for memcap-set
// It requires a name field of the memcap
// It also requires the memcap value which can be unlimited or > "6mb" (string)
// Not used in v3
type MemCapSetRequest struct {
	Name  string      `json:"config"`
	Value interface{} `json:"memcap"`
}

// MemCapShowRequest is the json object needed for memcap-show
// It requires a name field of the memcap
// Not used in v3
type MemCapShowRequest struct {
	Name string `json:"config"`
}

// MemCapListResponse is message from memcap-list
type MemCapListResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// MemCapShowResponse is message from memcap-show of a specific memcap
type MemCapShowResponse struct {
	Value string `json:"value"`
}

// MemCapSetResponse is message from memcap-set for a specific memcat
type MemCapSetResponse StringResponse

// String is a helper method to turn go type into struct
func (m MemCapSetResponse) String() string {
	return string(m)
}

// MemCapSetCommand sets memcap
// Not implemented in v3
func (s *Socket) MemCapSetCommand(ctx context.Context, memcapName, memcapValue interface{}) (string, error) {
	if unlimited, ok := memcapValue.(string); ok {
		if strings.ToLower(unlimited) != "unlimited" {
			return "", errors.New("only unlimited can be passed in if passing string memcap value")
		}
	} else if memcapIntVal, ok := memcapValue.(int); ok {
		// TODO memcap value minimums aren't explicitly defined anywhere
		// 1000000 is 1mb which is usually the bare miniumum (TODO)
		if memcapIntVal < 1000000 {
			return "", errors.New("memcap int value is too small, must be over 1000000")
		}
	}
	// create and marshal the "memcap-set" socket message with
	memcapSetResp := new(MemCapSetResponse)
	err := s.DoCommand(ctx, memcapSet, nil, memcapSetResp)
	return memcapSetResp.String(), err
}

// MemCapShowCommand does "memcap-show"
// Not implemented in v3
func (s *Socket) MemCapShowCommand(ctx context.Context, req MemCapShowRequest) (MemCapShowResponse, error) {
	// create and marshal the "memcap-show" socket message with
	memCapShowResp := MemCapShowResponse{}
	err := s.DoCommand(ctx, memcapShow, req, &memCapShowResp)
	return memCapShowResp, err
}

// MemCapListCommand does "memcap-list"
// Not implemented in v3
func (s *Socket) MemCapListCommand(ctx context.Context) ([]MemCapListResponse, error) {
	// create and marshal the "memcap-show" socket message with the ifaceStatRequest arg
	memcapListResp := []MemCapListResponse{}
	err := s.DoCommand(ctx, memcapList, nil, &memcapListResp)
	return memcapListResp, err
}
