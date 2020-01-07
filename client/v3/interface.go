package v3

import "github.com/ks2211/go-suricata/client"

// Client is the interface that holds the suricata methods
type Client interface {
	client.Client
	// Suricata v3 commands
	CaptureModeCommand() (string, error)
	CommandListCommand() ([]string, error)
	ConfGetCommand(string) (string, error)
	DumpCountersCommand() (client.DumpCountersResponse, error)
	IFaceListCommand() (client.IFaceListResponse, error)
	IFaceStatCommand(string) (client.IFaceStatResponse, error)
	ReloadRulesCommand() (string, error)
	RunningModeCommand() (string, error)
	ShutdownCommand() (string, error)
	UptimeCommand() (int, error)
	VersionCommand() (string, error)
}
