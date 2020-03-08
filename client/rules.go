package client

const (
	reloadRules string = "reload-rules"
)

// ReloadRulesResponse is response from "reload-rules"
type ReloadRulesResponse StringResponse

// String is a helper method to convert a go type into string
func (r ReloadRulesResponse) String() string {
	return string(r)
}

// ReloadRulesCommand performs a reload of rules without restarting suricata
func (s *Socket) ReloadRulesCommand() (string, error) {
	reloadRulesResp := new(ReloadRulesResponse)
	err := s.DoCommand(reloadRules, nil, reloadRulesResp)
	return reloadRulesResp.String(), err
}
