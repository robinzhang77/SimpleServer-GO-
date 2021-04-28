// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

package test

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CreateAgentReq struct {
	Plr    string  `protobuf:"bytes,1,req,name=plr" json:"plr"`
	X      float32 `protobuf:"fixed32,2,req,name=x" json:"x"`
	Y      float32 `protobuf:"fixed32,3,req,name=y" json:"y"`
	Z      float32 `protobuf:"fixed32,4,req,name=z" json:"z"`
	Radius float32 `protobuf:"fixed32,5,req,name=radius" json:"radius"`
	Speed  float32 `protobuf:"fixed32,6,req,name=speed" json:"speed"`
}

func (m *CreateAgentReq) Reset()         { *m = CreateAgentReq{} }
func (m *CreateAgentReq) String() string { return proto.CompactTextString(m) }
func (*CreateAgentReq) ProtoMessage()    {}
func (*CreateAgentReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}
func (m *CreateAgentReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateAgentReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateAgentReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateAgentReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentReq.Merge(m, src)
}
func (m *CreateAgentReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateAgentReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentReq proto.InternalMessageInfo

func (m *CreateAgentReq) GetPlr() string {
	if m != nil {
		return m.Plr
	}
	return ""
}

func (m *CreateAgentReq) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CreateAgentReq) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CreateAgentReq) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *CreateAgentReq) GetRadius() float32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *CreateAgentReq) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

type CreateAgentAck struct {
	X      float32 `protobuf:"fixed32,2,req,name=x" json:"x"`
	Y      float32 `protobuf:"fixed32,3,req,name=y" json:"y"`
	Z      float32 `protobuf:"fixed32,4,req,name=z" json:"z"`
	Radius float32 `protobuf:"fixed32,5,req,name=radius" json:"radius"`
	Speed  float32 `protobuf:"fixed32,6,req,name=speed" json:"speed"`
}

func (m *CreateAgentAck) Reset()         { *m = CreateAgentAck{} }
func (m *CreateAgentAck) String() string { return proto.CompactTextString(m) }
func (*CreateAgentAck) ProtoMessage()    {}
func (*CreateAgentAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}
func (m *CreateAgentAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateAgentAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateAgentAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreateAgentAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentAck.Merge(m, src)
}
func (m *CreateAgentAck) XXX_Size() int {
	return m.Size()
}
func (m *CreateAgentAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentAck.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentAck proto.InternalMessageInfo

func (m *CreateAgentAck) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *CreateAgentAck) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *CreateAgentAck) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *CreateAgentAck) GetRadius() float32 {
	if m != nil {
		return m.Radius
	}
	return 0
}

func (m *CreateAgentAck) GetSpeed() float32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateAgentReq)(nil), "test.CreateAgentReq")
	proto.RegisterType((*CreateAgentAck)(nil), "test.CreateAgentAck")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x1a, 0x19, 0xb9, 0xf8, 0x9c,
	0x8b, 0x52, 0x13, 0x4b, 0x52, 0x1d, 0xd3, 0x53, 0xf3, 0x4a, 0x82, 0x52, 0x0b, 0x85, 0x04, 0xb9,
	0x98, 0x0b, 0x72, 0x8a, 0x24, 0x18, 0x15, 0x98, 0x34, 0x38, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67,
	0x10, 0xe2, 0xe7, 0x62, 0xac, 0x90, 0x60, 0x52, 0x60, 0xd2, 0x60, 0x42, 0x08, 0x54, 0x4a, 0x30,
	0xa3, 0x0a, 0x54, 0x49, 0xb0, 0x20, 0x09, 0x88, 0x70, 0xb1, 0x15, 0x25, 0xa6, 0x64, 0x96, 0x16,
	0x4b, 0xb0, 0x22, 0x89, 0x0a, 0x73, 0xb1, 0x16, 0x17, 0xa4, 0xa6, 0xa6, 0x48, 0xb0, 0x21, 0x04,
	0x95, 0xf2, 0x50, 0x9c, 0xe0, 0x98, 0x9c, 0x4d, 0x5b, 0xfb, 0x9c, 0x24, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96,
	0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x01, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x39, 0xe4, 0x49, 0x20,
	0x01, 0x00, 0x00,
}

func (m *CreateAgentReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateAgentReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAgentReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Speed))))
	i--
	dAtA[i] = 0x35
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Radius))))
	i--
	dAtA[i] = 0x2d
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Z))))
	i--
	dAtA[i] = 0x25
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Y))))
	i--
	dAtA[i] = 0x1d
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.X))))
	i--
	dAtA[i] = 0x15
	i -= len(m.Plr)
	copy(dAtA[i:], m.Plr)
	i = encodeVarintTest(dAtA, i, uint64(len(m.Plr)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *CreateAgentAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateAgentAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreateAgentAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Speed))))
	i--
	dAtA[i] = 0x35
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Radius))))
	i--
	dAtA[i] = 0x2d
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Z))))
	i--
	dAtA[i] = 0x25
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.Y))))
	i--
	dAtA[i] = 0x1d
	i -= 4
	encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.X))))
	i--
	dAtA[i] = 0x15
	return len(dAtA) - i, nil
}

func encodeVarintTest(dAtA []byte, offset int, v uint64) int {
	offset -= sovTest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreateAgentReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Plr)
	n += 1 + l + sovTest(uint64(l))
	n += 5
	n += 5
	n += 5
	n += 5
	n += 5
	return n
}

func (m *CreateAgentAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 5
	n += 5
	n += 5
	n += 5
	n += 5
	return n
}

func sovTest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTest(x uint64) (n int) {
	return sovTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateAgentReq) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateAgentReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateAgentReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Plr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.X = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Y = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000004)
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Z = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000008)
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Radius", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Radius = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000010)
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Speed", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Speed = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000020)
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("plr")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("x")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("y")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("z")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("radius")
	}
	if hasFields[0]&uint64(0x00000020) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("speed")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateAgentAck) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateAgentAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateAgentAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.X = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000001)
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Y = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000002)
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Z", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Z = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000004)
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Radius", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Radius = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000008)
		case 6:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Speed", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.Speed = float32(math.Float32frombits(v))
			hasFields[0] |= uint64(0x00000010)
		default:
			iNdEx = preIndex
			skippy, err := skipTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTest
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("x")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("y")
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("z")
	}
	if hasFields[0]&uint64(0x00000008) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("radius")
	}
	if hasFields[0]&uint64(0x00000010) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("speed")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTest
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTest
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTest = fmt.Errorf("proto: unexpected end of group")
)