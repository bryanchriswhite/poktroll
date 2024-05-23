// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/session/session.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/pokt-network/poktroll/x/application/types"
	types "github.com/pokt-network/poktroll/x/shared/types"
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

// SessionHeader is a lightweight header for a session that can be passed around.
// It is the minimal amount of data required to hydrate & retrieve all data relevant to the session.
type SessionHeader struct {
	ApplicationAddress string         `protobuf:"bytes,1,opt,name=application_address,json=applicationAddress,proto3" json:"application_address,omitempty"`
	Service            *types.Service `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// NOTE: session_id can be derived from the above values using on-chain but is included in the header for convenience
	SessionId               string `protobuf:"bytes,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionStartBlockHeight int64  `protobuf:"varint,4,opt,name=session_start_block_height,json=sessionStartBlockHeight,proto3" json:"session_start_block_height,omitempty"`
	// Note that`session_end_block_height` is a derivative of (`start` + `num_blocks_per_session`)
	// as goverened by on-chain params at the time of the session start.
	// It is stored as an additional field to simplofy business logic in case
	// the number of blocks_per_session changes during the session.
	SessionEndBlockHeight int64 `protobuf:"varint,5,opt,name=session_end_block_height,json=sessionEndBlockHeight,proto3" json:"session_end_block_height,omitempty"`
}

func (m *SessionHeader) Reset()         { *m = SessionHeader{} }
func (m *SessionHeader) String() string { return proto.CompactTextString(m) }
func (*SessionHeader) ProtoMessage()    {}
func (*SessionHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c21b78e7ee4068a, []int{0}
}
func (m *SessionHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SessionHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SessionHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionHeader.Merge(m, src)
}
func (m *SessionHeader) XXX_Size() int {
	return m.Size()
}
func (m *SessionHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionHeader.DiscardUnknown(m)
}

var xxx_messageInfo_SessionHeader proto.InternalMessageInfo

func (m *SessionHeader) GetApplicationAddress() string {
	if m != nil {
		return m.ApplicationAddress
	}
	return ""
}

func (m *SessionHeader) GetService() *types.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *SessionHeader) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *SessionHeader) GetSessionStartBlockHeight() int64 {
	if m != nil {
		return m.SessionStartBlockHeight
	}
	return 0
}

func (m *SessionHeader) GetSessionEndBlockHeight() int64 {
	if m != nil {
		return m.SessionEndBlockHeight
	}
	return 0
}

// Session is a fully hydrated session object that contains all the information for the Session
// and its parcipants.
type Session struct {
	Header              *SessionHeader      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	SessionId           string              `protobuf:"bytes,2,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	SessionNumber       int64               `protobuf:"varint,3,opt,name=session_number,json=sessionNumber,proto3" json:"session_number,omitempty"`
	NumBlocksPerSession int64               `protobuf:"varint,4,opt,name=num_blocks_per_session,json=numBlocksPerSession,proto3" json:"num_blocks_per_session,omitempty"`
	Application         *types1.Application `protobuf:"bytes,5,opt,name=application,proto3" json:"application,omitempty"`
	Suppliers           []*types.Supplier   `protobuf:"bytes,6,rep,name=suppliers,proto3" json:"suppliers,omitempty"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c21b78e7ee4068a, []int{1}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Session.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return m.Size()
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetHeader() *SessionHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Session) GetSessionNumber() int64 {
	if m != nil {
		return m.SessionNumber
	}
	return 0
}

func (m *Session) GetNumBlocksPerSession() int64 {
	if m != nil {
		return m.NumBlocksPerSession
	}
	return 0
}

func (m *Session) GetApplication() *types1.Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *Session) GetSuppliers() []*types.Supplier {
	if m != nil {
		return m.Suppliers
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionHeader)(nil), "poktroll.session.SessionHeader")
	proto.RegisterType((*Session)(nil), "poktroll.session.Session")
}

func init() { proto.RegisterFile("poktroll/session/session.proto", fileDescriptor_6c21b78e7ee4068a) }

var fileDescriptor_6c21b78e7ee4068a = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0x04, 0x52, 0x65, 0xad, 0x22, 0xb4, 0xe5, 0xc3, 0x8d, 0x54, 0x13, 0x2a, 0x81,
	0x72, 0xa9, 0x8d, 0xdc, 0x43, 0x0e, 0x9c, 0x1a, 0x84, 0xd4, 0x72, 0x40, 0xc8, 0xb9, 0x71, 0xb1,
	0x1c, 0xef, 0x2a, 0xb6, 0x62, 0xef, 0x5a, 0xbb, 0x6b, 0x3e, 0xde, 0x82, 0x77, 0xe0, 0x15, 0xb8,
	0xf1, 0x02, 0x1c, 0x2b, 0x4e, 0x1c, 0x51, 0xf2, 0x22, 0x95, 0xd7, 0xb3, 0xb1, 0x9b, 0x9c, 0xb2,
	0x99, 0xdf, 0x7f, 0x66, 0xff, 0x33, 0x3b, 0x46, 0x6e, 0xc9, 0xd7, 0x4a, 0xf0, 0x3c, 0xf7, 0x25,
	0x95, 0x32, 0xe3, 0xcc, 0xfc, 0x7a, 0xa5, 0xe0, 0x8a, 0xe3, 0xc7, 0x86, 0x7b, 0x10, 0x1f, 0x9f,
	0x26, 0x5c, 0x16, 0x5c, 0x46, 0x9a, 0xfb, 0xcd, 0x9f, 0x46, 0x3c, 0x3e, 0x6b, 0x8b, 0xa5, 0xb1,
	0xa0, 0xc4, 0x97, 0x54, 0x7c, 0xc9, 0x12, 0x0a, 0xf8, 0xf5, 0x0e, 0xc7, 0x65, 0x99, 0x67, 0x49,
	0xac, 0xea, 0xfb, 0x3a, 0x67, 0xd0, 0xb9, 0x07, 0x65, 0xaa, 0x5a, 0x43, 0x45, 0xc3, 0xcf, 0x7f,
	0xf6, 0xd1, 0xf1, 0xa2, 0x71, 0x73, 0x4d, 0x63, 0x42, 0x05, 0xbe, 0x41, 0x27, 0x9d, 0x32, 0x51,
	0x4c, 0x88, 0xa0, 0x52, 0x3a, 0xd6, 0xc4, 0x9a, 0x8e, 0xe6, 0xce, 0xdf, 0x5f, 0x17, 0x4f, 0xc0,
	0xe7, 0x55, 0x43, 0x16, 0x4a, 0x64, 0x6c, 0x15, 0xe2, 0x4e, 0x12, 0x10, 0x1c, 0xa0, 0x23, 0x70,
	0xed, 0xf4, 0x27, 0xd6, 0xd4, 0x0e, 0x1c, 0xaf, 0x1d, 0x81, 0xb6, 0xe3, 0x2d, 0x1a, 0x1e, 0x1a,
	0x21, 0x3e, 0x43, 0x08, 0xa6, 0x13, 0x65, 0xc4, 0x19, 0xd4, 0xb7, 0x86, 0x23, 0x88, 0xdc, 0x10,
	0xfc, 0x16, 0x8d, 0x0d, 0x96, 0x2a, 0x16, 0x2a, 0x5a, 0xe6, 0x3c, 0x59, 0x47, 0x29, 0xcd, 0x56,
	0xa9, 0x72, 0x1e, 0x4c, 0xac, 0xe9, 0x20, 0x7c, 0x0e, 0x8a, 0x45, 0x2d, 0x98, 0xd7, 0xfc, 0x5a,
	0x63, 0x3c, 0x43, 0x8e, 0x49, 0xa6, 0x8c, 0xdc, 0x4f, 0x7d, 0xa8, 0x53, 0x9f, 0x02, 0x7f, 0xcf,
	0x48, 0x27, 0xf1, 0xfc, 0x77, 0x1f, 0x1d, 0xc1, 0x94, 0xf0, 0x0c, 0x0d, 0x53, 0x3d, 0x29, 0x3d,
	0x12, 0x3b, 0x78, 0xe1, 0xed, 0x3f, 0xab, 0x77, 0x6f, 0xa0, 0x21, 0xc8, 0xf7, 0x3a, 0xeb, 0xef,
	0x77, 0xf6, 0x0a, 0x3d, 0x32, 0x98, 0x55, 0xc5, 0x92, 0x0a, 0xdd, 0xfc, 0x20, 0x3c, 0x86, 0xe8,
	0x47, 0x1d, 0xc4, 0x97, 0xe8, 0x19, 0xab, 0x8a, 0xc6, 0xbb, 0x8c, 0x4a, 0x2a, 0x22, 0xe0, 0xd0,
	0xfc, 0x09, 0xab, 0x0a, 0x6d, 0x5d, 0x7e, 0xa2, 0xc2, 0x78, 0x7e, 0x87, 0xec, 0xce, 0xf3, 0xe8,
	0x5e, 0xed, 0xe0, 0x65, 0x6b, 0xbc, 0xbb, 0x37, 0x57, 0xed, 0x39, 0xec, 0x66, 0xe1, 0x19, 0x1a,
	0x99, 0xe5, 0x91, 0xce, 0x70, 0x32, 0x98, 0xda, 0xc1, 0xe9, 0xe1, 0x7b, 0x82, 0x22, 0x6c, 0xb5,
	0xf3, 0x0f, 0x7f, 0x36, 0xae, 0x75, 0xbb, 0x71, 0xad, 0xff, 0x1b, 0xd7, 0xfa, 0xb1, 0x75, 0x7b,
	0xb7, 0x5b, 0xb7, 0xf7, 0x6f, 0xeb, 0xf6, 0x3e, 0xbf, 0x59, 0x65, 0x2a, 0xad, 0x96, 0x5e, 0xc2,
	0x0b, 0xbf, 0xae, 0x74, 0xc1, 0xa8, 0xfa, 0xca, 0xc5, 0xda, 0xdf, 0x6d, 0xed, 0xb7, 0xdd, 0xb7,
	0xa4, 0xbe, 0x97, 0x54, 0x2e, 0x87, 0x7a, 0x6d, 0x2f, 0xef, 0x02, 0x00, 0x00, 0xff, 0xff, 0x87,
	0x46, 0x6f, 0xb0, 0x6c, 0x03, 0x00, 0x00,
}

func (m *SessionHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SessionHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SessionEndBlockHeight != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.SessionEndBlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.SessionStartBlockHeight != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.SessionStartBlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.SessionId) > 0 {
		i -= len(m.SessionId)
		copy(dAtA[i:], m.SessionId)
		i = encodeVarintSession(dAtA, i, uint64(len(m.SessionId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Service != nil {
		{
			size, err := m.Service.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ApplicationAddress) > 0 {
		i -= len(m.ApplicationAddress)
		copy(dAtA[i:], m.ApplicationAddress)
		i = encodeVarintSession(dAtA, i, uint64(len(m.ApplicationAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Session) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Session) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Session) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Suppliers) > 0 {
		for iNdEx := len(m.Suppliers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Suppliers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintSession(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Application != nil {
		{
			size, err := m.Application.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.NumBlocksPerSession != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.NumBlocksPerSession))
		i--
		dAtA[i] = 0x20
	}
	if m.SessionNumber != 0 {
		i = encodeVarintSession(dAtA, i, uint64(m.SessionNumber))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SessionId) > 0 {
		i -= len(m.SessionId)
		copy(dAtA[i:], m.SessionId)
		i = encodeVarintSession(dAtA, i, uint64(len(m.SessionId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSession(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSession(dAtA []byte, offset int, v uint64) int {
	offset -= sovSession(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SessionHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ApplicationAddress)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	if m.Service != nil {
		l = m.Service.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	if m.SessionStartBlockHeight != 0 {
		n += 1 + sovSession(uint64(m.SessionStartBlockHeight))
	}
	if m.SessionEndBlockHeight != 0 {
		n += 1 + sovSession(uint64(m.SessionEndBlockHeight))
	}
	return n
}

func (m *Session) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	l = len(m.SessionId)
	if l > 0 {
		n += 1 + l + sovSession(uint64(l))
	}
	if m.SessionNumber != 0 {
		n += 1 + sovSession(uint64(m.SessionNumber))
	}
	if m.NumBlocksPerSession != 0 {
		n += 1 + sovSession(uint64(m.NumBlocksPerSession))
	}
	if m.Application != nil {
		l = m.Application.Size()
		n += 1 + l + sovSession(uint64(l))
	}
	if len(m.Suppliers) > 0 {
		for _, e := range m.Suppliers {
			l = e.Size()
			n += 1 + l + sovSession(uint64(l))
		}
	}
	return n
}

func sovSession(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSession(x uint64) (n int) {
	return sovSession(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SessionHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: SessionHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Service == nil {
				m.Service = &types.Service{}
			}
			if err := m.Service.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionStartBlockHeight", wireType)
			}
			m.SessionStartBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionStartBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionEndBlockHeight", wireType)
			}
			m.SessionEndBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionEndBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Session) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSession
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
			return fmt.Errorf("proto: Session: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Session: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &SessionHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
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
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SessionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionNumber", wireType)
			}
			m.SessionNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumBlocksPerSession", wireType)
			}
			m.NumBlocksPerSession = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumBlocksPerSession |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Application == nil {
				m.Application = &types1.Application{}
			}
			if err := m.Application.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suppliers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSession
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSession
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSession
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Suppliers = append(m.Suppliers, &types.Supplier{})
			if err := m.Suppliers[len(m.Suppliers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSession(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSession
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSession(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
					return 0, ErrIntOverflowSession
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
				return 0, ErrInvalidLengthSession
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSession
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSession
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSession        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSession          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSession = fmt.Errorf("proto: unexpected end of group")
)
