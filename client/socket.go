package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"time"
)

// TODO V3
// V3: register-tenant-handler, unregister-tenant-handler, register-tenant, reload-tenant, unregister-tenant
// V5: all commands
// pcap commands

// Errors
var (
	ErrUnimplemented  error = errors.New("error unimplemented in library")
	ErrUnknownCommand error = errors.New("Unknown command")
)

// Constants
const (
	responseOK    string = "OK"
	responseNOK   string = "NOK"
	clientVersion string = "0.1"

	// DefaultSocketPathV3 is default path of the socket
	DefaultSocketPathV3 string = "/var/run/suricata-command.socket"
	// DefaultSocketPathV4 is default socket path of suricata 4
	DefaultSocketPathV4 string = "/var/run/suricata/suricata-command.socket"
	// DefaultSocketPathV5 is default socket path of suricata 5
	DefaultSocketPathV5 string = "/var/run/suricata/suricata-command.socket"
)

// Socket holds the net conn
type socket struct {
	path string
	conn net.Conn
}

// CreateSocket returns a socket with the path set
func CreateSocket(path string) (Socket, error) {
	// create struct
	sock := socket{
		path: path,
	}

	// establish conn to socket
	if err := sock.Dial(); err != nil {
		return nil, err
	}
	return &sock, nil
}

// Dial does a net dial to the unix socket and sets the struct conn
func (s *socket) Dial() error {
	// establish conn to socket
	conn, err := net.DialTimeout("unix", s.path, time.Second*3)
	if err != nil {
		return err
	}
	s.conn = conn
	// marshal connect message to bytes
	data, err := json.Marshal(map[string]string{
		"version": clientVersion,
	})
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

// Path returns socket path
func (s *socket) Path() string {
	return s.path
}

// Close is a helper func to close the net conn (use in a defer)
func (s *socket) Close() {
	s.conn.Close()
}

// DoCommand is a helper function to run a command and get the unmarshaled high level response (the message is json.RawMessage)
func (s *socket) DoCommand(command string, args interface{}) (*Response, error) {
	// send the command and marshal the args before sending
	if err := s.SendMessage(command, args); err != nil {
		return nil, err
	}
	// get response
	response, err := s.ReadResponse(command)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// ReadResponse gets response from the tcp socket read
func (s *socket) ReadResponse(commandName string) (*Response, error) {
	// read from socket
	response, err := s.receive()
	if err != nil {
		return nil, err
	}
	// if response not ok, check the message and return error
	if response.Return != responseOK {
		var errMessage InformationResponse
		if err := json.Unmarshal(response.Message, &errMessage); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Error doing command %s, %v", commandName, errMessage)
	}
	return response, nil
}

// write is a helper method to write bytes to the socket
func (s *socket) write(data []byte) (int, error) {
	return s.conn.Write(data)
}

// receive is a helper method to receive a response from the socket
func (s *socket) receive() (*Response, error) {
	response := Response{}
	decoder := json.NewDecoder(s.conn)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
