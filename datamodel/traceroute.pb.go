// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/datamodel/traceroute.proto
// DO NOT EDIT!

package datamodel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TracerouteMeasurement struct {
	Staleness    int64  `protobuf:"varint,1,opt,name=staleness" json:"staleness,omitempty"`
	Dst          uint32 `protobuf:"varint,3,opt,name=dst" json:"dst,omitempty"`
	Confidence   string `protobuf:"bytes,4,opt,name=confidence" json:"confidence,omitempty"`
	Dport        string `protobuf:"bytes,5,opt,name=dport" json:"dport,omitempty"`
	FirstHop     string `protobuf:"bytes,6,opt,name=first_hop" json:"first_hop,omitempty"`
	GapLimit     string `protobuf:"bytes,7,opt,name=gap_limit" json:"gap_limit,omitempty"`
	GapAction    string `protobuf:"bytes,8,opt,name=gap_action" json:"gap_action,omitempty"`
	MaxTtl       string `protobuf:"bytes,9,opt,name=max_ttl" json:"max_ttl,omitempty"`
	PathDiscov   bool   `protobuf:"varint,10,opt,name=path_discov" json:"path_discov,omitempty"`
	Loops        string `protobuf:"bytes,11,opt,name=loops" json:"loops,omitempty"`
	LoopAction   string `protobuf:"bytes,12,opt,name=loop_action" json:"loop_action,omitempty"`
	Payload      string `protobuf:"bytes,13,opt,name=payload" json:"payload,omitempty"`
	Method       string `protobuf:"bytes,14,opt,name=method" json:"method,omitempty"`
	Attempts     string `protobuf:"bytes,15,opt,name=attempts" json:"attempts,omitempty"`
	SendAll      bool   `protobuf:"varint,16,opt,name=send_all" json:"send_all,omitempty"`
	Sport        string `protobuf:"bytes,17,opt,name=sport" json:"sport,omitempty"`
	Src          uint32 `protobuf:"varint,18,opt,name=src" json:"src,omitempty"`
	Tos          string `protobuf:"bytes,19,opt,name=tos" json:"tos,omitempty"`
	TimeExceeded bool   `protobuf:"varint,20,opt,name=time_exceeded" json:"time_exceeded,omitempty"`
	UserId       string `protobuf:"bytes,21,opt,name=user_id" json:"user_id,omitempty"`
	Wait         string `protobuf:"bytes,22,opt,name=wait" json:"wait,omitempty"`
	WaitProbe    string `protobuf:"bytes,23,opt,name=wait_probe" json:"wait_probe,omitempty"`
	GssEntry     string `protobuf:"bytes,24,opt,name=gss_entry" json:"gss_entry,omitempty"`
	LssName      string `protobuf:"bytes,25,opt,name=lss_name" json:"lss_name,omitempty"`
	Timeout      int64  `protobuf:"varint,26,opt,name=timeout" json:"timeout,omitempty"`
	CheckCache   bool   `protobuf:"varint,27,opt,name=check_cache" json:"check_cache,omitempty"`
	CheckDb      bool   `protobuf:"varint,28,opt,name=check_db" json:"check_db,omitempty"`
}

func (m *TracerouteMeasurement) Reset()                    { *m = TracerouteMeasurement{} }
func (m *TracerouteMeasurement) String() string            { return proto.CompactTextString(m) }
func (*TracerouteMeasurement) ProtoMessage()               {}
func (*TracerouteMeasurement) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type TracerouteArg struct {
	Traceroutes []*TracerouteMeasurement `protobuf:"bytes,1,rep,name=traceroutes" json:"traceroutes,omitempty"`
}

func (m *TracerouteArg) Reset()                    { *m = TracerouteArg{} }
func (m *TracerouteArg) String() string            { return proto.CompactTextString(m) }
func (*TracerouteArg) ProtoMessage()               {}
func (*TracerouteArg) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *TracerouteArg) GetTraceroutes() []*TracerouteMeasurement {
	if m != nil {
		return m.Traceroutes
	}
	return nil
}

type TracerouteArgResp struct {
	Traceroutes []*Traceroute `protobuf:"bytes,1,rep,name=traceroutes" json:"traceroutes,omitempty"`
}

func (m *TracerouteArgResp) Reset()                    { *m = TracerouteArgResp{} }
func (m *TracerouteArgResp) String() string            { return proto.CompactTextString(m) }
func (*TracerouteArgResp) ProtoMessage()               {}
func (*TracerouteArgResp) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *TracerouteArgResp) GetTraceroutes() []*Traceroute {
	if m != nil {
		return m.Traceroutes
	}
	return nil
}

type TracerouteHop struct {
	Addr      uint32 `protobuf:"varint,1,opt,name=addr" json:"addr,omitempty"`
	ProbeTtl  uint32 `protobuf:"varint,2,opt,name=probe_ttl" json:"probe_ttl,omitempty"`
	ProbeId   uint32 `protobuf:"varint,3,opt,name=probe_id" json:"probe_id,omitempty"`
	ProbeSize uint32 `protobuf:"varint,4,opt,name=probe_size" json:"probe_size,omitempty"`
	Rtt       *RTT   `protobuf:"bytes,5,opt,name=rtt" json:"rtt,omitempty"`
	ReplyTtl  uint32 `protobuf:"varint,6,opt,name=reply_ttl" json:"reply_ttl,omitempty"`
	ReplyTos  uint32 `protobuf:"varint,7,opt,name=reply_tos" json:"reply_tos,omitempty"`
	ReplySize uint32 `protobuf:"varint,8,opt,name=reply_size" json:"reply_size,omitempty"`
	ReplyIpid uint32 `protobuf:"varint,9,opt,name=reply_ipid" json:"reply_ipid,omitempty"`
	IcmpType  uint32 `protobuf:"varint,10,opt,name=icmp_type" json:"icmp_type,omitempty"`
	IcmpCode  uint32 `protobuf:"varint,11,opt,name=icmp_code" json:"icmp_code,omitempty"`
	IcmpQTtl  uint32 `protobuf:"varint,12,opt,name=icmp_q_ttl" json:"icmp_q_ttl,omitempty"`
	IcmpQIpl  uint32 `protobuf:"varint,13,opt,name=icmp_q_ipl" json:"icmp_q_ipl,omitempty"`
	IcmpQTos  uint32 `protobuf:"varint,14,opt,name=icmp_q_tos" json:"icmp_q_tos,omitempty"`
}

func (m *TracerouteHop) Reset()                    { *m = TracerouteHop{} }
func (m *TracerouteHop) String() string            { return proto.CompactTextString(m) }
func (*TracerouteHop) ProtoMessage()               {}
func (*TracerouteHop) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *TracerouteHop) GetRtt() *RTT {
	if m != nil {
		return m.Rtt
	}
	return nil
}

type Traceroute struct {
	Type       string           `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	UserId     uint32           `protobuf:"varint,2,opt,name=user_id" json:"user_id,omitempty"`
	Method     string           `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	Src        uint32           `protobuf:"varint,4,opt,name=src" json:"src,omitempty"`
	Dst        uint32           `protobuf:"varint,5,opt,name=dst" json:"dst,omitempty"`
	Sport      uint32           `protobuf:"varint,6,opt,name=sport" json:"sport,omitempty"`
	Dport      uint32           `protobuf:"varint,7,opt,name=dport" json:"dport,omitempty"`
	StopReason string           `protobuf:"bytes,8,opt,name=stop_reason" json:"stop_reason,omitempty"`
	StopData   uint32           `protobuf:"varint,9,opt,name=stop_data" json:"stop_data,omitempty"`
	Start      *TracerouteTime  `protobuf:"bytes,10,opt,name=start" json:"start,omitempty"`
	HopCount   uint32           `protobuf:"varint,11,opt,name=hop_count" json:"hop_count,omitempty"`
	Attempts   uint32           `protobuf:"varint,12,opt,name=attempts" json:"attempts,omitempty"`
	Hoplimit   uint32           `protobuf:"varint,13,opt,name=hoplimit" json:"hoplimit,omitempty"`
	Firsthop   uint32           `protobuf:"varint,14,opt,name=firsthop" json:"firsthop,omitempty"`
	Wait       uint32           `protobuf:"varint,15,opt,name=wait" json:"wait,omitempty"`
	WaitProbe  uint32           `protobuf:"varint,16,opt,name=wait_probe" json:"wait_probe,omitempty"`
	Tos        uint32           `protobuf:"varint,17,opt,name=tos" json:"tos,omitempty"`
	ProbeSize  uint32           `protobuf:"varint,18,opt,name=probe_size" json:"probe_size,omitempty"`
	Hops       []*TracerouteHop `protobuf:"bytes,19,rep,name=hops" json:"hops,omitempty"`
	Error      string           `protobuf:"bytes,20,opt,name=error" json:"error,omitempty"`
	Version    string           `protobuf:"bytes,21,opt,name=version" json:"version,omitempty"`
	GapLimit   uint32           `protobuf:"varint,22,opt,name=gap_limit" json:"gap_limit,omitempty"`
	Id         int64            `protobuf:"varint,23,opt,name=id" json:"id,omitempty"`
}

func (m *Traceroute) Reset()                    { *m = Traceroute{} }
func (m *Traceroute) String() string            { return proto.CompactTextString(m) }
func (*Traceroute) ProtoMessage()               {}
func (*Traceroute) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *Traceroute) GetStart() *TracerouteTime {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Traceroute) GetHops() []*TracerouteHop {
	if m != nil {
		return m.Hops
	}
	return nil
}

type TracerouteTime struct {
	Sec   int64  `protobuf:"varint,1,opt,name=sec" json:"sec,omitempty"`
	Usec  int64  `protobuf:"varint,2,opt,name=usec" json:"usec,omitempty"`
	Ftime string `protobuf:"bytes,3,opt,name=ftime" json:"ftime,omitempty"`
}

func (m *TracerouteTime) Reset()                    { *m = TracerouteTime{} }
func (m *TracerouteTime) String() string            { return proto.CompactTextString(m) }
func (*TracerouteTime) ProtoMessage()               {}
func (*TracerouteTime) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func init() {
	proto.RegisterType((*TracerouteMeasurement)(nil), "datamodel.TracerouteMeasurement")
	proto.RegisterType((*TracerouteArg)(nil), "datamodel.TracerouteArg")
	proto.RegisterType((*TracerouteArgResp)(nil), "datamodel.TracerouteArgResp")
	proto.RegisterType((*TracerouteHop)(nil), "datamodel.TracerouteHop")
	proto.RegisterType((*Traceroute)(nil), "datamodel.Traceroute")
	proto.RegisterType((*TracerouteTime)(nil), "datamodel.TracerouteTime")
}

func init() {
	proto.RegisterFile("github.com/NEU-SNS/ReverseTraceroute/datamodel/traceroute.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 761 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6e, 0xdb, 0x3a,
	0x10, 0x45, 0x2c, 0xc7, 0xb1, 0x69, 0xcb, 0x89, 0x95, 0xef, 0x7c, 0x26, 0xf9, 0x8b, 0xc0, 0x8b,
	0x8f, 0xa0, 0x40, 0x6d, 0x20, 0x45, 0x17, 0xed, 0x26, 0x68, 0x81, 0x16, 0xdd, 0x34, 0x8b, 0xc4,
	0xdd, 0x74, 0x23, 0xd0, 0xd2, 0x24, 0x16, 0x2a, 0x89, 0x2c, 0x49, 0xa7, 0x71, 0xaf, 0xd3, 0x4b,
	0xf4, 0x5e, 0xbd, 0x40, 0x87, 0x23, 0x53, 0x76, 0x52, 0x6f, 0xba, 0x12, 0xe6, 0x89, 0x7c, 0xc3,
	0x79, 0xf3, 0x66, 0xd8, 0xe5, 0x5d, 0x66, 0xe7, 0x8b, 0xd9, 0x38, 0x91, 0xc5, 0xe4, 0xea, 0xdd,
	0xa7, 0xe7, 0x37, 0x57, 0x37, 0x93, 0x6b, 0xb8, 0x07, 0x6d, 0x60, 0xaa, 0x45, 0x02, 0x5a, 0x2e,
	0x2c, 0x4c, 0x52, 0x61, 0x45, 0x21, 0x53, 0xc8, 0x27, 0xb6, 0x06, 0xc7, 0x4a, 0x4b, 0x2b, 0xa3,
	0x4e, 0xfd, 0xef, 0xe4, 0xd5, 0xdf, 0x72, 0x65, 0xc5, 0x8a, 0x65, 0xf4, 0x2b, 0x60, 0xc3, 0xf5,
	0x99, 0x8f, 0x20, 0xcc, 0x42, 0x43, 0x01, 0xa5, 0x8d, 0x06, 0xac, 0x63, 0xac, 0xc8, 0xa1, 0x04,
	0x63, 0xf8, 0xce, 0xd9, 0xce, 0x79, 0x10, 0x75, 0x59, 0x90, 0x1a, 0xcb, 0x03, 0x0c, 0xc2, 0x28,
	0x62, 0x2c, 0x91, 0xe5, 0x6d, 0x96, 0x42, 0x99, 0x00, 0x6f, 0x22, 0xd6, 0x89, 0x42, 0xb6, 0x9b,
	0x2a, 0xa9, 0x2d, 0xdf, 0xa5, 0x10, 0x29, 0x6e, 0x33, 0x6d, 0x6c, 0x3c, 0x97, 0x8a, 0xb7, 0x3c,
	0x74, 0x27, 0x54, 0x9c, 0x67, 0x45, 0x66, 0xf9, 0x1e, 0x41, 0x48, 0xe4, 0x20, 0x91, 0xd8, 0x4c,
	0x96, 0xbc, 0x4d, 0xd8, 0x3e, 0xdb, 0x2b, 0xc4, 0x43, 0x6c, 0x6d, 0xce, 0x3b, 0x04, 0x1c, 0xb2,
	0xae, 0x12, 0x76, 0x1e, 0xa7, 0x99, 0x49, 0xe4, 0x3d, 0x67, 0x08, 0xb6, 0x5d, 0xba, 0x5c, 0x4a,
	0x65, 0x78, 0xd7, 0x9f, 0x71, 0xa1, 0x67, 0xea, 0x79, 0x26, 0x25, 0x96, 0xb9, 0x14, 0x29, 0x0f,
	0x09, 0xe8, 0xb3, 0x56, 0x01, 0x76, 0x2e, 0x53, 0xde, 0xa7, 0xf8, 0x80, 0xb5, 0x85, 0xb5, 0x50,
	0x28, 0x6b, 0xf8, 0xbe, 0x47, 0x0c, 0x94, 0x69, 0x2c, 0xf2, 0x9c, 0x1f, 0xf8, 0x44, 0x86, 0xea,
	0x1a, 0xd0, 0x01, 0xd4, 0xc1, 0xe8, 0x84, 0x47, 0xa4, 0x03, 0x06, 0x56, 0x1a, 0x7e, 0x48, 0x7f,
	0x86, 0x2c, 0x74, 0xe2, 0xc6, 0xf0, 0x90, 0x00, 0xa4, 0x90, 0xf2, 0x7f, 0xe8, 0x3e, 0x3e, 0x62,
	0x61, 0x40, 0xc7, 0x59, 0xca, 0x87, 0x74, 0xae, 0xc7, 0x9a, 0xdf, 0x04, 0x2a, 0x70, 0xe4, 0x15,
	0x70, 0x51, 0x8c, 0x2d, 0x99, 0x01, 0xff, 0xb7, 0x16, 0xca, 0x98, 0x18, 0x3b, 0xa1, 0x97, 0x9c,
	0xfb, 0x77, 0xe5, 0x08, 0x95, 0xa2, 0x00, 0x7e, 0xec, 0x8b, 0x73, 0xe9, 0xb0, 0x75, 0xfc, 0x84,
	0x3a, 0x84, 0x12, 0x24, 0x73, 0x48, 0xbe, 0xc4, 0x89, 0xc0, 0x2f, 0x3f, 0xa5, 0xec, 0x78, 0xaf,
	0x02, 0xd3, 0x19, 0xff, 0xcf, 0x21, 0xa3, 0xf7, 0x2c, 0x5c, 0x37, 0xfd, 0x8d, 0xbe, 0x8b, 0x5e,
	0xb2, 0xee, 0xda, 0x60, 0xae, 0xdd, 0xc1, 0x79, 0xf7, 0xe2, 0x6c, 0x5c, 0x5b, 0x66, 0xbc, 0xd5,
	0x23, 0xa3, 0x4b, 0x36, 0x78, 0xc4, 0x73, 0x0d, 0x46, 0x45, 0xcf, 0xb6, 0x71, 0x0d, 0xb7, 0x72,
	0x8d, 0x7e, 0x34, 0x36, 0x5f, 0xf2, 0x41, 0x2a, 0xa7, 0x8c, 0x48, 0x53, 0x4d, 0x8e, 0x0b, 0x9d,
	0x0a, 0x24, 0x0a, 0x39, 0xa1, 0x41, 0x10, 0x56, 0x53, 0x41, 0x28, 0x66, 0xed, 0xc4, 0x0a, 0x31,
	0xd9, 0xf7, 0xca, 0x89, 0x61, 0x74, 0xca, 0x02, 0x6d, 0x2b, 0x1f, 0x76, 0x2f, 0xfa, 0x1b, 0xc9,
	0xaf, 0xa7, 0x53, 0xc7, 0xaa, 0x41, 0xe5, 0x4b, 0x62, 0x6d, 0xf9, 0x44, 0x2b, 0x08, 0x7b, 0xb9,
	0xe7, 0x69, 0x2b, 0x88, 0x68, 0xdb, 0x8f, 0xb1, 0x4c, 0x61, 0xfa, 0x8e, 0xbf, 0x9a, 0x25, 0x85,
	0x8a, 0xed, 0x52, 0x01, 0x19, 0x73, 0x0d, 0x25, 0x98, 0x92, 0xcc, 0x49, 0x37, 0x09, 0xfa, 0x4a,
	0x49, 0x7b, 0x4f, 0xb0, 0x4c, 0xe5, 0x64, 0xcf, 0x47, 0xe7, 0xf0, 0x25, 0xce, 0xa2, 0xe1, 0xe8,
	0x67, 0xc0, 0xd8, 0x5a, 0x25, 0x27, 0x11, 0xe5, 0xda, 0xf1, 0x1e, 0xf0, 0xde, 0xaa, 0x04, 0x5a,
	0x1b, 0x3c, 0xd8, 0x74, 0x6b, 0xd3, 0xbb, 0xd5, 0x8d, 0xf0, 0x2e, 0x05, 0xb5, 0xad, 0x5b, 0x3e,
	0xac, 0xa6, 0xb7, 0xaa, 0x1f, 0xbd, 0x64, 0x2c, 0x8e, 0x93, 0xc6, 0x86, 0xd7, 0x83, 0x49, 0x5b,
	0x01, 0x41, 0x27, 0xe8, 0xaa, 0xfe, 0x73, 0x64, 0xb1, 0x02, 0xaf, 0x31, 0x12, 0xfb, 0x78, 0x6b,
	0xa7, 0xa7, 0x68, 0x53, 0x77, 0x19, 0x37, 0x01, 0xaa, 0xb2, 0x28, 0xed, 0x4a, 0x96, 0xcd, 0xe9,
	0xeb, 0x79, 0x04, 0x0f, 0x55, 0x0b, 0x22, 0xf4, 0x08, 0xad, 0x11, 0xb7, 0x45, 0x48, 0x90, 0x7a,
	0x7c, 0xf6, 0xbd, 0x64, 0x1b, 0xe3, 0x73, 0xb0, 0x39, 0x95, 0x83, 0x2d, 0x06, 0xa9, 0xc6, 0xf6,
	0x7f, 0xd6, 0x9c, 0xbb, 0xd5, 0x71, 0x48, 0xf6, 0xe4, 0x5b, 0x1f, 0xed, 0xfc, 0x88, 0xa2, 0x80,
	0xd6, 0x52, 0xd3, 0x24, 0x93, 0xda, 0x6e, 0xaf, 0xba, 0xfd, 0x32, 0xfc, 0x73, 0xa1, 0x1d, 0x11,
	0x35, 0x63, 0x0d, 0x6c, 0x86, 0x1b, 0xe3, 0x60, 0xf4, 0x9a, 0xf5, 0x9f, 0x88, 0xe0, 0xda, 0x01,
	0xc9, 0x6a, 0xa3, 0x62, 0x21, 0x0b, 0x17, 0x35, 0x28, 0xc2, 0x5c, 0xb7, 0x6e, 0x9e, 0xab, 0xc6,
	0xbd, 0xed, 0x7e, 0x5e, 0xef, 0xf8, 0x59, 0x8b, 0xf6, 0xf5, 0x8b, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x85, 0x45, 0x0b, 0x76, 0x38, 0x06, 0x00, 0x00,
}
