package v5

import v4 "github.com/ks2211/go-suricata/client/v4"

// Client holds v5 methods and implements v3,v4 methods
type Client interface {
	v4.Client // v4 client to implement v3 and v4 methods
	// TODO v5 methods
}
