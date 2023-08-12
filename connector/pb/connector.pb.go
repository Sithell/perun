// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: connector.proto

package pb

import (
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

type ServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Body:
	//
	//	*ServerRequest_RunContainer
	Body isServerRequest_Body `protobuf_oneof:"body"`
}

func (x *ServerRequest) Reset() {
	*x = ServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerRequest) ProtoMessage() {}

func (x *ServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerRequest.ProtoReflect.Descriptor instead.
func (*ServerRequest) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{0}
}

func (x *ServerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *ServerRequest) GetBody() isServerRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *ServerRequest) GetRunContainer() *RunContainerRequest {
	if x, ok := x.GetBody().(*ServerRequest_RunContainer); ok {
		return x.RunContainer
	}
	return nil
}

type isServerRequest_Body interface {
	isServerRequest_Body()
}

type ServerRequest_RunContainer struct {
	RunContainer *RunContainerRequest `protobuf:"bytes,3,opt,name=run_container,json=runContainer,proto3,oneof"`
}

func (*ServerRequest_RunContainer) isServerRequest_Body() {}

type ClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseTo string `protobuf:"bytes,1,opt,name=response_to,json=responseTo,proto3" json:"response_to,omitempty"`
	// Types that are assignable to Body:
	//
	//	*ClientResponse_RunContainer
	Body isClientResponse_Body `protobuf_oneof:"body"`
}

func (x *ClientResponse) Reset() {
	*x = ClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientResponse) ProtoMessage() {}

func (x *ClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientResponse.ProtoReflect.Descriptor instead.
func (*ClientResponse) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{1}
}

func (x *ClientResponse) GetResponseTo() string {
	if x != nil {
		return x.ResponseTo
	}
	return ""
}

func (m *ClientResponse) GetBody() isClientResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (x *ClientResponse) GetRunContainer() *RunContainerResponse {
	if x, ok := x.GetBody().(*ClientResponse_RunContainer); ok {
		return x.RunContainer
	}
	return nil
}

type isClientResponse_Body interface {
	isClientResponse_Body()
}

type ClientResponse_RunContainer struct {
	RunContainer *RunContainerResponse `protobuf:"bytes,3,opt,name=run_container,json=runContainer,proto3,oneof"`
}

func (*ClientResponse_RunContainer) isClientResponse_Body() {}

type RunContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Cmd   string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *RunContainerRequest) Reset() {
	*x = RunContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunContainerRequest) ProtoMessage() {}

func (x *RunContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunContainerRequest.ProtoReflect.Descriptor instead.
func (*RunContainerRequest) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{2}
}

func (x *RunContainerRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *RunContainerRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type RunContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RunContainerResponse) Reset() {
	*x = RunContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunContainerResponse) ProtoMessage() {}

func (x *RunContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunContainerResponse.ProtoReflect.Descriptor instead.
func (*RunContainerResponse) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{3}
}

func (x *RunContainerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RunContainerParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image      string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Cmd        string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	ProviderId uint64 `protobuf:"varint,3,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
}

func (x *RunContainerParams) Reset() {
	*x = RunContainerParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunContainerParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunContainerParams) ProtoMessage() {}

func (x *RunContainerParams) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunContainerParams.ProtoReflect.Descriptor instead.
func (*RunContainerParams) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{4}
}

func (x *RunContainerParams) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *RunContainerParams) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *RunContainerParams) GetProviderId() uint64 {
	if x != nil {
		return x.ProviderId
	}
	return 0
}

type ContainerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ContainerInfo) Reset() {
	*x = ContainerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerInfo) ProtoMessage() {}

func (x *ContainerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerInfo.ProtoReflect.Descriptor instead.
func (*ContainerInfo) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{5}
}

func (x *ContainerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{6}
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{7}
}

type ProviderConnections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connections []*ProviderConnection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections,omitempty"`
}

func (x *ProviderConnections) Reset() {
	*x = ProviderConnections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderConnections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderConnections) ProtoMessage() {}

func (x *ProviderConnections) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderConnections.ProtoReflect.Descriptor instead.
func (*ProviderConnections) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{8}
}

func (x *ProviderConnections) GetConnections() []*ProviderConnection {
	if x != nil {
		return x.Connections
	}
	return nil
}

type ProviderConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderId uint64 `protobuf:"varint,1,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
}

func (x *ProviderConnection) Reset() {
	*x = ProviderConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connector_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderConnection) ProtoMessage() {}

func (x *ProviderConnection) ProtoReflect() protoreflect.Message {
	mi := &file_connector_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderConnection.ProtoReflect.Descriptor instead.
func (*ProviderConnection) Descriptor() ([]byte, []int) {
	return file_connector_proto_rawDescGZIP(), []int{9}
}

func (x *ProviderConnection) GetProviderId() uint64 {
	if x != nil {
		return x.ProviderId
	}
	return 0
}

var File_connector_proto protoreflect.FileDescriptor

var file_connector_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x6e, 0x0a, 0x0d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a,
	0x0d, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x81, 0x01, 0x0a,
	0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x6f,
	0x12, 0x46, 0x0a, 0x0d, 0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x75, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x22, 0x3d, 0x0a, 0x13, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x22,
	0x26, 0x0a, 0x14, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x12, 0x52, 0x75, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x0e, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x56, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x35, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32,
	0x8c, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x49, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x18,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x32, 0xd1,
	0x01, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x33, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x69, 0x74, 0x68, 0x65, 0x6c, 0x6c, 0x2f, 0x70, 0x65, 0x72, 0x75, 0x6e, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_connector_proto_rawDescOnce sync.Once
	file_connector_proto_rawDescData = file_connector_proto_rawDesc
)

func file_connector_proto_rawDescGZIP() []byte {
	file_connector_proto_rawDescOnce.Do(func() {
		file_connector_proto_rawDescData = protoimpl.X.CompressGZIP(file_connector_proto_rawDescData)
	})
	return file_connector_proto_rawDescData
}

var file_connector_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_connector_proto_goTypes = []interface{}{
	(*ServerRequest)(nil),        // 0: connector.ServerRequest
	(*ClientResponse)(nil),       // 1: connector.ClientResponse
	(*RunContainerRequest)(nil),  // 2: connector.RunContainerRequest
	(*RunContainerResponse)(nil), // 3: connector.RunContainerResponse
	(*RunContainerParams)(nil),   // 4: connector.RunContainerParams
	(*ContainerInfo)(nil),        // 5: connector.ContainerInfo
	(*Empty)(nil),                // 6: connector.Empty
	(*PingResponse)(nil),         // 7: connector.PingResponse
	(*ProviderConnections)(nil),  // 8: connector.ProviderConnections
	(*ProviderConnection)(nil),   // 9: connector.ProviderConnection
}
var file_connector_proto_depIdxs = []int32{
	2, // 0: connector.ServerRequest.run_container:type_name -> connector.RunContainerRequest
	3, // 1: connector.ClientResponse.run_container:type_name -> connector.RunContainerResponse
	9, // 2: connector.ProviderConnections.connections:type_name -> connector.ProviderConnection
	6, // 3: connector.Provider.Ping:input_type -> connector.Empty
	1, // 4: connector.Provider.InitConnection:input_type -> connector.ClientResponse
	6, // 5: connector.Api.Ping:input_type -> connector.Empty
	6, // 6: connector.Api.GetActiveConnections:input_type -> connector.Empty
	4, // 7: connector.Api.RunContainer:input_type -> connector.RunContainerParams
	7, // 8: connector.Provider.Ping:output_type -> connector.PingResponse
	0, // 9: connector.Provider.InitConnection:output_type -> connector.ServerRequest
	7, // 10: connector.Api.Ping:output_type -> connector.PingResponse
	8, // 11: connector.Api.GetActiveConnections:output_type -> connector.ProviderConnections
	5, // 12: connector.Api.RunContainer:output_type -> connector.ContainerInfo
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_connector_proto_init() }
func file_connector_proto_init() {
	if File_connector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerRequest); i {
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
		file_connector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientResponse); i {
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
		file_connector_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunContainerRequest); i {
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
		file_connector_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunContainerResponse); i {
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
		file_connector_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunContainerParams); i {
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
		file_connector_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerInfo); i {
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
		file_connector_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_connector_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResponse); i {
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
		file_connector_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderConnections); i {
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
		file_connector_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderConnection); i {
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
	file_connector_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ServerRequest_RunContainer)(nil),
	}
	file_connector_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ClientResponse_RunContainer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_connector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_connector_proto_goTypes,
		DependencyIndexes: file_connector_proto_depIdxs,
		MessageInfos:      file_connector_proto_msgTypes,
	}.Build()
	File_connector_proto = out.File
	file_connector_proto_rawDesc = nil
	file_connector_proto_goTypes = nil
	file_connector_proto_depIdxs = nil
}
