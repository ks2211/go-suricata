package client

const (
	version string = "version"
)

// VersionResponse is response from "version"
type VersionResponse StringResponse

// String is a helper method to convert a go type into string
func (v VersionResponse) String() string {
	return string(v)
}

// VersionCommand gets version of suricata and sets the version in the pointer
func (s *Socket) VersionCommand() (string, error) {
	versionResp := new(VersionResponse)
	err := s.DoCommand(version, nil, versionResp)
	return versionResp.String(), err
}
