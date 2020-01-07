package v4

import (
	"encoding/json"

	v3 "github.com/ks2211/go-suricata/client/v3"
)

// v4 commands not in v3
// ruleset-reload-rules, ruleset-reload-nonblocking, ruleset-reload-time, ruleset-stats, ruleset-failed-rules
// memcap-set, memcap-show, memcap-list ---> validate memcap-set value
// add-hostbit, remove-hostbit, list-hostbit
// reopen-log-files

// Constants
const (
	// DefaultSocketPathV4 is default socket path of suricata 4
	DefaultSocketPathV4 string = "/var/run/suricata/suricata-command.socket"
)

// SocketV4 is suricata V4+ commands
type SocketV4 struct {
	v3.Client
}

// Constants
const (
	reopenLogFilesCommand string = "reopen-log-files"
)

// NewSocketV4 creates the socket v3 and v4 structs
func NewSocketV4(path string) (Client, error) {
	v3Client, err := v3.NewSocketV3(path)
	if err != nil {
		return nil, err
	}
	return &SocketV4{
		v3Client,
	}, nil
}

// ReopenLogFilesCommand reopens log files via the reopen-logfiles command
func (s *SocketV4) ReopenLogFilesCommand() (string, error) {
	// create and marshal the "reopen-log-files" socket message with
	response, err := s.DoCommand(reopenLogFilesCommand, nil)
	if err != nil {
		return "", err
	}
	// unmarshal into the iface stat struct
	var reopenLogFiles ReopenLogFilesResponse
	if err := json.Unmarshal(response.Message, &reopenLogFiles); err != nil {
		return "", err
	}
	return reopenLogFiles.String(), nil
}
