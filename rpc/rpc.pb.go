// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PeerTopicInfo struct {
	Peer                 string   `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`
	Topics               []string `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerTopicInfo) Reset()         { *m = PeerTopicInfo{} }
func (m *PeerTopicInfo) String() string { return proto.CompactTextString(m) }
func (*PeerTopicInfo) ProtoMessage()    {}
func (*PeerTopicInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *PeerTopicInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerTopicInfo.Unmarshal(m, b)
}
func (m *PeerTopicInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerTopicInfo.Marshal(b, m, deterministic)
}
func (m *PeerTopicInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerTopicInfo.Merge(m, src)
}
func (m *PeerTopicInfo) XXX_Size() int {
	return xxx_messageInfo_PeerTopicInfo.Size(m)
}
func (m *PeerTopicInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerTopicInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerTopicInfo proto.InternalMessageInfo

func (m *PeerTopicInfo) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *PeerTopicInfo) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type PublishData struct {
	Info                 *PeerTopicInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Data                 *Data          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PublishData) Reset()         { *m = PublishData{} }
func (m *PublishData) String() string { return proto.CompactTextString(m) }
func (*PublishData) ProtoMessage()    {}
func (*PublishData) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *PublishData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishData.Unmarshal(m, b)
}
func (m *PublishData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishData.Marshal(b, m, deterministic)
}
func (m *PublishData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishData.Merge(m, src)
}
func (m *PublishData) XXX_Size() int {
	return xxx_messageInfo_PublishData.Size(m)
}
func (m *PublishData) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishData.DiscardUnknown(m)
}

var xxx_messageInfo_PublishData proto.InternalMessageInfo

func (m *PublishData) GetInfo() *PeerTopicInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PublishData) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Response struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Data struct {
	Topic                []string `protobuf:"bytes,1,rep,name=topic,proto3" json:"topic,omitempty"`
	Raw                  []byte   `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetTopic() []string {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Data) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerTopicInfo)(nil), "rpc.PeerTopicInfo")
	proto.RegisterType((*PublishData)(nil), "rpc.PublishData")
	proto.RegisterType((*Response)(nil), "rpc.Response")
	proto.RegisterType((*Data)(nil), "rpc.Data")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x9b, 0x26, 0x5f, 0xfb, 0x65, 0x6a, 0xa1, 0x0e, 0x22, 0x8b, 0x20, 0x94, 0x45, 0xa4,
	0x07, 0x09, 0x52, 0x2f, 0x82, 0x57, 0x2f, 0x1e, 0x84, 0xb2, 0xf6, 0x05, 0x36, 0x71, 0xda, 0x2e,
	0x68, 0x76, 0xd9, 0xdd, 0xe2, 0xdb, 0xf8, 0xac, 0xb2, 0xd3, 0x84, 0x22, 0xf6, 0x36, 0x33, 0xf9,
	0x65, 0xfe, 0xbf, 0x61, 0xa1, 0xf4, 0xae, 0xa9, 0x9c, 0xb7, 0xd1, 0x62, 0xee, 0x5d, 0x23, 0x9f,
	0x60, 0xba, 0x22, 0xf2, 0x6b, 0xeb, 0x4c, 0xf3, 0xd2, 0x6e, 0x2c, 0x22, 0x14, 0x8e, 0xc8, 0x8b,
	0x6c, 0x9e, 0x2d, 0x4a, 0xc5, 0x35, 0x5e, 0xc2, 0x28, 0x26, 0x20, 0x88, 0xe1, 0x3c, 0x5f, 0x94,
	0xaa, 0xeb, 0xe4, 0x1a, 0x26, 0xab, 0x7d, 0xfd, 0x61, 0xc2, 0xee, 0x59, 0x47, 0x8d, 0xb7, 0x50,
	0x98, 0x76, 0x63, 0xf9, 0xd7, 0xc9, 0x12, 0xab, 0x14, 0xf5, 0x6b, 0xb9, 0xe2, 0xef, 0x78, 0x0d,
	0xc5, 0xbb, 0x8e, 0x5a, 0x0c, 0x99, 0x2b, 0x99, 0x4b, 0x0b, 0x14, 0x8f, 0xe5, 0x0d, 0xfc, 0x57,
	0x14, 0x9c, 0x6d, 0x03, 0xa1, 0x80, 0xf1, 0x27, 0x85, 0xa0, 0xb7, 0xd4, 0x09, 0xf5, 0xad, 0xac,
	0xa0, 0xe0, 0xd0, 0x0b, 0xf8, 0xc7, 0x36, 0x22, 0x63, 0xb5, 0x43, 0x83, 0x33, 0xc8, 0xbd, 0xfe,
	0xe2, 0x84, 0x33, 0x95, 0xca, 0xe5, 0x77, 0x06, 0xc5, 0x2b, 0x85, 0x1d, 0x3e, 0xc2, 0xb9, 0xa2,
	0xad, 0x09, 0x31, 0x89, 0x75, 0xfa, 0x78, 0x42, 0xf6, 0x6a, 0xca, 0xb3, 0x5e, 0x45, 0x0e, 0xf0,
	0x0e, 0xc6, 0x3d, 0x3f, 0x3b, 0xf0, 0xc7, 0xe3, 0xff, 0xd2, 0x15, 0x94, 0x6f, 0xfb, 0x3a, 0x34,
	0xde, 0xd4, 0x74, 0x72, 0xff, 0xf1, 0x70, 0x39, 0xb8, 0xcf, 0xea, 0x11, 0xbf, 0xca, 0xc3, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xc3, 0xd3, 0xe2, 0xa2, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MeshClient is the client API for Mesh service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MeshClient interface {
	RegisterToPublish(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error)
	Publish(ctx context.Context, in *PublishData, opts ...grpc.CallOption) (*Response, error)
	Subscribe(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (Mesh_SubscribeClient, error)
}

type meshClient struct {
	cc *grpc.ClientConn
}

func NewMeshClient(cc *grpc.ClientConn) MeshClient {
	return &meshClient{cc}
}

func (c *meshClient) RegisterToPublish(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Mesh/RegisterToPublish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) Publish(ctx context.Context, in *PublishData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/rpc.Mesh/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshClient) Subscribe(ctx context.Context, in *PeerTopicInfo, opts ...grpc.CallOption) (Mesh_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Mesh_serviceDesc.Streams[0], "/rpc.Mesh/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &meshSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mesh_SubscribeClient interface {
	Recv() (*Data, error)
	grpc.ClientStream
}

type meshSubscribeClient struct {
	grpc.ClientStream
}

func (x *meshSubscribeClient) Recv() (*Data, error) {
	m := new(Data)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MeshServer is the server API for Mesh service.
type MeshServer interface {
	RegisterToPublish(context.Context, *PeerTopicInfo) (*Response, error)
	Publish(context.Context, *PublishData) (*Response, error)
	Subscribe(*PeerTopicInfo, Mesh_SubscribeServer) error
}

func RegisterMeshServer(s *grpc.Server, srv MeshServer) {
	s.RegisterService(&_Mesh_serviceDesc, srv)
}

func _Mesh_RegisterToPublish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerTopicInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).RegisterToPublish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Mesh/RegisterToPublish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).RegisterToPublish(ctx, req.(*PeerTopicInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Mesh/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshServer).Publish(ctx, req.(*PublishData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mesh_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PeerTopicInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeshServer).Subscribe(m, &meshSubscribeServer{stream})
}

type Mesh_SubscribeServer interface {
	Send(*Data) error
	grpc.ServerStream
}

type meshSubscribeServer struct {
	grpc.ServerStream
}

func (x *meshSubscribeServer) Send(m *Data) error {
	return x.ServerStream.SendMsg(m)
}

var _Mesh_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Mesh",
	HandlerType: (*MeshServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterToPublish",
			Handler:    _Mesh_RegisterToPublish_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _Mesh_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Mesh_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc.proto",
}
