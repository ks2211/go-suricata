/*

Package client is a Go client library for interacting with suricata via unix socket.

The type-safe library implements most (if not all) commands for suricata v3, v4, and v5 along with the pcap command.

The full list of v3 commands implemented are (these commands are supported in v4 and v5):

	capture-mode
	command-list
	conf-get
	dump-counters (TODO - Response fix)
	iface-list
	iface-stat
	reload-rules
	register-tenant (TODO)
	register-tenant-handler (TODO)
	reload-tenant (TODO)
	running-mode
	shutdown
	unregister-tenant (TODO)
	unregister-tenant-handler (TODO)
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

The full list of pcap commands implemented are:

	pcap-file
	pcap-file-continuous
	pcap-file-number
	pcap-file-list
	pcap-last-processed
	pcap-interrupt
	pcap-current

The full list of v5 commands implemented are:

	iface-bypassed-stat (TODO)
	ebpf-bypassed-stat (TODO)
	dataset-add (TODO)

If you have any suggestion or comment, please feel free to open an issue on
this GitHub page

By Kaushik Shanadi
*/
package client
