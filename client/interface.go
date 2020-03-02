package client

// Socket is a interface for all base suricata (v3) commands
type Socket interface {
	// Connection helpers
	Close()
	Dial() error

	// Helper methods
	Path() string
	DoCommand(string, interface{}) (*Response, error)
	ReadResponse(string) (*Response, error)
	SendMessage(string, interface{}) error
	write([]byte) (int, error)
	receive() (*Response, error)

	// v3
	CaptureModeCommand() (string, error)
	CommandListCommand() (*CommandListResponse, error)
	ConfGetCommand(string) (string, error)
	DumpCountersCommand() (*DumpCountersResponse, error)
	IFaceListCommand() (*IFaceListResponse, error)
	IFaceStatCommand(string) (*IFaceStatResponse, error)
	ReloadRulesCommand() (string, error)
	RunningModeCommand() (string, error)
	ShutdownCommand() (string, error)
	UptimeCommand() (int, error)
	VersionCommand() (string, error)

	// v4
	AddHostBitCommand(string, string, int) (string, error)
	ListHostBitCommand(string) (*ListHostBitsResponse, error)
	MemCapSetCommand(interface{}, interface{}) (string, error)
	MemCapShowCommand(string) (*MemCapShowResponse, error)
	MemCapListCommand() ([]MemCapListResponse, error)
	RemoveHostBitCommand(string, string) (string, error)
	ReopenLogFilesCommand() (string, error)
	RulesetReloadRulesCommand() (string, error)
	RulesetReloadNonBlockingCommand() (string, error)
	RulesetReloadTimeCommand() ([]RulesetReloadTimeResponse, error)
	RulesetStatsCommand() ([]RulesetStatsResponse, error)
	RulesetFailedRulesCommand() ([]RulesetFailedRulesResponse, error)
}
