package client

import "context"

// TODO
// 	pcap-file, pcap-file-continuous, pcap-file-number, pcap-file-list, pcap-last-processed, pcap-interrupt, pcap-current

const (
	pcapFile           string = "pcap-file"
	pcapFileContinuous string = "pcap-file-continuous"
	pcapFileNumber     string = "pcap-file-number"
	pcapFileList       string = "pcap-file-list"
	pcapLastProcessed  string = "pcap-last-processed"
	pcapInterrupt      string = "pcap-interrupt"
	pcapCurrent        string = "pcap-current"
)

// PcapFileRequest holds the pcap file name and output directory
type PcapFileRequest struct {
	PcapFile        string `json:"filename"`
	OutputDirectory string `json:"output-dir"`
}

// PcapFileResponse is response from "pcap-file" and "pcap-file-continuous"
type PcapFileResponse StringResponse

// String is a helper method to convert a go type into string
func (p PcapFileResponse) String() string {
	return string(p)
}

// PcapFileNumberResponse is response from "pcap-file-number"
type PcapFileNumberResponse int

// PcapLastProcessed is response from "pcap-last-processed"
type PcapLastProcessed int

// PcapInterruptResponse is response from "pcap-interrupt"
type PcapInterruptResponse StringResponse

// String is a helper method to convert a go type into string
func (p PcapInterruptResponse) String() string {
	return string(p)
}

// PcapCurrentResponse is response from "pcap-current"
type PcapCurrentResponse StringResponse

// String is a helper method to convert a go type into string
func (p PcapCurrentResponse) String() string {
	return string(p)
}

// PcapFileListResponse holds the count and files
type PcapFileListResponse struct {
	Count int      `json:"count"`
	Files []string `json:"files"`
}

// PcapFileCommand sends a pcap file and output directory to be processed
func (s *Socket) PcapFileCommand(ctx context.Context, req PcapFileRequest) (string, error) {
	pcapFileResp := new(PcapFileResponse)
	err := s.DoCommand(ctx, pcapFile, req, pcapFileResp)
	return pcapFileResp.String(), err
}

// PcapFileContinuousCommand sends a pcap file and output directory to be processed
func (s *Socket) PcapFileContinuousCommand(ctx context.Context, req PcapFileRequest) (string, error) {
	pcapFileResp := new(PcapFileResponse)
	err := s.DoCommand(ctx, pcapFileContinuous, req, pcapFileResp)
	return pcapFileResp.String(), err
}

// PcapFileNumberCommand gets number of pcap files being processed
func (s *Socket) PcapFileNumberCommand(ctx context.Context) (int, error) {
	pcapFileNumberResp := new(PcapFileNumberResponse)
	err := s.DoCommand(ctx, pcapFileNumber, nil, pcapFileNumberResp)
	return int(*pcapFileNumberResp), err
}

// PcapLastProcessedCommand gets last pcap procssed
func (s *Socket) PcapLastProcessedCommand(ctx context.Context) (int, error) {
	pcapLastProcessedResp := new(PcapLastProcessed)
	err := s.DoCommand(ctx, pcapLastProcessed, nil, pcapLastProcessedResp)
	return int(*pcapLastProcessedResp), err
}

// PcapInterruptCommand does an interrupt on pcap processing
func (s *Socket) PcapFileListCommand(ctx context.Context) (PcapFileListResponse, error) {
	pcapFileListResp := PcapFileListResponse{}
	err := s.DoCommand(ctx, pcapFileList, nil, &pcapFileListResp)
	return pcapFileListResp, err
}

// PcapInterruptCommand does an interrupt on pcap processing
func (s *Socket) PcapInterruptCommand(ctx context.Context) (string, error) {
	pcapInterruptResp := new(PcapInterruptResponse)
	err := s.DoCommand(ctx, pcapInterrupt, nil, pcapInterruptResp)
	return pcapInterruptResp.String(), err
}

// PcapCurrentCommand gets current pcap being processed
func (s *Socket) PcapCurrentCommand(ctx context.Context) (string, error) {
	pcapCurrentResp := new(PcapCurrentResponse)
	err := s.DoCommand(ctx, pcapCurrent, nil, pcapCurrentResp)
	return pcapCurrentResp.String(), err
}
