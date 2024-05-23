// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/shared/service.proto

// NOTE that the `shared` package is not a Cosmos module,
// but rather a manually created package to resolve circular type dependencies.

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// Enum to define RPC types
type RPCType int32

const (
	RPCType_UNKNOWN_RPC RPCType = 0
	RPCType_GRPC        RPCType = 1
	RPCType_WEBSOCKET   RPCType = 2
	RPCType_JSON_RPC    RPCType = 3
	RPCType_REST        RPCType = 4
)

var RPCType_name = map[int32]string{
	0: "UNKNOWN_RPC",
	1: "GRPC",
	2: "WEBSOCKET",
	3: "JSON_RPC",
	4: "REST",
}

var RPCType_value = map[string]int32{
	"UNKNOWN_RPC": 0,
	"GRPC":        1,
	"WEBSOCKET":   2,
	"JSON_RPC":    3,
	"REST":        4,
}

func (x RPCType) String() string {
	return proto.EnumName(RPCType_name, int32(x))
}

func (RPCType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_302c2f793a11ae1e, []int{0}
}

// Enum to define configuration options
// TODO_RESEARCH: Should these be configs, SLAs or something else? There will be more discussion once we get closer to implementing on-chain QoS.
type ConfigOptions int32

const (
	ConfigOptions_UNKNOWN_CONFIG ConfigOptions = 0
	ConfigOptions_TIMEOUT        ConfigOptions = 1
)

var ConfigOptions_name = map[int32]string{
	0: "UNKNOWN_CONFIG",
	1: "TIMEOUT",
}

var ConfigOptions_value = map[string]int32{
	"UNKNOWN_CONFIG": 0,
	"TIMEOUT":        1,
}

func (x ConfigOptions) String() string {
	return proto.EnumName(ConfigOptions_name, int32(x))
}

func (ConfigOptions) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_302c2f793a11ae1e, []int{1}
}

// Service message to encapsulate unique and semantic identifiers for a service on the network
type Service struct {
	// For example, what if we want to request a session for a certain service but with some additional configs that identify it?
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// TODO_TECHDEBT: Name is currently unused but acts as a reminder that an optional onchain representation of the service is necessary
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_302c2f793a11ae1e, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Service.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return m.Size()
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// ApplicationServiceConfig holds the service configuration the application stakes for
type ApplicationServiceConfig struct {
	Service *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (m *ApplicationServiceConfig) Reset()         { *m = ApplicationServiceConfig{} }
func (m *ApplicationServiceConfig) String() string { return proto.CompactTextString(m) }
func (*ApplicationServiceConfig) ProtoMessage()    {}
func (*ApplicationServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_302c2f793a11ae1e, []int{1}
}
func (m *ApplicationServiceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApplicationServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApplicationServiceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApplicationServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationServiceConfig.Merge(m, src)
}
func (m *ApplicationServiceConfig) XXX_Size() int {
	return m.Size()
}
func (m *ApplicationServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationServiceConfig proto.InternalMessageInfo

func (m *ApplicationServiceConfig) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

// SupplierServiceConfig holds the service configuration the supplier stakes for
type SupplierServiceConfig struct {
	Service   *Service            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Endpoints []*SupplierEndpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (m *SupplierServiceConfig) Reset()         { *m = SupplierServiceConfig{} }
func (m *SupplierServiceConfig) String() string { return proto.CompactTextString(m) }
func (*SupplierServiceConfig) ProtoMessage()    {}
func (*SupplierServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_302c2f793a11ae1e, []int{2}
}
func (m *SupplierServiceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplierServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupplierServiceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupplierServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplierServiceConfig.Merge(m, src)
}
func (m *SupplierServiceConfig) XXX_Size() int {
	return m.Size()
}
func (m *SupplierServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplierServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SupplierServiceConfig proto.InternalMessageInfo

func (m *SupplierServiceConfig) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *SupplierServiceConfig) GetEndpoints() []*SupplierEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// SupplierEndpoint message to hold service configuration details
type SupplierEndpoint struct {
	Url     string          `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	RpcType RPCType         `protobuf:"varint,2,opt,name=rpc_type,json=rpcType,proto3,enum=poktroll.shared.RPCType" json:"rpc_type,omitempty"`
	Configs []*ConfigOption `protobuf:"bytes,3,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (m *SupplierEndpoint) Reset()         { *m = SupplierEndpoint{} }
func (m *SupplierEndpoint) String() string { return proto.CompactTextString(m) }
func (*SupplierEndpoint) ProtoMessage()    {}
func (*SupplierEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_302c2f793a11ae1e, []int{3}
}
func (m *SupplierEndpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SupplierEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SupplierEndpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SupplierEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupplierEndpoint.Merge(m, src)
}
func (m *SupplierEndpoint) XXX_Size() int {
	return m.Size()
}
func (m *SupplierEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_SupplierEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_SupplierEndpoint proto.InternalMessageInfo

func (m *SupplierEndpoint) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SupplierEndpoint) GetRpcType() RPCType {
	if m != nil {
		return m.RpcType
	}
	return RPCType_UNKNOWN_RPC
}

func (m *SupplierEndpoint) GetConfigs() []*ConfigOption {
	if m != nil {
		return m.Configs
	}
	return nil
}

// Key-value wrapper for config options, as proto maps can't be keyed by enums
type ConfigOption struct {
	Key   ConfigOptions `protobuf:"varint,1,opt,name=key,proto3,enum=poktroll.shared.ConfigOptions" json:"key,omitempty"`
	Value string        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ConfigOption) Reset()         { *m = ConfigOption{} }
func (m *ConfigOption) String() string { return proto.CompactTextString(m) }
func (*ConfigOption) ProtoMessage()    {}
func (*ConfigOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_302c2f793a11ae1e, []int{4}
}
func (m *ConfigOption) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigOption.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigOption.Merge(m, src)
}
func (m *ConfigOption) XXX_Size() int {
	return m.Size()
}
func (m *ConfigOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigOption.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigOption proto.InternalMessageInfo

func (m *ConfigOption) GetKey() ConfigOptions {
	if m != nil {
		return m.Key
	}
	return ConfigOptions_UNKNOWN_CONFIG
}

func (m *ConfigOption) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("poktroll.shared.RPCType", RPCType_name, RPCType_value)
	proto.RegisterEnum("poktroll.shared.ConfigOptions", ConfigOptions_name, ConfigOptions_value)
	proto.RegisterType((*Service)(nil), "poktroll.shared.Service")
	proto.RegisterType((*ApplicationServiceConfig)(nil), "poktroll.shared.ApplicationServiceConfig")
	proto.RegisterType((*SupplierServiceConfig)(nil), "poktroll.shared.SupplierServiceConfig")
	proto.RegisterType((*SupplierEndpoint)(nil), "poktroll.shared.SupplierEndpoint")
	proto.RegisterType((*ConfigOption)(nil), "poktroll.shared.ConfigOption")
}

func init() { proto.RegisterFile("poktroll/shared/service.proto", fileDescriptor_302c2f793a11ae1e) }

var fileDescriptor_302c2f793a11ae1e = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xf6, 0xa3, 0xe0, 0x64, 0xd2, 0xa6, 0xd6, 0x08, 0x24, 0x5f, 0x6a, 0x15, 0x9f, 0xaa, 0x4a,
	0xb5, 0xab, 0xf4, 0xc0, 0x11, 0x51, 0xcb, 0x54, 0x21, 0xc2, 0xae, 0xd6, 0x2e, 0x95, 0xb8, 0x54,
	0xa9, 0xb3, 0xb4, 0x56, 0x5c, 0xef, 0x6a, 0xed, 0x14, 0xf2, 0x1f, 0x38, 0x20, 0x7e, 0x15, 0xc7,
	0x1e, 0x39, 0xa2, 0xe4, 0x8f, 0xa0, 0xf5, 0x83, 0x47, 0x8b, 0xb8, 0x70, 0x1b, 0x7b, 0xbe, 0xd7,
	0xcc, 0x0e, 0xec, 0x70, 0x36, 0xaf, 0x04, 0xcb, 0x73, 0xaf, 0xbc, 0x9e, 0x0a, 0x3a, 0xf3, 0x4a,
	0x2a, 0x6e, 0xb3, 0x94, 0xba, 0x5c, 0xb0, 0x8a, 0xe1, 0x76, 0xd7, 0x76, 0x9b, 0xb6, 0x73, 0x00,
	0x46, 0xdc, 0x20, 0x70, 0x08, 0x5a, 0x36, 0xb3, 0xd4, 0x5d, 0x75, 0xaf, 0x4f, 0xb4, 0x6c, 0x86,
	0x08, 0x1b, 0xc5, 0xf4, 0x86, 0x5a, 0x5a, 0xfd, 0xa7, 0xae, 0x9d, 0x10, 0xac, 0x97, 0x9c, 0xe7,
	0x59, 0x3a, 0xad, 0x32, 0x56, 0xb4, 0x4c, 0x9f, 0x15, 0xef, 0xb3, 0x2b, 0x1c, 0x81, 0xd1, 0x9a,
	0xd5, 0x22, 0x83, 0x91, 0xe5, 0xde, 0x73, 0x73, 0x5b, 0x02, 0xe9, 0x80, 0xce, 0x27, 0x15, 0x9e,
	0xc6, 0x0b, 0xa9, 0x48, 0xc5, 0x7f, 0xab, 0xe1, 0x0b, 0xe8, 0xd3, 0x62, 0xc6, 0x59, 0x56, 0x54,
	0xa5, 0xa5, 0xed, 0xea, 0x7b, 0x83, 0xd1, 0xb3, 0x87, 0xac, 0xd6, 0x2e, 0x68, 0x91, 0xe4, 0x17,
	0xc7, 0xf9, 0xa2, 0x82, 0x79, 0xbf, 0x8f, 0x26, 0xe8, 0x0b, 0x91, 0xb7, 0x8b, 0x91, 0x25, 0x1e,
	0x41, 0x4f, 0xf0, 0xf4, 0xa2, 0x5a, 0xf2, 0x66, 0x3b, 0xc3, 0xbf, 0x84, 0x23, 0xa7, 0x7e, 0xb2,
	0xe4, 0x94, 0x18, 0x82, 0xa7, 0xb2, 0xc0, 0xe7, 0x60, 0xa4, 0xf5, 0x68, 0xa5, 0xa5, 0xd7, 0xd1,
	0x76, 0x1e, 0x70, 0x9a, 0xd1, 0x23, 0x2e, 0x77, 0x4b, 0x3a, 0xb4, 0xf3, 0x16, 0x36, 0x7f, 0x6f,
	0xe0, 0x21, 0xe8, 0x73, 0xba, 0xac, 0xf3, 0x0c, 0x47, 0xf6, 0x3f, 0x45, 0x4a, 0x22, 0xa1, 0xf8,
	0x04, 0x1e, 0xdd, 0x4e, 0xf3, 0x45, 0xf7, 0x94, 0xcd, 0xc7, 0xfe, 0x04, 0x8c, 0x36, 0x24, 0x6e,
	0xc3, 0xe0, 0x2c, 0x9c, 0x84, 0xd1, 0x79, 0x78, 0x41, 0x4e, 0x7d, 0x53, 0xc1, 0x1e, 0x6c, 0x9c,
	0xc8, 0x4a, 0xc5, 0x2d, 0xe8, 0x9f, 0x07, 0xc7, 0x71, 0xe4, 0x4f, 0x82, 0xc4, 0xd4, 0x70, 0x13,
	0x7a, 0xaf, 0xe3, 0xa8, 0x81, 0xe9, 0x12, 0x46, 0x82, 0x38, 0x31, 0x37, 0xf6, 0x0f, 0x61, 0xeb,
	0x0f, 0x63, 0x44, 0x18, 0x76, 0x92, 0x7e, 0x14, 0xbe, 0x1a, 0x9f, 0x98, 0x0a, 0x0e, 0xc0, 0x48,
	0xc6, 0x6f, 0x82, 0xe8, 0x2c, 0x31, 0xd5, 0xe3, 0xf1, 0xd7, 0x95, 0xad, 0xde, 0xad, 0x6c, 0xf5,
	0xfb, 0xca, 0x56, 0x3f, 0xaf, 0x6d, 0xe5, 0x6e, 0x6d, 0x2b, 0xdf, 0xd6, 0xb6, 0xf2, 0xce, 0xbb,
	0xca, 0xaa, 0xeb, 0xc5, 0xa5, 0x9b, 0xb2, 0x1b, 0x4f, 0x4e, 0x77, 0x50, 0xd0, 0xea, 0x03, 0x13,
	0x73, 0xef, 0xe7, 0x6d, 0x7f, 0xec, 0xae, 0x5b, 0x3e, 0x41, 0x79, 0xf9, 0xb8, 0x3e, 0xee, 0xa3,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x1d, 0x9c, 0x9c, 0xfd, 0x02, 0x00, 0x00,
}

func (m *Service) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Service) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Service) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintService(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintService(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ApplicationServiceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApplicationServiceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApplicationServiceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Service != nil {
		{
			size, err := m.Service.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SupplierServiceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplierServiceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplierServiceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Service != nil {
		{
			size, err := m.Service.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SupplierEndpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SupplierEndpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SupplierEndpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for iNdEx := len(m.Configs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Configs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.RpcType != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.RpcType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintService(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigOption) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigOption) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigOption) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintService(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Key != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Key))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Service) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *ApplicationServiceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Service != nil {
		l = m.Service.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *SupplierServiceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Service != nil {
		l = m.Service.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	return n
}

func (m *SupplierEndpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.RpcType != 0 {
		n += 1 + sovService(uint64(m.RpcType))
	}
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	return n
}

func (m *ConfigOption) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Key != 0 {
		n += 1 + sovService(uint64(m.Key))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Service) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Service: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Service: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *ApplicationServiceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ApplicationServiceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApplicationServiceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Service == nil {
				m.Service = &Service{}
			}
			if err := m.Service.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *SupplierServiceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: SupplierServiceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplierServiceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Service", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Service == nil {
				m.Service = &Service{}
			}
			if err := m.Service.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, &SupplierEndpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *SupplierEndpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: SupplierEndpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SupplierEndpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RpcType", wireType)
			}
			m.RpcType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RpcType |= RPCType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Configs = append(m.Configs, &ConfigOption{})
			if err := m.Configs[len(m.Configs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *ConfigOption) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ConfigOption: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigOption: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			m.Key = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Key |= ConfigOptions(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
