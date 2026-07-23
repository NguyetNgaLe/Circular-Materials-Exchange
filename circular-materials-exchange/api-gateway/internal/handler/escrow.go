package handler

import (
	orderpb "api-gateway/internal/pb/order"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type EscrowHandler struct {
	order orderpb.OrderServiceClient
}

func NewEscrowHandler(order orderpb.OrderServiceClient) *EscrowHandler {
	return &EscrowHandler{order: order}
}

func (h *EscrowHandler) CreateEscrow(c *gin.Context) {
	var req struct {
		TransactionID string  `json:"transactionId" binding:"required"`
		BuyerID       string  `json:"buyerId" binding:"required"`
		BuyerName     string  `json:"buyerName"`
		SellerID      string  `json:"sellerId" binding:"required"`
		SellerName    string  `json:"sellerName"`
		Amount        float64 `json:"amount" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	escrow, err := h.order.CreateEscrow(ctx, &orderpb.CreateEscrowRequest{
		TransactionId: req.TransactionID, BuyerId: req.BuyerID, BuyerName: req.BuyerName,
		SellerId: req.SellerID, SellerName: req.SellerName, Amount: req.Amount, HoldDays: 3,
	})
	if err != nil {
		writeRPCError(c, err, "Loi tao escrow")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"id": escrow.GetId(), "amount": escrow.GetAmount(), "feeAmount": escrow.GetFeeAmount(),
		"sellerAmount": escrow.GetSellerAmount(), "status": escrow.GetStatus(), "holdUntil": escrow.GetHoldUntil(),
	}})
}

func (h *EscrowHandler) ReleaseEscrow(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	escrow, err := h.order.ReleaseEscrow(ctx, &orderpb.ReleaseEscrowRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Escrow not found or already released")
		return
	}
	wallet, _ := h.order.GetSellerWallet(ctx, &orderpb.SellerRequest{SellerId: escrow.GetSellerId()})
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"escrowId": escrow.GetId(), "status": escrow.GetStatus(),
		"sellerAmount": escrow.GetSellerAmount(), "feeAmount": escrow.GetFeeAmount(),
		"sellerBalance": wallet.GetBalance(),
	}})
}

func (h *EscrowHandler) GetSellerWallet(c *gin.Context) {
	userID, _ := c.Get("user_id")
	ctx, cancel := rpcContext(c)
	defer cancel()
	wallet, err := h.order.GetSellerWallet(ctx, &orderpb.SellerRequest{SellerId: stringValue(userID)})
	if err != nil {
		writeRPCError(c, err, "Loi lay vi seller")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"balance": wallet.GetBalance(), "totalEarned": wallet.GetTotalEarned(),
		"totalFeesPaid": wallet.GetTotalFeesPaid(), "totalWithdrawn": wallet.GetTotalWithdrawn(),
	}})
}

func (h *EscrowHandler) GetSellerWalletTransactions(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListSellerWalletTransactions(ctx, &orderpb.SellerRequest{
		SellerId: stringValue(userID), Page: page, PageSize: pageSize,
	})
	if err != nil {
		writeRPCError(c, err, "Loi lay lich su vi seller")
		return
	}
	items := make([]gin.H, 0, len(response.GetTransactions()))
	for _, transaction := range response.GetTransactions() {
		items = append(items, gin.H{
			"id": transaction.GetId(), "type": transaction.GetType(), "amount": transaction.GetAmount(),
			"balanceAfter": transaction.GetBalanceAfter(), "referenceType": transaction.GetReferenceType(),
			"description": transaction.GetDescription(), "createdAt": transaction.GetCreatedAt(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": items}})
}

func (h *EscrowHandler) CreateWithdrawal(c *gin.Context) {
	var req struct {
		Amount      float64 `json:"amount" binding:"required"`
		BankName    string  `json:"bankName" binding:"required"`
		BankAccount string  `json:"bankAccount" binding:"required"`
		BankOwner   string  `json:"bankOwner" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	userName, _ := c.Get("name")
	if stringValue(userName) == "" {
		userName, _ = c.Get("email")
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	item, err := h.order.CreateWithdrawal(ctx, &orderpb.CreateWithdrawalRequest{
		SellerId: stringValue(userID), SellerName: stringValue(userName), Amount: req.Amount,
		BankName: req.BankName, BankAccount: req.BankAccount, BankOwner: req.BankOwner,
	})
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "insufficient balance") {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "So du khong du"})
			return
		}
		writeRPCError(c, err, "Loi tao yeu cau rut tien")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"id": item.GetId(), "amount": item.GetAmount(), "status": item.GetStatus(),
		"bankName": item.GetBankName(), "bankAccount": item.GetBankAccount(),
	}})
}

func (h *EscrowHandler) GetSellerWithdrawals(c *gin.Context) {
	userID, _ := c.Get("user_id")
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListSellerWithdrawals(ctx, &orderpb.SellerRequest{
		SellerId: stringValue(userID), Page: page, PageSize: pageSize,
	})
	if err != nil {
		writeRPCError(c, err, "Loi lay yeu cau rut tien")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": withdrawalsJSON(response.GetWithdrawals(), false)}})
}

func (h *EscrowHandler) ListEscrows(c *gin.Context) {
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListEscrows(ctx, &orderpb.PageRequest{Page: page, PageSize: pageSize})
	if err != nil {
		writeRPCError(c, err, "Loi lay danh sach escrow")
		return
	}
	totalHolding := 0.0
	items := make([]gin.H, 0, len(response.GetEscrows()))
	for _, escrow := range response.GetEscrows() {
		if escrow.GetStatus() == "holding" {
			totalHolding += escrow.GetAmount()
		}
		items = append(items, gin.H{
			"id": escrow.GetId(), "transactionId": escrow.GetTransactionId(),
			"buyerId": escrow.GetBuyerId(), "buyerName": escrow.GetBuyerName(),
			"sellerId": escrow.GetSellerId(), "sellerName": escrow.GetSellerName(),
			"amount": escrow.GetAmount(), "feeAmount": escrow.GetFeeAmount(),
			"sellerAmount": escrow.GetSellerAmount(), "status": escrow.GetStatus(),
			"holdUntil": escrow.GetHoldUntil(), "releasedAt": escrow.GetReleasedAt(),
			"createdAt": escrow.GetCreatedAt(),
		})
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"escrows": items, "totalHolding": totalHolding}})
}

func (h *EscrowHandler) ListWithdrawals(c *gin.Context) {
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListWithdrawals(ctx, &orderpb.PageRequest{Page: page, PageSize: pageSize})
	if err != nil {
		writeRPCError(c, err, "Loi lay yeu cau rut tien")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"withdrawals": withdrawalsJSON(response.GetWithdrawals(), true)}})
}

func (h *EscrowHandler) ApproveWithdrawal(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	item, err := h.order.ApproveWithdrawal(ctx, &orderpb.ProcessWithdrawalRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Not found or already processed")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": item.GetId(), "status": item.GetStatus()}})
}

func (h *EscrowHandler) RejectWithdrawal(c *gin.Context) {
	var req struct {
		Reason string `json:"reason"`
	}
	_ = c.ShouldBindJSON(&req)
	ctx, cancel := rpcContext(c)
	defer cancel()
	item, err := h.order.RejectWithdrawal(ctx, &orderpb.ProcessWithdrawalRequest{Id: c.Param("id"), Reason: req.Reason})
	if err != nil {
		writeRPCError(c, err, "Not found or already processed")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": item.GetId(), "status": item.GetStatus()}})
}

func withdrawalsJSON(items []*orderpb.WithdrawalRequest, includeSeller bool) []gin.H {
	result := make([]gin.H, 0, len(items))
	for _, item := range items {
		value := gin.H{
			"id": item.GetId(), "amount": item.GetAmount(), "bankName": item.GetBankName(),
			"bankAccount": item.GetBankAccount(), "bankOwner": item.GetBankOwner(),
			"status": item.GetStatus(), "adminNote": item.GetAdminNote(),
			"processedAt": item.GetProcessedAt(), "createdAt": item.GetCreatedAt(),
		}
		if includeSeller {
			value["sellerId"] = item.GetSellerId()
			value["sellerName"] = item.GetSellerName()
		}
		result = append(result, value)
	}
	return result
}
