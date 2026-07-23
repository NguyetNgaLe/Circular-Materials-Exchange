package handler

import (
	orderpb "api-gateway/internal/pb/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FinanceHandler struct {
	order orderpb.OrderServiceClient
}

func NewFinanceHandler(order orderpb.OrderServiceClient) *FinanceHandler {
	return &FinanceHandler{order: order}
}

func (h *FinanceHandler) GetOverview(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.GetFinanceOverview(ctx, &orderpb.Empty{})
	if err != nil {
		writeRPCError(c, err, "Loi lay tong quan tai chinh")
		return
	}
	monthly := make([]gin.H, 0, len(response.GetMonthlyRevenue()))
	for _, item := range response.GetMonthlyRevenue() {
		monthly = append(monthly, gin.H{"month": item.GetMonth(), "total": item.GetRevenue()})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"totalIncome": response.GetTotalIncome(), "monthIncome": response.GetMonthIncome(),
		"totalTransactions": response.GetTotalTransactions(), "walletBalance": response.GetWalletBalance(),
		"monthlyData": monthly,
	}})
}

func (h *FinanceHandler) ListFees(c *gin.Context) {
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListPlatformFees(ctx, &orderpb.PageRequest{Page: page, PageSize: pageSize})
	if err != nil {
		writeRPCError(c, err, "Loi lay danh sach phi")
		return
	}
	items := make([]gin.H, 0, len(response.GetFees()))
	for _, fee := range response.GetFees() {
		items = append(items, gin.H{
			"id": fee.GetId(), "transactionId": fee.GetTransactionId(),
			"sellerId": fee.GetSellerId(), "buyerId": fee.GetBuyerId(),
			"transactionAmount": fee.GetTransactionAmount(), "feeRate": fee.GetFeeRate(),
			"feeAmount": fee.GetFeeAmount(), "feeType": fee.GetFeeType(),
			"status": fee.GetStatus(), "collectedAt": fee.GetCollectedAt(),
			"createdAt": fee.GetCreatedAt(), "listingTitle": fee.GetListingTitle(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"fees": items, "total": response.GetTotal()}})
}

func (h *FinanceHandler) GetWallet(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.GetPlatformWallet(ctx, &orderpb.Empty{})
	if err != nil {
		writeRPCError(c, err, "Loi lay vi san")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"balance": response.GetBalance(), "totalIncome": response.GetTotalIncome(),
		"totalExpense": response.GetTotalExpense(),
	}})
}

func (h *FinanceHandler) ListWalletTransactions(c *gin.Context) {
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListPlatformWalletTransactions(ctx, &orderpb.PageRequest{Page: page, PageSize: pageSize})
	if err != nil {
		writeRPCError(c, err, "Loi lay lich su vi san")
		return
	}
	items := make([]gin.H, 0, len(response.GetTransactions()))
	for _, transaction := range response.GetTransactions() {
		items = append(items, gin.H{
			"id": transaction.GetId(), "type": transaction.GetType(),
			"amount": transaction.GetAmount(), "referenceType": transaction.GetReferenceType(),
			"referenceId": transaction.GetReferenceId(), "description": transaction.GetDescription(),
			"balanceAfter": transaction.GetBalanceAfter(), "createdAt": transaction.GetCreatedAt(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": items, "total": response.GetTotal()}})
}

func (h *FinanceHandler) CollectFee(c *gin.Context) {
	var req struct {
		TransactionID string  `json:"transactionId" binding:"required"`
		Amount        float64 `json:"amount" binding:"required"`
		SellerID      string  `json:"sellerId" binding:"required"`
		BuyerID       string  `json:"buyerId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.CollectFee(ctx, &orderpb.CollectFeeRequest{
		TransactionId: req.TransactionID, TransactionAmount: req.Amount,
		SellerId: req.SellerID, BuyerId: req.BuyerID, FeeRate: 0.02,
	})
	if err != nil {
		writeRPCError(c, err, "Loi thu phi")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"feeId": response.GetFeeId(), "feeAmount": response.GetFeeAmount(),
		"feeRate": 0.02, "balance": response.GetWalletBalance(),
	}})
}
