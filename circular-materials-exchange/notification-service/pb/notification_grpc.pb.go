package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationServiceServer interface {
	CreateNotification(context.Context, *CreateNotificationRequest) (*Notification, error)
	ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error)
	MarkRead(context.Context, *MarkReadRequest) (*Notification, error)
	MarkAllRead(context.Context, *MarkAllReadRequest) (*Empty, error)
	GetUnreadCount(context.Context, *GetUnreadCountRequest) (*UnreadCountResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

type UnimplementedNotificationServiceServer struct{}

func (*UnimplementedNotificationServiceServer) CreateNotification(context.Context, *CreateNotificationRequest) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotification not implemented")
}
func (*UnimplementedNotificationServiceServer) ListNotifications(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotifications not implemented")
}
func (*UnimplementedNotificationServiceServer) MarkRead(context.Context, *MarkReadRequest) (*Notification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkRead not implemented")
}
func (*UnimplementedNotificationServiceServer) MarkAllRead(context.Context, *MarkAllReadRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAllRead not implemented")
}
func (*UnimplementedNotificationServiceServer) GetUnreadCount(context.Context, *GetUnreadCountRequest) (*UnreadCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnreadCount not implemented")
}
func (*UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_CreateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).CreateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/notification.NotificationService/CreateNotification"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).CreateNotification(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_ListNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).ListNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/notification.NotificationService/ListNotifications"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).ListNotifications(ctx, req.(*ListNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/notification.NotificationService/MarkRead"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkRead(ctx, req.(*MarkReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_MarkAllRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAllReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).MarkAllRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/notification.NotificationService/MarkAllRead"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).MarkAllRead(ctx, req.(*MarkAllReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_GetUnreadCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnreadCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).GetUnreadCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/notification.NotificationService/GetUnreadCount"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).GetUnreadCount(ctx, req.(*GetUnreadCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateNotification", Handler: _NotificationService_CreateNotification_Handler},
		{MethodName: "ListNotifications", Handler: _NotificationService_ListNotifications_Handler},
		{MethodName: "MarkRead", Handler: _NotificationService_MarkRead_Handler},
		{MethodName: "MarkAllRead", Handler: _NotificationService_MarkAllRead_Handler},
		{MethodName: "GetUnreadCount", Handler: _NotificationService_GetUnreadCount_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification.proto",
}
