package client

// Client is a interface for all base suricata (v3) commands
type Client interface {
	// Connection helpers
	ConnectSocket() error
	Close()
	Dial() error

	// Helper methods
	DoCommand(string, interface{}) (*Response, error)
	GetResponseFromRead(string) (*Response, error)
	SendMessage(string, interface{}) error
	write([]byte) (int, error)
	receive() (*Response, error)
}
