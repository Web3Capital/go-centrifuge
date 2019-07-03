// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/service.proto

package accountpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UpdateAccountRequest struct {
	AccountId            string       `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Data                 *AccountData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateAccountRequest) Reset()         { *m = UpdateAccountRequest{} }
func (m *UpdateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAccountRequest) ProtoMessage()    {}
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a506582bbbbe51d0, []int{0}
}

func (m *UpdateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAccountRequest.Unmarshal(m, b)
}
func (m *UpdateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAccountRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAccountRequest.Merge(m, src)
}
func (m *UpdateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAccountRequest.Size(m)
}
func (m *UpdateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAccountRequest proto.InternalMessageInfo

func (m *UpdateAccountRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *UpdateAccountRequest) GetData() *AccountData {
	if m != nil {
		return m.Data
	}
	return nil
}

type EthereumAccount struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumAccount) Reset()         { *m = EthereumAccount{} }
func (m *EthereumAccount) String() string { return proto.CompactTextString(m) }
func (*EthereumAccount) ProtoMessage()    {}
func (*EthereumAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a506582bbbbe51d0, []int{1}
}

func (m *EthereumAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthereumAccount.Unmarshal(m, b)
}
func (m *EthereumAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthereumAccount.Marshal(b, m, deterministic)
}
func (m *EthereumAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumAccount.Merge(m, src)
}
func (m *EthereumAccount) XXX_Size() int {
	return xxx_messageInfo_EthereumAccount.Size(m)
}
func (m *EthereumAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumAccount proto.InternalMessageInfo

func (m *EthereumAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EthereumAccount) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *EthereumAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type KeyPair struct {
	Pub                  string   `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
	Pvt                  string   `protobuf:"bytes,2,opt,name=pvt,proto3" json:"pvt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyPair) Reset()         { *m = KeyPair{} }
func (m *KeyPair) String() string { return proto.CompactTextString(m) }
func (*KeyPair) ProtoMessage()    {}
func (*KeyPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a506582bbbbe51d0, []int{2}
}

func (m *KeyPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyPair.Unmarshal(m, b)
}
func (m *KeyPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyPair.Marshal(b, m, deterministic)
}
func (m *KeyPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyPair.Merge(m, src)
}
func (m *KeyPair) XXX_Size() int {
	return xxx_messageInfo_KeyPair.Size(m)
}
func (m *KeyPair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyPair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyPair proto.InternalMessageInfo

func (m *KeyPair) GetPub() string {
	if m != nil {
		return m.Pub
	}
	return ""
}

func (m *KeyPair) GetPvt() string {
	if m != nil {
		return m.Pvt
	}
	return ""
}

type AccountData struct {
	EthAccount                       *EthereumAccount `protobuf:"bytes,1,opt,name=eth_account,json=ethAccount,proto3" json:"eth_account,omitempty"`
	EthDefaultAccountName            string           `protobuf:"bytes,2,opt,name=eth_default_account_name,json=ethDefaultAccountName,proto3" json:"eth_default_account_name,omitempty"`
	ReceiveEventNotificationEndpoint string           `protobuf:"bytes,3,opt,name=receive_event_notification_endpoint,json=receiveEventNotificationEndpoint,proto3" json:"receive_event_notification_endpoint,omitempty"`
	IdentityId                       string           `protobuf:"bytes,4,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`
	SigningKeyPair                   *KeyPair         `protobuf:"bytes,5,opt,name=signing_key_pair,json=signingKeyPair,proto3" json:"signing_key_pair,omitempty"`
	P2PKeyPair                       *KeyPair         `protobuf:"bytes,7,opt,name=p2p_key_pair,json=p2pKeyPair,proto3" json:"p2p_key_pair,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}         `json:"-"`
	XXX_unrecognized                 []byte           `json:"-"`
	XXX_sizecache                    int32            `json:"-"`
}

func (m *AccountData) Reset()         { *m = AccountData{} }
func (m *AccountData) String() string { return proto.CompactTextString(m) }
func (*AccountData) ProtoMessage()    {}
func (*AccountData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a506582bbbbe51d0, []int{3}
}

func (m *AccountData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountData.Unmarshal(m, b)
}
func (m *AccountData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountData.Marshal(b, m, deterministic)
}
func (m *AccountData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountData.Merge(m, src)
}
func (m *AccountData) XXX_Size() int {
	return xxx_messageInfo_AccountData.Size(m)
}
func (m *AccountData) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountData.DiscardUnknown(m)
}

var xxx_messageInfo_AccountData proto.InternalMessageInfo

func (m *AccountData) GetEthAccount() *EthereumAccount {
	if m != nil {
		return m.EthAccount
	}
	return nil
}

func (m *AccountData) GetEthDefaultAccountName() string {
	if m != nil {
		return m.EthDefaultAccountName
	}
	return ""
}

func (m *AccountData) GetReceiveEventNotificationEndpoint() string {
	if m != nil {
		return m.ReceiveEventNotificationEndpoint
	}
	return ""
}

func (m *AccountData) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *AccountData) GetSigningKeyPair() *KeyPair {
	if m != nil {
		return m.SigningKeyPair
	}
	return nil
}

func (m *AccountData) GetP2PKeyPair() *KeyPair {
	if m != nil {
		return m.P2PKeyPair
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateAccountRequest)(nil), "account.UpdateAccountRequest")
	proto.RegisterType((*EthereumAccount)(nil), "account.EthereumAccount")
	proto.RegisterType((*KeyPair)(nil), "account.KeyPair")
	proto.RegisterType((*AccountData)(nil), "account.AccountData")
}

func init() { proto.RegisterFile("account/service.proto", fileDescriptor_a506582bbbbe51d0) }

var fileDescriptor_a506582bbbbe51d0 = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0x96, 0xd3, 0xfe, 0x7f, 0xe8, 0x38, 0x2d, 0x61, 0x95, 0x4a, 0xc6, 0xa2, 0x10, 0x99, 0x4b,
	0x54, 0x91, 0x58, 0x98, 0x03, 0x22, 0x9c, 0x5a, 0x9a, 0x43, 0x85, 0xa8, 0x22, 0x23, 0x0e, 0x70,
	0xb1, 0x36, 0xf6, 0xd4, 0x59, 0xa5, 0xd9, 0x5d, 0xec, 0x8d, 0xab, 0x08, 0x71, 0x28, 0x8f, 0x00,
	0x8f, 0xc6, 0x13, 0x20, 0xf1, 0x20, 0xc8, 0xeb, 0x75, 0x12, 0x68, 0x7a, 0xf2, 0x7a, 0xe6, 0xfb,
	0xbe, 0xd9, 0x6f, 0x76, 0x06, 0x0e, 0x69, 0x1c, 0x8b, 0x05, 0x57, 0x7e, 0x8e, 0x59, 0xc1, 0x62,
	0x1c, 0xc8, 0x4c, 0x28, 0x41, 0x9a, 0x26, 0xec, 0x3e, 0x4a, 0x85, 0x48, 0xaf, 0xd0, 0xa7, 0x92,
	0xf9, 0x94, 0x73, 0xa1, 0xa8, 0x62, 0x82, 0xe7, 0x15, 0xcc, 0x7d, 0xa6, 0x3f, 0x71, 0x3f, 0x45,
	0xde, 0xcf, 0xaf, 0x69, 0x9a, 0x62, 0xe6, 0x0b, 0xa9, 0x11, 0xb7, 0xd1, 0x5e, 0x04, 0x9d, 0x0f,
	0x32, 0xa1, 0x0a, 0x4f, 0x2a, 0xf1, 0x10, 0x3f, 0x2f, 0x30, 0x57, 0xe4, 0x08, 0xc0, 0x94, 0x8b,
	0x58, 0xe2, 0x58, 0x5d, 0xab, 0xb7, 0x17, 0xee, 0x99, 0xc8, 0x79, 0x42, 0x7a, 0xb0, 0x9b, 0x50,
	0x45, 0x9d, 0x46, 0xd7, 0xea, 0xd9, 0x41, 0x67, 0x60, 0x32, 0x03, 0xa3, 0x72, 0x46, 0x15, 0x0d,
	0x35, 0xc2, 0xfb, 0x08, 0xf7, 0x47, 0x6a, 0x8a, 0x19, 0x2e, 0xe6, 0x26, 0x49, 0x1c, 0x68, 0xd2,
	0x24, 0xc9, 0x30, 0xcf, 0x8d, 0x70, 0xfd, 0x4b, 0xda, 0xb0, 0x33, 0xc3, 0xa5, 0x56, 0xdd, 0x0b,
	0xcb, 0x23, 0x71, 0xe1, 0x9e, 0xa4, 0x79, 0x7e, 0x2d, 0xb2, 0xc4, 0xd9, 0xd1, 0xe1, 0xd5, 0xbf,
	0xd7, 0x87, 0xe6, 0x5b, 0x5c, 0x8e, 0x29, 0xcb, 0x4a, 0xa2, 0x5c, 0x4c, 0x8c, 0x5c, 0x79, 0xd4,
	0x91, 0x42, 0xd5, 0x52, 0xb2, 0x50, 0xde, 0xaf, 0x06, 0xd8, 0x1b, 0xf7, 0x23, 0xaf, 0xc0, 0x46,
	0x35, 0x8d, 0xcc, 0xd5, 0x35, 0xd7, 0x0e, 0x9c, 0x95, 0x95, 0x7f, 0x6e, 0x1d, 0x02, 0xaa, 0x69,
	0xed, 0xe0, 0x25, 0x38, 0x25, 0x35, 0xc1, 0x4b, 0xba, 0xb8, 0x52, 0xb5, 0x44, 0xc4, 0xe9, 0x1c,
	0x4d, 0xc5, 0x43, 0x54, 0xd3, 0xb3, 0x2a, 0x6d, 0x48, 0x17, 0x74, 0x8e, 0xe4, 0x1d, 0x3c, 0xcd,
	0x30, 0x46, 0x56, 0x60, 0x84, 0x05, 0x96, 0x14, 0xa1, 0xd8, 0x25, 0x8b, 0xf5, 0x9b, 0x44, 0xc8,
	0x13, 0x29, 0x18, 0x57, 0xc6, 0x69, 0xd7, 0x40, 0x47, 0x25, 0xf2, 0x62, 0x03, 0x38, 0x32, 0x38,
	0xf2, 0x04, 0x6c, 0x96, 0x20, 0x57, 0x4c, 0x2d, 0xcb, 0x67, 0xda, 0xd5, 0x34, 0xa8, 0x43, 0xe7,
	0x09, 0x19, 0x42, 0x3b, 0x67, 0x29, 0x67, 0x3c, 0x8d, 0x66, 0xb8, 0x8c, 0x24, 0x65, 0x99, 0xf3,
	0x9f, 0x36, 0xda, 0x5e, 0x19, 0x35, 0x3d, 0x0c, 0x0f, 0x0c, 0xb2, 0xee, 0x69, 0x00, 0x2d, 0x19,
	0xc8, 0x35, 0xaf, 0x79, 0x07, 0x0f, 0x64, 0x20, 0xcd, 0x39, 0xb8, 0x69, 0xc0, 0x81, 0xf1, 0xfb,
	0xbe, 0x1a, 0x5e, 0x32, 0x83, 0xfd, 0x37, 0x19, 0xae, 0x27, 0x8c, 0x6c, 0x9d, 0x16, 0x77, 0x6b,
	0xd4, 0x1b, 0x7c, 0x3f, 0xe9, 0xb8, 0xa4, 0xe2, 0xe7, 0x5d, 0xca, 0xbb, 0x26, 0xf9, 0xed, 0xe7,
	0xef, 0x1f, 0x8d, 0x07, 0x5e, 0xcb, 0x2f, 0x9e, 0xfb, 0x86, 0x97, 0x0f, 0xad, 0x63, 0x72, 0x63,
	0xc1, 0xfe, 0x5f, 0xf3, 0x4c, 0x8e, 0x56, 0xba, 0xdb, 0xe6, 0xfc, 0x8e, 0xb2, 0xaf, 0x75, 0xd9,
	0x8a, 0x70, 0xab, 0xec, 0x63, 0xf7, 0xe1, 0x66, 0x59, 0xff, 0xcb, 0x7a, 0x51, 0xbe, 0x0e, 0xad,
	0xe3, 0xd3, 0x1e, 0xd8, 0xb1, 0x98, 0xd7, 0xba, 0xa7, 0x2d, 0xd3, 0x88, 0x71, 0xb9, 0x6f, 0x63,
	0xeb, 0x53, 0xbd, 0x45, 0x72, 0x32, 0xf9, 0x5f, 0xef, 0xe0, 0x8b, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x17, 0x9b, 0xdd, 0x8c, 0xf1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountServiceClient interface {
	CreateAccount(ctx context.Context, in *AccountData, opts ...grpc.CallOption) (*AccountData, error)
	UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*AccountData, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *AccountData, opts ...grpc.CallOption) (*AccountData, error) {
	out := new(AccountData)
	err := c.cc.Invoke(ctx, "/account.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *UpdateAccountRequest, opts ...grpc.CallOption) (*AccountData, error) {
	out := new(AccountData)
	err := c.cc.Invoke(ctx, "/account.AccountService/UpdateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
type AccountServiceServer interface {
	CreateAccount(context.Context, *AccountData) (*AccountData, error)
	UpdateAccount(context.Context, *UpdateAccountRequest) (*AccountData, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*AccountData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*UpdateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/service.proto",
}
