package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	OrderService_CreateOffer_FullMethodName            = "/order.OrderService/CreateOffer"
	OrderService_GetOffer_FullMethodName                = "/order.OrderService/GetOffer"
	OrderService_ListOffers_FullMethodName              = "/order.OrderService/ListOffers"
	OrderService_AcceptOffer_FullMethodName             = "/order.OrderService/AcceptOffer"
	OrderService_RejectOffer_FullMethodName             = "/order.OrderService/RejectOffer"
	OrderService_CancelOffer_FullMethodName             = "/order.OrderService/CancelOffer"
	OrderService_GetTransaction_FullMethodName          = "/order.OrderService/GetTransaction"
	OrderService_ListTransactions_FullMethodName        = "/order.OrderService/ListTransactions"
	OrderService_UpdateTransactionStatus_FullMethodName = "/order.OrderService/UpdateTransactionStatus"
)

type OrderServiceServer interface {
	CreateOffer(context.Context, *CreateOfferRequest) (*Offer, error)
	GetOffer(context.Context, *GetOfferRequest) (*Offer, error)
	ListOffers(context.Context, *ListOffersRequest) (*ListOffersResponse, error)
	AcceptOffer(context.Context, *AcceptOfferRequest) (*Transaction, error)
	RejectOffer(context.Context, *RejectOfferRequest) (*Offer, error)
	CancelOffer(context.Context, *CancelOfferRequest) (*Offer, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error)
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	UpdateTransactionStatus(context.Context, *UpdateStatusRequest) (*Transaction, error)
	mustEmbedUnimplementedOrderServiceServer()
}

type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) CreateOffer(context.Context, *CreateOfferRequest) (*Offer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOffer not implemented")
}
func (UnimplementedOrderServiceServer) GetOffer(context.Context, *GetOfferRequest) (*Offer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffer not implemented")
}
func (UnimplementedOrderServiceServer) ListOffers(context.Context, *ListOffersRequest) (*ListOffersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOffers not implemented")
}
func (UnimplementedOrderServiceServer) AcceptOffer(context.Context, *AcceptOfferRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptOffer not implemented")
}
func (UnimplementedOrderServiceServer) RejectOffer(context.Context, *RejectOfferRequest) (*Offer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectOffer not implemented")
}
func (UnimplementedOrderServiceServer) CancelOffer(context.Context, *CancelOfferRequest) (*Offer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOffer not implemented")
}
func (UnimplementedOrderServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedOrderServiceServer) ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedOrderServiceServer) UpdateTransactionStatus(context.Context, *UpdateStatusRequest) (*Transaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTransactionStatus not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_CreateOffer_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOffer(ctx, req.(*CreateOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_GetOffer_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOffer(ctx, req.(*GetOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOffers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOffersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOffers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_ListOffers_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOffers(ctx, req.(*ListOffersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_AcceptOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AcceptOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_AcceptOffer_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AcceptOffer(ctx, req.(*AcceptOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_RejectOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).RejectOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_RejectOffer_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).RejectOffer(ctx, req.(*RejectOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CancelOffer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOfferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CancelOffer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_CancelOffer_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CancelOffer(ctx, req.(*CancelOfferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_GetTransaction_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_ListTransactions_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateTransactionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateTransactionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: OrderService_UpdateTransactionStatus_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateTransactionStatus(ctx, req.(*UpdateStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateOffer", Handler: _OrderService_CreateOffer_Handler},
		{MethodName: "GetOffer", Handler: _OrderService_GetOffer_Handler},
		{MethodName: "ListOffers", Handler: _OrderService_ListOffers_Handler},
		{MethodName: "AcceptOffer", Handler: _OrderService_AcceptOffer_Handler},
		{MethodName: "RejectOffer", Handler: _OrderService_RejectOffer_Handler},
		{MethodName: "CancelOffer", Handler: _OrderService_CancelOffer_Handler},
		{MethodName: "GetTransaction", Handler: _OrderService_GetTransaction_Handler},
		{MethodName: "ListTransactions", Handler: _OrderService_ListTransactions_Handler},
		{MethodName: "UpdateTransactionStatus", Handler: _OrderService_UpdateTransactionStatus_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
