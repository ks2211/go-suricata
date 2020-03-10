package client

import (
	"errors"
	"fmt"
	"strconv"
)

const errCommandArgMissing = "error command arg(s) are missing %s"

var (
	errGetArgType = errors.New("error getting argument type")
)

// commandArgsMap holds the map of commands that require args
var commandArgsMap = map[string]interface{}{
	confGet:                 ConfGetRequest{},
	ifaceStat:               IFaceStatRequest{},
	registerTenant:          RegisterTenantRequest{},
	unregisterTenant:        UnRegisterTenantRequest{},
	registerTenantHandler:   RegisterTenantHandlerRequest{},
	unregisterTenantHandler: UnRegisterTenantHandlerRequest{},
	reloadTenant:            ReloadTenantRequest{},
	addHostBit:              AddHostBitRequest{},
	listHostBit:             ListHostBitRequest{},
	removeHostBit:           RemoveHostBitRequest{},
	memcapShow:              MemCapShowRequest{},
}

// BuildCommandArgs checks the command name and builds the appropriate input struct and returns as interface
func BuildCommandArgs(commandName string, args []string) (interface{}, error) {
	switch commandName {
	// conf-get
	case confGet:
		if len(args) != 1 || args[0] == "" {
			return nil, fmt.Errorf(errCommandArgMissing, "variable")
		}
		v, ok := commandArgsMap[confGet].(ConfGetRequest)
		if !ok {
			return nil, errGetArgType
		}
		v.Variable = args[0]
		return v, nil
	// iface-stat
	case ifaceStat:
		if len(args) != 1 || args[0] == "" {
			return nil, fmt.Errorf(errCommandArgMissing, "iface")
		}
		v, ok := commandArgsMap[ifaceStat].(IFaceStatRequest)
		if !ok {
			return nil, errGetArgType
		}
		v.IFace = args[0]
		return v, nil
	// add-hostbit (v4+)
	case addHostBit:
		if len(args) != 3 {
			return nil, fmt.Errorf(errCommandArgMissing, "ipaddress, hostbit, expire")
		}
		v, ok := commandArgsMap[addHostBit].(AddHostBitRequest)
		if !ok {
			return nil, errGetArgType
		}
		v.IPAddr = args[0]
		v.BitName = args[1]
		expire, err := strconv.Atoi(args[2])
		if err != nil {
			return nil, errors.New("error expire must be a number")
		}
		v.Expire = expire
		return v, nil
	// list-hostbit (v4+)
	case listHostBit:
		if len(args) != 1 || args[0] == "" {
			return nil, fmt.Errorf(errCommandArgMissing, "ipaddress")
		}
		v, ok := commandArgsMap[listHostBit].(ListHostBitRequest)
		if !ok {
			return nil, errGetArgType
		}
		v.IPAddr = args[0]
		return v, nil
	// remove-hostbit (v4+)
	case removeHostBit:
		if len(args) != 2 {
			return nil, fmt.Errorf(errCommandArgMissing, "ipaddress, hostbit")
		}
		v, ok := commandArgsMap[removeHostBit].(RemoveHostBitRequest)
		if !ok {
			return nil, errGetArgType
		}
		v.IPAddr = args[0]
		v.BitName = args[1]
		return v, nil
	default:
		return nil, errors.New("error unknown command or currently unimplemented")
	}
}
