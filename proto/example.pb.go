// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

/*
Package rrfp is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	Message
	Head
	Body
	ExampleEchoRequest
	ExampleEchoResponse
*/
package rrfp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Hd *Head `protobuf:"bytes,10,opt,name=hd" json:"hd,omitempty"`
	By *Body `protobuf:"bytes,20,opt,name=by" json:"by,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetHd() *Head {
	if m != nil {
		return m.Hd
	}
	return nil
}

func (m *Message) GetBy() *Body {
	if m != nil {
		return m.By
	}
	return nil
}

type Head struct {
	SesionNo string `protobuf:"bytes,10,opt,name=sesionNo" json:"sesionNo,omitempty"`
	UniqueId string `protobuf:"bytes,20,opt,name=uniqueId" json:"uniqueId,omitempty"`
}

func (m *Head) Reset()                    { *m = Head{} }
func (m *Head) String() string            { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()               {}
func (*Head) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Head) GetSesionNo() string {
	if m != nil {
		return m.SesionNo
	}
	return ""
}

func (m *Head) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

type Body struct {
	// Types that are valid to be assigned to MsgType:
	//	*Body_ExampleEchoRequest
	//	*Body_ExampleEchoResponse
	MsgType isBody_MsgType `protobuf_oneof:"msgType"`
}

func (m *Body) Reset()                    { *m = Body{} }
func (m *Body) String() string            { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()               {}
func (*Body) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isBody_MsgType interface {
	isBody_MsgType()
}

type Body_ExampleEchoRequest struct {
	ExampleEchoRequest *ExampleEchoRequest `protobuf:"bytes,10,opt,name=example_echo_request,json=exampleEchoRequest,oneof"`
}
type Body_ExampleEchoResponse struct {
	ExampleEchoResponse *ExampleEchoResponse `protobuf:"bytes,20,opt,name=example_echo_response,json=exampleEchoResponse,oneof"`
}

func (*Body_ExampleEchoRequest) isBody_MsgType()  {}
func (*Body_ExampleEchoResponse) isBody_MsgType() {}

func (m *Body) GetMsgType() isBody_MsgType {
	if m != nil {
		return m.MsgType
	}
	return nil
}

func (m *Body) GetExampleEchoRequest() *ExampleEchoRequest {
	if x, ok := m.GetMsgType().(*Body_ExampleEchoRequest); ok {
		return x.ExampleEchoRequest
	}
	return nil
}

func (m *Body) GetExampleEchoResponse() *ExampleEchoResponse {
	if x, ok := m.GetMsgType().(*Body_ExampleEchoResponse); ok {
		return x.ExampleEchoResponse
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Body) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Body_OneofMarshaler, _Body_OneofUnmarshaler, _Body_OneofSizer, []interface{}{
		(*Body_ExampleEchoRequest)(nil),
		(*Body_ExampleEchoResponse)(nil),
	}
}

func _Body_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Body)
	// msgType
	switch x := m.MsgType.(type) {
	case *Body_ExampleEchoRequest:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExampleEchoRequest); err != nil {
			return err
		}
	case *Body_ExampleEchoResponse:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExampleEchoResponse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Body.MsgType has unexpected type %T", x)
	}
	return nil
}

func _Body_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Body)
	switch tag {
	case 10: // msgType.example_echo_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExampleEchoRequest)
		err := b.DecodeMessage(msg)
		m.MsgType = &Body_ExampleEchoRequest{msg}
		return true, err
	case 20: // msgType.example_echo_response
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExampleEchoResponse)
		err := b.DecodeMessage(msg)
		m.MsgType = &Body_ExampleEchoResponse{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Body_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Body)
	// msgType
	switch x := m.MsgType.(type) {
	case *Body_ExampleEchoRequest:
		s := proto.Size(x.ExampleEchoRequest)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Body_ExampleEchoResponse:
		s := proto.Size(x.ExampleEchoResponse)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ExampleEchoRequest struct {
	Msg string `protobuf:"bytes,10,opt,name=msg" json:"msg,omitempty"`
}

func (m *ExampleEchoRequest) Reset()                    { *m = ExampleEchoRequest{} }
func (m *ExampleEchoRequest) String() string            { return proto.CompactTextString(m) }
func (*ExampleEchoRequest) ProtoMessage()               {}
func (*ExampleEchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExampleEchoRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ExampleEchoResponse struct {
	Msg string `protobuf:"bytes,10,opt,name=msg" json:"msg,omitempty"`
}

func (m *ExampleEchoResponse) Reset()                    { *m = ExampleEchoResponse{} }
func (m *ExampleEchoResponse) String() string            { return proto.CompactTextString(m) }
func (*ExampleEchoResponse) ProtoMessage()               {}
func (*ExampleEchoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExampleEchoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "rrfp.Message")
	proto.RegisterType((*Head)(nil), "rrfp.Head")
	proto.RegisterType((*Body)(nil), "rrfp.Body")
	proto.RegisterType((*ExampleEchoRequest)(nil), "rrfp.ExampleEchoRequest")
	proto.RegisterType((*ExampleEchoResponse)(nil), "rrfp.ExampleEchoResponse")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0xdb, 0x18, 0xac, 0x19, 0x11, 0x64, 0x5a, 0x21, 0xf6, 0x24, 0x7b, 0x50, 0x4f, 0x39,
	0xe8, 0x5d, 0xb0, 0x50, 0x88, 0xe0, 0x1f, 0x58, 0xbc, 0x97, 0xa4, 0x3b, 0x26, 0x05, 0x93, 0xd9,
	0x66, 0x1a, 0x30, 0x6f, 0xe5, 0x23, 0xca, 0x26, 0xa1, 0x48, 0xd3, 0xdb, 0xce, 0x7c, 0x3f, 0xbe,
	0x99, 0x61, 0xe1, 0x82, 0x7e, 0x92, 0xc2, 0x7e, 0x53, 0x64, 0x2b, 0xde, 0x31, 0xfa, 0x55, 0xf5,
	0x65, 0xd5, 0x33, 0x4c, 0xde, 0x48, 0x24, 0xc9, 0x08, 0xe7, 0xe0, 0xe5, 0x26, 0x84, 0x9b, 0xf1,
	0xfd, 0xf9, 0x03, 0x44, 0x8e, 0x46, 0x31, 0x25, 0x46, 0x7b, 0xb9, 0x71, 0x2c, 0x6d, 0xc2, 0xd9,
	0x7f, 0xb6, 0x60, 0xd3, 0x68, 0x2f, 0x6d, 0xd4, 0x13, 0xf8, 0x2e, 0x87, 0x73, 0x38, 0x13, 0x92,
	0x0d, 0x97, 0xef, 0xdc, 0x5a, 0x02, 0xbd, 0xaf, 0x1d, 0xab, 0xcb, 0xcd, 0xb6, 0xa6, 0x17, 0xd3,
	0x5a, 0x02, 0xbd, 0xaf, 0xd5, 0xef, 0x18, 0x7c, 0x27, 0xc3, 0x57, 0x98, 0xf5, 0x2b, 0xae, 0x68,
	0x9d, 0xf3, 0xaa, 0xa2, 0x6d, 0x4d, 0xb2, 0xeb, 0x57, 0x0a, 0xbb, 0xb1, 0xcb, 0x2e, 0xb1, 0x5c,
	0xe7, 0xac, 0x3b, 0x1e, 0x8f, 0x34, 0xd2, 0xa0, 0x8b, 0x1f, 0x70, 0x75, 0x60, 0x13, 0xcb, 0xa5,
	0x50, 0x7f, 0xc5, 0xf5, 0x11, 0x5d, 0x17, 0x88, 0x47, 0x7a, 0x4a, 0xc3, 0xf6, 0x22, 0x80, 0x49,
	0x21, 0xd9, 0x67, 0x63, 0x49, 0xdd, 0x02, 0x0e, 0xf7, 0xc0, 0x4b, 0x38, 0x29, 0x24, 0xeb, 0x6f,
	0x77, 0x4f, 0x75, 0x07, 0xd3, 0x23, 0x03, 0x86, 0xc1, 0xf4, 0xb4, 0xfd, 0x93, 0xc7, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x53, 0x0f, 0xa9, 0x0c, 0xa4, 0x01, 0x00, 0x00,
}
