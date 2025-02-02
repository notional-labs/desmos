// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/msg_server.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/msg_server.proto", fileDescriptor_e042115035680605)
}

var fileDescriptor_e042115035680605 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x5b, 0x21, 0x21, 0xe1, 0x69, 0x02, 0x0c, 0x12, 0x5a, 0x0e, 0xb9, 0x21, 0xfe, 0xad,
	0x09, 0xed, 0x18, 0x9c, 0x38, 0x74, 0xdd, 0x0d, 0x26, 0xa1, 0xb2, 0x5e, 0xb8, 0x54, 0x4e, 0xfa,
	0xd6, 0x8d, 0xea, 0xda, 0x21, 0x76, 0x2b, 0x2a, 0x21, 0xf1, 0x15, 0xe0, 0x5b, 0xed, 0xb8, 0x23,
	0x47, 0xd4, 0x7e, 0x11, 0x94, 0xa4, 0x35, 0xee, 0x9f, 0xb8, 0x09, 0xb7, 0xb8, 0xfe, 0x3d, 0xef,
	0xf3, 0xf8, 0xb5, 0x6b, 0xa3, 0xe7, 0x03, 0x90, 0x13, 0x21, 0xfd, 0x38, 0x11, 0xc3, 0x88, 0x81,
	0xf4, 0x67, 0xcd, 0x00, 0x14, 0x69, 0xfa, 0x13, 0x49, 0xfb, 0x12, 0x92, 0x19, 0x24, 0x5e, 0x9c,
	0x08, 0x25, 0xf0, 0x93, 0x9c, 0xf4, 0xd6, 0xa4, 0xb7, 0x22, 0x9d, 0xc7, 0x54, 0x50, 0x91, 0x31,
	0x7e, 0xfa, 0x95, 0xe3, 0xce, 0x09, 0x15, 0x82, 0x32, 0xf0, 0xb3, 0x51, 0x30, 0x1d, 0xfa, 0x84,
	0xcf, 0xd7, 0x53, 0xa1, 0x48, 0x2b, 0xf5, 0x73, 0x4d, 0x3e, 0x58, 0x4d, 0x9d, 0x16, 0xc6, 0x11,
	0x03, 0x60, 0x99, 0x24, 0xfd, 0x7d, 0x45, 0xb7, 0x0e, 0xd0, 0x09, 0x30, 0xa2, 0x22, 0xc1, 0xe5,
	0x28, 0x8a, 0x65, 0x49, 0xcd, 0x40, 0x11, 0xda, 0x4f, 0xe0, 0xeb, 0x14, 0xa4, 0x5a, 0x6b, 0x5e,
	0x5a, 0x9a, 0xb4, 0x9d, 0xe9, 0xb5, 0x95, 0xdd, 0x97, 0xc8, 0xae, 0xd8, 0x97, 0xc7, 0xb3, 0x2a,
	0xc2, 0x11, 0x89, 0x78, 0x9f, 0x45, 0x7c, 0x7c, 0xb8, 0xab, 0x29, 0x4f, 0xe2, 0xd8, 0xa4, 0x5b,
	0x37, 0xc7, 0xe8, 0xce, 0x95, 0xa4, 0x98, 0xa2, 0xa3, 0xcf, 0x64, 0x06, 0x9f, 0x72, 0x11, 0x7e,
	0xe6, 0x15, 0x1c, 0x00, 0xef, 0x4a, 0x52, 0x03, 0x74, 0xfc, 0x92, 0x60, 0x17, 0x64, 0x2c, 0xb8,
	0x04, 0x3c, 0x41, 0xc7, 0x97, 0xc0, 0x40, 0x69, 0xab, 0x17, 0xb6, 0x0a, 0x1b, 0xa8, 0xd3, 0x2c,
	0x8d, 0x6a, 0xbb, 0x1f, 0xe8, 0x51, 0x37, 0xef, 0xe7, 0xe5, 0x35, 0xa1, 0xd7, 0x09, 0xe1, 0x72,
	0x08, 0x09, 0xb6, 0xc6, 0xde, 0x23, 0x70, 0xde, 0x55, 0x14, 0xe8, 0x00, 0xbf, 0xea, 0xe8, 0xa4,
	0x43, 0x78, 0x08, 0x6c, 0x73, 0x3a, 0x53, 0xe0, 0x73, 0x5b, 0xd9, 0x42, 0x99, 0xf3, 0xfe, 0xbf,
	0x64, 0x1b, 0x99, 0xda, 0x61, 0x08, 0xb1, 0xaa, 0x9c, 0xa9, 0x50, 0x66, 0xcf, 0x54, 0x28, 0xdb,
	0xc8, 0xd4, 0x85, 0xe1, 0x54, 0x42, 0xe5, 0x4c, 0x85, 0x32, 0x7b, 0xa6, 0x42, 0x99, 0xce, 0xf4,
	0x1d, 0xe1, 0x4e, 0x02, 0x44, 0x41, 0xd7, 0xf8, 0x27, 0x63, 0xcf, 0xda, 0xfc, 0x1d, 0xde, 0x79,
	0x5b, 0x8d, 0x37, 0xdd, 0xf3, 0x33, 0x5d, 0xde, 0x7d, 0x97, 0xb7, 0xbb, 0xef, 0xf2, 0xda, 0x9d,
	0xa0, 0x7b, 0x17, 0x4c, 0x84, 0xe3, 0x9e, 0x84, 0x04, 0x3f, 0xb5, 0x15, 0xd1, 0x98, 0xd3, 0x28,
	0x85, 0x69, 0x0b, 0x8a, 0x8e, 0x7a, 0x3c, 0xd0, 0x26, 0xd6, 0x3b, 0xc7, 0x00, 0xed, 0x77, 0x8e,
	0x01, 0x6a, 0xa3, 0x19, 0x7a, 0xf0, 0x31, 0xe2, 0xe3, 0x4e, 0x7a, 0x57, 0xb6, 0xc3, 0x50, 0x4c,
	0xb9, 0xc2, 0xa7, 0xb6, 0x22, 0xdb, 0xb4, 0xf3, 0xa6, 0x0a, 0x6d, 0xee, 0x60, 0x8f, 0xb3, 0x6d,
	0x67, 0xcf, 0x1e, 0x7f, 0x9b, 0xb7, 0xef, 0xe0, 0x2e, 0xaf, 0xdd, 0x25, 0xba, 0x9f, 0x26, 0x6b,
	0xc7, 0x31, 0x8b, 0xc2, 0x6c, 0x83, 0xf1, 0xab, 0x43, 0xcb, 0x30, 0x60, 0xe7, 0xac, 0x02, 0xac,
	0x4d, 0xe7, 0xe8, 0x61, 0x1e, 0xc9, 0xb4, 0x6d, 0x1c, 0x5e, 0x81, 0x69, 0x7c, 0x5e, 0x09, 0x5f,
	0x5b, 0x5f, 0x7c, 0xb8, 0x59, 0xb8, 0xf5, 0xdb, 0x85, 0x5b, 0xff, 0xb3, 0x70, 0xeb, 0x3f, 0x97,
	0x6e, 0xed, 0x76, 0xe9, 0xd6, 0x7e, 0x2f, 0xdd, 0xda, 0x97, 0x26, 0x8d, 0xd4, 0x68, 0x1a, 0x78,
	0xa1, 0x98, 0xf8, 0x79, 0xe9, 0x06, 0x23, 0x81, 0x5c, 0x7d, 0xfb, 0xb3, 0x96, 0xff, 0xed, 0xdf,
	0x73, 0xa9, 0xe6, 0x31, 0xc8, 0xe0, 0x6e, 0xf6, 0x3c, 0x9e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0x9a, 0x07, 0x97, 0x33, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SaveProfile defines the method to save a profile
	SaveProfile(ctx context.Context, in *MsgSaveProfile, opts ...grpc.CallOption) (*MsgSaveProfileResponse, error)
	// DeleteProfile defines the method to delete an existing profile
	DeleteProfile(ctx context.Context, in *MsgDeleteProfile, opts ...grpc.CallOption) (*MsgDeleteProfileResponse, error)
	// RequestDTagTransfer defines the method to request another user to transfer
	// their DTag to you
	RequestDTagTransfer(ctx context.Context, in *MsgRequestDTagTransfer, opts ...grpc.CallOption) (*MsgRequestDTagTransferResponse, error)
	// CancelDTagTransferRequest defines the method to cancel an outgoing DTag
	// transfer request
	CancelDTagTransferRequest(ctx context.Context, in *MsgCancelDTagTransferRequest, opts ...grpc.CallOption) (*MsgCancelDTagTransferRequestResponse, error)
	// AcceptDTagTransferRequest defines the method to accept an incoming DTag
	// transfer request
	AcceptDTagTransferRequest(ctx context.Context, in *MsgAcceptDTagTransferRequest, opts ...grpc.CallOption) (*MsgAcceptDTagTransferRequestResponse, error)
	// RefuseDTagTransferRequest defines the method to refuse an incoming DTag
	// transfer request
	RefuseDTagTransferRequest(ctx context.Context, in *MsgRefuseDTagTransferRequest, opts ...grpc.CallOption) (*MsgRefuseDTagTransferRequestResponse, error)
	// CreateRelationship defines a method for creating a new relationship
	CreateRelationship(ctx context.Context, in *MsgCreateRelationship, opts ...grpc.CallOption) (*MsgCreateRelationshipResponse, error)
	// DeleteRelationship defines a method for deleting a relationship
	DeleteRelationship(ctx context.Context, in *MsgDeleteRelationship, opts ...grpc.CallOption) (*MsgDeleteRelationshipResponse, error)
	// BlockUser defines a method for blocking a user
	BlockUser(ctx context.Context, in *MsgBlockUser, opts ...grpc.CallOption) (*MsgBlockUserResponse, error)
	// UnblockUser defines a method for unblocking a user
	UnblockUser(ctx context.Context, in *MsgUnblockUser, opts ...grpc.CallOption) (*MsgUnblockUserResponse, error)
	// LinkChainAccount defines a method to link an external chain account to a
	// profile
	LinkChainAccount(ctx context.Context, in *MsgLinkChainAccount, opts ...grpc.CallOption) (*MsgLinkChainAccountResponse, error)
	// UnlinkChainAccount defines a method to unlink an external chain account
	// from a profile
	UnlinkChainAccount(ctx context.Context, in *MsgUnlinkChainAccount, opts ...grpc.CallOption) (*MsgUnlinkChainAccountResponse, error)
	// LinkApplication defines a method to create a centralized application
	// link
	LinkApplication(ctx context.Context, in *MsgLinkApplication, opts ...grpc.CallOption) (*MsgLinkApplicationResponse, error)
	// UnlinkApplication defines a method to remove a centralized application
	UnlinkApplication(ctx context.Context, in *MsgUnlinkApplication, opts ...grpc.CallOption) (*MsgUnlinkApplicationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SaveProfile(ctx context.Context, in *MsgSaveProfile, opts ...grpc.CallOption) (*MsgSaveProfileResponse, error) {
	out := new(MsgSaveProfileResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/SaveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteProfile(ctx context.Context, in *MsgDeleteProfile, opts ...grpc.CallOption) (*MsgDeleteProfileResponse, error) {
	out := new(MsgDeleteProfileResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/DeleteProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestDTagTransfer(ctx context.Context, in *MsgRequestDTagTransfer, opts ...grpc.CallOption) (*MsgRequestDTagTransferResponse, error) {
	out := new(MsgRequestDTagTransferResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/RequestDTagTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelDTagTransferRequest(ctx context.Context, in *MsgCancelDTagTransferRequest, opts ...grpc.CallOption) (*MsgCancelDTagTransferRequestResponse, error) {
	out := new(MsgCancelDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/CancelDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptDTagTransferRequest(ctx context.Context, in *MsgAcceptDTagTransferRequest, opts ...grpc.CallOption) (*MsgAcceptDTagTransferRequestResponse, error) {
	out := new(MsgAcceptDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/AcceptDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RefuseDTagTransferRequest(ctx context.Context, in *MsgRefuseDTagTransferRequest, opts ...grpc.CallOption) (*MsgRefuseDTagTransferRequestResponse, error) {
	out := new(MsgRefuseDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/RefuseDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateRelationship(ctx context.Context, in *MsgCreateRelationship, opts ...grpc.CallOption) (*MsgCreateRelationshipResponse, error) {
	out := new(MsgCreateRelationshipResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/CreateRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteRelationship(ctx context.Context, in *MsgDeleteRelationship, opts ...grpc.CallOption) (*MsgDeleteRelationshipResponse, error) {
	out := new(MsgDeleteRelationshipResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/DeleteRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BlockUser(ctx context.Context, in *MsgBlockUser, opts ...grpc.CallOption) (*MsgBlockUserResponse, error) {
	out := new(MsgBlockUserResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnblockUser(ctx context.Context, in *MsgUnblockUser, opts ...grpc.CallOption) (*MsgUnblockUserResponse, error) {
	out := new(MsgUnblockUserResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/UnblockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LinkChainAccount(ctx context.Context, in *MsgLinkChainAccount, opts ...grpc.CallOption) (*MsgLinkChainAccountResponse, error) {
	out := new(MsgLinkChainAccountResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/LinkChainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnlinkChainAccount(ctx context.Context, in *MsgUnlinkChainAccount, opts ...grpc.CallOption) (*MsgUnlinkChainAccountResponse, error) {
	out := new(MsgUnlinkChainAccountResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/UnlinkChainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LinkApplication(ctx context.Context, in *MsgLinkApplication, opts ...grpc.CallOption) (*MsgLinkApplicationResponse, error) {
	out := new(MsgLinkApplicationResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/LinkApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnlinkApplication(ctx context.Context, in *MsgUnlinkApplication, opts ...grpc.CallOption) (*MsgUnlinkApplicationResponse, error) {
	out := new(MsgUnlinkApplicationResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Msg/UnlinkApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SaveProfile defines the method to save a profile
	SaveProfile(context.Context, *MsgSaveProfile) (*MsgSaveProfileResponse, error)
	// DeleteProfile defines the method to delete an existing profile
	DeleteProfile(context.Context, *MsgDeleteProfile) (*MsgDeleteProfileResponse, error)
	// RequestDTagTransfer defines the method to request another user to transfer
	// their DTag to you
	RequestDTagTransfer(context.Context, *MsgRequestDTagTransfer) (*MsgRequestDTagTransferResponse, error)
	// CancelDTagTransferRequest defines the method to cancel an outgoing DTag
	// transfer request
	CancelDTagTransferRequest(context.Context, *MsgCancelDTagTransferRequest) (*MsgCancelDTagTransferRequestResponse, error)
	// AcceptDTagTransferRequest defines the method to accept an incoming DTag
	// transfer request
	AcceptDTagTransferRequest(context.Context, *MsgAcceptDTagTransferRequest) (*MsgAcceptDTagTransferRequestResponse, error)
	// RefuseDTagTransferRequest defines the method to refuse an incoming DTag
	// transfer request
	RefuseDTagTransferRequest(context.Context, *MsgRefuseDTagTransferRequest) (*MsgRefuseDTagTransferRequestResponse, error)
	// CreateRelationship defines a method for creating a new relationship
	CreateRelationship(context.Context, *MsgCreateRelationship) (*MsgCreateRelationshipResponse, error)
	// DeleteRelationship defines a method for deleting a relationship
	DeleteRelationship(context.Context, *MsgDeleteRelationship) (*MsgDeleteRelationshipResponse, error)
	// BlockUser defines a method for blocking a user
	BlockUser(context.Context, *MsgBlockUser) (*MsgBlockUserResponse, error)
	// UnblockUser defines a method for unblocking a user
	UnblockUser(context.Context, *MsgUnblockUser) (*MsgUnblockUserResponse, error)
	// LinkChainAccount defines a method to link an external chain account to a
	// profile
	LinkChainAccount(context.Context, *MsgLinkChainAccount) (*MsgLinkChainAccountResponse, error)
	// UnlinkChainAccount defines a method to unlink an external chain account
	// from a profile
	UnlinkChainAccount(context.Context, *MsgUnlinkChainAccount) (*MsgUnlinkChainAccountResponse, error)
	// LinkApplication defines a method to create a centralized application
	// link
	LinkApplication(context.Context, *MsgLinkApplication) (*MsgLinkApplicationResponse, error)
	// UnlinkApplication defines a method to remove a centralized application
	UnlinkApplication(context.Context, *MsgUnlinkApplication) (*MsgUnlinkApplicationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SaveProfile(ctx context.Context, req *MsgSaveProfile) (*MsgSaveProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProfile not implemented")
}
func (*UnimplementedMsgServer) DeleteProfile(ctx context.Context, req *MsgDeleteProfile) (*MsgDeleteProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (*UnimplementedMsgServer) RequestDTagTransfer(ctx context.Context, req *MsgRequestDTagTransfer) (*MsgRequestDTagTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDTagTransfer not implemented")
}
func (*UnimplementedMsgServer) CancelDTagTransferRequest(ctx context.Context, req *MsgCancelDTagTransferRequest) (*MsgCancelDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) AcceptDTagTransferRequest(ctx context.Context, req *MsgAcceptDTagTransferRequest) (*MsgAcceptDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) RefuseDTagTransferRequest(ctx context.Context, req *MsgRefuseDTagTransferRequest) (*MsgRefuseDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefuseDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) CreateRelationship(ctx context.Context, req *MsgCreateRelationship) (*MsgCreateRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRelationship not implemented")
}
func (*UnimplementedMsgServer) DeleteRelationship(ctx context.Context, req *MsgDeleteRelationship) (*MsgDeleteRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationship not implemented")
}
func (*UnimplementedMsgServer) BlockUser(ctx context.Context, req *MsgBlockUser) (*MsgBlockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (*UnimplementedMsgServer) UnblockUser(ctx context.Context, req *MsgUnblockUser) (*MsgUnblockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockUser not implemented")
}
func (*UnimplementedMsgServer) LinkChainAccount(ctx context.Context, req *MsgLinkChainAccount) (*MsgLinkChainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkChainAccount not implemented")
}
func (*UnimplementedMsgServer) UnlinkChainAccount(ctx context.Context, req *MsgUnlinkChainAccount) (*MsgUnlinkChainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkChainAccount not implemented")
}
func (*UnimplementedMsgServer) LinkApplication(ctx context.Context, req *MsgLinkApplication) (*MsgLinkApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkApplication not implemented")
}
func (*UnimplementedMsgServer) UnlinkApplication(ctx context.Context, req *MsgUnlinkApplication) (*MsgUnlinkApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkApplication not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SaveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSaveProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SaveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/SaveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SaveProfile(ctx, req.(*MsgSaveProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/DeleteProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteProfile(ctx, req.(*MsgDeleteProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestDTagTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestDTagTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestDTagTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/RequestDTagTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestDTagTransfer(ctx, req.(*MsgRequestDTagTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/CancelDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelDTagTransferRequest(ctx, req.(*MsgCancelDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/AcceptDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptDTagTransferRequest(ctx, req.(*MsgAcceptDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RefuseDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRefuseDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RefuseDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/RefuseDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RefuseDTagTransferRequest(ctx, req.(*MsgRefuseDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/CreateRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRelationship(ctx, req.(*MsgCreateRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/DeleteRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteRelationship(ctx, req.(*MsgDeleteRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBlockUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BlockUser(ctx, req.(*MsgBlockUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnblockUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/UnblockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnblockUser(ctx, req.(*MsgUnblockUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LinkChainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLinkChainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LinkChainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/LinkChainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LinkChainAccount(ctx, req.(*MsgLinkChainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnlinkChainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnlinkChainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnlinkChainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/UnlinkChainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnlinkChainAccount(ctx, req.(*MsgUnlinkChainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LinkApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLinkApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LinkApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/LinkApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LinkApplication(ctx, req.(*MsgLinkApplication))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnlinkApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnlinkApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnlinkApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Msg/UnlinkApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnlinkApplication(ctx, req.(*MsgUnlinkApplication))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.profiles.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveProfile",
			Handler:    _Msg_SaveProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _Msg_DeleteProfile_Handler,
		},
		{
			MethodName: "RequestDTagTransfer",
			Handler:    _Msg_RequestDTagTransfer_Handler,
		},
		{
			MethodName: "CancelDTagTransferRequest",
			Handler:    _Msg_CancelDTagTransferRequest_Handler,
		},
		{
			MethodName: "AcceptDTagTransferRequest",
			Handler:    _Msg_AcceptDTagTransferRequest_Handler,
		},
		{
			MethodName: "RefuseDTagTransferRequest",
			Handler:    _Msg_RefuseDTagTransferRequest_Handler,
		},
		{
			MethodName: "CreateRelationship",
			Handler:    _Msg_CreateRelationship_Handler,
		},
		{
			MethodName: "DeleteRelationship",
			Handler:    _Msg_DeleteRelationship_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _Msg_BlockUser_Handler,
		},
		{
			MethodName: "UnblockUser",
			Handler:    _Msg_UnblockUser_Handler,
		},
		{
			MethodName: "LinkChainAccount",
			Handler:    _Msg_LinkChainAccount_Handler,
		},
		{
			MethodName: "UnlinkChainAccount",
			Handler:    _Msg_UnlinkChainAccount_Handler,
		},
		{
			MethodName: "LinkApplication",
			Handler:    _Msg_LinkApplication_Handler,
		},
		{
			MethodName: "UnlinkApplication",
			Handler:    _Msg_UnlinkApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/profiles/v1beta1/msg_server.proto",
}
