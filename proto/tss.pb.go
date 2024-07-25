// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: tss.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 키 생성 요청
type KeygenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Threshold    int32      `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty"`
	TotalParties int32      `protobuf:"varint,2,opt,name=total_parties,json=totalParties,proto3" json:"total_parties,omitempty"`
	PartyId      string     `protobuf:"bytes,3,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Parties      []*PartyID `protobuf:"bytes,4,rep,name=parties,proto3" json:"parties,omitempty"`
}

func (x *KeygenRequest) Reset() {
	*x = KeygenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeygenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeygenRequest) ProtoMessage() {}

func (x *KeygenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeygenRequest.ProtoReflect.Descriptor instead.
func (*KeygenRequest) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{0}
}

func (x *KeygenRequest) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *KeygenRequest) GetTotalParties() int32 {
	if x != nil {
		return x.TotalParties
	}
	return 0
}

func (x *KeygenRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *KeygenRequest) GetParties() []*PartyID {
	if x != nil {
		return x.Parties
	}
	return nil
}

// 키 생성 응답
type KeygenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	PublicKey    string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *KeygenResponse) Reset() {
	*x = KeygenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeygenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeygenResponse) ProtoMessage() {}

func (x *KeygenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeygenResponse.ProtoReflect.Descriptor instead.
func (*KeygenResponse) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{1}
}

func (x *KeygenResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *KeygenResponse) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *KeygenResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Party ID
type PartyID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Moniker string `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Key     []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PartyID) Reset() {
	*x = PartyID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartyID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartyID) ProtoMessage() {}

func (x *PartyID) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartyID.ProtoReflect.Descriptor instead.
func (*PartyID) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{2}
}

func (x *PartyID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PartyID) GetMoniker() string {
	if x != nil {
		return x.Moniker
	}
	return ""
}

func (x *PartyID) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

// 메시지 래퍼
type MessageWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From        *PartyID   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To          []*PartyID `protobuf:"bytes,2,rep,name=to,proto3" json:"to,omitempty"`
	IsBroadcast bool       `protobuf:"varint,3,opt,name=is_broadcast,json=isBroadcast,proto3" json:"is_broadcast,omitempty"`
	SessionId   string     `protobuf:"bytes,4,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Content     *anypb.Any `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MessageWrapper) Reset() {
	*x = MessageWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWrapper) ProtoMessage() {}

func (x *MessageWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWrapper.ProtoReflect.Descriptor instead.
func (*MessageWrapper) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{3}
}

func (x *MessageWrapper) GetFrom() *PartyID {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *MessageWrapper) GetTo() []*PartyID {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MessageWrapper) GetIsBroadcast() bool {
	if x != nil {
		return x.IsBroadcast
	}
	return false
}

func (x *MessageWrapper) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *MessageWrapper) GetContent() *anypb.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

// 메시지 응답
type MessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{4}
}

func (x *MessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// 키 생성 라운드 메시지들
type KGRound1Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Commitment []byte   `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	PaillierN  []byte   `protobuf:"bytes,2,opt,name=paillier_n,json=paillierN,proto3" json:"paillier_n,omitempty"`
	NTilde     []byte   `protobuf:"bytes,3,opt,name=n_tilde,json=nTilde,proto3" json:"n_tilde,omitempty"`
	H1         []byte   `protobuf:"bytes,4,opt,name=h1,proto3" json:"h1,omitempty"`
	H2         []byte   `protobuf:"bytes,5,opt,name=h2,proto3" json:"h2,omitempty"`
	Dlnproof_1 [][]byte `protobuf:"bytes,6,rep,name=dlnproof_1,json=dlnproof1,proto3" json:"dlnproof_1,omitempty"`
	Dlnproof_2 [][]byte `protobuf:"bytes,7,rep,name=dlnproof_2,json=dlnproof2,proto3" json:"dlnproof_2,omitempty"`
}

func (x *KGRound1Message) Reset() {
	*x = KGRound1Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound1Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound1Message) ProtoMessage() {}

func (x *KGRound1Message) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound1Message.ProtoReflect.Descriptor instead.
func (*KGRound1Message) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{5}
}

func (x *KGRound1Message) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *KGRound1Message) GetPaillierN() []byte {
	if x != nil {
		return x.PaillierN
	}
	return nil
}

func (x *KGRound1Message) GetNTilde() []byte {
	if x != nil {
		return x.NTilde
	}
	return nil
}

func (x *KGRound1Message) GetH1() []byte {
	if x != nil {
		return x.H1
	}
	return nil
}

func (x *KGRound1Message) GetH2() []byte {
	if x != nil {
		return x.H2
	}
	return nil
}

func (x *KGRound1Message) GetDlnproof_1() [][]byte {
	if x != nil {
		return x.Dlnproof_1
	}
	return nil
}

func (x *KGRound1Message) GetDlnproof_2() [][]byte {
	if x != nil {
		return x.Dlnproof_2
	}
	return nil
}

type KGRound2Message1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Share    []byte   `protobuf:"bytes,1,opt,name=share,proto3" json:"share,omitempty"`
	FacProof [][]byte `protobuf:"bytes,2,rep,name=fac_proof,json=facProof,proto3" json:"fac_proof,omitempty"`
}

func (x *KGRound2Message1) Reset() {
	*x = KGRound2Message1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound2Message1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound2Message1) ProtoMessage() {}

func (x *KGRound2Message1) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound2Message1.ProtoReflect.Descriptor instead.
func (*KGRound2Message1) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{6}
}

func (x *KGRound2Message1) GetShare() []byte {
	if x != nil {
		return x.Share
	}
	return nil
}

func (x *KGRound2Message1) GetFacProof() [][]byte {
	if x != nil {
		return x.FacProof
	}
	return nil
}

type KGRound2Message2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeCommitment [][]byte `protobuf:"bytes,1,rep,name=de_commitment,json=deCommitment,proto3" json:"de_commitment,omitempty"`
	ModProof     [][]byte `protobuf:"bytes,2,rep,name=mod_proof,json=modProof,proto3" json:"mod_proof,omitempty"`
}

func (x *KGRound2Message2) Reset() {
	*x = KGRound2Message2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound2Message2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound2Message2) ProtoMessage() {}

func (x *KGRound2Message2) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound2Message2.ProtoReflect.Descriptor instead.
func (*KGRound2Message2) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{7}
}

func (x *KGRound2Message2) GetDeCommitment() [][]byte {
	if x != nil {
		return x.DeCommitment
	}
	return nil
}

func (x *KGRound2Message2) GetModProof() [][]byte {
	if x != nil {
		return x.ModProof
	}
	return nil
}

type KGRound3Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaillierProof [][]byte `protobuf:"bytes,1,rep,name=paillier_proof,json=paillierProof,proto3" json:"paillier_proof,omitempty"`
}

func (x *KGRound3Message) Reset() {
	*x = KGRound3Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tss_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KGRound3Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KGRound3Message) ProtoMessage() {}

func (x *KGRound3Message) ProtoReflect() protoreflect.Message {
	mi := &file_tss_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KGRound3Message.ProtoReflect.Descriptor instead.
func (*KGRound3Message) Descriptor() ([]byte, []int) {
	return file_tss_proto_rawDescGZIP(), []int{8}
}

func (x *KGRound3Message) GetPaillierProof() [][]byte {
	if x != nil {
		return x.PaillierProof
	}
	return nil
}

var File_tss_proto protoreflect.FileDescriptor

var file_tss_proto_rawDesc = []byte{
	0x0a, 0x09, 0x74, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x74, 0x73, 0x73,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x0d,
	0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74,
	0x73, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x6b, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xc2, 0x01, 0x0a, 0x0e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x73,
	0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x1c, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x73,
	0x73, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x49, 0x44, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x21, 0x0a,
	0x0c, 0x69, 0x73, 0x5f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0x2b, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xc7, 0x01, 0x0a,
	0x0f, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x31, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x4e, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x5f, 0x74, 0x69, 0x6c, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x6e, 0x54, 0x69, 0x6c, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x31, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x68, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x32, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x68, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6c, 0x6e, 0x70,
	0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x31, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x6c,
	0x6e, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x31, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6c, 0x6e, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x5f, 0x32, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x64, 0x6c, 0x6e,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x32, 0x22, 0x45, 0x0a, 0x10, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e,
	0x64, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x61, 0x63, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x54, 0x0a,
	0x10, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x32, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x6f, 0x64, 0x5f, 0x70, 0x72,
	0x6f, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x22, 0x38, 0x0a, 0x0f, 0x4b, 0x47, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x33, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69,
	0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d,
	0x70, 0x61, 0x69, 0x6c, 0x6c, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x32, 0x89, 0x01,
	0x0a, 0x0a, 0x54, 0x53, 0x53, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0e,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x12, 0x12,
	0x2e, 0x74, 0x73, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0f, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x2e, 0x74,
	0x73, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x1a, 0x14, 0x2e, 0x74, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x74, 0x73, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tss_proto_rawDescOnce sync.Once
	file_tss_proto_rawDescData = file_tss_proto_rawDesc
)

func file_tss_proto_rawDescGZIP() []byte {
	file_tss_proto_rawDescOnce.Do(func() {
		file_tss_proto_rawDescData = protoimpl.X.CompressGZIP(file_tss_proto_rawDescData)
	})
	return file_tss_proto_rawDescData
}

var file_tss_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tss_proto_goTypes = []any{
	(*KeygenRequest)(nil),    // 0: tss.KeygenRequest
	(*KeygenResponse)(nil),   // 1: tss.KeygenResponse
	(*PartyID)(nil),          // 2: tss.PartyID
	(*MessageWrapper)(nil),   // 3: tss.MessageWrapper
	(*MessageResponse)(nil),  // 4: tss.MessageResponse
	(*KGRound1Message)(nil),  // 5: tss.KGRound1Message
	(*KGRound2Message1)(nil), // 6: tss.KGRound2Message1
	(*KGRound2Message2)(nil), // 7: tss.KGRound2Message2
	(*KGRound3Message)(nil),  // 8: tss.KGRound3Message
	(*anypb.Any)(nil),        // 9: google.protobuf.Any
}
var file_tss_proto_depIdxs = []int32{
	2, // 0: tss.KeygenRequest.parties:type_name -> tss.PartyID
	2, // 1: tss.MessageWrapper.from:type_name -> tss.PartyID
	2, // 2: tss.MessageWrapper.to:type_name -> tss.PartyID
	9, // 3: tss.MessageWrapper.content:type_name -> google.protobuf.Any
	0, // 4: tss.TSSService.InitiateKeygen:input_type -> tss.KeygenRequest
	3, // 5: tss.TSSService.ExchangeMessage:input_type -> tss.MessageWrapper
	1, // 6: tss.TSSService.InitiateKeygen:output_type -> tss.KeygenResponse
	4, // 7: tss.TSSService.ExchangeMessage:output_type -> tss.MessageResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tss_proto_init() }
func file_tss_proto_init() {
	if File_tss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tss_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KeygenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*KeygenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PartyID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MessageWrapper); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*MessageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*KGRound1Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*KGRound2Message1); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*KGRound2Message2); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tss_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*KGRound3Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tss_proto_goTypes,
		DependencyIndexes: file_tss_proto_depIdxs,
		MessageInfos:      file_tss_proto_msgTypes,
	}.Build()
	File_tss_proto = out.File
	file_tss_proto_rawDesc = nil
	file_tss_proto_goTypes = nil
	file_tss_proto_depIdxs = nil
}
