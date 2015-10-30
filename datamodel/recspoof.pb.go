/*
Copyright (c) 2015, Northeastern University
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:
     * Redistributions of source code must retain the above copyright
       notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above copyright
       notice, this list of conditions and the following disclaimer in the
       documentation and/or other materials provided with the distribution.
     * Neither the name of the Northeastern University nor the
       names of its contributors may be used to endorse or promote products
       derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL Northeastern University BE LIABLE FOR ANY
 DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
// Code generated by protoc-gen-go.
// source: github.com/NEU-SNS/ReverseTraceroute/datamodel/recspoof.proto
// DO NOT EDIT!

package datamodel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TSType int32

const (
	TSType_TSOnly    TSType = 0
	TSType_TSAndAddr TSType = 1
	TSType_TSPreSpec TSType = 3
)

var TSType_name = map[int32]string{
	0: "TSOnly",
	1: "TSAndAddr",
	3: "TSPreSpec",
}
var TSType_value = map[string]int32{
	"TSOnly":    0,
	"TSAndAddr": 1,
	"TSPreSpec": 3,
}

func (x TSType) String() string {
	return proto.EnumName(TSType_name, int32(x))
}

type RecSpoof struct {
	Spoofs []*Spoof `protobuf:"bytes,1,rep,name=spoofs" json:"spoofs,omitempty"`
}

func (m *RecSpoof) Reset()         { *m = RecSpoof{} }
func (m *RecSpoof) String() string { return proto.CompactTextString(m) }
func (*RecSpoof) ProtoMessage()    {}

func (m *RecSpoof) GetSpoofs() []*Spoof {
	if m != nil {
		return m.Spoofs
	}
	return nil
}

type Spoof struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Id uint32 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *Spoof) Reset()         { *m = Spoof{} }
func (m *Spoof) String() string { return proto.CompactTextString(m) }
func (*Spoof) ProtoMessage()    {}

type SpoofedProbes struct {
	Probes []*Probe `protobuf:"bytes,1,rep,name=probes" json:"probes,omitempty"`
}

func (m *SpoofedProbes) Reset()         { *m = SpoofedProbes{} }
func (m *SpoofedProbes) String() string { return proto.CompactTextString(m) }
func (*SpoofedProbes) ProtoMessage()    {}

func (m *SpoofedProbes) GetProbes() []*Probe {
	if m != nil {
		return m.Probes
	}
	return nil
}

type SpoofedProbesResponse struct {
}

func (m *SpoofedProbesResponse) Reset()         { *m = SpoofedProbesResponse{} }
func (m *SpoofedProbesResponse) String() string { return proto.CompactTextString(m) }
func (*SpoofedProbesResponse) ProtoMessage()    {}

type Probe struct {
	SpooferIp uint32       `protobuf:"varint,1,opt,name=spoofer_ip" json:"spoofer_ip,omitempty"`
	ProbeId   uint32       `protobuf:"varint,2,opt,name=probe_id" json:"probe_id,omitempty"`
	Src       uint32       `protobuf:"varint,4,opt,name=src" json:"src,omitempty"`
	Dst       uint32       `protobuf:"varint,5,opt,name=dst" json:"dst,omitempty"`
	Id        uint32       `protobuf:"varint,6,opt,name=id" json:"id,omitempty"`
	SeqNum    uint32       `protobuf:"varint,7,opt,name=seq_num" json:"seq_num,omitempty"`
	RR        *RecordRoute `protobuf:"bytes,8,opt,name=r_r" json:"r_r,omitempty"`
	Ts        *TimeStamp   `protobuf:"bytes,9,opt,name=ts" json:"ts,omitempty"`
	SenderIp  string       `protobuf:"bytes,10,opt,name=sender_ip" json:"sender_ip,omitempty"`
}

func (m *Probe) Reset()         { *m = Probe{} }
func (m *Probe) String() string { return proto.CompactTextString(m) }
func (*Probe) ProtoMessage()    {}

func (m *Probe) GetRR() *RecordRoute {
	if m != nil {
		return m.RR
	}
	return nil
}

func (m *Probe) GetTs() *TimeStamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

type RecordRoute struct {
	Hops []uint32 `protobuf:"varint,1,rep,packed,name=hops" json:"hops,omitempty"`
}

func (m *RecordRoute) Reset()         { *m = RecordRoute{} }
func (m *RecordRoute) String() string { return proto.CompactTextString(m) }
func (*RecordRoute) ProtoMessage()    {}

type TimeStamp struct {
	Type   TSType   `protobuf:"varint,1,opt,name=type,enum=datamodel.TSType" json:"type,omitempty"`
	Stamps []*Stamp `protobuf:"bytes,2,rep,name=stamps" json:"stamps,omitempty"`
}

func (m *TimeStamp) Reset()         { *m = TimeStamp{} }
func (m *TimeStamp) String() string { return proto.CompactTextString(m) }
func (*TimeStamp) ProtoMessage()    {}

func (m *TimeStamp) GetStamps() []*Stamp {
	if m != nil {
		return m.Stamps
	}
	return nil
}

type Stamp struct {
	Time uint32 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Ip   uint32 `protobuf:"varint,2,opt,name=ip" json:"ip,omitempty"`
}

func (m *Stamp) Reset()         { *m = Stamp{} }
func (m *Stamp) String() string { return proto.CompactTextString(m) }
func (*Stamp) ProtoMessage()    {}

type NotifyRecSpoofResponse struct {
	Error string `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *NotifyRecSpoofResponse) Reset()         { *m = NotifyRecSpoofResponse{} }
func (m *NotifyRecSpoofResponse) String() string { return proto.CompactTextString(m) }
func (*NotifyRecSpoofResponse) ProtoMessage()    {}

type ReceiveSpoofedProbesResponse struct {
}

func (m *ReceiveSpoofedProbesResponse) Reset()         { *m = ReceiveSpoofedProbesResponse{} }
func (m *ReceiveSpoofedProbesResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiveSpoofedProbesResponse) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("datamodel.TSType", TSType_name, TSType_value)
}
