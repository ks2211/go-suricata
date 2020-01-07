[![Documentation](https://godoc.org/github.com/ks2211/go-suricata?status.svg)](http://godoc.org/github.com/ks2211/go-suricata)
[![Go Report Card](https://goreportcard.com/badge/github.com/ks2211/go-suricata)](https://goreportcard.com/report/github.com/ks2211/go-suricata)

# Go Suricata Client 

Gosuricata is a Go client library for interacting with suricata using the unix socket

## Prerequisites

* Go 1.12+ installed

* Suricata installed and running

* Suricata socket enabled in the suricata config along with the socket path

    * See this [link](https://suricata.readthedocs.io/en/suricata-4.1.3/unix-socket.html#introduction) for enabling the socket

## Installation

Go get the library

```bash
go get github.com/ks2211/go-suricata
```

## Usage

```go
package main

// import the specific client library you need
import (
    "github.com/ks2211/go-suricata/client/v3"
    "log"
    //v4 "github.com/ks2211/go-suricata/client/v4"
    //v5 "github.com/ks2211/go-suricata/client/v5"
)

func main() {
    // create the client
    s, err := socket.NewSocketV3(v3.DefaultSocketPathV3)
    //s, err := v4.NewSocketV4(v4.DefaultSocketPathV4)
    if err != nil {
        log.Fatalf("Error conn %v", err)
    }
    // connect the socket and defer clsoe
    if err := s.ConnectSocket(); err != nil {
		log.Fatalf("error connect %v", err)
	}
    defer s.Close()
    // use the client to run command methods
    commands, err := s.GetCommandList()
	if err != nil {
		log.Fatalf("Error command list %v", err)
    }
    log.Println("COMMANDS", commands)
    runningMode, err := s.RunningModeCommand()
    if err != nil {
        log.Fatalf("error running mode %v", err)
    }
    log.Println("RUN MODE", runningMode)
    // run a command manually--note you will have to pass in a struct/map/interface
    // the type can be marshalled into json
    r, err := s.DoCommand("some-command", struct{
        Field string `json:"field"`
    }{
        "test"
    })
    if err != nil {
        log.Fatalf("error running command %v", err)
    }
    // handle response
    retData := map[string]interface{} // or struct
    if err := json.Unmarshal(r.Message, &retData); err != nil {
        log.Fatalf("error unmarshal data %v", err)
    }
    log.Println("response", r.Status, retData)
}
```

## Design

The way the library/client is set up:

* The client directory holds the client connection and base command methods (sending, receving, parsing)

* Suricata 4 implements all of the methods from Suricata 3 (and subsequently Suricata 5 implements all methods from both 3 and 4)

* If you are using Suricata 4 or 5, you can import those packages and call the client constructor methods (e.g `NewSocketV4()` or `NewSocketV5()`)

    * This will automatically implement the v3 commands as well give you access to the v4 or v5 commands

* The clients implements most (if not all) commands for that specific Suricata version. The command methods that are not implemented by the library, the user has the option of using the `DoCommand` method to run the command manually.

* Each command returns a JSON response (shown below) with a status and a "message" that is either a JSON object or a string with the response from the command

```json
{
    "status": "OK|NOK",
    "message": {}|""
}
```

* For implemented methods, the library will parse the message into Go types (either string for string messages or a Go struct for json message)

* For unimplemented methods, the `message` field will be  converted to Go's `json.RawMessage`. A user can unmarshal this field into their own Go type. 

```go
resp, err := s.DoCommand("command-list", nil)
if err != nil {
    // handle err
}
fmt.Println(string(resp.Message)) // prints string of json
response := struct{
    FieldA string `json:"field_a"`
}{}
if err := json.Unmarshal(resp.Message, &response); err != nil {
    // handle err
}
// work with the response
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Testing

*IN PROGRESS*

* In order to run tests, you will need to either run them as sudo/root or have access to the suricata socket

* To run tests, run `go test`. You can pass in the `-run=<TEST_CASE_NAME>` flag to run a specific test

* The tests for the base client relies on Suricata Version 3.2

* *TODO*: Most of the test case expected values (e.g interface names, suricata config items) have been hardcoded to my personal machine

* *TODO*: There is a lot of duplicate/repeated code around creating the socket client/closing the connection for each testing method--need test helpers

* *TODO*: Look into using table test cases for more of the command test cases


## TODO

* Testing (In progress)

* Tenant and Pcap Commands

* V5 Commands

* General cleanup (idomatic patterns, gofmt, lint, etc)

* Fill this in