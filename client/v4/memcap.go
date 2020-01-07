package v4

import (
	"encoding/json"
	"errors"
	"strings"
)

// Constants
const (
	memcapSetCommand  string = "memcap-set"
	memcapShowCommand string = "memcap-show"
	memcapListCommand string = "memcap-list"
)

// MemCapSetRequest is the json object needed for memcap-set
// It requires a name field of the memcap
// It also requires the memcap value which can be unlimited or > "6mb" (string)
type MemCapSetRequest struct {
	Name  string      `json:"config"`
	Value interface{} `json:"memcap"`
}

// MemCapShowRequest is the json object needed for memcap-show
// It requires a name field of the memcap
type MemCapShowRequest struct {
	Name string `json:"config"`
}

// MemCapSetCommand sets memcap
func (s *SocketV4) MemCapSetCommand(memcapName, memcapValue interface{}) (string, error) {
	// TODO validate the memcap value--can be "unlimited" or > 6mb
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
	response, err := s.DoCommand(memcapSetCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the iface stat struct
	var memcapSet MemCapSetResponse
	if err := json.Unmarshal(response.Message, &memcapSet); err != nil {
		return "", err
	}
	return memcapSet.String(), nil
}

// MemCapShowCommand does "memcap-show"
func (s *SocketV4) MemCapShowCommand(memcapName string) (MemCapShowResponse, error) {
	// create and marshal the "memcap-show" socket message with
	response, err := s.DoCommand(memcapShowCommand, MemCapShowRequest{
		Name: memcapName,
	})
	if err != nil {
		return MemCapShowResponse{}, err
	}
	// unmarshal into the iface stat struct
	var memcapShow MemCapShowResponse
	if err := json.Unmarshal(response.Message, &memcapShow); err != nil {
		return MemCapShowResponse{}, err
	}
	return memcapShow, nil
}

// MemCapListCommand does "memcap-list"
func (s *SocketV4) MemCapListCommand() ([]MemCapListResponse, error) {
	// create and marshal the "memcap-show" socket message with the ifaceStatRequest arg
	response, err := s.DoCommand(memcapListCommand, nil)
	if err != nil {
		return nil, err
	}
	// unmarshal into the iface stat struct
	var memcapList []MemCapListResponse
	if err := json.Unmarshal(response.Message, &memcapList); err != nil {
		return nil, err
	}
	return memcapList, nil
}
