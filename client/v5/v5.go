package v5

import (
	v4 "github.com/ks2211/go-suricata/client/v4"
)

// v5 commands not in v3, v4

// Constants
const (
	// DefaultSocketPathV5 is default socket path of suricata 5
	DefaultSocketPathV5 string = "/var/run/suricata/suricata-command.socket"
)

// SocketV5 is suricata V4+ commands
// Embed the v4 socket client (which embeds the v3 socket client)
type SocketV5 struct {
	v4.Client
}

// NewSocketV5 is constructor func that takes in the initial v3+v4 socket and adds additional methods for v5
func NewSocketV5(path string) (Client, error) {
	v4Client, err := v4.NewSocketV4(path)
	if err != nil {
		return nil, err
	}
	return &SocketV5{
		v4Client,
	}, nil
}
