package client

import (
	"testing"
)

func TestIFaceListCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedOutput := []string{
		"wlan0",
	}

	ifaceList, err := s.IFaceListCommand()
	if err != nil {
		t.Fatalf("error doing iface-list command %v", err)
	}
	if ifaceList.Count != len(expectedOutput) {
		t.Fatalf("expected iface-list count %v got %v", len(expectedOutput), ifaceList.Count)
	}
	if expectedOutput[0] != ifaceList.Ifaces[0] {
		t.Fatalf("expected iface-list output %v got %v", expectedOutput[0], ifaceList.Ifaces[0])
	}
}

func TestIFaceStatCommand(t *testing.T) {
	s, err := CreateSocket(DefaultSocketPathV3)
	if err != nil {
		t.Fatalf("error create socket client %v", err)
	}
	defer s.Close()

	expectedOutput := IFaceStatResponse{
		Drop:             0,
		InvalidChecksums: 1,
		Pkts:             -1,
	}

	ifaceStat, err := s.IFaceStatCommand("wlan0")
	if err != nil {
		t.Fatalf("error doing iface-stat command %v", err)
	}
	if ifaceStat.Drop != expectedOutput.Drop {
		t.Fatalf("expected iface-stat drop %v got %v", expectedOutput.Drop, ifaceStat.Drop)
	}
	if ifaceStat.InvalidChecksums != expectedOutput.InvalidChecksums {
		t.Fatalf("expected iface-stat invalid checksums %v got %v", expectedOutput.InvalidChecksums, ifaceStat.InvalidChecksums)
	}
}
