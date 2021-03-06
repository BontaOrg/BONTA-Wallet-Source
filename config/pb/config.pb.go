// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package configpb is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	Config
	P2PConfig
	APIConfig
	NetworkConfig
	LogConfig
	AppConfig
	DataConfig
	AdvancedConfig
*/
package configpb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	App      *AppConfig      `protobuf:"bytes,1,opt,name=app" json:"app"`
	Network  *NetworkConfig  `protobuf:"bytes,2,opt,name=network" json:"network"`
	Log      *LogConfig      `protobuf:"bytes,3,opt,name=log" json:"log"`
	Data     *DataConfig     `protobuf:"bytes,4,opt,name=data" json:"data"`
	Advanced *AdvancedConfig `protobuf:"bytes,5,opt,name=advanced" json:"advanced"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetApp() *AppConfig {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Config) GetNetwork() *NetworkConfig {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Config) GetLog() *LogConfig {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Config) GetData() *DataConfig {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Config) GetAdvanced() *AdvancedConfig {
	if m != nil {
		return m.Advanced
	}
	return nil
}

type P2PConfig struct {
	Seeds            string   `protobuf:"bytes,1,opt,name=seeds,proto3" json:"seeds"`
	AddPeer          []string `protobuf:"bytes,2,rep,name=add_peer,json=addPeer" json:"add_peer"`
	SkipUpnp         bool     `protobuf:"varint,3,opt,name=skip_upnp,json=skipUpnp,proto3" json:"skip_upnp"`
	HandshakeTimeout uint32   `protobuf:"varint,4,opt,name=handshake_timeout,json=handshakeTimeout,proto3" json:"handshake_timeout"`
	DialTimeout      uint32   `protobuf:"varint,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout"`
	VaultMode        bool     `protobuf:"varint,6,opt,name=vault_mode,json=vaultMode,proto3" json:"vault_mode"`
	ListenAddress    string   `protobuf:"bytes,7,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address"`
}

func (m *P2PConfig) Reset()                    { *m = P2PConfig{} }
func (m *P2PConfig) String() string            { return proto.CompactTextString(m) }
func (*P2PConfig) ProtoMessage()               {}
func (*P2PConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{1} }

func (m *P2PConfig) GetSeeds() string {
	if m != nil {
		return m.Seeds
	}
	return ""
}

func (m *P2PConfig) GetAddPeer() []string {
	if m != nil {
		return m.AddPeer
	}
	return nil
}

func (m *P2PConfig) GetSkipUpnp() bool {
	if m != nil {
		return m.SkipUpnp
	}
	return false
}

func (m *P2PConfig) GetHandshakeTimeout() uint32 {
	if m != nil {
		return m.HandshakeTimeout
	}
	return 0
}

func (m *P2PConfig) GetDialTimeout() uint32 {
	if m != nil {
		return m.DialTimeout
	}
	return 0
}

func (m *P2PConfig) GetVaultMode() bool {
	if m != nil {
		return m.VaultMode
	}
	return false
}

func (m *P2PConfig) GetListenAddress() string {
	if m != nil {
		return m.ListenAddress
	}
	return ""
}

type APIConfig struct {
	Host         string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host"`
	GRPCPort     string   `protobuf:"bytes,2,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port"`
	HttpPort     string   `protobuf:"bytes,3,opt,name=http_port,json=httpPort,proto3" json:"http_port"`
	HttpCORSAddr []string `protobuf:"bytes,4,rep,name=http_cors_addr,json=httpCORSAddr" json:"http_cors_addr"`
	DisableTls   bool     `protobuf:"varint,5,opt,name=disable_tls,json=disableTls,proto3" json:"disable_tls"`
	RpcCert      string   `protobuf:"bytes,6,opt,name=rpc_cert,json=rpcCert,proto3" json:"rpc_cert"`
	RpcKey       string   `protobuf:"bytes,7,opt,name=rpc_key,json=rpcKey,proto3" json:"rpc_key"`
}

func (m *APIConfig) Reset()                    { *m = APIConfig{} }
func (m *APIConfig) String() string            { return proto.CompactTextString(m) }
func (*APIConfig) ProtoMessage()               {}
func (*APIConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{2} }

func (m *APIConfig) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *APIConfig) GetGRPCPort() string {
	if m != nil {
		return m.GRPCPort
	}
	return ""
}

func (m *APIConfig) GetHttpPort() string {
	if m != nil {
		return m.HttpPort
	}
	return ""
}

func (m *APIConfig) GetHttpCORSAddr() []string {
	if m != nil {
		return m.HttpCORSAddr
	}
	return nil
}

func (m *APIConfig) GetDisableTls() bool {
	if m != nil {
		return m.DisableTls
	}
	return false
}

func (m *APIConfig) GetRpcCert() string {
	if m != nil {
		return m.RpcCert
	}
	return ""
}

func (m *APIConfig) GetRpcKey() string {
	if m != nil {
		return m.RpcKey
	}
	return ""
}

type NetworkConfig struct {
	P2P *P2PConfig `protobuf:"bytes,1,opt,name=p2p" json:"p2p"`
	API *APIConfig `protobuf:"bytes,2,opt,name=api" json:"api"`
}

func (m *NetworkConfig) Reset()                    { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string            { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()               {}
func (*NetworkConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{3} }

func (m *NetworkConfig) GetP2P() *P2PConfig {
	if m != nil {
		return m.P2P
	}
	return nil
}

func (m *NetworkConfig) GetAPI() *APIConfig {
	if m != nil {
		return m.API
	}
	return nil
}

type LogConfig struct {
	LogDir   string `protobuf:"bytes,1,opt,name=log_dir,json=logDir,proto3" json:"log_dir"`
	LogLevel string `protobuf:"bytes,2,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
}

func (m *LogConfig) Reset()                    { *m = LogConfig{} }
func (m *LogConfig) String() string            { return proto.CompactTextString(m) }
func (*LogConfig) ProtoMessage()               {}
func (*LogConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{4} }

func (m *LogConfig) GetLogDir() string {
	if m != nil {
		return m.LogDir
	}
	return ""
}

func (m *LogConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

type AppConfig struct {
	Profile            string `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile"`
	CPUProfile         string `protobuf:"bytes,2,opt,name=cpu_profile,json=cpuProfile,proto3" json:"cpu_profile"`
	NoPeerBloomFilters bool   `protobuf:"varint,3,opt,name=no_peer_bloom_filters,json=noPeerBloomFilters,proto3" json:"no_peer_bloom_filters"`
}

func (m *AppConfig) Reset()                    { *m = AppConfig{} }
func (m *AppConfig) String() string            { return proto.CompactTextString(m) }
func (*AppConfig) ProtoMessage()               {}
func (*AppConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{5} }

func (m *AppConfig) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *AppConfig) GetCPUProfile() string {
	if m != nil {
		return m.CPUProfile
	}
	return ""
}

func (m *AppConfig) GetNoPeerBloomFilters() bool {
	if m != nil {
		return m.NoPeerBloomFilters
	}
	return false
}

type DataConfig struct {
	DbType        string `protobuf:"bytes,1,opt,name=db_type,json=dbType,proto3" json:"db_type"`
	DbDir         string `protobuf:"bytes,2,opt,name=db_dir,json=dbDir,proto3" json:"db_dir"`
	WalletPubPass string `protobuf:"bytes,3,opt,name=wallet_pub_pass,json=walletPubPass,proto3" json:"wallet_pub_pass"`
}

func (m *DataConfig) Reset()                    { *m = DataConfig{} }
func (m *DataConfig) String() string            { return proto.CompactTextString(m) }
func (*DataConfig) ProtoMessage()               {}
func (*DataConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{6} }

func (m *DataConfig) GetDbType() string {
	if m != nil {
		return m.DbType
	}
	return ""
}

func (m *DataConfig) GetDbDir() string {
	if m != nil {
		return m.DbDir
	}
	return ""
}

func (m *DataConfig) GetWalletPubPass() string {
	if m != nil {
		return m.WalletPubPass
	}
	return ""
}

type AdvancedConfig struct {
	AddressGapLimit         uint32 `protobuf:"varint,1,opt,name=address_gap_limit,json=addressGapLimit,proto3" json:"address_gap_limit"`
	MaxUnusedStakingAddress uint32 `protobuf:"varint,2,opt,name=max_unused_staking_address,json=maxUnusedStakingAddress,proto3" json:"max_unused_staking_address"`
	MaxTxFee                string `protobuf:"bytes,3,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee"`
}

func (m *AdvancedConfig) Reset()                    { *m = AdvancedConfig{} }
func (m *AdvancedConfig) String() string            { return proto.CompactTextString(m) }
func (*AdvancedConfig) ProtoMessage()               {}
func (*AdvancedConfig) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{7} }

func (m *AdvancedConfig) GetAddressGapLimit() uint32 {
	if m != nil {
		return m.AddressGapLimit
	}
	return 0
}

func (m *AdvancedConfig) GetMaxUnusedStakingAddress() uint32 {
	if m != nil {
		return m.MaxUnusedStakingAddress
	}
	return 0
}

func (m *AdvancedConfig) GetMaxTxFee() string {
	if m != nil {
		return m.MaxTxFee
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "configpb.Config")
	proto.RegisterType((*P2PConfig)(nil), "configpb.P2PConfig")
	proto.RegisterType((*APIConfig)(nil), "configpb.APIConfig")
	proto.RegisterType((*NetworkConfig)(nil), "configpb.NetworkConfig")
	proto.RegisterType((*LogConfig)(nil), "configpb.LogConfig")
	proto.RegisterType((*AppConfig)(nil), "configpb.AppConfig")
	proto.RegisterType((*DataConfig)(nil), "configpb.DataConfig")
	proto.RegisterType((*AdvancedConfig)(nil), "configpb.AdvancedConfig")
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdb, 0x4e, 0xdc, 0x3a,
	0x14, 0xd5, 0x30, 0xd7, 0x6c, 0x18, 0x38, 0xf8, 0x80, 0xc8, 0xb9, 0xe9, 0x70, 0xa2, 0x43, 0x85,
	0x5a, 0x09, 0x09, 0xda, 0xb7, 0x3e, 0x4d, 0x41, 0x54, 0x55, 0x69, 0x15, 0xa5, 0xc3, 0x63, 0x65,
	0x39, 0xb1, 0x27, 0x13, 0x8d, 0x27, 0xb6, 0x6c, 0x07, 0x66, 0xbe, 0xa4, 0x3f, 0xd6, 0x9f, 0xe8,
	0x53, 0x7f, 0xa1, 0xf2, 0x25, 0x03, 0x54, 0x7d, 0x8b, 0xd7, 0x5a, 0xde, 0xde, 0x6b, 0x7b, 0x39,
	0xb0, 0x53, 0x88, 0x7a, 0x56, 0x95, 0x67, 0x52, 0x09, 0x23, 0xd0, 0xc8, 0xaf, 0x64, 0x9e, 0x7c,
	0xeb, 0xc0, 0xe0, 0xd2, 0x2d, 0xd0, 0x09, 0x74, 0x89, 0x94, 0x71, 0xe7, 0xb8, 0x73, 0xba, 0x7d,
	0xf1, 0xfb, 0x59, 0x2b, 0x39, 0x9b, 0x48, 0xe9, 0x15, 0x99, 0xe5, 0xd1, 0x39, 0x0c, 0x6b, 0x66,
	0xee, 0x85, 0x5a, 0xc4, 0x5b, 0x4e, 0x7a, 0xf4, 0x20, 0xfd, 0xe8, 0x89, 0x20, 0x6f, 0x75, 0xb6,
	0x32, 0x17, 0x65, 0xdc, 0xfd, 0xb9, 0xf2, 0x8d, 0x28, 0xdb, 0xca, 0x5c, 0x94, 0xe8, 0x14, 0x7a,
	0x94, 0x18, 0x12, 0xf7, 0x9c, 0xee, 0xe0, 0x41, 0x77, 0x45, 0x0c, 0x09, 0x42, 0xa7, 0x40, 0xaf,
	0x60, 0x44, 0xe8, 0x1d, 0xa9, 0x0b, 0x46, 0xe3, 0xbe, 0x53, 0xc7, 0x8f, 0xfa, 0x0d, 0x4c, 0xd8,
	0xb1, 0x51, 0x26, 0xdf, 0x3b, 0x10, 0xa5, 0x17, 0x69, 0xb0, 0x7b, 0x00, 0x7d, 0xcd, 0x18, 0xd5,
	0xce, 0x70, 0x94, 0xf9, 0x05, 0xfa, 0xc3, 0x56, 0xa6, 0x58, 0x32, 0xa6, 0xe2, 0xad, 0xe3, 0xee,
	0x69, 0x94, 0x0d, 0x09, 0xa5, 0x29, 0x63, 0x0a, 0xfd, 0x05, 0x91, 0x5e, 0x54, 0x12, 0x37, 0xb2,
	0x96, 0xce, 0xcb, 0x28, 0x1b, 0x59, 0xe0, 0x56, 0xd6, 0x12, 0xbd, 0x80, 0xfd, 0x39, 0xa9, 0xa9,
	0x9e, 0x93, 0x05, 0xc3, 0xa6, 0x5a, 0x32, 0xd1, 0x18, 0x67, 0x64, 0x9c, 0xfd, 0xb6, 0x21, 0xa6,
	0x1e, 0x47, 0xff, 0xc1, 0x0e, 0xad, 0x08, 0xdf, 0xe8, 0xfa, 0x4e, 0xb7, 0x6d, 0xb1, 0x56, 0xf2,
	0x0f, 0xc0, 0x1d, 0x69, 0xb8, 0xc1, 0x4b, 0x41, 0x59, 0x3c, 0x70, 0xa7, 0x45, 0x0e, 0xf9, 0x20,
	0x28, 0x43, 0x27, 0xb0, 0xcb, 0x2b, 0x6d, 0x58, 0x8d, 0x09, 0xa5, 0x8a, 0x69, 0x1d, 0x0f, 0x9d,
	0x8b, 0xb1, 0x47, 0x27, 0x1e, 0x4c, 0xbe, 0x76, 0x20, 0x9a, 0xa4, 0xef, 0x82, 0x63, 0x04, 0xbd,
	0xb9, 0xd0, 0x26, 0x18, 0x76, 0xdf, 0xd6, 0x54, 0xa9, 0x64, 0x81, 0xa5, 0x50, 0xc6, 0xdd, 0x67,
	0x94, 0x8d, 0x2c, 0x90, 0x0a, 0xe5, 0xc8, 0xb9, 0x31, 0xd2, 0x93, 0x5d, 0x4f, 0x5a, 0xc0, 0x91,
	0xff, 0xc3, 0xae, 0x23, 0x0b, 0xa1, 0xb4, 0xeb, 0x22, 0xee, 0xb9, 0x79, 0xed, 0x58, 0xf4, 0x52,
	0x28, 0x6d, 0x9b, 0x40, 0xff, 0xc2, 0x36, 0xad, 0x34, 0xc9, 0x39, 0xc3, 0x86, 0x6b, 0xe7, 0x74,
	0x94, 0x41, 0x80, 0xa6, 0xdc, 0x0d, 0xdc, 0x9e, 0x5f, 0x30, 0x65, 0x9c, 0xcd, 0x28, 0x1b, 0x2a,
	0x59, 0x5c, 0x32, 0x65, 0xd0, 0x11, 0xd8, 0x4f, 0xbc, 0x60, 0xeb, 0xe0, 0x6e, 0xa0, 0x64, 0xf1,
	0x9e, 0xad, 0x93, 0xcf, 0x30, 0x7e, 0x92, 0x34, 0x1b, 0x30, 0x79, 0xf1, 0x8b, 0xe8, 0x6e, 0x6e,
	0x3b, 0xb3, 0xbc, 0x4f, 0x78, 0x15, 0x62, 0xfb, 0x38, 0xe1, 0xed, 0x88, 0x6c, 0xc2, 0xab, 0x64,
	0x02, 0xd1, 0x26, 0x99, 0xb6, 0x09, 0x2e, 0x4a, 0x4c, 0x2b, 0x15, 0xe6, 0x36, 0xe0, 0xa2, 0xbc,
	0xaa, 0x5c, 0x1c, 0x2c, 0xc1, 0xd9, 0x1d, 0xe3, 0xed, 0xe4, 0xb8, 0x28, 0x6f, 0xec, 0x3a, 0x59,
	0x43, 0xb4, 0x79, 0x36, 0x28, 0x86, 0xa1, 0x54, 0x62, 0x56, 0x71, 0x16, 0x4a, 0xb4, 0x4b, 0x3b,
	0x9d, 0x42, 0x36, 0xb8, 0x65, 0x7d, 0x15, 0x28, 0x64, 0x93, 0x06, 0xc1, 0x39, 0x1c, 0xd6, 0xc2,
	0xa5, 0x11, 0xe7, 0x5c, 0x88, 0x25, 0x9e, 0x55, 0xdc, 0x30, 0xa5, 0x43, 0xfe, 0x50, 0x2d, 0x6c,
	0x34, 0xdf, 0x58, 0xea, 0xda, 0x33, 0x09, 0x05, 0x78, 0x78, 0x2f, 0xb6, 0x7d, 0x9a, 0x63, 0xb3,
	0x96, 0xed, 0xd9, 0x03, 0x9a, 0x4f, 0xd7, 0x92, 0xa1, 0x43, 0x18, 0xd0, 0xdc, 0xd9, 0xf2, 0xa7,
	0xf6, 0x69, 0x6e, 0x5d, 0x3d, 0x83, 0xbd, 0x7b, 0xc2, 0x39, 0x33, 0x58, 0x36, 0x39, 0x96, 0x44,
	0xeb, 0x70, 0xf1, 0x63, 0x0f, 0xa7, 0x4d, 0x9e, 0x12, 0xad, 0x93, 0x2f, 0x1d, 0xd8, 0x7d, 0xfa,
	0xd0, 0xd0, 0x73, 0xd8, 0x0f, 0x61, 0xc4, 0x25, 0x91, 0x98, 0x57, 0xcb, 0xca, 0x67, 0x6d, 0x9c,
	0xed, 0x05, 0xe2, 0x2d, 0x91, 0x37, 0x16, 0x46, 0xaf, 0xe1, 0xcf, 0x25, 0x59, 0xe1, 0xa6, 0x6e,
	0x34, 0xa3, 0x58, 0x1b, 0xb2, 0xa8, 0xea, 0x72, 0x93, 0xe5, 0x2d, 0xb7, 0xe9, 0x68, 0x49, 0x56,
	0xb7, 0x4e, 0xf0, 0xc9, 0xf3, 0x21, 0xd5, 0xe8, 0x6f, 0x00, 0xbb, 0xd9, 0xac, 0xf0, 0x8c, 0xb1,
	0x36, 0x97, 0x4b, 0xb2, 0x9a, 0xae, 0xae, 0x19, 0xcb, 0x07, 0xee, 0x17, 0xf7, 0xf2, 0x47, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfc, 0x04, 0x15, 0x60, 0xf2, 0x04, 0x00, 0x00,
}
