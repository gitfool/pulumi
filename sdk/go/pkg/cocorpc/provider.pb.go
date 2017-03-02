// Code generated by protoc-gen-go.
// source: provider.proto
// DO NOT EDIT!

package cocorpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import google_protobuf1 "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NameRequest struct {
	Type       string                   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *NameRequest) Reset()                    { *m = NameRequest{} }
func (m *NameRequest) String() string            { return proto.CompactTextString(m) }
func (*NameRequest) ProtoMessage()               {}
func (*NameRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *NameRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NameRequest) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type NameResponse struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *NameResponse) Reset()                    { *m = NameResponse{} }
func (m *NameResponse) String() string            { return proto.CompactTextString(m) }
func (*NameResponse) ProtoMessage()               {}
func (*NameResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *NameResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateRequest struct {
	Type       string                   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Properties *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=properties" json:"properties,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *CreateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRequest) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type CreateResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ReadResponse struct {
	Properties *google_protobuf1.Struct `protobuf:"bytes,1,opt,name=properties" json:"properties,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ReadResponse) GetProperties() *google_protobuf1.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

type UpdateRequest struct {
	Id   string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string                   `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Olds *google_protobuf1.Struct `protobuf:"bytes,3,opt,name=olds" json:"olds,omitempty"`
	News *google_protobuf1.Struct `protobuf:"bytes,4,opt,name=news" json:"news,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateRequest) GetOlds() *google_protobuf1.Struct {
	if m != nil {
		return m.Olds
	}
	return nil
}

func (m *UpdateRequest) GetNews() *google_protobuf1.Struct {
	if m != nil {
		return m.News
	}
	return nil
}

type UpdateImpactResponse struct {
	Replace bool                     `protobuf:"varint,1,opt,name=replace" json:"replace,omitempty"`
	Impacts *google_protobuf1.Struct `protobuf:"bytes,2,opt,name=impacts" json:"impacts,omitempty"`
}

func (m *UpdateImpactResponse) Reset()                    { *m = UpdateImpactResponse{} }
func (m *UpdateImpactResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateImpactResponse) ProtoMessage()               {}
func (*UpdateImpactResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *UpdateImpactResponse) GetReplace() bool {
	if m != nil {
		return m.Replace
	}
	return false
}

func (m *UpdateImpactResponse) GetImpacts() *google_protobuf1.Struct {
	if m != nil {
		return m.Impacts
	}
	return nil
}

type DeleteRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*NameRequest)(nil), "cocorpc.NameRequest")
	proto.RegisterType((*NameResponse)(nil), "cocorpc.NameResponse")
	proto.RegisterType((*CreateRequest)(nil), "cocorpc.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "cocorpc.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "cocorpc.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "cocorpc.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "cocorpc.UpdateRequest")
	proto.RegisterType((*UpdateImpactResponse)(nil), "cocorpc.UpdateImpactResponse")
	proto.RegisterType((*DeleteRequest)(nil), "cocorpc.DeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceProvider service

type ResourceProviderClient interface {
	// Name names a given resource.  Sometimes this will be assigned by a developer, and so the provider
	// simply fetches it from the property bag; other times, the provider will assign this based on its own algorithm.
	// In any case, resources with the same name must be safe to use interchangeably with one another.
	Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read reads the instance state identified by ID, returning a populated resource object, or an error if not found.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update updates an existing resource with new values.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// UpdateImpact checks what impacts a hypothetical update will have on the resource's properties.
	UpdateImpact(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateImpactResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type resourceProviderClient struct {
	cc *grpc.ClientConn
}

func NewResourceProviderClient(cc *grpc.ClientConn) ResourceProviderClient {
	return &resourceProviderClient{cc}
}

func (c *resourceProviderClient) Name(ctx context.Context, in *NameRequest, opts ...grpc.CallOption) (*NameResponse, error) {
	out := new(NameResponse)
	err := grpc.Invoke(ctx, "/cocorpc.ResourceProvider/Name", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/cocorpc.ResourceProvider/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/cocorpc.ResourceProvider/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/cocorpc.ResourceProvider/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) UpdateImpact(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateImpactResponse, error) {
	out := new(UpdateImpactResponse)
	err := grpc.Invoke(ctx, "/cocorpc.ResourceProvider/UpdateImpact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceProviderClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/cocorpc.ResourceProvider/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceProvider service

type ResourceProviderServer interface {
	// Name names a given resource.  Sometimes this will be assigned by a developer, and so the provider
	// simply fetches it from the property bag; other times, the provider will assign this based on its own algorithm.
	// In any case, resources with the same name must be safe to use interchangeably with one another.
	Name(context.Context, *NameRequest) (*NameResponse, error)
	// Create allocates a new instance of the provided resource and returns its unique ID afterwards.  (The input ID
	// must be blank.)  If this call fails, the resource must not have been created (i.e., it is "transacational").
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read reads the instance state identified by ID, returning a populated resource object, or an error if not found.
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update updates an existing resource with new values.
	Update(context.Context, *UpdateRequest) (*google_protobuf.Empty, error)
	// UpdateImpact checks what impacts a hypothetical update will have on the resource's properties.
	UpdateImpact(context.Context, *UpdateRequest) (*UpdateImpactResponse, error)
	// Delete tears down an existing resource with the given ID.  If it fails, the resource is assumed to still exist.
	Delete(context.Context, *DeleteRequest) (*google_protobuf.Empty, error)
}

func RegisterResourceProviderServer(s *grpc.Server, srv ResourceProviderServer) {
	s.RegisterService(&_ResourceProvider_serviceDesc, srv)
}

func _ResourceProvider_Name_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Name(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cocorpc.ResourceProvider/Name",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Name(ctx, req.(*NameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cocorpc.ResourceProvider/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cocorpc.ResourceProvider/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cocorpc.ResourceProvider/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_UpdateImpact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).UpdateImpact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cocorpc.ResourceProvider/UpdateImpact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).UpdateImpact(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceProvider_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceProviderServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cocorpc.ResourceProvider/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceProviderServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cocorpc.ResourceProvider",
	HandlerType: (*ResourceProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Name",
			Handler:    _ResourceProvider_Name_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResourceProvider_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ResourceProvider_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceProvider_Update_Handler,
		},
		{
			MethodName: "UpdateImpact",
			Handler:    _ResourceProvider_UpdateImpact_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceProvider_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}

func init() { proto.RegisterFile("provider.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x6d, 0x62, 0x68, 0xf5, 0xf6, 0x03, 0x19, 0x9e, 0xef, 0x85, 0xa8, 0x50, 0x66, 0xf5, 0x40,
	0xc8, 0xa3, 0x2d, 0x45, 0xd0, 0xa5, 0x4a, 0x71, 0x23, 0x12, 0x71, 0x23, 0x6e, 0xd2, 0xc9, 0xb5,
	0x04, 0x92, 0xce, 0x38, 0x33, 0x51, 0xfa, 0x23, 0xfc, 0xcb, 0x22, 0xc9, 0x24, 0x71, 0x92, 0x52,
	0x5b, 0x17, 0x6f, 0x97, 0xdc, 0x39, 0x73, 0xee, 0xb9, 0x73, 0xce, 0x85, 0x99, 0x90, 0xfc, 0x47,
	0x9a, 0xa0, 0x0c, 0x85, 0xe4, 0x9a, 0x93, 0x11, 0xe3, 0x8c, 0x4b, 0xc1, 0x82, 0xa7, 0x3b, 0xce,
	0x77, 0x19, 0xde, 0x55, 0xe5, 0x6d, 0xf1, 0xed, 0x0e, 0x73, 0xa1, 0x0f, 0x06, 0x15, 0x3c, 0xeb,
	0x1f, 0x2a, 0x2d, 0x0b, 0xa6, 0xcd, 0x29, 0xfd, 0x02, 0xe3, 0x0f, 0x71, 0x8e, 0x11, 0x7e, 0x2f,
	0x50, 0x69, 0x42, 0xc0, 0xd3, 0x07, 0x81, 0xbe, 0x33, 0x77, 0x6e, 0x1f, 0x45, 0xd5, 0x37, 0x79,
	0x09, 0x20, 0x24, 0x17, 0x28, 0x75, 0x8a, 0xca, 0x77, 0xe7, 0xce, 0xed, 0x78, 0x79, 0x13, 0x1a,
	0xd6, 0xb0, 0x61, 0x0d, 0x3f, 0x55, 0xac, 0x91, 0x05, 0xa5, 0x14, 0x26, 0x86, 0x5b, 0x09, 0xbe,
	0x57, 0x58, 0x92, 0xef, 0xe3, 0xbc, 0x25, 0x2f, 0xbf, 0xe9, 0x57, 0x98, 0xbe, 0x91, 0x18, 0xeb,
	0xfb, 0x51, 0x30, 0x87, 0x59, 0xc3, 0x5e, 0x6b, 0x98, 0x81, 0x9b, 0x26, 0x35, 0xb9, 0x9b, 0x26,
	0x74, 0x01, 0xe3, 0x08, 0xe3, 0xa4, 0xe9, 0xde, 0x3b, 0x6e, 0xd5, 0xb8, 0x7f, 0xd5, 0xd0, 0x0d,
	0x4c, 0xcc, 0x95, 0x9a, 0xb2, 0xab, 0xce, 0xb9, 0x5c, 0xdd, 0x2f, 0x07, 0xa6, 0x9f, 0x45, 0x62,
	0x0d, 0x7f, 0x41, 0x7b, 0xf2, 0x02, 0x3c, 0x9e, 0x25, 0xca, 0x7f, 0xf0, 0xef, 0x46, 0x15, 0xa8,
	0x04, 0xef, 0xf1, 0xa7, 0xf2, 0xbd, 0x33, 0xe0, 0x12, 0x44, 0x19, 0x5c, 0x19, 0x39, 0xef, 0x73,
	0x11, 0x33, 0xdd, 0x0e, 0xe8, 0xc3, 0x48, 0xa2, 0xc8, 0x62, 0x66, 0x5c, 0x79, 0x18, 0x35, 0xbf,
	0x64, 0x01, 0xa3, 0xb4, 0xc2, 0x9e, 0x75, 0xa5, 0xc1, 0xd1, 0x15, 0x4c, 0xdf, 0x62, 0x86, 0xff,
	0x35, 0xf3, 0xf2, 0xb7, 0x0b, 0x8f, 0x23, 0x54, 0xbc, 0x90, 0x0c, 0x3f, 0xd6, 0x4b, 0x40, 0xd6,
	0xe0, 0x95, 0xf1, 0x22, 0x57, 0x61, 0xbd, 0x07, 0xa1, 0x95, 0xe4, 0xe0, 0x49, 0xaf, 0x6a, 0x66,
	0xa1, 0x03, 0xf2, 0x1a, 0x86, 0x26, 0x13, 0xe4, 0xba, 0x85, 0x74, 0x22, 0x18, 0xdc, 0x1c, 0xd5,
	0xdb, 0xcb, 0x6b, 0xf0, 0x4a, 0xef, 0xad, 0x9e, 0x56, 0x7a, 0xac, 0x9e, 0x76, 0x40, 0xe8, 0x80,
	0xbc, 0x82, 0xa1, 0x79, 0x59, 0xab, 0x67, 0xc7, 0xf9, 0xe0, 0xfa, 0xe8, 0xe1, 0xde, 0x95, 0x3b,
	0x4c, 0x07, 0x64, 0x03, 0x13, 0xdb, 0x95, 0x93, 0x0c, 0xcf, 0x7b, 0xf5, 0xae, 0x89, 0x46, 0x84,
	0x79, 0x79, 0x8b, 0xa2, 0x63, 0xc5, 0x69, 0x11, 0xdb, 0x61, 0x55, 0x59, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xf3, 0xeb, 0x1f, 0x7c, 0x83, 0x04, 0x00, 0x00,
}
