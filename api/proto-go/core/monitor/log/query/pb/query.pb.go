// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: query.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source    string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Stream    string `protobuf:"bytes,3,opt,name=stream,proto3" json:"stream,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Offset    int64  `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	Content   string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
	Level     string `protobuf:"bytes,8,opt,name=level,proto3" json:"level,omitempty"`
	RequestId string `protobuf:"bytes,9,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *LogItem) Reset() {
	*x = LogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogItem) ProtoMessage() {}

func (x *LogItem) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogItem.ProtoReflect.Descriptor instead.
func (*LogItem) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{0}
}

func (x *LogItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LogItem) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *LogItem) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *LogItem) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogItem) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *LogItem) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *LogItem) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *LogItem) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type GetLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source    string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"` // job or container
	Stream    string `protobuf:"bytes,3,opt,name=stream,proto3" json:"stream,omitempty"`
	RequestId string `protobuf:"bytes,4,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Start     int64  `protobuf:"varint,5,opt,name=start,proto3" json:"start,omitempty"`
	End       int64  `protobuf:"varint,6,opt,name=end,proto3" json:"end,omitempty"`
	Count     int64  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	Pattern   string `protobuf:"bytes,8,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Debug     bool   `protobuf:"varint,10,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *GetLogRequest) Reset() {
	*x = GetLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogRequest) ProtoMessage() {}

func (x *GetLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogRequest.ProtoReflect.Descriptor instead.
func (*GetLogRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetLogRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLogRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetLogRequest) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *GetLogRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetLogRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetLogRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *GetLogRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetLogRequest) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *GetLogRequest) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type GetLogByRuntimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source        string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"` // job or container
	Stream        string `protobuf:"bytes,3,opt,name=stream,proto3" json:"stream,omitempty"`
	RequestId     string `protobuf:"bytes,4,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Start         int64  `protobuf:"varint,5,opt,name=start,proto3" json:"start,omitempty"`
	End           int64  `protobuf:"varint,6,opt,name=end,proto3" json:"end,omitempty"`
	Count         int64  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	ApplicationId string `protobuf:"bytes,8,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	Pattern       string `protobuf:"bytes,9,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Debug         bool   `protobuf:"varint,10,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *GetLogByRuntimeRequest) Reset() {
	*x = GetLogByRuntimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogByRuntimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogByRuntimeRequest) ProtoMessage() {}

func (x *GetLogByRuntimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogByRuntimeRequest.ProtoReflect.Descriptor instead.
func (*GetLogByRuntimeRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{2}
}

func (x *GetLogByRuntimeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLogByRuntimeRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetLogByRuntimeRequest) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *GetLogByRuntimeRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetLogByRuntimeRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetLogByRuntimeRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *GetLogByRuntimeRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetLogByRuntimeRequest) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

func (x *GetLogByRuntimeRequest) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *GetLogByRuntimeRequest) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type GetLogByOrganizationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Source      string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"` // job or container
	Stream      string `protobuf:"bytes,3,opt,name=stream,proto3" json:"stream,omitempty"`
	RequestId   string `protobuf:"bytes,4,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Start       int64  `protobuf:"varint,5,opt,name=start,proto3" json:"start,omitempty"`
	End         int64  `protobuf:"varint,6,opt,name=end,proto3" json:"end,omitempty"`
	Count       int64  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	ClusterName string `protobuf:"bytes,8,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Pattern     string `protobuf:"bytes,9,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Debug       bool   `protobuf:"varint,11,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *GetLogByOrganizationRequest) Reset() {
	*x = GetLogByOrganizationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogByOrganizationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogByOrganizationRequest) ProtoMessage() {}

func (x *GetLogByOrganizationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogByOrganizationRequest.ProtoReflect.Descriptor instead.
func (*GetLogByOrganizationRequest) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{3}
}

func (x *GetLogByOrganizationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetLogByOrganizationRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *GetLogByOrganizationRequest) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *GetLogByOrganizationRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetLogByOrganizationRequest) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetLogByOrganizationRequest) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *GetLogByOrganizationRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetLogByOrganizationRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *GetLogByOrganizationRequest) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *GetLogByOrganizationRequest) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type GetLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []*LogItem `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *GetLogResponse) Reset() {
	*x = GetLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogResponse) ProtoMessage() {}

func (x *GetLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogResponse.ProtoReflect.Descriptor instead.
func (*GetLogResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{4}
}

func (x *GetLogResponse) GetLines() []*LogItem {
	if x != nil {
		return x.Lines
	}
	return nil
}

type GetLogByRuntimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []*LogItem `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *GetLogByRuntimeResponse) Reset() {
	*x = GetLogByRuntimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogByRuntimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogByRuntimeResponse) ProtoMessage() {}

func (x *GetLogByRuntimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogByRuntimeResponse.ProtoReflect.Descriptor instead.
func (*GetLogByRuntimeResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{5}
}

func (x *GetLogByRuntimeResponse) GetLines() []*LogItem {
	if x != nil {
		return x.Lines
	}
	return nil
}

type GetLogByOrganizationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []*LogItem `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *GetLogByOrganizationResponse) Reset() {
	*x = GetLogByOrganizationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogByOrganizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogByOrganizationResponse) ProtoMessage() {}

func (x *GetLogByOrganizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogByOrganizationResponse.ProtoReflect.Descriptor instead.
func (*GetLogByOrganizationResponse) Descriptor() ([]byte, []int) {
	return file_query_proto_rawDescGZIP(), []int{6}
}

func (x *GetLogByOrganizationResponse) GetLines() []*LogItem {
	if x != nil {
		return x.Lines
	}
	return nil
}

var File_query_proto protoreflect.FileDescriptor

var file_query_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x07, 0x4c, 0x6f, 0x67,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0xdb, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x22, 0x8a, 0x02, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x22, 0x8b, 0x02, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75,
	0x67, 0x22, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22,
	0x55, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x5a, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x32, 0xcc, 0x03, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x2a, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x65,
	0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0b, 0x12, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x97, 0x01, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x33, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0xa8, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x65, 0x72, 0x64, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x42, 0x79,
	0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67,
	0x73, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64,
	0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_query_proto_rawDescOnce sync.Once
	file_query_proto_rawDescData = file_query_proto_rawDesc
)

func file_query_proto_rawDescGZIP() []byte {
	file_query_proto_rawDescOnce.Do(func() {
		file_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_query_proto_rawDescData)
	})
	return file_query_proto_rawDescData
}

var file_query_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_query_proto_goTypes = []interface{}{
	(*LogItem)(nil),                      // 0: erda.core.monitor.log.query.LogItem
	(*GetLogRequest)(nil),                // 1: erda.core.monitor.log.query.GetLogRequest
	(*GetLogByRuntimeRequest)(nil),       // 2: erda.core.monitor.log.query.GetLogByRuntimeRequest
	(*GetLogByOrganizationRequest)(nil),  // 3: erda.core.monitor.log.query.GetLogByOrganizationRequest
	(*GetLogResponse)(nil),               // 4: erda.core.monitor.log.query.GetLogResponse
	(*GetLogByRuntimeResponse)(nil),      // 5: erda.core.monitor.log.query.GetLogByRuntimeResponse
	(*GetLogByOrganizationResponse)(nil), // 6: erda.core.monitor.log.query.GetLogByOrganizationResponse
}
var file_query_proto_depIdxs = []int32{
	0, // 0: erda.core.monitor.log.query.GetLogResponse.lines:type_name -> erda.core.monitor.log.query.LogItem
	0, // 1: erda.core.monitor.log.query.GetLogByRuntimeResponse.lines:type_name -> erda.core.monitor.log.query.LogItem
	0, // 2: erda.core.monitor.log.query.GetLogByOrganizationResponse.lines:type_name -> erda.core.monitor.log.query.LogItem
	1, // 3: erda.core.monitor.log.query.LogQueryService.GetLog:input_type -> erda.core.monitor.log.query.GetLogRequest
	2, // 4: erda.core.monitor.log.query.LogQueryService.GetLogByRuntime:input_type -> erda.core.monitor.log.query.GetLogByRuntimeRequest
	3, // 5: erda.core.monitor.log.query.LogQueryService.GetLogByOrganization:input_type -> erda.core.monitor.log.query.GetLogByOrganizationRequest
	4, // 6: erda.core.monitor.log.query.LogQueryService.GetLog:output_type -> erda.core.monitor.log.query.GetLogResponse
	5, // 7: erda.core.monitor.log.query.LogQueryService.GetLogByRuntime:output_type -> erda.core.monitor.log.query.GetLogByRuntimeResponse
	6, // 8: erda.core.monitor.log.query.LogQueryService.GetLogByOrganization:output_type -> erda.core.monitor.log.query.GetLogByOrganizationResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_query_proto_init() }
func file_query_proto_init() {
	if File_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogItem); i {
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
		file_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogRequest); i {
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
		file_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogByRuntimeRequest); i {
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
		file_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogByOrganizationRequest); i {
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
		file_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogResponse); i {
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
		file_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogByRuntimeResponse); i {
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
		file_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogByOrganizationResponse); i {
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
			RawDescriptor: file_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_query_proto_goTypes,
		DependencyIndexes: file_query_proto_depIdxs,
		MessageInfos:      file_query_proto_msgTypes,
	}.Build()
	File_query_proto = out.File
	file_query_proto_rawDesc = nil
	file_query_proto_goTypes = nil
	file_query_proto_depIdxs = nil
}
