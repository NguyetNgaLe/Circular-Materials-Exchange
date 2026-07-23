package handler

import (
	"context"
	"order-service/internal/repository"
	"order-service/pb"
)

func (h *OrderHandler) GetFinanceOverview(ctx context.Context, _ *pb.Empty) (*pb.FinanceOverview, error) {
	item, err := h.svc.GetFinanceOverview()
	if err != nil {
		return nil, err
	}
	monthly := make([]*pb.MonthlyRevenue, 0, len(item.MonthlyRevenue))
	for _, value := range item.MonthlyRevenue {
		monthly = append(monthly, &pb.MonthlyRevenue{Month: value.Month, Revenue: value.Revenue})
	}
	return &pb.FinanceOverview{
		TotalIncome: item.TotalIncome, MonthIncome: item.MonthIncome,
		TotalTransactions: item.TotalTransactions, WalletBalance: item.WalletBalance,
		MonthlyRevenue: monthly,
	}, nil
}

func (h *OrderHandler) ListPlatformFees(ctx context.Context, req *pb.PageRequest) (*pb.ListPlatformFeesResponse, error) {
	items, total, err := h.svc.ListPlatformFees(int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.PlatformFee, 0, len(items))
	for i := range items {
		result = append(result, platformFeeToProto(&items[i]))
	}
	return &pb.ListPlatformFeesResponse{Fees: result, Total: int32(total)}, nil
}

func (h *OrderHandler) GetPlatformWallet(ctx context.Context, _ *pb.Empty) (*pb.PlatformWallet, error) {
	item, err := h.svc.GetPlatformWallet()
	if err != nil {
		return nil, err
	}
	return &pb.PlatformWallet{Balance: item.Balance, TotalIncome: item.TotalIncome, TotalExpense: item.TotalExpense}, nil
}

func (h *OrderHandler) ListPlatformWalletTransactions(ctx context.Context, req *pb.PageRequest) (*pb.ListPlatformWalletTransactionsResponse, error) {
	items, total, err := h.svc.ListPlatformWalletTransactions(int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.PlatformWalletTransaction, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.PlatformWalletTransaction{
			Id: item.ID, Type: item.Type, Amount: item.Amount, BalanceAfter: item.BalanceAfter,
			ReferenceType: item.ReferenceType, ReferenceId: item.ReferenceID,
			Description: item.Description, CreatedAt: item.CreatedAt,
		})
	}
	return &pb.ListPlatformWalletTransactionsResponse{Transactions: result, Total: int32(total)}, nil
}

func (h *OrderHandler) CollectFee(ctx context.Context, req *pb.CollectFeeRequest) (*pb.CollectFeeResponse, error) {
	id, feeAmount, balance, err := h.svc.CollectFee(
		req.GetTransactionId(), req.GetSellerId(), req.GetBuyerId(),
		req.GetTransactionAmount(), req.GetFeeRate(),
	)
	if err != nil {
		return nil, err
	}
	return &pb.CollectFeeResponse{FeeId: id, FeeAmount: feeAmount, WalletBalance: balance}, nil
}

func (h *OrderHandler) CreateEscrow(ctx context.Context, req *pb.CreateEscrowRequest) (*pb.EscrowTransaction, error) {
	item, err := h.svc.CreateEscrow(
		req.GetTransactionId(), req.GetBuyerId(), req.GetBuyerName(),
		req.GetSellerId(), req.GetSellerName(), req.GetAmount(), int(req.GetHoldDays()),
	)
	if err != nil {
		return nil, err
	}
	return escrowToProto(item), nil
}

func (h *OrderHandler) ListEscrows(ctx context.Context, req *pb.PageRequest) (*pb.ListEscrowsResponse, error) {
	items, total, err := h.svc.ListEscrows(int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.EscrowTransaction, 0, len(items))
	for i := range items {
		result = append(result, escrowToProto(&items[i]))
	}
	return &pb.ListEscrowsResponse{Escrows: result, Total: int32(total)}, nil
}

func (h *OrderHandler) ReleaseEscrow(ctx context.Context, req *pb.ReleaseEscrowRequest) (*pb.EscrowTransaction, error) {
	item, err := h.svc.ReleaseEscrow(req.GetId())
	if err != nil {
		return nil, err
	}
	return escrowToProto(item), nil
}

func (h *OrderHandler) GetSellerWallet(ctx context.Context, req *pb.SellerRequest) (*pb.SellerWallet, error) {
	item, err := h.svc.GetSellerWallet(req.GetSellerId())
	if err != nil {
		return nil, err
	}
	return sellerWalletToProto(item), nil
}

func (h *OrderHandler) ListSellerWalletTransactions(ctx context.Context, req *pb.SellerRequest) (*pb.ListSellerWalletTransactionsResponse, error) {
	items, total, err := h.svc.ListSellerWalletTransactions(req.GetSellerId(), int(req.GetPage()), int(req.GetPageSize()))
	if err != nil {
		return nil, err
	}
	result := make([]*pb.SellerWalletTransaction, 0, len(items))
	for _, item := range items {
		result = append(result, &pb.SellerWalletTransaction{
			Id: item.ID, Type: item.Type, Amount: item.Amount, BalanceAfter: item.BalanceAfter,
			ReferenceType: item.ReferenceType, ReferenceId: item.ReferenceID,
			Description: item.Description, CreatedAt: item.CreatedAt,
		})
	}
	return &pb.ListSellerWalletTransactionsResponse{Transactions: result, Total: int32(total)}, nil
}

func (h *OrderHandler) CreateWithdrawal(ctx context.Context, req *pb.CreateWithdrawalRequest) (*pb.WithdrawalRequest, error) {
	item, err := h.svc.CreateWithdrawal(
		req.GetSellerId(), req.GetSellerName(), req.GetAmount(),
		req.GetBankName(), req.GetBankAccount(), req.GetBankOwner(),
	)
	if err != nil {
		return nil, err
	}
	return withdrawalToProto(item), nil
}

func (h *OrderHandler) ListSellerWithdrawals(ctx context.Context, req *pb.SellerRequest) (*pb.ListWithdrawalsResponse, error) {
	return h.listWithdrawals(req.GetSellerId(), int(req.GetPage()), int(req.GetPageSize()))
}

func (h *OrderHandler) ListWithdrawals(ctx context.Context, req *pb.PageRequest) (*pb.ListWithdrawalsResponse, error) {
	return h.listWithdrawals("", int(req.GetPage()), int(req.GetPageSize()))
}

func (h *OrderHandler) ApproveWithdrawal(ctx context.Context, req *pb.ProcessWithdrawalRequest) (*pb.WithdrawalRequest, error) {
	item, err := h.svc.ApproveWithdrawal(req.GetId())
	if err != nil {
		return nil, err
	}
	return withdrawalToProto(item), nil
}

func (h *OrderHandler) RejectWithdrawal(ctx context.Context, req *pb.ProcessWithdrawalRequest) (*pb.WithdrawalRequest, error) {
	item, err := h.svc.RejectWithdrawal(req.GetId(), req.GetReason())
	if err != nil {
		return nil, err
	}
	return withdrawalToProto(item), nil
}

func (h *OrderHandler) listWithdrawals(sellerID string, page, pageSize int) (*pb.ListWithdrawalsResponse, error) {
	items, total, err := h.svc.ListWithdrawals(sellerID, page, pageSize)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.WithdrawalRequest, 0, len(items))
	for i := range items {
		result = append(result, withdrawalToProto(&items[i]))
	}
	return &pb.ListWithdrawalsResponse{Withdrawals: result, Total: int32(total)}, nil
}

func platformFeeToProto(item *repository.PlatformFee) *pb.PlatformFee {
	return &pb.PlatformFee{
		Id: item.ID, TransactionId: item.TransactionID, SellerId: item.SellerID,
		BuyerId: item.BuyerID, TransactionAmount: item.TransactionAmount,
		FeeRate: item.FeeRate, FeeAmount: item.FeeAmount, FeeType: item.FeeType,
		Status: item.Status, CollectedAt: item.CollectedAt, CreatedAt: item.CreatedAt,
		ListingTitle: item.ListingTitle,
	}
}

func escrowToProto(item *repository.EscrowTransaction) *pb.EscrowTransaction {
	return &pb.EscrowTransaction{
		Id: item.ID, TransactionId: item.TransactionID, BuyerId: item.BuyerID,
		BuyerName: item.BuyerName, SellerId: item.SellerID, SellerName: item.SellerName,
		Amount: item.Amount, FeeRate: item.FeeRate, FeeAmount: item.FeeAmount,
		SellerAmount: item.SellerAmount, Status: item.Status, HoldUntil: item.HoldUntil,
		ReleasedAt: item.ReleasedAt, CreatedAt: item.CreatedAt,
	}
}

func sellerWalletToProto(item *repository.SellerWallet) *pb.SellerWallet {
	return &pb.SellerWallet{
		SellerId: item.SellerID, SellerName: item.SellerName, Balance: item.Balance,
		TotalEarned: item.TotalEarned, TotalFeesPaid: item.TotalFees,
		TotalWithdrawn: item.TotalWithdrawn,
	}
}

func withdrawalToProto(item *repository.WithdrawalRequest) *pb.WithdrawalRequest {
	return &pb.WithdrawalRequest{
		Id: item.ID, SellerId: item.SellerID, SellerName: item.SellerName,
		Amount: item.Amount, BankName: item.BankName, BankAccount: item.BankAccount,
		BankOwner: item.BankOwner, Status: item.Status, AdminNote: item.AdminNote,
		ProcessedAt: item.ProcessedAt, CreatedAt: item.CreatedAt,
	}
}
