// protoc --go_out=plugins=grpc:. *.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: blog/v1/blog.proto

package v1

import (
	reflect "reflect"
	sync "sync"
	pbentity "yijunqiang/gf-micro/blog/api/pbentity"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty" v:"required"`       // v: required
	Content  string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty" v:"required"`   // v: required
	Nickname string `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty" v:"required"` // v: required
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{0}
}

func (x *CreateReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type CreateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRes) Reset() {
	*x = CreateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRes) ProtoMessage() {}

func (x *CreateRes) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRes.ProtoReflect.Descriptor instead.
func (*CreateRes) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{1}
}

type EditReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty" v:"required"`       // v: required
	Content  string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty" v:"required"`   // v: required
	Nickname string `protobuf:"bytes,3,opt,name=Nickname,proto3" json:"Nickname,omitempty" v:"required"` // v: required
}

func (x *EditReq) Reset() {
	*x = EditReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditReq) ProtoMessage() {}

func (x *EditReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditReq.ProtoReflect.Descriptor instead.
func (*EditReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{2}
}

func (x *EditReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *EditReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

type EditRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditRes) Reset() {
	*x = EditRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRes) ProtoMessage() {}

func (x *EditRes) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditRes.ProtoReflect.Descriptor instead.
func (*EditRes) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{3}
}

type GetOneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty" v:"required"` // v: required
}

func (x *GetOneReq) Reset() {
	*x = GetOneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneReq) ProtoMessage() {}

func (x *GetOneReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneReq.ProtoReflect.Descriptor instead.
func (*GetOneReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{4}
}

func (x *GetOneReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetOneRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blog *pbentity.Blog `protobuf:"bytes,1,opt,name=Blog,proto3" json:"Blog,omitempty"`
}

func (x *GetOneRes) Reset() {
	*x = GetOneRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneRes) ProtoMessage() {}

func (x *GetOneRes) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneRes.ProtoReflect.Descriptor instead.
func (*GetOneRes) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{5}
}

func (x *GetOneRes) GetBlog() *pbentity.Blog {
	if x != nil {
		return x.Blog
	}
	return nil
}

type GetListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListReq) Reset() {
	*x = GetListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListReq) ProtoMessage() {}

func (x *GetListReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListReq.ProtoReflect.Descriptor instead.
func (*GetListReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{6}
}

type GetListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*pbentity.Blog `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *GetListRes) Reset() {
	*x = GetListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListRes) ProtoMessage() {}

func (x *GetListRes) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListRes.ProtoReflect.Descriptor instead.
func (*GetListRes) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{7}
}

func (x *GetListRes) GetList() []*pbentity.Blog {
	if x != nil {
		return x.List
	}
	return nil
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty" v:"required"` // v: required
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteRes) Reset() {
	*x = DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRes) ProtoMessage() {}

func (x *DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRes.ProtoReflect.Descriptor instead.
func (*DeleteRes) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{9}
}

type BatDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=Ids,proto3" json:"Ids,omitempty" v:"required"` // v: required
}

func (x *BatDeleteReq) Reset() {
	*x = BatDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatDeleteReq) ProtoMessage() {}

func (x *BatDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatDeleteReq.ProtoReflect.Descriptor instead.
func (*BatDeleteReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{10}
}

func (x *BatDeleteReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type BatDeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatNo string `protobuf:"bytes,1,opt,name=BatNo,proto3" json:"BatNo,omitempty"`
}

func (x *BatDeleteRes) Reset() {
	*x = BatDeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatDeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatDeleteRes) ProtoMessage() {}

func (x *BatDeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatDeleteRes.ProtoReflect.Descriptor instead.
func (*BatDeleteRes) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{11}
}

func (x *BatDeleteRes) GetBatNo() string {
	if x != nil {
		return x.BatNo
	}
	return ""
}

type GetBatDeleteStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatNo string `protobuf:"bytes,1,opt,name=BatNo,proto3" json:"BatNo,omitempty" v:"required"` // v: required
}

func (x *GetBatDeleteStatusReq) Reset() {
	*x = GetBatDeleteStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatDeleteStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatDeleteStatusReq) ProtoMessage() {}

func (x *GetBatDeleteStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatDeleteStatusReq.ProtoReflect.Descriptor instead.
func (*GetBatDeleteStatusReq) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{12}
}

func (x *GetBatDeleteStatusReq) GetBatNo() string {
	if x != nil {
		return x.BatNo
	}
	return ""
}

type GetBatDeleteStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *GetBatDeleteStatusRes) Reset() {
	*x = GetBatDeleteStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blog_v1_blog_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBatDeleteStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBatDeleteStatusRes) ProtoMessage() {}

func (x *GetBatDeleteStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file_blog_v1_blog_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBatDeleteStatusRes.ProtoReflect.Descriptor instead.
func (*GetBatDeleteStatusRes) Descriptor() ([]byte, []int) {
	return file_blog_v1_blog_proto_rawDescGZIP(), []int{13}
}

func (x *GetBatDeleteStatusRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_blog_v1_blog_proto protoreflect.FileDescriptor

var file_blog_v1_blog_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x6c, 0x6f, 0x67, 0x1a, 0x13, 0x70, 0x62, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x57, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x09, 0x0a, 0x07,
	0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x04, 0x42, 0x6c, 0x6f, 0x67, 0x22, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x22, 0x30, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x12, 0x22, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22,
	0x20, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x49, 0x64,
	0x73, 0x22, 0x24, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x42, 0x61, 0x74, 0x4e, 0x6f, 0x22, 0x2d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x14, 0x0a, 0x05, 0x42, 0x61, 0x74, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x42, 0x61, 0x74, 0x4e, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xf2, 0x02, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67,
	0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x62, 0x6c, 0x6f,
	0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x26,
	0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x45, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x45, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x12, 0x0f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x0f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x0f, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x09, 0x42, 0x61, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x42, 0x61, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1b, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24,
	0x79, 0x69, 0x6a, 0x75, 0x6e, 0x71, 0x69, 0x61, 0x6e, 0x67, 0x2f, 0x67, 0x66, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6c, 0x6f,
	0x67, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blog_v1_blog_proto_rawDescOnce sync.Once
	file_blog_v1_blog_proto_rawDescData = file_blog_v1_blog_proto_rawDesc
)

func file_blog_v1_blog_proto_rawDescGZIP() []byte {
	file_blog_v1_blog_proto_rawDescOnce.Do(func() {
		file_blog_v1_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_blog_v1_blog_proto_rawDescData)
	})
	return file_blog_v1_blog_proto_rawDescData
}

var file_blog_v1_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_blog_v1_blog_proto_goTypes = []interface{}{
	(*CreateReq)(nil),             // 0: blog.CreateReq
	(*CreateRes)(nil),             // 1: blog.CreateRes
	(*EditReq)(nil),               // 2: blog.EditReq
	(*EditRes)(nil),               // 3: blog.EditRes
	(*GetOneReq)(nil),             // 4: blog.GetOneReq
	(*GetOneRes)(nil),             // 5: blog.GetOneRes
	(*GetListReq)(nil),            // 6: blog.GetListReq
	(*GetListRes)(nil),            // 7: blog.GetListRes
	(*DeleteReq)(nil),             // 8: blog.DeleteReq
	(*DeleteRes)(nil),             // 9: blog.DeleteRes
	(*BatDeleteReq)(nil),          // 10: blog.BatDeleteReq
	(*BatDeleteRes)(nil),          // 11: blog.BatDeleteRes
	(*GetBatDeleteStatusReq)(nil), // 12: blog.GetBatDeleteStatusReq
	(*GetBatDeleteStatusRes)(nil), // 13: blog.GetBatDeleteStatusRes
	(*pbentity.Blog)(nil),         // 14: pbentity.Blog
}
var file_blog_v1_blog_proto_depIdxs = []int32{
	14, // 0: blog.GetOneRes.Blog:type_name -> pbentity.Blog
	14, // 1: blog.GetListRes.List:type_name -> pbentity.Blog
	0,  // 2: blog.Blog.Create:input_type -> blog.CreateReq
	2,  // 3: blog.Blog.Edit:input_type -> blog.EditReq
	4,  // 4: blog.Blog.GetOne:input_type -> blog.GetOneReq
	6,  // 5: blog.Blog.GetList:input_type -> blog.GetListReq
	8,  // 6: blog.Blog.Delete:input_type -> blog.DeleteReq
	10, // 7: blog.Blog.BatDelete:input_type -> blog.BatDeleteReq
	12, // 8: blog.Blog.GetBatDeleteStatus:input_type -> blog.GetBatDeleteStatusReq
	1,  // 9: blog.Blog.Create:output_type -> blog.CreateRes
	3,  // 10: blog.Blog.Edit:output_type -> blog.EditRes
	5,  // 11: blog.Blog.GetOne:output_type -> blog.GetOneRes
	7,  // 12: blog.Blog.GetList:output_type -> blog.GetListRes
	9,  // 13: blog.Blog.Delete:output_type -> blog.DeleteRes
	11, // 14: blog.Blog.BatDelete:output_type -> blog.BatDeleteRes
	13, // 15: blog.Blog.GetBatDeleteStatus:output_type -> blog.GetBatDeleteStatusRes
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_blog_v1_blog_proto_init() }
func file_blog_v1_blog_proto_init() {
	if File_blog_v1_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blog_v1_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_blog_v1_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRes); i {
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
		file_blog_v1_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditReq); i {
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
		file_blog_v1_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditRes); i {
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
		file_blog_v1_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneReq); i {
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
		file_blog_v1_blog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneRes); i {
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
		file_blog_v1_blog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListReq); i {
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
		file_blog_v1_blog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListRes); i {
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
		file_blog_v1_blog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_blog_v1_blog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRes); i {
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
		file_blog_v1_blog_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatDeleteReq); i {
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
		file_blog_v1_blog_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatDeleteRes); i {
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
		file_blog_v1_blog_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatDeleteStatusReq); i {
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
		file_blog_v1_blog_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBatDeleteStatusRes); i {
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
			RawDescriptor: file_blog_v1_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blog_v1_blog_proto_goTypes,
		DependencyIndexes: file_blog_v1_blog_proto_depIdxs,
		MessageInfos:      file_blog_v1_blog_proto_msgTypes,
	}.Build()
	File_blog_v1_blog_proto = out.File
	file_blog_v1_blog_proto_rawDesc = nil
	file_blog_v1_blog_proto_goTypes = nil
	file_blog_v1_blog_proto_depIdxs = nil
}
