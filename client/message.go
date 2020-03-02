package client

import "encoding/json"

// socketMessage is the json message thats being sent to the suricata socket
type socketMessage struct {
	Command   string      `json:"command"`
	Arguments interface{} `json:"arguments"`
}

// createSocketMessage creates socket message
func createSocketMessage(command string, args interface{}) socketMessage {
	return socketMessage{
		command,
		args,
	}
}

// marshalSocketMessage marshals the socket message into json bytes
func (msg socketMessage) marshalSocketMessage() ([]byte, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// SendMessage sends a message command to the socket
func (s *socket) SendMessage(command string, args interface{}) error {
	// create and marshal the "version" socket message with no args
	msg := createSocketMessage(command, args)
	socketMsg, err := msg.marshalSocketMessage()
	if err != nil {
		return err
	}
	// write to the socket
	if _, err := s.write(socketMsg); err != nil {
		return err
	}
	return nil
}
