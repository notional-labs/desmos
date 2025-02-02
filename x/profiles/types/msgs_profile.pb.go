// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/msgs_profile.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// MsgSaveProfile represents a message to save a profile.
type MsgSaveProfile struct {
	DTag           string `protobuf:"bytes,1,opt,name=dtag,proto3" json:"dtag,omitempty" yaml:"dtag"`
	Nickname       string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty" yaml:"nickname"`
	Bio            string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty" yaml:"bio"`
	ProfilePicture string `protobuf:"bytes,4,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture" yaml:"profile_picture"`
	CoverPicture   string `protobuf:"bytes,5,opt,name=cover_picture,json=coverPicture,proto3" json:"cover_picture" yaml:"cover_picture"`
	Creator        string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *MsgSaveProfile) Reset()         { *m = MsgSaveProfile{} }
func (m *MsgSaveProfile) String() string { return proto.CompactTextString(m) }
func (*MsgSaveProfile) ProtoMessage()    {}
func (*MsgSaveProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba44ae1080fac42f, []int{0}
}
func (m *MsgSaveProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveProfile.Merge(m, src)
}
func (m *MsgSaveProfile) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveProfile proto.InternalMessageInfo

// MsgSaveProfileResponse defines the Msg/SaveProfile response type.
type MsgSaveProfileResponse struct {
}

func (m *MsgSaveProfileResponse) Reset()         { *m = MsgSaveProfileResponse{} }
func (m *MsgSaveProfileResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSaveProfileResponse) ProtoMessage()    {}
func (*MsgSaveProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba44ae1080fac42f, []int{1}
}
func (m *MsgSaveProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveProfileResponse.Merge(m, src)
}
func (m *MsgSaveProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveProfileResponse proto.InternalMessageInfo

// MsgDeleteProfile represents the message used to delete an existing profile.
type MsgDeleteProfile struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *MsgDeleteProfile) Reset()         { *m = MsgDeleteProfile{} }
func (m *MsgDeleteProfile) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProfile) ProtoMessage()    {}
func (*MsgDeleteProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba44ae1080fac42f, []int{2}
}
func (m *MsgDeleteProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProfile.Merge(m, src)
}
func (m *MsgDeleteProfile) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProfile proto.InternalMessageInfo

// MsgDeleteProfileResponse defines the Msg/DeleteProfile response type.
type MsgDeleteProfileResponse struct {
}

func (m *MsgDeleteProfileResponse) Reset()         { *m = MsgDeleteProfileResponse{} }
func (m *MsgDeleteProfileResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProfileResponse) ProtoMessage()    {}
func (*MsgDeleteProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba44ae1080fac42f, []int{3}
}
func (m *MsgDeleteProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProfileResponse.Merge(m, src)
}
func (m *MsgDeleteProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProfileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSaveProfile)(nil), "desmos.profiles.v1beta1.MsgSaveProfile")
	proto.RegisterType((*MsgSaveProfileResponse)(nil), "desmos.profiles.v1beta1.MsgSaveProfileResponse")
	proto.RegisterType((*MsgDeleteProfile)(nil), "desmos.profiles.v1beta1.MsgDeleteProfile")
	proto.RegisterType((*MsgDeleteProfileResponse)(nil), "desmos.profiles.v1beta1.MsgDeleteProfileResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/msgs_profile.proto", fileDescriptor_ba44ae1080fac42f)
}

var fileDescriptor_ba44ae1080fac42f = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x93, 0xad, 0x8c, 0x61, 0xa0, 0x43, 0x61, 0xda, 0x42, 0x0f, 0xf1, 0xe4, 0x13, 0x7f,
	0xb6, 0x5a, 0x1d, 0xb7, 0x1d, 0xa7, 0x9d, 0x40, 0x43, 0x53, 0x40, 0x1c, 0xb8, 0x54, 0x4e, 0xfa,
	0xce, 0x8b, 0x48, 0xe2, 0x10, 0xbb, 0x15, 0xfd, 0x06, 0x1c, 0xf9, 0x08, 0xfb, 0x38, 0x1c, 0x77,
	0xe4, 0x64, 0xa1, 0xf6, 0x82, 0x7a, 0x42, 0xf9, 0x04, 0x28, 0x76, 0x52, 0x68, 0x25, 0x04, 0xdc,
	0xfc, 0xbe, 0xcf, 0xef, 0x79, 0x5e, 0x27, 0x7e, 0xd1, 0xd3, 0x11, 0xc8, 0x4c, 0x48, 0x5a, 0x94,
	0xe2, 0x32, 0x49, 0x41, 0xd2, 0xc9, 0x20, 0x02, 0xc5, 0x06, 0x34, 0x93, 0x5c, 0x0e, 0x9b, 0x6e,
	0xbf, 0x28, 0x85, 0x12, 0xde, 0xbe, 0x65, 0xfb, 0x2d, 0xdb, 0x6f, 0xd8, 0xde, 0x2e, 0x17, 0x5c,
	0x18, 0x86, 0xd6, 0x27, 0x8b, 0xf7, 0x1e, 0x71, 0x21, 0x78, 0x0a, 0xd4, 0x54, 0xd1, 0xf8, 0x92,
	0xb2, 0x7c, 0xda, 0x4a, 0xb1, 0xa8, 0x93, 0x86, 0xd6, 0x63, 0x8b, 0x46, 0x3a, 0xfc, 0xe3, 0x85,
	0xc4, 0x08, 0xd2, 0xb5, 0x2b, 0xf5, 0x8e, 0xff, 0x42, 0x97, 0x90, 0x32, 0x95, 0x88, 0x5c, 0x5e,
	0x25, 0x85, 0xfc, 0x47, 0xcf, 0x48, 0x31, 0x3e, 0x2c, 0xe1, 0xc3, 0x18, 0xa4, 0x6a, 0x3c, 0xe4,
	0xc7, 0x06, 0xea, 0x9e, 0x4b, 0xfe, 0x9a, 0x4d, 0xe0, 0xc2, 0xda, 0xbc, 0x67, 0xa8, 0x53, 0x93,
	0xbe, 0x7b, 0xe0, 0x3e, 0xbe, 0x73, 0xba, 0x3f, 0xd3, 0xb8, 0x73, 0xf6, 0x86, 0xf1, 0x4a, 0xe3,
	0xbb, 0x53, 0x96, 0xa5, 0x27, 0xa4, 0x56, 0x49, 0x68, 0x20, 0x8f, 0xa2, 0xed, 0x3c, 0x89, 0xdf,
	0xe7, 0x2c, 0x03, 0x7f, 0xc3, 0x18, 0x1e, 0x56, 0x1a, 0xef, 0x58, 0xb0, 0x55, 0x48, 0xb8, 0x84,
	0xbc, 0x03, 0xb4, 0x19, 0x25, 0xc2, 0xdf, 0x34, 0x6c, 0xb7, 0xd2, 0x18, 0x59, 0x36, 0x4a, 0x04,
	0x09, 0x6b, 0xc9, 0x7b, 0x8b, 0x76, 0x9a, 0x2f, 0x18, 0x16, 0x49, 0xac, 0xc6, 0x25, 0xf8, 0x1d,
	0x43, 0x1f, 0x2d, 0x34, 0x5e, 0x97, 0x2a, 0x8d, 0xf7, 0x6c, 0xc0, 0x9a, 0x40, 0xc2, 0x6e, 0xd3,
	0xb9, 0xb0, 0x0d, 0xef, 0x15, 0xba, 0x1f, 0x8b, 0x09, 0x94, 0xcb, 0xd4, 0x5b, 0x26, 0xf5, 0xc9,
	0x42, 0xe3, 0x55, 0xa1, 0xd2, 0x78, 0xd7, 0x66, 0xae, 0xb4, 0x49, 0x78, 0xcf, 0xd4, 0x6d, 0xde,
	0x21, 0xba, 0x1d, 0x97, 0xc0, 0x94, 0x28, 0xfd, 0x2d, 0x93, 0xe4, 0x55, 0x1a, 0x77, 0x1b, 0xa3,
	0x15, 0x48, 0xd8, 0x22, 0x27, 0xdb, 0x9f, 0xae, 0xb1, 0xf3, 0xfd, 0x1a, 0x3b, 0xc4, 0x47, 0x7b,
	0xab, 0x7f, 0x3c, 0x04, 0x59, 0x88, 0x5c, 0x02, 0x79, 0x81, 0x1e, 0x9c, 0x4b, 0x7e, 0x06, 0x29,
	0xa8, 0xe5, 0x6b, 0xfc, 0x36, 0xc5, 0xfd, 0x9f, 0x29, 0x3d, 0xe4, 0xaf, 0x67, 0xb5, 0x73, 0x4e,
	0x5f, 0x7e, 0x99, 0x05, 0xee, 0xcd, 0x2c, 0x70, 0xbf, 0xcd, 0x02, 0xf7, 0xf3, 0x3c, 0x70, 0x6e,
	0xe6, 0x81, 0xf3, 0x75, 0x1e, 0x38, 0xef, 0x06, 0x3c, 0x51, 0x57, 0xe3, 0xa8, 0x1f, 0x8b, 0x8c,
	0xda, 0x6d, 0x3a, 0x4a, 0x59, 0x24, 0x9b, 0x33, 0x9d, 0x1c, 0xd3, 0x8f, 0xbf, 0xd6, 0x4b, 0x4d,
	0x0b, 0x90, 0xd1, 0x96, 0x59, 0xa4, 0xe7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x15, 0xa4, 0xa8,
	0xf6, 0x71, 0x03, 0x00, 0x00,
}

func (m *MsgSaveProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveProfile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveProfile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CoverPicture) > 0 {
		i -= len(m.CoverPicture)
		copy(dAtA[i:], m.CoverPicture)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.CoverPicture)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProfilePicture) > 0 {
		i -= len(m.ProfilePicture)
		copy(dAtA[i:], m.ProfilePicture)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.ProfilePicture)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bio) > 0 {
		i -= len(m.Bio)
		copy(dAtA[i:], m.Bio)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Bio)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DTag) > 0 {
		i -= len(m.DTag)
		copy(dAtA[i:], m.DTag)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.DTag)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSaveProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProfile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProfile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSaveProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DTag)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Bio)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.ProfilePicture)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.CoverPicture)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	return n
}

func (m *MsgSaveProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	return n
}

func (m *MsgDeleteProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsProfile(x uint64) (n int) {
	return sovMsgsProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSaveProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgSaveProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfilePicture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfilePicture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoverPicture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoverPicture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgSaveProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgSaveProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgDeleteProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgDeleteProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgDeleteProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgDeleteProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func skipMsgsProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsProfile
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
					return 0, ErrIntOverflowMsgsProfile
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
					return 0, ErrIntOverflowMsgsProfile
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
				return 0, ErrInvalidLengthMsgsProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsProfile = fmt.Errorf("proto: unexpected end of group")
)
