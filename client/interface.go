package client

// Socket is a interface for all base suricata (v3) commands
type SocketIFace interface {
	// Connection helpers
	Close()
	Dial() error

	// Helper methods
	Path() string
	DoCommand(string, interface{}, interface{}) error
	ReadResponse(string) (*Response, error)
	SendMessage(string, interface{}) error
	write([]byte) (int, error)
	receive() (*Response, error)

	// v3
	CaptureModeCommand() (string, error)
	CommandListCommand() (*CommandListResponse, error)
	ConfGetCommand(ConfGetRequest) (string, error)
	DumpCountersCommand() (*DumpCountersResponse, error)
	IFaceListCommand() (*IFaceListResponse, error)
	IFaceStatCommand(IFaceStatRequest) (*IFaceStatResponse, error)
	RegisterTenantCommand(RegisterTenantRequest) (*RegisterTenantResposne, error)
	ReloadRulesCommand() (string, error)
	RunningModeCommand() (string, error)
	ShutdownCommand() (string, error)
	UptimeCommand() (int, error)
	VersionCommand() (string, error)

	// v4
	AddHostBitCommand(AddHostBitRequest) (string, error)
	ListHostBitCommand(ListHostBitRequest) (*ListHostBitsResponse, error)
	MemCapSetCommand(interface{}, interface{}) (string, error)
	MemCapShowCommand(MemCapShowRequest) (*MemCapShowResponse, error)
	MemCapListCommand() (*[]MemCapListResponse, error)
	RemoveHostBitCommand(RemoveHostBitRequest) (string, error)
	ReopenLogFilesCommand() (string, error)
	RulesetReloadRulesCommand() (string, error)
	RulesetReloadNonBlockingCommand() (string, error)
	RulesetReloadTimeCommand() (*[]RulesetReloadTimeResponse, error)
	RulesetStatsCommand() (*[]RulesetStatsResponse, error)
	RulesetFailedRulesCommand() (*[]RulesetFailedRulesResponse, error)
}

// Iface check
var _ SocketIFace = (*Socket)(nil)
