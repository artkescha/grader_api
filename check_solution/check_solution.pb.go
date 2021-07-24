// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: check_solution.proto

package check_solution

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

type GradingResponse_Status int32

const (
	GradingResponse_OK   GradingResponse_Status = 0
	GradingResponse_FAIL GradingResponse_Status = 1
)

// Enum value maps for GradingResponse_Status.
var (
	GradingResponse_Status_name = map[int32]string{
		0: "OK",
		1: "FAIL",
	}
	GradingResponse_Status_value = map[string]int32{
		"OK":   0,
		"FAIL": 1,
	}
)

func (x GradingResponse_Status) Enum() *GradingResponse_Status {
	p := new(GradingResponse_Status)
	*p = x
	return p
}

func (x GradingResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GradingResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_check_solution_proto_enumTypes[0].Descriptor()
}

func (GradingResponse_Status) Type() protoreflect.EnumType {
	return &file_check_solution_proto_enumTypes[0]
}

func (x GradingResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GradingResponse_Status.Descriptor instead.
func (GradingResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_check_solution_proto_rawDescGZIP(), []int{3, 0}
}

type GradingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files   []*File  `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	Payload *Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *GradingRequest) Reset() {
	*x = GradingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_solution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradingRequest) ProtoMessage() {}

func (x *GradingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_check_solution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradingRequest.ProtoReflect.Descriptor instead.
func (*GradingRequest) Descriptor() ([]byte, []int) {
	return file_check_solution_proto_rawDescGZIP(), []int{0}
}

func (x *GradingRequest) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *GradingRequest) GetPayload() *Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label    string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_solution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_check_solution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_check_solution_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *File) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Container string `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	PartId    string `protobuf:"bytes,2,opt,name=partId,proto3" json:"partId,omitempty"`
	TaskId    int64  `protobuf:"varint,3,opt,name=taskId,proto3" json:"taskId,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_solution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_check_solution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_check_solution_proto_rawDescGZIP(), []int{2}
}

func (x *Payload) GetContainer() string {
	if x != nil {
		return x.Container
	}
	return ""
}

func (x *Payload) GetPartId() string {
	if x != nil {
		return x.PartId
	}
	return ""
}

func (x *Payload) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GradingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      GradingResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=check_solution.GradingResponse_Status" json:"status,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GradingResponse) Reset() {
	*x = GradingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_check_solution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GradingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GradingResponse) ProtoMessage() {}

func (x *GradingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_check_solution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GradingResponse.ProtoReflect.Descriptor instead.
func (*GradingResponse) Descriptor() ([]byte, []int) {
	return file_check_solution_proto_rawDescGZIP(), []int{3}
}

func (x *GradingResponse) GetStatus() GradingResponse_Status {
	if x != nil {
		return x.Status
	}
	return GradingResponse_OK
}

func (x *GradingResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_check_solution_proto protoreflect.FileDescriptor

var file_check_solution_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6f, 0x0a, 0x0e, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x38, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x57, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x47,
	0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x32, 0x5c, 0x0a, 0x0e,
	0x67, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a,
	0x0a, 0x07, 0x47, 0x72, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_check_solution_proto_rawDescOnce sync.Once
	file_check_solution_proto_rawDescData = file_check_solution_proto_rawDesc
)

func file_check_solution_proto_rawDescGZIP() []byte {
	file_check_solution_proto_rawDescOnce.Do(func() {
		file_check_solution_proto_rawDescData = protoimpl.X.CompressGZIP(file_check_solution_proto_rawDescData)
	})
	return file_check_solution_proto_rawDescData
}

var file_check_solution_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_check_solution_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_check_solution_proto_goTypes = []interface{}{
	(GradingResponse_Status)(0), // 0: check_solution.GradingResponse.Status
	(*GradingRequest)(nil),      // 1: check_solution.GradingRequest
	(*File)(nil),                // 2: check_solution.File
	(*Payload)(nil),             // 3: check_solution.Payload
	(*GradingResponse)(nil),     // 4: check_solution.GradingResponse
}
var file_check_solution_proto_depIdxs = []int32{
	2, // 0: check_solution.GradingRequest.files:type_name -> check_solution.File
	3, // 1: check_solution.GradingRequest.payload:type_name -> check_solution.Payload
	0, // 2: check_solution.GradingResponse.status:type_name -> check_solution.GradingResponse.Status
	1, // 3: check_solution.gradingService.Grading:input_type -> check_solution.GradingRequest
	4, // 4: check_solution.gradingService.Grading:output_type -> check_solution.GradingResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_check_solution_proto_init() }
func file_check_solution_proto_init() {
	if File_check_solution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_check_solution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradingRequest); i {
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
		file_check_solution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_check_solution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_check_solution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GradingResponse); i {
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
			RawDescriptor: file_check_solution_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_check_solution_proto_goTypes,
		DependencyIndexes: file_check_solution_proto_depIdxs,
		EnumInfos:         file_check_solution_proto_enumTypes,
		MessageInfos:      file_check_solution_proto_msgTypes,
	}.Build()
	File_check_solution_proto = out.File
	file_check_solution_proto_rawDesc = nil
	file_check_solution_proto_goTypes = nil
	file_check_solution_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GradingServiceClient is the client API for GradingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GradingServiceClient interface {
	Grading(ctx context.Context, in *GradingRequest, opts ...grpc.CallOption) (*GradingResponse, error)
}

type gradingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGradingServiceClient(cc grpc.ClientConnInterface) GradingServiceClient {
	return &gradingServiceClient{cc}
}

func (c *gradingServiceClient) Grading(ctx context.Context, in *GradingRequest, opts ...grpc.CallOption) (*GradingResponse, error) {
	out := new(GradingResponse)
	err := c.cc.Invoke(ctx, "/check_solution.gradingService/Grading", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GradingServiceServer is the server API for GradingService service.
type GradingServiceServer interface {
	Grading(context.Context, *GradingRequest) (*GradingResponse, error)
}

// UnimplementedGradingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGradingServiceServer struct {
}

func (*UnimplementedGradingServiceServer) Grading(context.Context, *GradingRequest) (*GradingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grading not implemented")
}

func RegisterGradingServiceServer(s *grpc.Server, srv GradingServiceServer) {
	s.RegisterService(&_GradingService_serviceDesc, srv)
}

func _GradingService_Grading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GradingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GradingServiceServer).Grading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/check_solution.gradingService/Grading",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GradingServiceServer).Grading(ctx, req.(*GradingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GradingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "check_solution.gradingService",
	HandlerType: (*GradingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grading",
			Handler:    _GradingService_Grading_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "check_solution.proto",
}
