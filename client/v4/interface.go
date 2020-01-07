package v4

import v3 "github.com/ks2211/go-suricata/client/v3"

// Client holds v4 methods and implements v3 methods
type Client interface {
	v3.Client
	// Suricata v4 commands (unimplemented by v3)
	AddHostBitCommand(string, string, int) (string, error)
	ListHostBitCommand(string) (ListHostBitsResponse, error)
	MemCapSetCommand(interface{}, interface{}) (string, error)
	MemCapShowCommand(string) (MemCapShowResponse, error)
	MemCapListCommand() ([]MemCapListResponse, error)
	RemoveHostBitCommand(string, string) (string, error)
	ReopenLogFilesCommand() (string, error)
	RulesetReloadRulesCommand() (string, error)
	RulesetReloadNonBlockingCommand() (string, error)
	RulesetReloadTimeCommand() ([]RulesetReloadTimeResponse, error)
	RulesetStatsCommand() ([]RulesetStatsResponse, error)
	RulesetFailedRulesCommand() ([]RulesetFailedRulesResponse, error)
}
