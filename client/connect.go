package client

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// Dial does a net dial to the unix socket and sets the struct conn
func (s *Socket) Dial() error {
	// establish conn to socket
	conn, err := net.DialTimeout("unix", s.path, time.Second*3)
	if err != nil {
		return err
	}
	s.conn = conn
	return nil
}

// ConnectSocket does the initial client version send message to connect to the socket
func (s *Socket) ConnectSocket() error {
	// marshal connect message to bytes
	data, err := json.Marshal(socketInitMessage)
	if err != nil {
		return err
	}
	// write message
	if _, err := s.write(data); err != nil {
		return err
	}
	// get response
	response, err := s.receive()
	if err != nil {
		return err
	}
	// if response not ok, check the message and return error
	if response.Return != responseOK {
		var connMessage ConnectResponse
		if err := json.Unmarshal(response.Message, &connMessage); err != nil {
			return err
		}
		return fmt.Errorf("Could not connect to socket %v", connMessage)
	}
	return nil
}

// Close is a helper func to close the net conn (use in a defer)
func (s *Socket) Close() {
	s.conn.Close()
}
