// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.1
// source: user.proto

package __

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

// UserRPCClient is the client API for UserRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserRPCClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*ID, error)
	UpdateProfile(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*Nothing, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Nothing, error)
	GetFavorites(ctx context.Context, in *ID, opts ...grpc.CallOption) (*GetFavoritesResponse, error)
	SetFavorite(ctx context.Context, in *HandleFavorite, opts ...grpc.CallOption) (*Nothing, error)
	ResetFavorite(ctx context.Context, in *HandleFavorite, opts ...grpc.CallOption) (*Nothing, error)
	CheckFavorite(ctx context.Context, in *HandleFavorite, opts ...grpc.CallOption) (*Nothing, error)
	FindByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*UserData, error)
	FindByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*UserData, error)
	Subscribe(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionID, error)
	UpdateSubscribtionStatus(ctx context.Context, in *SubscriptionID, opts ...grpc.CallOption) (*Nothing, error)
}

type userRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRPCClient(cc grpc.ClientConnInterface) UserRPCClient {
	return &userRPCClient{cc}
}

func (c *userRPCClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/user.UserRPC/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) UpdateProfile(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/user.UserRPC/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/user.UserRPC/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) GetFavorites(ctx context.Context, in *ID, opts ...grpc.CallOption) (*GetFavoritesResponse, error) {
	out := new(GetFavoritesResponse)
	err := c.cc.Invoke(ctx, "/user.UserRPC/GetFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) SetFavorite(ctx context.Context, in *HandleFavorite, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/user.UserRPC/SetFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) ResetFavorite(ctx context.Context, in *HandleFavorite, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/user.UserRPC/ResetFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) CheckFavorite(ctx context.Context, in *HandleFavorite, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/user.UserRPC/CheckFavorite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) FindByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/user.UserRPC/FindByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) FindByEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/user.UserRPC/FindByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) Subscribe(ctx context.Context, in *CreateSubscriptionRequest, opts ...grpc.CallOption) (*SubscriptionID, error) {
	out := new(SubscriptionID)
	err := c.cc.Invoke(ctx, "/user.UserRPC/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRPCClient) UpdateSubscribtionStatus(ctx context.Context, in *SubscriptionID, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/user.UserRPC/UpdateSubscribtionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRPCServer is the server API for UserRPC service.
// All implementations must embed UnimplementedUserRPCServer
// for forward compatibility
type UserRPCServer interface {
	Create(context.Context, *CreateUserRequest) (*ID, error)
	UpdateProfile(context.Context, *UserData) (*Nothing, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*Nothing, error)
	GetFavorites(context.Context, *ID) (*GetFavoritesResponse, error)
	SetFavorite(context.Context, *HandleFavorite) (*Nothing, error)
	ResetFavorite(context.Context, *HandleFavorite) (*Nothing, error)
	CheckFavorite(context.Context, *HandleFavorite) (*Nothing, error)
	FindByID(context.Context, *ID) (*UserData, error)
	FindByEmail(context.Context, *Email) (*UserData, error)
	Subscribe(context.Context, *CreateSubscriptionRequest) (*SubscriptionID, error)
	UpdateSubscribtionStatus(context.Context, *SubscriptionID) (*Nothing, error)
	mustEmbedUnimplementedUserRPCServer()
}

// UnimplementedUserRPCServer must be embedded to have forward compatible implementations.
type UnimplementedUserRPCServer struct {
}

func (UnimplementedUserRPCServer) Create(context.Context, *CreateUserRequest) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserRPCServer) UpdateProfile(context.Context, *UserData) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserRPCServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedUserRPCServer) GetFavorites(context.Context, *ID) (*GetFavoritesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorites not implemented")
}
func (UnimplementedUserRPCServer) SetFavorite(context.Context, *HandleFavorite) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFavorite not implemented")
}
func (UnimplementedUserRPCServer) ResetFavorite(context.Context, *HandleFavorite) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetFavorite not implemented")
}
func (UnimplementedUserRPCServer) CheckFavorite(context.Context, *HandleFavorite) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckFavorite not implemented")
}
func (UnimplementedUserRPCServer) FindByID(context.Context, *ID) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedUserRPCServer) FindByEmail(context.Context, *Email) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEmail not implemented")
}
func (UnimplementedUserRPCServer) Subscribe(context.Context, *CreateSubscriptionRequest) (*SubscriptionID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedUserRPCServer) UpdateSubscribtionStatus(context.Context, *SubscriptionID) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSubscribtionStatus not implemented")
}
func (UnimplementedUserRPCServer) mustEmbedUnimplementedUserRPCServer() {}

// UnsafeUserRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserRPCServer will
// result in compilation errors.
type UnsafeUserRPCServer interface {
	mustEmbedUnimplementedUserRPCServer()
}

func RegisterUserRPCServer(s grpc.ServiceRegistrar, srv UserRPCServer) {
	s.RegisterService(&UserRPC_ServiceDesc, srv)
}

func _UserRPC_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).UpdateProfile(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_GetFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).GetFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/GetFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).GetFavorites(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_SetFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleFavorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).SetFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/SetFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).SetFavorite(ctx, req.(*HandleFavorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_ResetFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleFavorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).ResetFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/ResetFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).ResetFavorite(ctx, req.(*HandleFavorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_CheckFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleFavorite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).CheckFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/CheckFavorite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).CheckFavorite(ctx, req.(*HandleFavorite))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/FindByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).FindByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_FindByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).FindByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/FindByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).FindByEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).Subscribe(ctx, req.(*CreateSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRPC_UpdateSubscribtionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscriptionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRPCServer).UpdateSubscribtionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserRPC/UpdateSubscribtionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRPCServer).UpdateSubscribtionStatus(ctx, req.(*SubscriptionID))
	}
	return interceptor(ctx, in, info, handler)
}

// UserRPC_ServiceDesc is the grpc.ServiceDesc for UserRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserRPC",
	HandlerType: (*UserRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserRPC_Create_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserRPC_UpdateProfile_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserRPC_UpdatePassword_Handler,
		},
		{
			MethodName: "GetFavorites",
			Handler:    _UserRPC_GetFavorites_Handler,
		},
		{
			MethodName: "SetFavorite",
			Handler:    _UserRPC_SetFavorite_Handler,
		},
		{
			MethodName: "ResetFavorite",
			Handler:    _UserRPC_ResetFavorite_Handler,
		},
		{
			MethodName: "CheckFavorite",
			Handler:    _UserRPC_CheckFavorite_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _UserRPC_FindByID_Handler,
		},
		{
			MethodName: "FindByEmail",
			Handler:    _UserRPC_FindByEmail_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _UserRPC_Subscribe_Handler,
		},
		{
			MethodName: "UpdateSubscribtionStatus",
			Handler:    _UserRPC_UpdateSubscribtionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
