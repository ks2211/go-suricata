package client

import "context"

// Socket is a interface for all base suricata (v3) commands
type SocketIFace interface {
	// Connection helpers
	Close()
	Dial() error

	// Helper methods
	Path() string
	DoCommand(ctx context.Context, command string, args interface{}, resp interface{}) error
	DoCommandWithResponse(ctx context.Context, command string, args interface{}) (*Response, error)
	ReadResponse(string) (*Response, error)
	SendMessage(string, interface{}) error
	write([]byte) (int, error)
	receive() (*Response, error)

	// v3
	CaptureModeCommand(ctx context.Context) (string, error)
	CommandListCommand(ctx context.Context) (*CommandListResponse, error)
	ConfGetCommand(ctx context.Context, req ConfGetRequest) (string, error)
	DumpCountersCommand(ctx context.Context) (*DumpCountersResponse, error)
	HelpCommand(ctx context.Context) (*CommandListResponse, error)
	IFaceListCommand(ctx context.Context) (*IFaceListResponse, error)
	IFaceStatCommand(ctx context.Context, req IFaceStatRequest) (*IFaceStatResponse, error)

	ReloadTenantCommand(ctx context.Context, req ReloadTenantRequest) (*ReloadTenantResponse, error)
	RegisterTenantCommand(ctx context.Context, req RegisterTenantRequest) (*RegisterTenantResposne, error)
	RegisterTenantHandlerCommand(ctx context.Context, req RegisterTenantHandlerRequest) (*RegisterTenantHandlerResponse, error)
	UnRegisterTenantCommand(ctx context.Context, req UnRegisterTenantRequest) (*UnRegisterTenantResponse, error)
	UnRegisterTenantHandlerCommand(ctx context.Context, req UnRegisterTenantHandlerRequest) (*UnRegisterTenantHandlerResponse, error)

	ReloadRulesCommand(ctx context.Context) (string, error)
	RunningModeCommand(ctx context.Context) (string, error)
	ShutdownCommand(ctx context.Context) (string, error)
	UptimeCommand(ctx context.Context) (int, error)
	VersionCommand(ctx context.Context) (string, error)

	// v4
	AddHostBitCommand(ctx context.Context, req AddHostBitRequest) (string, error)
	ListHostBitCommand(ctx context.Context, req ListHostBitRequest) (*ListHostBitsResponse, error)
	MemCapSetCommand(ctx context.Context, memCapName interface{}, memCapValue interface{}) (string, error)
	MemCapShowCommand(ctx context.Context, req MemCapShowRequest) (*MemCapShowResponse, error)
	MemCapListCommand(ctx context.Context) (*[]MemCapListResponse, error)
	RemoveHostBitCommand(ctx context.Context, req RemoveHostBitRequest) (string, error)
	ReopenLogFilesCommand(ctx context.Context) (string, error)
	RulesetReloadRulesCommand(ctx context.Context) (string, error)
	RulesetReloadNonBlockingCommand(ctx context.Context) (string, error)
	RulesetReloadTimeCommand(ctx context.Context) (*[]RulesetReloadTimeResponse, error)
	RulesetStatsCommand(ctx context.Context) (*[]RulesetStatsResponse, error)
	RulesetFailedRulesCommand(ctx context.Context) (*[]RulesetFailedRulesResponse, error)

	// PCAP
	PcapFileCommand(ctx context.Context, req PcapFileRequest) (string, error)
	PcapFileContinuousCommand(ctx context.Context, req PcapFileRequest) (string, error)
	PcapFileNumberCommand(ctx context.Context) (int, error)
	PcapFileListCommand(ctx context.Context) (*PcapFileListResponse, error)
	PcapLastProcessedCommand(ctx context.Context) (int, error)
	PcapInterruptCommand(ctx context.Context) (string, error)
	PcapCurrentCommand(ctx context.Context) (string, error)

	// V5
	// IFaceByPassStatCommand() error
	// EbpfByPassStatCommand() error
	// DataSetAddCommand() error // 3 args
}

// Iface check
var _ SocketIFace = (*Socket)(nil)
