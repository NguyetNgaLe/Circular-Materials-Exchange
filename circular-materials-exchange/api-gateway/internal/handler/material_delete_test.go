package handler

import (
	materialpb "api-gateway/internal/pb/material"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type deleteListingMaterialClient struct {
	materialpb.MaterialServiceClient
	listing   *materialpb.SupplyListing
	deletedID string
}

func (f *deleteListingMaterialClient) GetListing(
	_ context.Context,
	_ *materialpb.GetListingRequest,
	_ ...grpc.CallOption,
) (*materialpb.SupplyListing, error) {
	return f.listing, nil
}

func (f *deleteListingMaterialClient) DeleteListing(
	_ context.Context,
	request *materialpb.DeleteListingRequest,
	_ ...grpc.CallOption,
) (*materialpb.Empty, error) {
	f.deletedID = request.GetId()
	return &materialpb.Empty{}, nil
}

func TestDeleteListingChecksOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name          string
		userID        string
		role          string
		wantStatus    int
		wantDeletedID string
	}{
		{
			name:          "owner can delete",
			userID:        "seller-1",
			role:          "business",
			wantStatus:    http.StatusOK,
			wantDeletedID: "listing-1",
		},
		{
			name:       "another business is forbidden",
			userID:     "seller-2",
			role:       "business",
			wantStatus: http.StatusForbidden,
		},
		{
			name:          "admin can delete",
			userID:        "admin-1",
			role:          "admin",
			wantStatus:    http.StatusOK,
			wantDeletedID: "listing-1",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			client := &deleteListingMaterialClient{
				listing: &materialpb.SupplyListing{Id: "listing-1", SellerId: "seller-1"},
			}
			handler := NewMaterialHandler(client, nil)
			recorder := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(recorder)
			ctx.Request = httptest.NewRequest(http.MethodDelete, "/api/listings/listing-1", nil)
			ctx.Params = gin.Params{{Key: "id", Value: "listing-1"}}
			ctx.Set("user_id", test.userID)
			ctx.Set("role", test.role)

			handler.DeleteListing(ctx)

			if recorder.Code != test.wantStatus {
				t.Fatalf("status = %d, want %d; body=%s", recorder.Code, test.wantStatus, recorder.Body.String())
			}
			if client.deletedID != test.wantDeletedID {
				t.Fatalf("deleted ID = %q, want %q", client.deletedID, test.wantDeletedID)
			}
		})
	}
}
