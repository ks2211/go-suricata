package v3

// TODO
// register-tenant-handler, unregister-tenant-handler, register-tenant, reload-tenant, unregister-tenant
// pcap commands

import (
	"github.com/ks2211/go-suricata/client"
)

const (
	// DefaultSocketPathV3 is default path of the socket
	DefaultSocketPathV3 string = "/var/run/suricata-command.socket"
)

// SocketV3 is suricata version 3 socket client
type SocketV3 struct {
	*client.Socket
}

// NewSocketV3 returns socket v3 client
func NewSocketV3(path string) (Client, error) {
	sock, err := client.CreateSocket(path)
	if err != nil {
		return nil, err
	}
	return &SocketV3{
		sock,
	}, nil
}
