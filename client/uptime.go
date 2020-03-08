package client

const (
	uptime string = "uptime"
)

// UptimeResponse is response from "uptime"
type UptimeResponse int

// UptimeCommand gets uptime of suricata
func (s *Socket) UptimeCommand() (int, error) {
	uptimeResp := new(UptimeResponse)
	err := s.DoCommand(uptime, nil, uptimeResp)
	return int(*uptimeResp), err
}
