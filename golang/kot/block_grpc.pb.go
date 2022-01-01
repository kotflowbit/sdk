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
	BlockHeight(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyBlock, error)
	Latest(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyLatest, error)
	CheckBroken(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyBool, error)
	AddressValue(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyValue, error)
	AddressBalance(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyBalance, error)
	GetTransaction(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyTransaction, error)
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*ReplyBool, error)
	Transaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*ReplyBool, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) BlockHeight(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyBlock, error) {
	out := new(ReplyBlock)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/BlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Latest(ctx context.Context, in *HeightRequest, opts ...grpc.CallOption) (*ReplyLatest, error) {
	out := new(ReplyLatest)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/Latest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) CheckBroken(ctx context.Context, in *HashRequest, opts ...grpc.CallOption) (*ReplyBool, error) {
	out := new(ReplyBool)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/CheckBroken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) AddressValue(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyValue, error) {
	out := new(ReplyValue)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/AddressValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) AddressBalance(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*ReplyBalance, error) {
	out := new(ReplyBalance)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/AddressBalance", in, out, opts...)
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

func (c *greeterClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*ReplyBool, error) {
	out := new(ReplyBool)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Transaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*ReplyBool, error) {
	out := new(ReplyBool)
	err := c.cc.Invoke(ctx, "/KotFlowBlock.Greeter/Transaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	BlockHeight(context.Context, *HeightRequest) (*ReplyBlock, error)
	Latest(context.Context, *HeightRequest) (*ReplyLatest, error)
	CheckBroken(context.Context, *HashRequest) (*ReplyBool, error)
	AddressValue(context.Context, *AddressRequest) (*ReplyValue, error)
	AddressBalance(context.Context, *AddressRequest) (*ReplyBalance, error)
	GetTransaction(context.Context, *HashRequest) (*ReplyTransaction, error)
	Verify(context.Context, *VerifyRequest) (*ReplyBool, error)
	Transaction(context.Context, *TransactionRequest) (*ReplyBool, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) BlockHeight(context.Context, *HeightRequest) (*ReplyBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockHeight not implemented")
}
func (UnimplementedGreeterServer) Latest(context.Context, *HeightRequest) (*ReplyLatest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Latest not implemented")
}
func (UnimplementedGreeterServer) CheckBroken(context.Context, *HashRequest) (*ReplyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBroken not implemented")
}
func (UnimplementedGreeterServer) AddressValue(context.Context, *AddressRequest) (*ReplyValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressValue not implemented")
}
func (UnimplementedGreeterServer) AddressBalance(context.Context, *AddressRequest) (*ReplyBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressBalance not implemented")
}
func (UnimplementedGreeterServer) GetTransaction(context.Context, *HashRequest) (*ReplyTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedGreeterServer) Verify(context.Context, *VerifyRequest) (*ReplyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedGreeterServer) Transaction(context.Context, *TransactionRequest) (*ReplyBool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transaction not implemented")
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

func _Greeter_BlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).BlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/BlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).BlockHeight(ctx, req.(*HeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Latest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Latest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/Latest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Latest(ctx, req.(*HeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_CheckBroken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CheckBroken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/CheckBroken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CheckBroken(ctx, req.(*HashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_AddressValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).AddressValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/AddressValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).AddressValue(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_AddressBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).AddressBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/AddressBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).AddressBalance(ctx, req.(*AddressRequest))
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

func _Greeter_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Transaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Transaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KotFlowBlock.Greeter/Transaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Transaction(ctx, req.(*TransactionRequest))
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
			MethodName: "BlockHeight",
			Handler:    _Greeter_BlockHeight_Handler,
		},
		{
			MethodName: "Latest",
			Handler:    _Greeter_Latest_Handler,
		},
		{
			MethodName: "CheckBroken",
			Handler:    _Greeter_CheckBroken_Handler,
		},
		{
			MethodName: "AddressValue",
			Handler:    _Greeter_AddressValue_Handler,
		},
		{
			MethodName: "AddressBalance",
			Handler:    _Greeter_AddressBalance_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Greeter_GetTransaction_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Greeter_Verify_Handler,
		},
		{
			MethodName: "Transaction",
			Handler:    _Greeter_Transaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/block.proto",
}
