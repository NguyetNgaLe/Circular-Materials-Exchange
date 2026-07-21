package pb

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion7

type MaterialServiceClient interface {
	ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*Category, error)
	CreateListing(ctx context.Context, in *CreateListingRequest, opts ...grpc.CallOption) (*SupplyListing, error)
	GetListing(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*SupplyListing, error)
	ListListings(ctx context.Context, in *ListListingsRequest, opts ...grpc.CallOption) (*ListListingsResponse, error)
	UpdateListing(ctx context.Context, in *UpdateListingRequest, opts ...grpc.CallOption) (*SupplyListing, error)
	DeleteListing(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateDemand(ctx context.Context, in *CreateDemandRequest, opts ...grpc.CallOption) (*DemandListing, error)
	GetDemand(ctx context.Context, in *GetDemandRequest, opts ...grpc.CallOption) (*DemandListing, error)
	ListDemands(ctx context.Context, in *ListDemandsRequest, opts ...grpc.CallOption) (*ListDemandsResponse, error)
}

type materialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaterialServiceClient(cc grpc.ClientConnInterface) MaterialServiceClient {
	return &materialServiceClient{cc}
}

func (c *materialServiceClient) ListCategories(ctx context.Context, in *ListCategoriesRequest, opts ...grpc.CallOption) (*ListCategoriesResponse, error) {
	out := new(ListCategoriesResponse)
	err := c.cc.Invoke(ctx, "/material.MaterialService/ListCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/material.MaterialService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) CreateListing(ctx context.Context, in *CreateListingRequest, opts ...grpc.CallOption) (*SupplyListing, error) {
	out := new(SupplyListing)
	err := c.cc.Invoke(ctx, "/material.MaterialService/CreateListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) GetListing(ctx context.Context, in *GetListingRequest, opts ...grpc.CallOption) (*SupplyListing, error) {
	out := new(SupplyListing)
	err := c.cc.Invoke(ctx, "/material.MaterialService/GetListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) ListListings(ctx context.Context, in *ListListingsRequest, opts ...grpc.CallOption) (*ListListingsResponse, error) {
	out := new(ListListingsResponse)
	err := c.cc.Invoke(ctx, "/material.MaterialService/ListListings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) UpdateListing(ctx context.Context, in *UpdateListingRequest, opts ...grpc.CallOption) (*SupplyListing, error) {
	out := new(SupplyListing)
	err := c.cc.Invoke(ctx, "/material.MaterialService/UpdateListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) DeleteListing(ctx context.Context, in *DeleteListingRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/material.MaterialService/DeleteListing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) CreateDemand(ctx context.Context, in *CreateDemandRequest, opts ...grpc.CallOption) (*DemandListing, error) {
	out := new(DemandListing)
	err := c.cc.Invoke(ctx, "/material.MaterialService/CreateDemand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) GetDemand(ctx context.Context, in *GetDemandRequest, opts ...grpc.CallOption) (*DemandListing, error) {
	out := new(DemandListing)
	err := c.cc.Invoke(ctx, "/material.MaterialService/GetDemand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *materialServiceClient) ListDemands(ctx context.Context, in *ListDemandsRequest, opts ...grpc.CallOption) (*ListDemandsResponse, error) {
	out := new(ListDemandsResponse)
	err := c.cc.Invoke(ctx, "/material.MaterialService/ListDemands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type MaterialServiceServer interface {
	ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error)
	CreateListing(context.Context, *CreateListingRequest) (*SupplyListing, error)
	GetListing(context.Context, *GetListingRequest) (*SupplyListing, error)
	ListListings(context.Context, *ListListingsRequest) (*ListListingsResponse, error)
	UpdateListing(context.Context, *UpdateListingRequest) (*SupplyListing, error)
	DeleteListing(context.Context, *DeleteListingRequest) (*Empty, error)
	CreateDemand(context.Context, *CreateDemandRequest) (*DemandListing, error)
	GetDemand(context.Context, *GetDemandRequest) (*DemandListing, error)
	ListDemands(context.Context, *ListDemandsRequest) (*ListDemandsResponse, error)
	mustEmbedUnimplementedMaterialServiceServer()
}

type UnimplementedMaterialServiceServer struct{}

func (UnimplementedMaterialServiceServer) ListCategories(context.Context, *ListCategoriesRequest) (*ListCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCategories not implemented")
}
func (UnimplementedMaterialServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedMaterialServiceServer) CreateListing(context.Context, *CreateListingRequest) (*SupplyListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateListing not implemented")
}
func (UnimplementedMaterialServiceServer) GetListing(context.Context, *GetListingRequest) (*SupplyListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListing not implemented")
}
func (UnimplementedMaterialServiceServer) ListListings(context.Context, *ListListingsRequest) (*ListListingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListListings not implemented")
}
func (UnimplementedMaterialServiceServer) UpdateListing(context.Context, *UpdateListingRequest) (*SupplyListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateListing not implemented")
}
func (UnimplementedMaterialServiceServer) DeleteListing(context.Context, *DeleteListingRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteListing not implemented")
}
func (UnimplementedMaterialServiceServer) CreateDemand(context.Context, *CreateDemandRequest) (*DemandListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDemand not implemented")
}
func (UnimplementedMaterialServiceServer) GetDemand(context.Context, *GetDemandRequest) (*DemandListing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDemand not implemented")
}
func (UnimplementedMaterialServiceServer) ListDemands(context.Context, *ListDemandsRequest) (*ListDemandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDemands not implemented")
}
func (UnimplementedMaterialServiceServer) mustEmbedUnimplementedMaterialServiceServer() {}

func RegisterMaterialServiceServer(s *grpc.Server, srv MaterialServiceServer) {
	s.RegisterService(&_MaterialService_serviceDesc, srv)
}

func _MaterialService_ListCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).ListCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/ListCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).ListCategories(ctx, req.(*ListCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_CreateListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).CreateListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/CreateListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).CreateListing(ctx, req.(*CreateListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_GetListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).GetListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/GetListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).GetListing(ctx, req.(*GetListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_ListListings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListListingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).ListListings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/ListListings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).ListListings(ctx, req.(*ListListingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_UpdateListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).UpdateListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/UpdateListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).UpdateListing(ctx, req.(*UpdateListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_DeleteListing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteListingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).DeleteListing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/DeleteListing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).DeleteListing(ctx, req.(*DeleteListingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_CreateDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDemandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).CreateDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/CreateDemand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).CreateDemand(ctx, req.(*CreateDemandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_GetDemand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDemandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).GetDemand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/GetDemand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).GetDemand(ctx, req.(*GetDemandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaterialService_ListDemands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDemandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaterialServiceServer).ListDemands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/material.MaterialService/ListDemands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaterialServiceServer).ListDemands(ctx, req.(*ListDemandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MaterialService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "material.MaterialService",
	HandlerType: (*MaterialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCategories",
			Handler:    _MaterialService_ListCategories_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _MaterialService_CreateCategory_Handler,
		},
		{
			MethodName: "CreateListing",
			Handler:    _MaterialService_CreateListing_Handler,
		},
		{
			MethodName: "GetListing",
			Handler:    _MaterialService_GetListing_Handler,
		},
		{
			MethodName: "ListListings",
			Handler:    _MaterialService_ListListings_Handler,
		},
		{
			MethodName: "UpdateListing",
			Handler:    _MaterialService_UpdateListing_Handler,
		},
		{
			MethodName: "DeleteListing",
			Handler:    _MaterialService_DeleteListing_Handler,
		},
		{
			MethodName: "CreateDemand",
			Handler:    _MaterialService_CreateDemand_Handler,
		},
		{
			MethodName: "GetDemand",
			Handler:    _MaterialService_GetDemand_Handler,
		},
		{
			MethodName: "ListDemands",
			Handler:    _MaterialService_ListDemands_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "material.proto",
}
