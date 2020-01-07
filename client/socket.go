package client

import (
	"encoding/json"
	"fmt"
	"net"
)

// Constants
const (
	responseOK    string = "OK"
	responseNOK   string = "NOK"
	clientVersion string = "0.1"
)

// socketInitMessage is the initial connect message
var socketInitMessage = map[string]string{
	"version": clientVersion,
}

// Socket holds the net conn
type Socket struct {
	path string
	conn net.Conn
}

// CreateSocket returns a socket with the path set
func CreateSocket(path string) (*Socket, error) {
	// create struct
	sock := &Socket{
		path: path,
	}

	// establish conn to socket
	if err := sock.Dial(); err != nil {
		return nil, err
	}
	return sock, nil
}

// DoCommand is a helper function to run a command and get the unmarshaled high level response (the message is json.RawMessage)
func (s *Socket) DoCommand(command string, args interface{}) (*Response, error) {
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
func (s *Socket) ReadResponse(commandName string) (*Response, error) {
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
