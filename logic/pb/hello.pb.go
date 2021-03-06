// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: hello.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Bundle            string   `protobuf:"bytes,2,opt,name=Bundle,proto3" json:"Bundle,omitempty"`
	DeveloperId       string   `protobuf:"bytes,3,opt,name=DeveloperId,proto3" json:"DeveloperId,omitempty"`
	Developer         string   `protobuf:"bytes,4,opt,name=Developer,proto3" json:"Developer,omitempty"`
	Title             string   `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	Categories        string   `protobuf:"bytes,6,opt,name=Categories,proto3" json:"Categories,omitempty"`
	Price             string   `protobuf:"bytes,7,opt,name=Price,proto3" json:"Price,omitempty"`
	Picture           string   `protobuf:"bytes,8,opt,name=Picture,proto3" json:"Picture,omitempty"`
	Screenshots       []string `protobuf:"bytes,9,rep,name=Screenshots,proto3" json:"Screenshots,omitempty"`
	Rating            string   `protobuf:"bytes,10,opt,name=Rating,proto3" json:"Rating,omitempty"`
	ReviewCount       string   `protobuf:"bytes,11,opt,name=ReviewCount,proto3" json:"ReviewCount,omitempty"`
	RatingHistogram   []string `protobuf:"bytes,12,rep,name=RatingHistogram,proto3" json:"RatingHistogram,omitempty"`
	Description       string   `protobuf:"bytes,13,opt,name=Description,proto3" json:"Description,omitempty"`
	ShortDescription  string   `protobuf:"bytes,14,opt,name=ShortDescription,proto3" json:"ShortDescription,omitempty"`
	RecentChanges     string   `protobuf:"bytes,15,opt,name=RecentChanges,proto3" json:"RecentChanges,omitempty"`
	ReleaseDate       string   `protobuf:"bytes,16,opt,name=ReleaseDate,proto3" json:"ReleaseDate,omitempty"`
	LastUpdateDate    string   `protobuf:"bytes,17,opt,name=LastUpdateDate,proto3" json:"LastUpdateDate,omitempty"`
	AppSize           string   `protobuf:"bytes,18,opt,name=AppSize,proto3" json:"AppSize,omitempty"`
	Installs          string   `protobuf:"bytes,19,opt,name=Installs,proto3" json:"Installs,omitempty"`
	Version           string   `protobuf:"bytes,20,opt,name=Version,proto3" json:"Version,omitempty"`
	AndroidVersion    string   `protobuf:"bytes,21,opt,name=AndroidVersion,proto3" json:"AndroidVersion,omitempty"`
	ContentRating     string   `protobuf:"bytes,22,opt,name=ContentRating,proto3" json:"ContentRating,omitempty"`
	DeveloperContacts string   `protobuf:"bytes,23,opt,name=DeveloperContacts,proto3" json:"DeveloperContacts,omitempty"`
	Geo               string   `protobuf:"bytes,24,opt,name=Geo,proto3" json:"Geo,omitempty"`
	Date              int64    `protobuf:"varint,25,opt,name=Date,proto3" json:"Date,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{0}
}

func (x *App) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *App) GetBundle() string {
	if x != nil {
		return x.Bundle
	}
	return ""
}

func (x *App) GetDeveloperId() string {
	if x != nil {
		return x.DeveloperId
	}
	return ""
}

func (x *App) GetDeveloper() string {
	if x != nil {
		return x.Developer
	}
	return ""
}

func (x *App) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *App) GetCategories() string {
	if x != nil {
		return x.Categories
	}
	return ""
}

func (x *App) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *App) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *App) GetScreenshots() []string {
	if x != nil {
		return x.Screenshots
	}
	return nil
}

func (x *App) GetRating() string {
	if x != nil {
		return x.Rating
	}
	return ""
}

func (x *App) GetReviewCount() string {
	if x != nil {
		return x.ReviewCount
	}
	return ""
}

func (x *App) GetRatingHistogram() []string {
	if x != nil {
		return x.RatingHistogram
	}
	return nil
}

func (x *App) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *App) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *App) GetRecentChanges() string {
	if x != nil {
		return x.RecentChanges
	}
	return ""
}

func (x *App) GetReleaseDate() string {
	if x != nil {
		return x.ReleaseDate
	}
	return ""
}

func (x *App) GetLastUpdateDate() string {
	if x != nil {
		return x.LastUpdateDate
	}
	return ""
}

func (x *App) GetAppSize() string {
	if x != nil {
		return x.AppSize
	}
	return ""
}

func (x *App) GetInstalls() string {
	if x != nil {
		return x.Installs
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *App) GetAndroidVersion() string {
	if x != nil {
		return x.AndroidVersion
	}
	return ""
}

func (x *App) GetContentRating() string {
	if x != nil {
		return x.ContentRating
	}
	return ""
}

func (x *App) GetDeveloperContacts() string {
	if x != nil {
		return x.DeveloperContacts
	}
	return ""
}

func (x *App) GetGeo() string {
	if x != nil {
		return x.Geo
	}
	return ""
}

func (x *App) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type GetTempAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (x *GetTempAppsRequest) Reset() {
	*x = GetTempAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTempAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTempAppsRequest) ProtoMessage() {}

func (x *GetTempAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTempAppsRequest.ProtoReflect.Descriptor instead.
func (*GetTempAppsRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{1}
}

func (x *GetTempAppsRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetTempAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*App `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *GetTempAppsResponse) Reset() {
	*x = GetTempAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTempAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTempAppsResponse) ProtoMessage() {}

func (x *GetTempAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTempAppsResponse.ProtoReflect.Descriptor instead.
func (*GetTempAppsResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{2}
}

func (x *GetTempAppsResponse) GetApps() []*App {
	if x != nil {
		return x.Apps
	}
	return nil
}

func (x *GetTempAppsResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type TerminateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminateRequest) Reset() {
	*x = TerminateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateRequest) ProtoMessage() {}

func (x *TerminateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateRequest.ProtoReflect.Descriptor instead.
func (*TerminateRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{3}
}

type TerminateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  string `protobuf:"bytes,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *TerminateResponse) Reset() {
	*x = TerminateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateResponse) ProtoMessage() {}

func (x *TerminateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateResponse.ProtoReflect.Descriptor instead.
func (*TerminateResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{4}
}

func (x *TerminateResponse) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

func (x *TerminateResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lang     string `protobuf:"bytes,1,opt,name=Lang,proto3" json:"Lang,omitempty"`
	Parallel int32  `protobuf:"varint,2,opt,name=Parallel,proto3" json:"Parallel,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{5}
}

func (x *StartRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *StartRequest) GetParallel() int32 {
	if x != nil {
		return x.Parallel
	}
	return 0
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok  string `protobuf:"bytes,1,opt,name=Ok,proto3" json:"Ok,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_hello_proto_rawDescGZIP(), []int{6}
}

func (x *StartResponse) GetOk() string {
	if x != nil {
		return x.Ok
	}
	return ""
}

func (x *StartResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_hello_proto protoreflect.FileDescriptor

var file_hello_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x89, 0x06, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x63, 0x72, 0x65, 0x65,
	0x6e, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x63,
	0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x52,
	0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x70, 0x70, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70,
	0x70, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x73, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x41,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x47, 0x65, 0x6f, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x47, 0x65, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x6d, 0x70, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x22,
	0x12, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x22, 0x3e, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x61,
	0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x22, 0x31, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x45,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x32, 0xb9, 0x01,
	0x0a, 0x09, 0x52, 0x69, 0x67, 0x68, 0x74, 0x48, 0x61, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x6d, 0x70, 0x41, 0x70, 0x70, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x6d, 0x70, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x41, 0x70, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_hello_proto_rawDescOnce sync.Once
	file_hello_proto_rawDescData = file_hello_proto_rawDesc
)

func file_hello_proto_rawDescGZIP() []byte {
	file_hello_proto_rawDescOnce.Do(func() {
		file_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_proto_rawDescData)
	})
	return file_hello_proto_rawDescData
}

var file_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_hello_proto_goTypes = []interface{}{
	(*App)(nil),                 // 0: pb.App
	(*GetTempAppsRequest)(nil),  // 1: pb.GetTempAppsRequest
	(*GetTempAppsResponse)(nil), // 2: pb.GetTempAppsResponse
	(*TerminateRequest)(nil),    // 3: pb.TerminateRequest
	(*TerminateResponse)(nil),   // 4: pb.TerminateResponse
	(*StartRequest)(nil),        // 5: pb.StartRequest
	(*StartResponse)(nil),       // 6: pb.StartResponse
}
var file_hello_proto_depIdxs = []int32{
	0, // 0: pb.GetTempAppsResponse.apps:type_name -> pb.App
	5, // 1: pb.RightHand.Start:input_type -> pb.StartRequest
	3, // 2: pb.RightHand.Terminate:input_type -> pb.TerminateRequest
	1, // 3: pb.RightHand.GetTempApps:input_type -> pb.GetTempAppsRequest
	6, // 4: pb.RightHand.Start:output_type -> pb.StartResponse
	4, // 5: pb.RightHand.Terminate:output_type -> pb.TerminateResponse
	2, // 6: pb.RightHand.GetTempApps:output_type -> pb.GetTempAppsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hello_proto_init() }
func file_hello_proto_init() {
	if File_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*App); i {
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
		file_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTempAppsRequest); i {
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
		file_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTempAppsResponse); i {
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
		file_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateRequest); i {
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
		file_hello_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateResponse); i {
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
		file_hello_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_hello_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
			RawDescriptor: file_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_proto_goTypes,
		DependencyIndexes: file_hello_proto_depIdxs,
		MessageInfos:      file_hello_proto_msgTypes,
	}.Build()
	File_hello_proto = out.File
	file_hello_proto_rawDesc = nil
	file_hello_proto_goTypes = nil
	file_hello_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RightHandClient is the client API for RightHand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RightHandClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error)
	GetTempApps(ctx context.Context, in *GetTempAppsRequest, opts ...grpc.CallOption) (*GetTempAppsResponse, error)
}

type rightHandClient struct {
	cc grpc.ClientConnInterface
}

func NewRightHandClient(cc grpc.ClientConnInterface) RightHandClient {
	return &rightHandClient{cc}
}

func (c *rightHandClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/pb.RightHand/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightHandClient) Terminate(ctx context.Context, in *TerminateRequest, opts ...grpc.CallOption) (*TerminateResponse, error) {
	out := new(TerminateResponse)
	err := c.cc.Invoke(ctx, "/pb.RightHand/Terminate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rightHandClient) GetTempApps(ctx context.Context, in *GetTempAppsRequest, opts ...grpc.CallOption) (*GetTempAppsResponse, error) {
	out := new(GetTempAppsResponse)
	err := c.cc.Invoke(ctx, "/pb.RightHand/GetTempApps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RightHandServer is the server API for RightHand service.
type RightHandServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error)
	GetTempApps(context.Context, *GetTempAppsRequest) (*GetTempAppsResponse, error)
}

// UnimplementedRightHandServer can be embedded to have forward compatible implementations.
type UnimplementedRightHandServer struct {
}

func (*UnimplementedRightHandServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedRightHandServer) Terminate(context.Context, *TerminateRequest) (*TerminateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Terminate not implemented")
}
func (*UnimplementedRightHandServer) GetTempApps(context.Context, *GetTempAppsRequest) (*GetTempAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTempApps not implemented")
}

func RegisterRightHandServer(s *grpc.Server, srv RightHandServer) {
	s.RegisterService(&_RightHand_serviceDesc, srv)
}

func _RightHand_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightHandServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RightHand/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightHandServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RightHand_Terminate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightHandServer).Terminate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RightHand/Terminate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightHandServer).Terminate(ctx, req.(*TerminateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RightHand_GetTempApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTempAppsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RightHandServer).GetTempApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.RightHand/GetTempApps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RightHandServer).GetTempApps(ctx, req.(*GetTempAppsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RightHand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.RightHand",
	HandlerType: (*RightHandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _RightHand_Start_Handler,
		},
		{
			MethodName: "Terminate",
			Handler:    _RightHand_Terminate_Handler,
		},
		{
			MethodName: "GetTempApps",
			Handler:    _RightHand_GetTempApps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}
