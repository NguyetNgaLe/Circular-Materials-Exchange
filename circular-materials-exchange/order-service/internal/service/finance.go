package service

import "order-service/internal/repository"

func (s *OrderService) GetFinanceOverview() (*repository.FinanceOverview, error) {
	return s.repo.GetFinanceOverview()
}

func (s *OrderService) ListPlatformFees(page, pageSize int) ([]repository.PlatformFee, int64, error) {
	page, pageSize = normalizePage(page, pageSize)
	return s.repo.ListPlatformFees(page, pageSize)
}

func (s *OrderService) GetPlatformWallet() (*repository.PlatformWallet, error) {
	return s.repo.GetPlatformWallet()
}

func (s *OrderService) ListPlatformWalletTransactions(page, pageSize int) ([]repository.PlatformWalletTransaction, int64, error) {
	page, pageSize = normalizePage(page, pageSize)
	return s.repo.ListPlatformWalletTransactions(page, pageSize)
}

func (s *OrderService) CollectFee(transactionID, sellerID, buyerID string, amount, feeRate float64) (string, float64, float64, error) {
	return s.repo.CollectFee(transactionID, sellerID, buyerID, amount, feeRate)
}

func (s *OrderService) CreateEscrow(transactionID, buyerID, buyerName, sellerID, sellerName string, amount float64, holdDays int) (*repository.EscrowTransaction, error) {
	return s.repo.CreateEscrow(transactionID, buyerID, buyerName, sellerID, sellerName, amount, holdDays, true)
}

func (s *OrderService) ListEscrows(page, pageSize int) ([]repository.EscrowTransaction, int64, error) {
	page, pageSize = normalizePage(page, pageSize)
	return s.repo.ListEscrows(page, pageSize)
}

func (s *OrderService) ReleaseEscrow(id string) (*repository.EscrowTransaction, error) {
	return s.repo.ReleaseEscrow(id, false)
}

func (s *OrderService) GetSellerWallet(sellerID string) (*repository.SellerWallet, error) {
	return s.repo.GetSellerWallet(sellerID)
}

func (s *OrderService) ListSellerWalletTransactions(sellerID string, page, pageSize int) ([]repository.SellerWalletTransaction, int64, error) {
	page, pageSize = normalizePage(page, pageSize)
	return s.repo.ListSellerWalletTransactions(sellerID, page, pageSize)
}

func (s *OrderService) CreateWithdrawal(sellerID, sellerName string, amount float64, bankName, bankAccount, bankOwner string) (*repository.WithdrawalRequest, error) {
	return s.repo.CreateWithdrawal(sellerID, sellerName, amount, bankName, bankAccount, bankOwner)
}

func (s *OrderService) ListWithdrawals(sellerID string, page, pageSize int) ([]repository.WithdrawalRequest, int64, error) {
	page, pageSize = normalizePage(page, pageSize)
	return s.repo.ListWithdrawals(sellerID, page, pageSize)
}

func (s *OrderService) ApproveWithdrawal(id string) (*repository.WithdrawalRequest, error) {
	return s.repo.ProcessWithdrawal(id, "completed", "")
}

func (s *OrderService) RejectWithdrawal(id, reason string) (*repository.WithdrawalRequest, error) {
	return s.repo.ProcessWithdrawal(id, "rejected", reason)
}

func normalizePage(page, pageSize int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 50
	}
	return page, pageSize
}
