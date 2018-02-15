// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metadata.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Metadata struct {
	ResourceVersion string `protobuf:"bytes,1,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	Namespace       string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{0} }

func (m *Metadata) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *Metadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "v1.Metadata")
}
func (this *Metadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metadata)
	if !ok {
		that2, ok := that.(Metadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ResourceVersion != that1.ResourceVersion {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	return true
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptorMetadata) }

var fileDescriptorMetadata = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x49,
	0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x12,
	0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x73, 0xf5, 0x41, 0x2c, 0x88, 0x8c, 0x52, 0x30, 0x17, 0x87, 0x2f,
	0x54, 0xad, 0x90, 0x26, 0x97, 0x40, 0x51, 0x6a, 0x71, 0x7e, 0x69, 0x51, 0x72, 0x6a, 0x7c, 0x59,
	0x6a, 0x51, 0x71, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x3f, 0x4c, 0x3c,
	0x0c, 0x22, 0x2c, 0x24, 0xc3, 0xc5, 0x99, 0x97, 0x98, 0x9b, 0x5a, 0x5c, 0x90, 0x98, 0x9c, 0x2a,
	0xc1, 0x04, 0x56, 0x83, 0x10, 0x70, 0x62, 0x59, 0xf1, 0x48, 0x8e, 0x31, 0x89, 0x0d, 0x6c, 0x83,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x72, 0xea, 0xdb, 0xe5, 0x8d, 0x00, 0x00, 0x00,
}