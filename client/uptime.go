package client

import "context"

const (
	uptime string = "uptime"
)

// UptimeResponse is response from "uptime"
type UptimeResponse int

// UptimeCommand gets uptime of suricata
func (s *Socket) UptimeCommand(ctx context.Context) (int, error) {
	uptimeResp := new(UptimeResponse)
	err := s.DoCommand(ctx, uptime, nil, uptimeResp)
	return int(*uptimeResp), err
}
