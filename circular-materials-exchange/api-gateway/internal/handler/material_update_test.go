package handler

import (
	materialpb "api-gateway/internal/pb/material"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type updateListingMaterialClient struct {
	materialpb.MaterialServiceClient
	listing      *materialpb.SupplyListing
	update       *materialpb.UpdateListingRequest
	listResponse *materialpb.ListListingsResponse
}

func (f *updateListingMaterialClient) GetListing(
	_ context.Context,
	_ *materialpb.GetListingRequest,
	_ ...grpc.CallOption,
) (*materialpb.SupplyListing, error) {
	return f.listing, nil
}

func (f *updateListingMaterialClient) UpdateListing(
	_ context.Context,
	request *materialpb.UpdateListingRequest,
	_ ...grpc.CallOption,
) (*materialpb.SupplyListing, error) {
	f.update = request
	return &materialpb.SupplyListing{
		Id:               request.GetId(),
		Title:            request.GetTitle(),
		CategoryId:       request.GetCategoryId(),
		SellerId:         f.listing.GetSellerId(),
		Description:      request.GetDescription(),
		Specs:            request.GetSpecs(),
		Quantity:         request.GetQuantity(),
		Unit:             request.GetUnit(),
		PricePerUnit:     request.GetPricePerUnit(),
		Currency:         request.GetCurrency(),
		Location:         request.GetLocation(),
		MinOrderQuantity: request.GetMinOrderQuantity(),
		Packaging:        request.GetPackaging(),
		Status:           request.GetStatus(),
		Images:           request.GetImages(),
	}, nil
}

func (f *updateListingMaterialClient) ListListings(
	_ context.Context,
	_ *materialpb.ListListingsRequest,
	_ ...grpc.CallOption,
) (*materialpb.ListListingsResponse, error) {
	return f.listResponse, nil
}

func existingSupplyListing() *materialpb.SupplyListing {
	return &materialpb.SupplyListing{
		Id: "listing-1", Title: "PP cu", CategoryId: "plastic",
		SellerId: "seller-1", CompanyId: "company-1",
		Description: "Mo ta cu", Specs: map[string]string{"grade": "A"},
		Quantity: 20, Unit: "Tan", PricePerUnit: 1000, Currency: "VND",
		Location: "HCM", MinOrderQuantity: 2, Packaging: "Bao",
		Status: "active", Images: []string{"/images/old.jpg"},
	}
}

func updateContext(body, userID, role string) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/api/listings/listing-1", strings.NewReader(body))
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Params = gin.Params{{Key: "id", Value: "listing-1"}}
	ctx.Set("user_id", userID)
	ctx.Set("role", role)
	return ctx, recorder
}

func TestUpdateListingOwnerCanHideWithoutLosingFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	client := &updateListingMaterialClient{listing: existingSupplyListing()}
	handler := NewMaterialHandler(client, nil)
	ctx, recorder := updateContext(`{"status":"hidden"}`, "seller-1", "business")

	handler.UpdateListing(ctx)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	if client.update == nil {
		t.Fatal("material UpdateListing was not called")
	}
	if client.update.GetStatus() != "hidden" {
		t.Fatalf("status = %q, want hidden", client.update.GetStatus())
	}
	if client.update.GetTitle() != "PP cu" || client.update.GetCategoryId() != "plastic" {
		t.Fatalf("partial update lost current fields: %#v", client.update)
	}
	if got := client.update.GetImages(); len(got) != 1 || got[0] != "/images/old.jpg" {
		t.Fatalf("images = %#v, want current image", got)
	}
}

func TestUpdateListingChecksOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)
	client := &updateListingMaterialClient{listing: existingSupplyListing()}
	handler := NewMaterialHandler(client, nil)
	ctx, recorder := updateContext(`{"title":"Khong duoc sua"}`, "seller-2", "business")

	handler.UpdateListing(ctx)

	if recorder.Code != http.StatusForbidden {
		t.Fatalf("status = %d, want %d; body=%s", recorder.Code, http.StatusForbidden, recorder.Body.String())
	}
	if client.update != nil {
		t.Fatal("material UpdateListing must not be called for another business")
	}
}

func TestUpdateListingValidatesStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	client := &updateListingMaterialClient{listing: existingSupplyListing()}
	handler := NewMaterialHandler(client, nil)
	ctx, recorder := updateContext(`{"status":"deleted"}`, "seller-1", "business")

	handler.UpdateListing(ctx)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d; body=%s", recorder.Code, http.StatusBadRequest, recorder.Body.String())
	}
	if client.update != nil {
		t.Fatal("material UpdateListing must not be called for an invalid status")
	}
}

func TestListMyListingsIncludesHiddenOwnedListings(t *testing.T) {
	gin.SetMode(gin.TestMode)
	client := &updateListingMaterialClient{
		listResponse: &materialpb.ListListingsResponse{Listings: []*materialpb.SupplyListing{
			{Id: "owned-active", SellerId: "seller-1", Status: "active"},
			{Id: "owned-hidden", SellerId: "seller-1", Status: "hidden"},
			{Id: "other-hidden", SellerId: "seller-2", Status: "hidden"},
		}},
	}
	handler := NewMaterialHandler(client, nil)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/api/my/listings", nil)
	ctx.Set("user_id", "seller-1")

	handler.ListMyListings(ctx)

	if recorder.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", recorder.Code, http.StatusOK, recorder.Body.String())
	}
	body := recorder.Body.String()
	if !strings.Contains(body, "owned-active") || !strings.Contains(body, "owned-hidden") {
		t.Fatalf("owned listings missing from response: %s", body)
	}
	if strings.Contains(body, "other-hidden") {
		t.Fatalf("another seller's listing leaked in response: %s", body)
	}
}
