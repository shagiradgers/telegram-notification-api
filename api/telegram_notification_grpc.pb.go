// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: api/telegram_notification.proto

package notification

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

// TelegramNotificationServiceClient is the client API for TelegramNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TelegramNotificationServiceClient interface {
	SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error)
	GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error)
	GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUsersById(ctx context.Context, in *GetUsersByIdRequest, opts ...grpc.CallOption) (*GetUsersByIdResponse, error)
	GetUsersByFilter(ctx context.Context, in *GetUsersByFilterRequest, opts ...grpc.CallOption) (*GetUsersByFilterResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type telegramNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTelegramNotificationServiceClient(cc grpc.ClientConnInterface) TelegramNotificationServiceClient {
	return &telegramNotificationServiceClient{cc}
}

func (c *telegramNotificationServiceClient) SendNotification(ctx context.Context, in *SendNotificationRequest, opts ...grpc.CallOption) (*SendNotificationResponse, error) {
	out := new(SendNotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/SendNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) GetNotification(ctx context.Context, in *GetNotificationRequest, opts ...grpc.CallOption) (*GetNotificationResponse, error) {
	out := new(GetNotificationResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/GetNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) GetNotifications(ctx context.Context, in *GetNotificationsRequest, opts ...grpc.CallOption) (*GetNotificationsResponse, error) {
	out := new(GetNotificationsResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/GetNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) GetUsersById(ctx context.Context, in *GetUsersByIdRequest, opts ...grpc.CallOption) (*GetUsersByIdResponse, error) {
	out := new(GetUsersByIdResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/GetUsersById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) GetUsersByFilter(ctx context.Context, in *GetUsersByFilterRequest, opts ...grpc.CallOption) (*GetUsersByFilterResponse, error) {
	out := new(GetUsersByFilterResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/GetUsersByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) EditUser(ctx context.Context, in *EditUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error) {
	out := new(EditUserResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/EditUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *telegramNotificationServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/notification.v1.telegram_notification_service/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TelegramNotificationServiceServer is the server API for TelegramNotificationService service.
// All implementations must embed UnimplementedTelegramNotificationServiceServer
// for forward compatibility
type TelegramNotificationServiceServer interface {
	SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error)
	GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error)
	GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUsersById(context.Context, *GetUsersByIdRequest) (*GetUsersByIdResponse, error)
	GetUsersByFilter(context.Context, *GetUsersByFilterRequest) (*GetUsersByFilterResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	EditUser(context.Context, *EditUserRequest) (*EditUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	mustEmbedUnimplementedTelegramNotificationServiceServer()
}

// UnimplementedTelegramNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTelegramNotificationServiceServer struct {
}

func (UnimplementedTelegramNotificationServiceServer) SendNotification(context.Context, *SendNotificationRequest) (*SendNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) GetNotification(context.Context, *GetNotificationRequest) (*GetNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) GetNotifications(context.Context, *GetNotificationsRequest) (*GetNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotifications not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) GetUsersById(context.Context, *GetUsersByIdRequest) (*GetUsersByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersById not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) GetUsersByFilter(context.Context, *GetUsersByFilterRequest) (*GetUsersByFilterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByFilter not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) EditUser(context.Context, *EditUserRequest) (*EditUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedTelegramNotificationServiceServer) mustEmbedUnimplementedTelegramNotificationServiceServer() {
}

// UnsafeTelegramNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TelegramNotificationServiceServer will
// result in compilation errors.
type UnsafeTelegramNotificationServiceServer interface {
	mustEmbedUnimplementedTelegramNotificationServiceServer()
}

func RegisterTelegramNotificationServiceServer(s grpc.ServiceRegistrar, srv TelegramNotificationServiceServer) {
	s.RegisterService(&TelegramNotificationService_ServiceDesc, srv)
}

func _TelegramNotificationService_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/SendNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).SendNotification(ctx, req.(*SendNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/GetNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).GetNotification(ctx, req.(*GetNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_GetNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).GetNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/GetNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).GetNotifications(ctx, req.(*GetNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_GetUsersById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).GetUsersById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/GetUsersById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).GetUsersById(ctx, req.(*GetUsersByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_GetUsersByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersByFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).GetUsersByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/GetUsersByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).GetUsersByFilter(ctx, req.(*GetUsersByFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).EditUser(ctx, req.(*EditUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TelegramNotificationService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TelegramNotificationServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.v1.telegram_notification_service/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TelegramNotificationServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TelegramNotificationService_ServiceDesc is the grpc.ServiceDesc for TelegramNotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TelegramNotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.v1.telegram_notification_service",
	HandlerType: (*TelegramNotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendNotification",
			Handler:    _TelegramNotificationService_SendNotification_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _TelegramNotificationService_GetNotification_Handler,
		},
		{
			MethodName: "GetNotifications",
			Handler:    _TelegramNotificationService_GetNotifications_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _TelegramNotificationService_GetUser_Handler,
		},
		{
			MethodName: "GetUsersById",
			Handler:    _TelegramNotificationService_GetUsersById_Handler,
		},
		{
			MethodName: "GetUsersByFilter",
			Handler:    _TelegramNotificationService_GetUsersByFilter_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _TelegramNotificationService_DeleteUser_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _TelegramNotificationService_EditUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _TelegramNotificationService_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/telegram_notification.proto",
}
