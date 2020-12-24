// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: gun.proto

package proto

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

type Hunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Hunk) Reset() {
	*x = Hunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gun_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hunk) ProtoMessage() {}

func (x *Hunk) ProtoReflect() protoreflect.Message {
	mi := &file_gun_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hunk.ProtoReflect.Descriptor instead.
func (*Hunk) Descriptor() ([]byte, []int) {
	return file_gun_proto_rawDescGZIP(), []int{0}
}

func (x *Hunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_gun_proto protoreflect.FileDescriptor

var file_gun_proto_rawDesc = []byte{
	0x0a, 0x09, 0x67, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x04, 0x48,
	0x75, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x25, 0x0a, 0x0a, 0x47, 0x75, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x03, 0x54, 0x75, 0x6e, 0x12, 0x05, 0x2e, 0x48,
	0x75, 0x6e, 0x6b, 0x1a, 0x05, 0x2e, 0x48, 0x75, 0x6e, 0x6b, 0x28, 0x01, 0x30, 0x01, 0x42, 0x21,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x51, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2f, 0x67, 0x75, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gun_proto_rawDescOnce sync.Once
	file_gun_proto_rawDescData = file_gun_proto_rawDesc
)

func file_gun_proto_rawDescGZIP() []byte {
	file_gun_proto_rawDescOnce.Do(func() {
		file_gun_proto_rawDescData = protoimpl.X.CompressGZIP(file_gun_proto_rawDescData)
	})
	return file_gun_proto_rawDescData
}

var file_gun_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gun_proto_goTypes = []interface{}{
	(*Hunk)(nil), // 0: Hunk
}
var file_gun_proto_depIdxs = []int32{
	0, // 0: GunService.Tun:input_type -> Hunk
	0, // 1: GunService.Tun:output_type -> Hunk
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gun_proto_init() }
func file_gun_proto_init() {
	if File_gun_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gun_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hunk); i {
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
			RawDescriptor: file_gun_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gun_proto_goTypes,
		DependencyIndexes: file_gun_proto_depIdxs,
		MessageInfos:      file_gun_proto_msgTypes,
	}.Build()
	File_gun_proto = out.File
	file_gun_proto_rawDesc = nil
	file_gun_proto_goTypes = nil
	file_gun_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GunServiceClient is the client API for GunService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GunServiceClient interface {
	Tun(ctx context.Context, opts ...grpc.CallOption) (GunService_TunClient, error)
}

type gunServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGunServiceClient(cc grpc.ClientConnInterface) GunServiceClient {
	return &gunServiceClient{cc}
}

func (c *gunServiceClient) Tun(ctx context.Context, opts ...grpc.CallOption) (GunService_TunClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GunService_serviceDesc.Streams[0], "/GunService/Tun", opts...)
	if err != nil {
		return nil, err
	}
	x := &gunServiceTunClient{stream}
	return x, nil
}

type GunService_TunClient interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ClientStream
}

type gunServiceTunClient struct {
	grpc.ClientStream
}

func (x *gunServiceTunClient) Send(m *Hunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gunServiceTunClient) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GunServiceServer is the server API for GunService service.
type GunServiceServer interface {
	Tun(GunService_TunServer) error
}

// UnimplementedGunServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGunServiceServer struct {
}

func (*UnimplementedGunServiceServer) Tun(GunService_TunServer) error {
	return status.Errorf(codes.Unimplemented, "method Tun not implemented")
}

func RegisterGunServiceServer(s *grpc.Server, srv GunServiceServer) {
	s.RegisterService(&_GunService_serviceDesc, srv)
}

func _GunService_Tun_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GunServiceServer).Tun(&gunServiceTunServer{stream})
}

type GunService_TunServer interface {
	Send(*Hunk) error
	Recv() (*Hunk, error)
	grpc.ServerStream
}

type gunServiceTunServer struct {
	grpc.ServerStream
}

func (x *gunServiceTunServer) Send(m *Hunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gunServiceTunServer) Recv() (*Hunk, error) {
	m := new(Hunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GunService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "GunService",
	HandlerType: (*GunServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tun",
			Handler:       _GunService_Tun_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gun.proto",
}