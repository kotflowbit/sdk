// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kot

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

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	GetBlockHeight(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyBlock, error)
	GetLatest(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyLatest, error)
	GetBroken(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyBool, error)
	GetAddressValue(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyValue, error)
	GetAddressBalance(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyBalance, error)
	GetTransaction(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyTransaction, error)
	PutVerify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*ReplyBool, error)
	PutTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*ReplyBool, error)
	GetVerifyHeight(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyVerifys, error)
	HasTransaction(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyBool, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) GetBlockHeight(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyBlock, error) {
	out := new(ReplyBlock)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/GetBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetLatest(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyLatest, error) {
	out := new(ReplyLatest)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/GetLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetBroken(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyBool, error) {
	out := new(ReplyBool)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/GetBroken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetAddressValue(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyValue, error) {
	out := new(ReplyValue)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/GetAddressValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetAddressBalance(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyBalance, error) {
	out := new(ReplyBalance)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/GetAddressBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetTransaction(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyTransaction, error) {
	out := new(ReplyTransaction)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) PutVerify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*ReplyBool, error) {
	out := new(ReplyBool)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/PutVerify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) PutTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*ReplyBool, error) {
	out := new(ReplyBool)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/PutTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetVerifyHeight(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyVerifys, error) {
	out := new(ReplyVerifys)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/GetVerifyHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) HasTransaction(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyBool, error) {
	out := new(ReplyBool)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/HasTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	GetBlockHeight(context.Context, *HeightRequest) (*ReplyBlock, error)
	GetLatest(context.Context, *HeightRequest) (*ReplyLatest, error)
	GetBroken(context.Context, *HashRequest) (*ReplyBool, error)
	GetAddressValue(context.Context, *AddressRequest) (*ReplyValue, error)
	GetAddressBalance(context.Context, *AddressRequest) (*ReplyBalance, error)
	GetTransaction(context.Context, *HashRequest) (*ReplyTransaction, error)
	PutVerify(context.Context, *VerifyRequest) (*ReplyBool, error)
	PutTransaction(context.Context, *TransactionRequest) (*ReplyBool, error)
	GetVerifyHeight(context.Context, *HeightRequest) (*ReplyVerifys, error)
	HasTransaction(context.Context, *HashRequest) (*ReplyBool, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) GetBlockHeight(context.Context, *HeightRequest) (*ReplyBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeight not implemented")
}
func (UnimplementedGreeterServer) GetLatest(context.Context, *HeightRequest) (*ReplyLatest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatest not implemented")
}
func (UnimplementedGreeterServer) GetBroken(context.Context, *HashRequest) (*ReplyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBroken not implemented")
}
func (UnimplementedGreeterServer) GetAddressValue(context.Context, *AddressRequest) (*ReplyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressValue not implemented")
}
func (UnimplementedGreeterServer) GetAddressBalance(context.Context, *AddressRequest) (*ReplyBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressBalance not implemented")
}
func (UnimplementedGreeterServer) GetTransaction(context.Context, *HashRequest) (*ReplyTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedGreeterServer) PutVerify(context.Context, *VerifyRequest) (*ReplyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutVerify not implemented")
}
func (UnimplementedGreeterServer) PutTransaction(context.Context, *TransactionRequest) (*ReplyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutTransaction not implemented")
}
func (UnimplementedGreeterServer) GetVerifyHeight(context.Context, *HeightRequest) (*ReplyVerifys, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVerifyHeight not implemented")
}
func (UnimplementedGreeterServer) HasTransaction(context.Context, *HashRequest) (*ReplyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasTransaction not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_GetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/GetBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetBlockHeight(ctx, req.(*HeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/GetLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetLatest(ctx, req.(*HeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetBroken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetBroken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/GetBroken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetBroken(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetAddressValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetAddressValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/GetAddressValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetAddressValue(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetAddressBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetAddressBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/GetAddressBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetAddressBalance(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetTransaction(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_PutVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).PutVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/PutVerify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).PutVerify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_PutTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).PutTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/PutTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).PutTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetVerifyHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetVerifyHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/GetVerifyHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetVerifyHeight(ctx, req.(*HeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_HasTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).HasTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/HasTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).HasTransaction(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KotFlowBlock.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockHeight",
			Handler:    _Greeter_GetBlockHeight_Handler,
		},
		{
			MethodName: "GetLatest",
			Handler:    _Greeter_GetLatest_Handler,
		},
		{
			MethodName: "GetBroken",
			Handler:    _Greeter_GetBroken_Handler,
		},
		{
			MethodName: "GetAddressValue",
			Handler:    _Greeter_GetAddressValue_Handler,
		},
		{
			MethodName: "GetAddressBalance",
			Handler:    _Greeter_GetAddressBalance_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Greeter_GetTransaction_Handler,
		},
		{
			MethodName: "PutVerify",
			Handler:    _Greeter_PutVerify_Handler,
		},
		{
			MethodName: "PutTransaction",
			Handler:    _Greeter_PutTransaction_Handler,
		},
		{
			MethodName: "GetVerifyHeight",
			Handler:    _Greeter_GetVerifyHeight_Handler,
		},
		{
			MethodName: "HasTransaction",
			Handler:    _Greeter_HasTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/block.proto",
}
