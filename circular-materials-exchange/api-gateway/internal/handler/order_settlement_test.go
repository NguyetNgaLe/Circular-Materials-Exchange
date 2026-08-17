package handler

import (
	"testing"

	orderpb "api-gateway/internal/pb/order"
)

func TestCalculateCompanySettlement(t *testing.T) {
	const companyID = "company-1"
	transactions := []*orderpb.Transaction{
		{
			Id: "sale", ListingTitle: "Nhựa PP", SellerId: companyID, BuyerId: "buyer",
			BuyerName: "Doanh nghiệp mua", Quantity: 100, AgreedPrice: 1000, Status: "completed",
		},
		{
			Id: "purchase", ListingTitle: "Pallet gỗ", BuyerId: companyID, SellerId: "seller",
			SellerName: "Doanh nghiệp bán", Quantity: 10, AgreedPrice: 2000, Status: "completed",
		},
		{
			Id: "pending", SellerId: companyID, BuyerId: "buyer-2",
			Quantity: 999, AgreedPrice: 999, Status: "in_progress",
		},
	}

	result := calculateCompanySettlement(companyID, transactions)
	if result.TotalReceived != 98000 {
		t.Fatalf("TotalReceived = %v, want 98000", result.TotalReceived)
	}
	if result.TotalPaid != 20000 {
		t.Fatalf("TotalPaid = %v, want 20000", result.TotalPaid)
	}
	if len(result.Entries) != 2 {
		t.Fatalf("len(Entries) = %d, want 2", len(result.Entries))
	}
	if result.Entries[0].FeeAmount != 2000 || result.Entries[0].SettledAmount != 98000 {
		t.Fatalf("sale entry = %+v", result.Entries[0])
	}
	if result.Entries[1].FeeAmount != 0 || result.Entries[1].SettledAmount != 20000 {
		t.Fatalf("purchase entry = %+v", result.Entries[1])
	}
}
