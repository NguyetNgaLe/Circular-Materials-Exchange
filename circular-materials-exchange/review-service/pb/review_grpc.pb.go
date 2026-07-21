package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReviewServiceServer interface {
	CreateReview(context.Context, *CreateReviewRequest) (*Review, error)
	GetReview(context.Context, *GetReviewRequest) (*Review, error)
	ListReviews(context.Context, *ListReviewsRequest) (*ListReviewsResponse, error)
	GetUserRating(context.Context, *GetUserRatingRequest) (*UserRatingResponse, error)
	mustEmbedUnimplementedReviewServiceServer()
}

type UnimplementedReviewServiceServer struct{}

func (*UnimplementedReviewServiceServer) CreateReview(context.Context, *CreateReviewRequest) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (*UnimplementedReviewServiceServer) GetReview(context.Context, *GetReviewRequest) (*Review, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReview not implemented")
}
func (*UnimplementedReviewServiceServer) ListReviews(context.Context, *ListReviewsRequest) (*ListReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListReviews not implemented")
}
func (*UnimplementedReviewServiceServer) GetUserRating(context.Context, *GetUserRatingRequest) (*UserRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRating not implemented")
}
func (*UnimplementedReviewServiceServer) mustEmbedUnimplementedReviewServiceServer() {}

func RegisterReviewServiceServer(s *grpc.Server, srv ReviewServiceServer) {
	s.RegisterService(&_ReviewService_serviceDesc, srv)
}

func _ReviewService_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/review.ReviewService/CreateReview"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/review.ReviewService/GetReview"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetReview(ctx, req.(*GetReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_ListReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).ListReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/review.ReviewService/ListReviews"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).ListReviews(ctx, req.(*ListReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReviewService_GetUserRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReviewServiceServer).GetUserRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: "/review.ReviewService/GetUserRating"}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReviewServiceServer).GetUserRating(ctx, req.(*GetUserRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReviewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "review.ReviewService",
	HandlerType: (*ReviewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "CreateReview", Handler: _ReviewService_CreateReview_Handler},
		{MethodName: "GetReview", Handler: _ReviewService_GetReview_Handler},
		{MethodName: "ListReviews", Handler: _ReviewService_ListReviews_Handler},
		{MethodName: "GetUserRating", Handler: _ReviewService_GetUserRating_Handler},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "review.proto",
}
