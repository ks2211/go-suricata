/*

Package client is a Go client library for interacting with suricata via unix socket.

The type-safe library implements most (if not all) commands for suricata v3, v4, and v5 along with the pcap command.

The full list of v3 commands implemented are (these commands are supported in v4 and v5):

	capture-mode
	command-list
	conf-get
	dump-counters
	iface-list
	iface-stat
	reload-rules
	register-tenant
	register-tenant-handler
	reload-tenant
	running-mode
	shutdown
	unregister-tenant
	unregister-tenant-handler
	uptime
	version

The full list of v4 commands implemented are:

	add-hostbit
	remove-hostbit
	list-hostbit
	reopen-log-files
	memcap-set
	memcap-show
	memcap-list
	ruleset-reload-rules
	ruleset-reload-nonblocking
	ruleset-reload-time
	ruleset-stats
	ruleset-failed-rules

The full list of v5 commands implemented are:

If you have any suggestion or comment, please feel free to open an issue on
this GitHub page

By Kaushik Shanadi
*/
package client

// V5
// "dataset-add"

// PCAP (TBD On what version)
// "pcap-file" "pcap-file-continuous"
