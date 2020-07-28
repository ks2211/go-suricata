package client

import "context"

const (
	dumpCounters string = "dump-counters"
)

// DumpCountersCommand performs the dump-counters command
func (s *Socket) DumpCountersCommand(ctx context.Context) (*DumpCountersResponse, error) {
	// create and marshal the "dump-counters" socket message with no args
	counters := &DumpCountersResponse{}
	err := s.DoCommand(ctx, dumpCounters, nil, counters)
	return counters, err
}

// DumpCountersResponse is response from "dump-counters"
// TODO the iface map below is messed up
type DumpCountersResponse struct {
	AppLayer struct {
		Flow struct {
			DcerpcTCP int `json:"dcerpc_tcp"`
			DcerpcUDP int `json:"dcerpc_udp"`
			DNSTCP    int `json:"dns_tcp"`
			DNSUDP    int `json:"dns_udp"`
			FailedTCP int `json:"failed_tcp"`
			FailedUDP int `json:"failed_udp"`
			Ftp       int `json:"ftp"`
			HTTP      int `json:"http"`
			Imap      int `json:"imap"`
			Msn       int `json:"msn"`
			Smb       int `json:"smb"`
			SMTP      int `json:"smtp"`
			SSH       int `json:"ssh"`
			TLS       int `json:"tls"`
		} `json:"flow"`
		Tx struct {
			DNSTCP int `json:"dns_tcp"`
			DNSUDP int `json:"dns_udp"`
			HTTP   int `json:"http"`
			SMTP   int `json:"smtp"`
			TLS    int `json:"tls"`
		} `json:"tx"`
	} `json:"app_layer"`
	Capture struct {
		KernelDrops   int `json:"kernel_drops"`
		KernelPackets int `json:"kernel_packets"`
	} `json:"capture"`
	Decoder struct {
		AvgPktSize int `json:"avg_pkt_size"`
		Bytes      int `json:"bytes"`
		Dce        struct {
			PktTooSmall int `json:"pkt_too_small"`
		} `json:"dce"`
		Erspan   int `json:"erspan"`
		Ethernet int `json:"ethernet"`
		Gre      int `json:"gre"`
		Icmpv4   int `json:"icmpv4"`
		Icmpv6   int `json:"icmpv6"`
		Invalid  int `json:"invalid"`
		Ipraw    struct {
			InvalidIPVersion int `json:"invalid_ip_version"`
		} `json:"ipraw"`
		Ipv4       int `json:"ipv4"`
		Ipv4InIpv6 int `json:"ipv4_in_ipv6"`
		Ipv6       int `json:"ipv6"`
		Ipv6InIpv6 int `json:"ipv6_in_ipv6"`
		Ltnull     struct {
			PktTooSmall     int `json:"pkt_too_small"`
			UnsupportedType int `json:"unsupported_type"`
		} `json:"ltnull"`
		MaxPktSize int `json:"max_pkt_size"`
		Mpls       int `json:"mpls"`
		Null       int `json:"null"`
		Pkts       int `json:"pkts"`
		Ppp        int `json:"ppp"`
		Pppoe      int `json:"pppoe"`
		Raw        int `json:"raw"`
		Sctp       int `json:"sctp"`
		Sll        int `json:"sll"`
		TCP        int `json:"tcp"`
		Teredo     int `json:"teredo"`
		UDP        int `json:"udp"`
		Vlan       int `json:"vlan"`
		VlanQinq   int `json:"vlan_qinq"`
	} `json:"decoder"`
	Defrag struct {
		Ipv4 struct {
			Fragments   int `json:"fragments"`
			Reassembled int `json:"reassembled"`
			Timeouts    int `json:"timeouts"`
		} `json:"ipv4"`
		Ipv6 struct {
			Fragments   int `json:"fragments"`
			Reassembled int `json:"reassembled"`
			Timeouts    int `json:"timeouts"`
		} `json:"ipv6"`
		MaxFragHits int `json:"max_frag_hits"`
	} `json:"defrag"`
	Detect struct {
		Alert int `json:"alert"`
	} `json:"detect"`
	DNS struct {
		MemcapGlobal int `json:"memcap_global"`
		MemcapState  int `json:"memcap_state"`
		Memuse       int `json:"memuse"`
	} `json:"dns"`
	Flow struct {
		EmergModeEntered int `json:"emerg_mode_entered"`
		EmergModeOver    int `json:"emerg_mode_over"`
		Memcap           int `json:"memcap"`
		Memuse           int `json:"memuse"`
		Spare            int `json:"spare"`
		TCPReuse         int `json:"tcp_reuse"`
	} `json:"flow"`
	FlowMgr struct {
		BypassedPruned    int `json:"bypassed_pruned"`
		ClosedPruned      int `json:"closed_pruned"`
		EstPruned         int `json:"est_pruned"`
		FlowsChecked      int `json:"flows_checked"`
		FlowsNotimeout    int `json:"flows_notimeout"`
		FlowsRemoved      int `json:"flows_removed"`
		FlowsTimeout      int `json:"flows_timeout"`
		FlowsTimeoutInuse int `json:"flows_timeout_inuse"`
		NewPruned         int `json:"new_pruned"`
		RowsBusy          int `json:"rows_busy"`
		RowsChecked       int `json:"rows_checked"`
		RowsEmpty         int `json:"rows_empty"`
		RowsMaxlen        int `json:"rows_maxlen"`
		RowsSkipped       int `json:"rows_skipped"`
	} `json:"flow_mgr"`
	HTTP struct {
		Memcap int `json:"memcap"`
		Memuse int `json:"memuse"`
	} `json:"http"`
	TCP struct {
		InvalidChecksum    int `json:"invalid_checksum"`
		Memuse             int `json:"memuse"`
		NoFlow             int `json:"no_flow"`
		Pseudo             int `json:"pseudo"`
		PseudoFailed       int `json:"pseudo_failed"`
		ReassemblyGap      int `json:"reassembly_gap"`
		ReassemblyMemuse   int `json:"reassembly_memuse"`
		Rst                int `json:"rst"`
		SegmentMemcapDrop  int `json:"segment_memcap_drop"`
		Sessions           int `json:"sessions"`
		SsnMemcapDrop      int `json:"ssn_memcap_drop"`
		StreamDepthReached int `json:"stream_depth_reached"`
		Syn                int `json:"syn"`
		Synack             int `json:"synack"`
	} `json:"tcp"`
	Threads struct {
		FM01 struct {
			Flow struct {
				EmergModeEntered int `json:"emerg_mode_entered"`
				EmergModeOver    int `json:"emerg_mode_over"`
				Spare            int `json:"spare"`
				TCPReuse         int `json:"tcp_reuse"`
			} `json:"flow"`
			FlowMgr struct {
				BypassedPruned    int `json:"bypassed_pruned"`
				ClosedPruned      int `json:"closed_pruned"`
				EstPruned         int `json:"est_pruned"`
				FlowsChecked      int `json:"flows_checked"`
				FlowsNotimeout    int `json:"flows_notimeout"`
				FlowsRemoved      int `json:"flows_removed"`
				FlowsTimeout      int `json:"flows_timeout"`
				FlowsTimeoutInuse int `json:"flows_timeout_inuse"`
				NewPruned         int `json:"new_pruned"`
				RowsBusy          int `json:"rows_busy"`
				RowsChecked       int `json:"rows_checked"`
				RowsEmpty         int `json:"rows_empty"`
				RowsMaxlen        int `json:"rows_maxlen"`
				RowsSkipped       int `json:"rows_skipped"`
			} `json:"flow_mgr"`
		} `json:"FM#01"`
		Global struct {
			DNS struct {
				MemcapGlobal int `json:"memcap_global"`
				MemcapState  int `json:"memcap_state"`
				Memuse       int `json:"memuse"`
			} `json:"dns"`
			Flow struct {
				Memuse int `json:"memuse"`
			} `json:"flow"`
			HTTP struct {
				Memcap int `json:"memcap"`
				Memuse int `json:"memuse"`
			} `json:"http"`
			TCP struct {
				Memuse           int `json:"memuse"`
				ReassemblyMemuse int `json:"reassembly_memuse"`
			} `json:"tcp"`
		} `json:"Global"`
		Ifaces map[string]struct {
			AppLayer struct {
				Flow struct {
					DcerpcTCP int `json:"dcerpc_tcp"`
					DcerpcUDP int `json:"dcerpc_udp"`
					DNSTCP    int `json:"dns_tcp"`
					DNSUDP    int `json:"dns_udp"`
					FailedTCP int `json:"failed_tcp"`
					FailedUDP int `json:"failed_udp"`
					Ftp       int `json:"ftp"`
					HTTP      int `json:"http"`
					Imap      int `json:"imap"`
					Msn       int `json:"msn"`
					Smb       int `json:"smb"`
					SMTP      int `json:"smtp"`
					SSH       int `json:"ssh"`
					TLS       int `json:"tls"`
				} `json:"flow"`
				Tx struct {
					DNSTCP int `json:"dns_tcp"`
					DNSUDP int `json:"dns_udp"`
					HTTP   int `json:"http"`
					SMTP   int `json:"smtp"`
					TLS    int `json:"tls"`
				} `json:"tx"`
			} `json:"app_layer"`
			Capture struct {
				KernelDrops   int `json:"kernel_drops"`
				KernelPackets int `json:"kernel_packets"`
			} `json:"capture"`
			Decoder struct {
				AvgPktSize int `json:"avg_pkt_size"`
				Bytes      int `json:"bytes"`
				Dce        struct {
					PktTooSmall int `json:"pkt_too_small"`
				} `json:"dce"`
				Erspan   int `json:"erspan"`
				Ethernet int `json:"ethernet"`
				Gre      int `json:"gre"`
				Icmpv4   int `json:"icmpv4"`
				Icmpv6   int `json:"icmpv6"`
				Invalid  int `json:"invalid"`
				Ipraw    struct {
					InvalidIPVersion int `json:"invalid_ip_version"`
				} `json:"ipraw"`
				Ipv4       int `json:"ipv4"`
				Ipv4InIpv6 int `json:"ipv4_in_ipv6"`
				Ipv6       int `json:"ipv6"`
				Ipv6InIpv6 int `json:"ipv6_in_ipv6"`
				Ltnull     struct {
					PktTooSmall     int `json:"pkt_too_small"`
					UnsupportedType int `json:"unsupported_type"`
				} `json:"ltnull"`
				MaxPktSize int `json:"max_pkt_size"`
				Mpls       int `json:"mpls"`
				Null       int `json:"null"`
				Pkts       int `json:"pkts"`
				Ppp        int `json:"ppp"`
				Pppoe      int `json:"pppoe"`
				Raw        int `json:"raw"`
				Sctp       int `json:"sctp"`
				Sll        int `json:"sll"`
				TCP        int `json:"tcp"`
				Teredo     int `json:"teredo"`
				UDP        int `json:"udp"`
				Vlan       int `json:"vlan"`
				VlanQinq   int `json:"vlan_qinq"`
			} `json:"decoder"`
			Defrag struct {
				Ipv4 struct {
					Fragments   int `json:"fragments"`
					Reassembled int `json:"reassembled"`
					Timeouts    int `json:"timeouts"`
				} `json:"ipv4"`
				Ipv6 struct {
					Fragments   int `json:"fragments"`
					Reassembled int `json:"reassembled"`
					Timeouts    int `json:"timeouts"`
				} `json:"ipv6"`
				MaxFragHits int `json:"max_frag_hits"`
			} `json:"defrag"`
			Detect struct {
				Alert int `json:"alert"`
			} `json:"detect"`
			Flow struct {
				Memcap int `json:"memcap"`
			} `json:"flow"`
			TCP struct {
				InvalidChecksum    int `json:"invalid_checksum"`
				NoFlow             int `json:"no_flow"`
				Pseudo             int `json:"pseudo"`
				PseudoFailed       int `json:"pseudo_failed"`
				ReassemblyGap      int `json:"reassembly_gap"`
				Rst                int `json:"rst"`
				SegmentMemcapDrop  int `json:"segment_memcap_drop"`
				Sessions           int `json:"sessions"`
				SsnMemcapDrop      int `json:"ssn_memcap_drop"`
				StreamDepthReached int `json:"stream_depth_reached"`
				Syn                int `json:"syn"`
				Synack             int `json:"synack"`
			} `json:"tcp"`
		}
	} `json:"threads"`
	Uptime int `json:"uptime"`
}
