// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ItemServiceClient is the client API for ItemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ItemServiceClient interface {
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error)
	ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error)
}

type itemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewItemServiceClient(cc grpc.ClientConnInterface) ItemServiceClient {
	return &itemServiceClient{cc}
}

func (c *itemServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, "/mercari.mercari_microservices_example.item.ItemService/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*GetItemResponse, error) {
	out := new(GetItemResponse)
	err := c.cc.Invoke(ctx, "/mercari.mercari_microservices_example.item.ItemService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *itemServiceClient) ListItems(ctx context.Context, in *ListItemsRequest, opts ...grpc.CallOption) (*ListItemsResponse, error) {
	out := new(ListItemsResponse)
	err := c.cc.Invoke(ctx, "/mercari.mercari_microservices_example.item.ItemService/ListItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ItemServiceServer is the server API for ItemService service.
// All implementations must embed UnimplementedItemServiceServer
// for forward compatibility
type ItemServiceServer interface {
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
	GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error)
	ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error)
	mustEmbedUnimplementedItemServiceServer()
}

// UnimplementedItemServiceServer must be embedded to have forward compatible implementations.
type UnimplementedItemServiceServer struct {
}

func (UnimplementedItemServiceServer) CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedItemServiceServer) GetItem(context.Context, *GetItemRequest) (*GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}
func (UnimplementedItemServiceServer) ListItems(context.Context, *ListItemsRequest) (*ListItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedItemServiceServer) mustEmbedUnimplementedItemServiceServer() {}

// UnsafeItemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ItemServiceServer will
// result in compilation errors.
type UnsafeItemServiceServer interface {
	mustEmbedUnimplementedItemServiceServer()
}

func RegisterItemServiceServer(s grpc.ServiceRegistrar, srv ItemServiceServer) {
	s.RegisterService(&ItemService_ServiceDesc, srv)
}

func _ItemService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.mercari_microservices_example.item.ItemService/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.mercari_microservices_example.item.ItemService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).GetItem(ctx, req.(*GetItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ItemService_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ItemServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mercari.mercari_microservices_example.item.ItemService/ListItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ItemServiceServer).ListItems(ctx, req.(*ListItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ItemService_ServiceDesc is the grpc.ServiceDesc for ItemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ItemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mercari.mercari_microservices_example.item.ItemService",
	HandlerType: (*ItemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItem",
			Handler:    _ItemService_CreateItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _ItemService_GetItem_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _ItemService_ListItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/item/proto/item.proto",
}
