// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: flyteidl/cacheservice/cacheservice.proto

package cacheservice

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

const (
	CacheService_Get_FullMethodName                    = "/flyteidl.cacheservice.CacheService/Get"
	CacheService_Put_FullMethodName                    = "/flyteidl.cacheservice.CacheService/Put"
	CacheService_Delete_FullMethodName                 = "/flyteidl.cacheservice.CacheService/Delete"
	CacheService_GetOrExtendReservation_FullMethodName = "/flyteidl.cacheservice.CacheService/GetOrExtendReservation"
	CacheService_ReleaseReservation_FullMethodName     = "/flyteidl.cacheservice.CacheService/ReleaseReservation"
)

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheServiceClient interface {
	// Retrieves cached data by key.
	Get(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*GetCacheResponse, error)
	// Stores or updates cached data by key.
	Put(ctx context.Context, in *PutCacheRequest, opts ...grpc.CallOption) (*PutCacheResponse, error)
	// Deletes cached data by key.
	Delete(ctx context.Context, in *DeleteCacheRequest, opts ...grpc.CallOption) (*DeleteCacheResponse, error)
	// Get or extend a reservation for a cache key
	GetOrExtendReservation(ctx context.Context, in *GetOrExtendReservationRequest, opts ...grpc.CallOption) (*GetOrExtendReservationResponse, error)
	// Release the reservation for a cache key
	ReleaseReservation(ctx context.Context, in *ReleaseReservationRequest, opts ...grpc.CallOption) (*ReleaseReservationResponse, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Get(ctx context.Context, in *GetCacheRequest, opts ...grpc.CallOption) (*GetCacheResponse, error) {
	out := new(GetCacheResponse)
	err := c.cc.Invoke(ctx, CacheService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Put(ctx context.Context, in *PutCacheRequest, opts ...grpc.CallOption) (*PutCacheResponse, error) {
	out := new(PutCacheResponse)
	err := c.cc.Invoke(ctx, CacheService_Put_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Delete(ctx context.Context, in *DeleteCacheRequest, opts ...grpc.CallOption) (*DeleteCacheResponse, error) {
	out := new(DeleteCacheResponse)
	err := c.cc.Invoke(ctx, CacheService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) GetOrExtendReservation(ctx context.Context, in *GetOrExtendReservationRequest, opts ...grpc.CallOption) (*GetOrExtendReservationResponse, error) {
	out := new(GetOrExtendReservationResponse)
	err := c.cc.Invoke(ctx, CacheService_GetOrExtendReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) ReleaseReservation(ctx context.Context, in *ReleaseReservationRequest, opts ...grpc.CallOption) (*ReleaseReservationResponse, error) {
	out := new(ReleaseReservationResponse)
	err := c.cc.Invoke(ctx, CacheService_ReleaseReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
// All implementations should embed UnimplementedCacheServiceServer
// for forward compatibility
type CacheServiceServer interface {
	// Retrieves cached data by key.
	Get(context.Context, *GetCacheRequest) (*GetCacheResponse, error)
	// Stores or updates cached data by key.
	Put(context.Context, *PutCacheRequest) (*PutCacheResponse, error)
	// Deletes cached data by key.
	Delete(context.Context, *DeleteCacheRequest) (*DeleteCacheResponse, error)
	// Get or extend a reservation for a cache key
	GetOrExtendReservation(context.Context, *GetOrExtendReservationRequest) (*GetOrExtendReservationResponse, error)
	// Release the reservation for a cache key
	ReleaseReservation(context.Context, *ReleaseReservationRequest) (*ReleaseReservationResponse, error)
}

// UnimplementedCacheServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (UnimplementedCacheServiceServer) Get(context.Context, *GetCacheRequest) (*GetCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacheServiceServer) Put(context.Context, *PutCacheRequest) (*PutCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedCacheServiceServer) Delete(context.Context, *DeleteCacheRequest) (*DeleteCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCacheServiceServer) GetOrExtendReservation(context.Context, *GetOrExtendReservationRequest) (*GetOrExtendReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrExtendReservation not implemented")
}
func (UnimplementedCacheServiceServer) ReleaseReservation(context.Context, *ReleaseReservationRequest) (*ReleaseReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseReservation not implemented")
}

// UnsafeCacheServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheServiceServer will
// result in compilation errors.
type UnsafeCacheServiceServer interface {
	mustEmbedUnimplementedCacheServiceServer()
}

func RegisterCacheServiceServer(s grpc.ServiceRegistrar, srv CacheServiceServer) {
	s.RegisterService(&CacheService_ServiceDesc, srv)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*GetCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Put(ctx, req.(*PutCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Delete(ctx, req.(*DeleteCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_GetOrExtendReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrExtendReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).GetOrExtendReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_GetOrExtendReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).GetOrExtendReservation(ctx, req.(*GetOrExtendReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_ReleaseReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).ReleaseReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheService_ReleaseReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).ReleaseReservation(ctx, req.(*ReleaseReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheService_ServiceDesc is the grpc.ServiceDesc for CacheService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.cacheservice.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _CacheService_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CacheService_Delete_Handler,
		},
		{
			MethodName: "GetOrExtendReservation",
			Handler:    _CacheService_GetOrExtendReservation_Handler,
		},
		{
			MethodName: "ReleaseReservation",
			Handler:    _CacheService_ReleaseReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/cacheservice/cacheservice.proto",
}