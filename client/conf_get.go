package client

const (
	confGet string = "conf-get"
)

// ConfGetRequest holds the config key item for conf-get command
type ConfGetRequest struct {
	Variable string `json:"variable"`
}

// ConfGetResponse is resposne from "conf-get <config.key>"
// This command only allows for returning string/int/bool, it does not return arrays (TODO)
type ConfGetResponse StringResponse

// String is a helper method to convert a go type into string
func (c ConfGetResponse) String() string {
	return string(c)
}

// ConfGetCommand performs a fetch on a config item using a key string, only returns string/int/bool values
func (s *Socket) ConfGetCommand(confGetRequest ConfGetRequest) (string, error) {
	configResp := new(ConfGetResponse)
	err := s.DoCommand(confGet, confGetRequest, configResp)
	return configResp.String(), err
}
