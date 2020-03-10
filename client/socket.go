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

// https://github.com/OISF/suricata/blob/master/python/suricata/sc/specs.py

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

// Response is the response from connecting.
// The Message is a json.RawMessage that can be unmarshalled into individual types
type Response struct {
	Return  string          `json:"return"`
	Message json.RawMessage `json:"message"`
}

// ConnectResponse is the response string from connecting
type ConnectResponse string

// StringResponse is a string response from an informational string (instead of the json)
type StringResponse string

// String is a helper method to convert a go type into string
func (i StringResponse) String() string {
	return string(i)
}

// Socket holds the net conn
type Socket struct {
	path string
	conn net.Conn
}

// CreateSocket returns a socket with the path set
func CreateSocket(path string) (SocketIFace, error) {
	// create struct
	sock := Socket{
		path: path,
	}

	// establish conn to socket
	if err := sock.Dial(); err != nil {
		return nil, err
	}
	return &sock, nil
}

// Dial does a net dial to the unix socket and sets the struct conn
func (s *Socket) Dial() error {
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
func (s *Socket) Path() string {
	return s.path
}

// Close is a helper func to close the net conn (use in a defer)
func (s *Socket) Close() {
	s.conn.Close()
}

// DoCommand is a helper function to run a command with optional args and unmarshals into the supplied interface
func (s *Socket) DoCommand(command string, args, resp interface{}) error {
	// send message and get response
	response, err := s.DoCommandWithResponse(command, args)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response.Message, resp)
	if err != nil {
		return err
	}
	return nil
}

// DoCommandWithResponse is a helper function to run a command with optional args and returns the raw json.RawMessage response
func (s *Socket) DoCommandWithResponse(command string, args interface{}) (*Response, error) {
	// send the command and marshal the args before sending
	if err := s.SendMessage(command, args); err != nil {
		return nil, err
	}
	// get response
	return s.ReadResponse(command)
}

// ReadResponse gets response from the tcp socket read
func (s *Socket) ReadResponse(commandName string) (*Response, error) {
	// read from socket
	response, err := s.receive()
	if err != nil {
		return nil, err
	}
	// if response not ok, check the message and return error
	if response.Return != responseOK {
		var errMessage StringResponse
		if err := json.Unmarshal(response.Message, &errMessage); err != nil {
			return nil, err
		}
		if errMessage.String() == ErrUnknownCommand.Error() {
			return nil, fmt.Errorf("Error command %s unsupported in this version %v", commandName, errMessage)
		}
		return nil, fmt.Errorf("Error doing command %s, %v", commandName, errMessage)
	}
	return response, nil
}

// write is a helper method to write bytes to the socket
func (s *Socket) write(data []byte) (int, error) {
	return s.conn.Write(data)
}

// receive is a helper method to receive a response from the socket
func (s *Socket) receive() (*Response, error) {
	response := Response{}
	decoder := json.NewDecoder(s.conn)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}
