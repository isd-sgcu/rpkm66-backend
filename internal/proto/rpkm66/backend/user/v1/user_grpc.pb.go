// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: rpkm66/backend/user/v1/user.proto

package v1

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
	UserService_FindOne_FullMethodName               = "/rpkm66.backend.user.v1.UserService/FindOne"
	UserService_FindByStudentID_FullMethodName       = "/rpkm66.backend.user.v1.UserService/FindByStudentID"
	UserService_Create_FullMethodName                = "/rpkm66.backend.user.v1.UserService/Create"
	UserService_Update_FullMethodName                = "/rpkm66.backend.user.v1.UserService/Update"
	UserService_Verify_FullMethodName                = "/rpkm66.backend.user.v1.UserService/Verify"
	UserService_Delete_FullMethodName                = "/rpkm66.backend.user.v1.UserService/Delete"
	UserService_CreateOrUpdate_FullMethodName        = "/rpkm66.backend.user.v1.UserService/CreateOrUpdate"
	UserService_ConfirmEstamp_FullMethodName         = "/rpkm66.backend.user.v1.UserService/ConfirmEstamp"
	UserService_GetUserEstamp_FullMethodName         = "/rpkm66.backend.user.v1.UserService/GetUserEstamp"
	UserService_UpdatePersonalityGame_FullMethodName = "/rpkm66.backend.user.v1.UserService/UpdatePersonalityGame"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	FindOne(ctx context.Context, in *FindOneUserRequest, opts ...grpc.CallOption) (*FindOneUserResponse, error)
	FindByStudentID(ctx context.Context, in *FindByStudentIDUserRequest, opts ...grpc.CallOption) (*FindByStudentIDUserResponse, error)
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	Verify(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyUserResponse, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	CreateOrUpdate(ctx context.Context, in *CreateOrUpdateUserRequest, opts ...grpc.CallOption) (*CreateOrUpdateUserResponse, error)
	ConfirmEstamp(ctx context.Context, in *ConfirmEstampRequest, opts ...grpc.CallOption) (*ConfirmEstampResponse, error)
	GetUserEstamp(ctx context.Context, in *GetUserEstampRequest, opts ...grpc.CallOption) (*GetUserEstampResponse, error)
	UpdatePersonalityGame(ctx context.Context, in *UpdatePersonalityGameRequest, opts ...grpc.CallOption) (*UpdatePersonalityGameResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) FindOne(ctx context.Context, in *FindOneUserRequest, opts ...grpc.CallOption) (*FindOneUserResponse, error) {
	out := new(FindOneUserResponse)
	err := c.cc.Invoke(ctx, UserService_FindOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindByStudentID(ctx context.Context, in *FindByStudentIDUserRequest, opts ...grpc.CallOption) (*FindByStudentIDUserResponse, error) {
	out := new(FindByStudentIDUserResponse)
	err := c.cc.Invoke(ctx, UserService_FindByStudentID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Verify(ctx context.Context, in *VerifyUserRequest, opts ...grpc.CallOption) (*VerifyUserResponse, error) {
	out := new(VerifyUserResponse)
	err := c.cc.Invoke(ctx, UserService_Verify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateOrUpdate(ctx context.Context, in *CreateOrUpdateUserRequest, opts ...grpc.CallOption) (*CreateOrUpdateUserResponse, error) {
	out := new(CreateOrUpdateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateOrUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ConfirmEstamp(ctx context.Context, in *ConfirmEstampRequest, opts ...grpc.CallOption) (*ConfirmEstampResponse, error) {
	out := new(ConfirmEstampResponse)
	err := c.cc.Invoke(ctx, UserService_ConfirmEstamp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserEstamp(ctx context.Context, in *GetUserEstampRequest, opts ...grpc.CallOption) (*GetUserEstampResponse, error) {
	out := new(GetUserEstampResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserEstamp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePersonalityGame(ctx context.Context, in *UpdatePersonalityGameRequest, opts ...grpc.CallOption) (*UpdatePersonalityGameResponse, error) {
	out := new(UpdatePersonalityGameResponse)
	err := c.cc.Invoke(ctx, UserService_UpdatePersonalityGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	FindOne(context.Context, *FindOneUserRequest) (*FindOneUserResponse, error)
	FindByStudentID(context.Context, *FindByStudentIDUserRequest) (*FindByStudentIDUserResponse, error)
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	Verify(context.Context, *VerifyUserRequest) (*VerifyUserResponse, error)
	Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	CreateOrUpdate(context.Context, *CreateOrUpdateUserRequest) (*CreateOrUpdateUserResponse, error)
	ConfirmEstamp(context.Context, *ConfirmEstampRequest) (*ConfirmEstampResponse, error)
	GetUserEstamp(context.Context, *GetUserEstampRequest) (*GetUserEstampResponse, error)
	UpdatePersonalityGame(context.Context, *UpdatePersonalityGameRequest) (*UpdatePersonalityGameResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) FindOne(context.Context, *FindOneUserRequest) (*FindOneUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedUserServiceServer) FindByStudentID(context.Context, *FindByStudentIDUserRequest) (*FindByStudentIDUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByStudentID not implemented")
}
func (UnimplementedUserServiceServer) Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) Verify(context.Context, *VerifyUserRequest) (*VerifyUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceServer) CreateOrUpdate(context.Context, *CreateOrUpdateUserRequest) (*CreateOrUpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (UnimplementedUserServiceServer) ConfirmEstamp(context.Context, *ConfirmEstampRequest) (*ConfirmEstampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmEstamp not implemented")
}
func (UnimplementedUserServiceServer) GetUserEstamp(context.Context, *GetUserEstampRequest) (*GetUserEstampResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserEstamp not implemented")
}
func (UnimplementedUserServiceServer) UpdatePersonalityGame(context.Context, *UpdatePersonalityGameRequest) (*UpdatePersonalityGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonalityGame not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindOne(ctx, req.(*FindOneUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindByStudentID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByStudentIDUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindByStudentID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_FindByStudentID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindByStudentID(ctx, req.(*FindByStudentIDUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Verify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Verify(ctx, req.(*VerifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateOrUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateOrUpdate(ctx, req.(*CreateOrUpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ConfirmEstamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmEstampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ConfirmEstamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ConfirmEstamp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ConfirmEstamp(ctx, req.(*ConfirmEstampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserEstamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserEstampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserEstamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserEstamp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserEstamp(ctx, req.(*GetUserEstampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePersonalityGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonalityGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePersonalityGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdatePersonalityGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePersonalityGame(ctx, req.(*UpdatePersonalityGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpkm66.backend.user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOne",
			Handler:    _UserService_FindOne_Handler,
		},
		{
			MethodName: "FindByStudentID",
			Handler:    _UserService_FindByStudentID_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _UserService_Verify_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
		{
			MethodName: "CreateOrUpdate",
			Handler:    _UserService_CreateOrUpdate_Handler,
		},
		{
			MethodName: "ConfirmEstamp",
			Handler:    _UserService_ConfirmEstamp_Handler,
		},
		{
			MethodName: "GetUserEstamp",
			Handler:    _UserService_GetUserEstamp_Handler,
		},
		{
			MethodName: "UpdatePersonalityGame",
			Handler:    _UserService_UpdatePersonalityGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpkm66/backend/user/v1/user.proto",
}
