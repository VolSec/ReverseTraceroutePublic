// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/time.proto
// DO NOT EDIT!

package datamodel

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Time struct {
	Sec  int64 `protobuf:"varint,1,opt,name=sec" json:"sec,omitempty"`
	Usec int64 `protobuf:"varint,2,opt,name=usec" json:"usec,omitempty"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}

func init() {
}