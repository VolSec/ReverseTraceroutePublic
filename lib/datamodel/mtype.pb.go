// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/mtype.proto
// DO NOT EDIT!

/*
Package datamodel is a generated protocol buffer package.

It is generated from these files:
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/mtype.proto
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/ping.proto
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/returnt.proto
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/service.proto
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/stats.proto
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/time.proto
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/traceroute.proto
	github.com/NEU-SNS/ReverseTraceroute/lib/datamodel/vantagepoint.proto

It has these top-level messages:
*/
package datamodel

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type MType int32

const (
	MType_PING       MType = 0
	MType_STATS      MType = 1
	MType_TRACEROUTE MType = 2
)

var MType_name = map[int32]string{
	0: "PING",
	1: "STATS",
	2: "TRACEROUTE",
}
var MType_value = map[string]int32{
	"PING":       0,
	"STATS":      1,
	"TRACEROUTE": 2,
}

func (x MType) String() string {
	return proto.EnumName(MType_name, int32(x))
}

func init() {
	proto.RegisterEnum("datamodel.MType", MType_name, MType_value)
}
