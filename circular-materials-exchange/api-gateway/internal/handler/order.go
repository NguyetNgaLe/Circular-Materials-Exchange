package handler

import (
	companypb "api-gateway/internal/pb/company"
	orderpb "api-gateway/internal/pb/order"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	order        orderpb.OrderServiceClient
	company      companypb.CompanyServiceClient
	notification *NotificationHandler
}

func NewOrderHandler(order orderpb.OrderServiceClient, company companypb.CompanyServiceClient) *OrderHandler {
	return &OrderHandler{order: order, company: company}
}

func (h *OrderHandler) SetNotificationHandler(handler *NotificationHandler) {
	h.notification = handler
}

func (h *OrderHandler) CreateOffer(c *gin.Context) {
	var req struct {
		Type          string  `json:"type" binding:"required"`
		ListingID     string  `json:"listingId" binding:"required"`
		ListingTitle  string  `json:"listingTitle"`
		BuyerName     string  `json:"buyerName"`
		SellerID      string  `json:"sellerId" binding:"required"`
		SellerName    string  `json:"sellerName"`
		Quantity      float64 `json:"quantity" binding:"required"`
		Unit          string  `json:"unit"`
		ProposedPrice float64 `json:"proposedPrice" binding:"required"`
		Currency      string  `json:"currency"`
		Message       string  `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	ctx, cancel := rpcContext(c)
	defer cancel()
	company, err := getCompanyByOwner(ctx, h.company, stringValue(userID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Bạn cần có hồ sơ doanh nghiệp để gửi đề nghị mua"})
		return
	}
	if company.GetStatus() != "verified" {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "message": "Doanh nghiệp chưa được admin duyệt"})
		return
	}
	buyerName := req.BuyerName
	if buyerName == "" {
		buyerName = company.GetName()
	}
	offer, err := h.order.CreateOffer(ctx, &orderpb.CreateOfferRequest{
		Type: req.Type, ListingId: req.ListingID, ListingTitle: req.ListingTitle,
		BuyerId: stringValue(userID), BuyerName: buyerName, SellerId: req.SellerID,
		SellerName: req.SellerName, Quantity: req.Quantity, Unit: req.Unit,
		ProposedPrice: req.ProposedPrice, Currency: req.Currency, Message: req.Message,
	})
	if err != nil {
		writeRPCError(c, err, "Lỗi tạo đề nghị")
		return
	}
	if h.notification != nil {
		h.notification.CreateNotification(req.SellerID, "Đề nghị mua mới",
			fmt.Sprintf("%s muốn mua %s %.0f %s với giá %.0f VND/%s",
				buyerName, req.ListingTitle, req.Quantity, req.Unit, req.ProposedPrice, req.Unit),
			"offer", offer.GetId())
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": offerJSON(offer)})
}

func (h *OrderHandler) ListOffers(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	filterUserID := stringValue(userID)
	if stringValue(role) == "admin" {
		filterUserID = ""
	}
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListOffers(ctx, &orderpb.ListOffersRequest{
		UserId: filterUserID, Role: c.Query("role"), Status: c.Query("status"),
		Page: page, PageSize: pageSize,
	})
	if err != nil {
		writeRPCError(c, err, "Lỗi lấy danh sách đề nghị")
		return
	}
	items := make([]gin.H, 0, len(response.GetOffers()))
	for _, offer := range response.GetOffers() {
		items = append(items, offerJSON(offer))
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"offers": items, "total": response.GetTotal()}})
}

func (h *OrderHandler) AcceptOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	actorName, _ := c.Get("name")
	if stringValue(actorName) == "" {
		actorName, _ = c.Get("email")
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	transaction, err := h.order.AcceptOffer(ctx, &orderpb.AcceptOfferRequest{
		OfferId: c.Param("id"), ActorId: stringValue(userID), ActorName: stringValue(actorName),
	})
	if err != nil {
		writeRPCError(c, err, "Lỗi chấp nhận đề nghị")
		return
	}
	if h.notification != nil {
		h.notification.CreateNotification(transaction.GetBuyerId(), "Đề nghị đã được chấp nhận",
			fmt.Sprintf("%s đã chấp nhận đề nghị mua %s. Giao dịch đã được tạo.",
				transaction.GetSellerName(), transaction.GetListingTitle()),
			"offer_accepted", c.Param("id"))
	}
	escrowData := gin.H{}
	if list, listErr := h.order.ListEscrows(ctx, &orderpb.PageRequest{Page: 1, PageSize: 200}); listErr == nil {
		for _, escrow := range list.GetEscrows() {
			if escrow.GetTransactionId() == transaction.GetId() {
				escrowData = escrowJSON(escrow)
				break
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
		"id": c.Param("id"), "status": "accepted",
		"transaction": gin.H{"id": transaction.GetId(), "status": transaction.GetStatus()},
		"escrow":      escrowData,
	}})
}

func (h *OrderHandler) RejectOffer(c *gin.Context) {
	userID, _ := c.Get("user_id")
	ctx, cancel := rpcContext(c)
	defer cancel()
	offer, err := h.order.RejectOffer(ctx, &orderpb.RejectOfferRequest{OfferId: c.Param("id"), ActorId: stringValue(userID)})
	if err != nil {
		writeRPCError(c, err, "Lỗi từ chối đề nghị")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": offer.GetId(), "status": offer.GetStatus()}})
}

func (h *OrderHandler) ListTransactions(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	filterUserID := stringValue(userID)
	if stringValue(role) == "admin" {
		filterUserID = ""
	}
	page, pageSize := pagination(c)
	ctx, cancel := rpcContext(c)
	defer cancel()
	response, err := h.order.ListTransactions(ctx, &orderpb.ListTransactionsRequest{
		UserId: filterUserID, Status: c.Query("status"), Page: page, PageSize: pageSize,
	})
	if err != nil {
		writeRPCError(c, err, "Lỗi lấy danh sách giao dịch")
		return
	}
	items := make([]gin.H, 0, len(response.GetTransactions()))
	for _, transaction := range response.GetTransactions() {
		items = append(items, transactionJSON(transaction))
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"transactions": items, "total": response.GetTotal()}})
}

func (h *OrderHandler) GetTransaction(c *gin.Context) {
	ctx, cancel := rpcContext(c)
	defer cancel()
	transaction, err := h.order.GetTransaction(ctx, &orderpb.GetTransactionRequest{Id: c.Param("id")})
	if err != nil {
		writeRPCError(c, err, "Không tìm thấy giao dịch")
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": transactionJSON(transaction)})
}

func (h *OrderHandler) UpdateTransactionStatus(c *gin.Context) {
	var req struct {
		Status string `json:"status" binding:"required"`
		Note   string `json:"note"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
		return
	}
	userID, _ := c.Get("user_id")
	actorName, _ := c.Get("name")
	if stringValue(actorName) == "" {
		actorName = userID
	}
	ctx, cancel := rpcContext(c)
	defer cancel()
	transaction, err := h.order.UpdateTransactionStatus(ctx, &orderpb.UpdateStatusRequest{
		TransactionId: c.Param("id"), NewStatus: req.Status,
		ActorId: stringValue(userID), ActorName: stringValue(actorName), Note: req.Note,
	})
	if err != nil {
		writeRPCError(c, err, "Lỗi cập nhật giao dịch")
		return
	}
	if h.notification != nil {
		switch req.Status {
		case "in_progress":
			h.notification.CreateNotification(transaction.GetBuyerId(), "Người bán đã giao hàng",
				fmt.Sprintf("Sản phẩm %s đã được giao. Vui lòng xác nhận nhận hàng.", transaction.GetListingTitle()),
				"transaction", transaction.GetId())
		case "completed":
			h.notification.CreateNotification(transaction.GetSellerId(), "Giao dịch hoàn tất",
				fmt.Sprintf("Giao dịch %s đã hoàn tất. Tiền đã chuyển vào ví của bạn.", transaction.GetListingTitle()),
				"transaction", transaction.GetId())
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": transaction.GetId(), "status": transaction.GetStatus()}})
}

func offerJSON(offer *orderpb.Offer) gin.H {
	return gin.H{
		"id": offer.GetId(), "type": offer.GetType(), "listingId": offer.GetListingId(),
		"listingTitle": offer.GetListingTitle(), "buyerId": offer.GetBuyerId(),
		"buyerName": offer.GetBuyerName(), "sellerId": offer.GetSellerId(),
		"sellerName": offer.GetSellerName(), "quantity": offer.GetQuantity(),
		"unit": offer.GetUnit(), "proposedPrice": offer.GetProposedPrice(),
		"currency": offer.GetCurrency(), "message": offer.GetMessage(),
		"status": offer.GetStatus(), "createdAt": offer.GetCreatedAt(),
	}
}

func transactionJSON(transaction *orderpb.Transaction) gin.H {
	events := make([]gin.H, 0, len(transaction.GetEvents()))
	for _, event := range transaction.GetEvents() {
		events = append(events, gin.H{
			"id": event.GetId(), "actorId": event.GetActorId(), "actorName": event.GetActorName(),
			"fromStatus": event.GetFromStatus(), "toStatus": event.GetToStatus(),
			"note": event.GetNote(), "createdAt": event.GetCreatedAt(),
		})
	}
	return gin.H{
		"id": transaction.GetId(), "offerId": transaction.GetOfferId(),
		"listingTitle": transaction.GetListingTitle(), "buyerId": transaction.GetBuyerId(),
		"buyerName": transaction.GetBuyerName(), "sellerId": transaction.GetSellerId(),
		"sellerName": transaction.GetSellerName(), "quantity": transaction.GetQuantity(),
		"unit": transaction.GetUnit(), "agreedPrice": transaction.GetAgreedPrice(),
		"currency": transaction.GetCurrency(), "paymentStatus": transaction.GetPaymentStatus(),
		"paymentMethod": transaction.GetPaymentMethod(), "settlementNote": transaction.GetSettlementNote(),
		"status": transaction.GetStatus(), "createdAt": transaction.GetCreatedAt(), "events": events,
	}
}

func escrowJSON(escrow *orderpb.EscrowTransaction) gin.H {
	return gin.H{
		"id": escrow.GetId(), "amount": escrow.GetAmount(), "feeAmount": escrow.GetFeeAmount(),
		"sellerAmount": escrow.GetSellerAmount(), "holdUntil": escrow.GetHoldUntil(),
	}
}
