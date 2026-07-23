import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { ArrowDownLeft, Loader2, Building2 } from 'lucide-react'

export default function SellerWalletPage() {
  const store = useStore()
  const [wallet, setWallet] = useState<any>(null)
  const [transactions, setTransactions] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => { loadData() }, [])

  const loadData = async () => {
    setLoading(true)
    try {
      const token = localStorage.getItem('token')
      const headers = { 'Authorization': `Bearer ${token}` }
      const [walletRes, txRes] = await Promise.all([
        fetch('/api/seller/wallet', { headers }).then(r => r.json()),
        fetch('/api/seller/wallet/transactions', { headers }).then(r => r.json()),
      ])
      if (walletRes.success) setWallet(walletRes.data)
      if (txRes.success) setTransactions(txRes.data.transactions || [])
    } catch (err) {
      console.error('Failed to load data:', err)
    }
    setLoading(false)
  }

  if (loading) {
    return <Layout showSidebar><div className="empty-state"><Loader2 className="animate-spin" size={32} /><p>Đang tải...</p></div></Layout>
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Tai khoan</p>
          <h1>Lịch sử Giao dịch</h1>
        </div>
      </div>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-icon green"><ArrowDownLeft size={22} /></div>
          <div><strong>{(wallet?.totalEarned || 0).toLocaleString()}đ</strong><span>Tổng tiền đã nhận</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-icon orange"><ArrowDownLeft size={22} /></div>
          <div><strong>{(wallet?.totalFeesPaid || 0).toLocaleString()}đ</strong><span>Tổng phí đã trả</span></div>
        </div>
      </div>

      <div className="panel">
        <div className="panel-header">
          <h2>Lịch sử tien chuyen ve tai khoan</h2>
        </div>
        <div style={{ padding: 12, background: '#eff6ff', borderRadius: 8, marginBottom: 16 }}>
          <p style={{ fontSize: 13, color: '#1e40af' }}>
            Tien tu dong chuyen ve tai khoan Ngân hang da lien ket khi Giao dịch Hoàn tất. Phi Giao dịch 2% da duoc tru.
          </p>
        </div>
        {transactions.length === 0 ? (
          <p className="muted" style={{ padding: 20 }}>Chưa có Giao dịch nao</p>
        ) : (
          <div className="table-responsive">
            <table className="table">
              <thead><tr><th>Ngay</th><th>Số tiền nhan</th><th>Phi</th><th>Thuc nhan</th><th>Mô tả</th></tr></thead>
              <tbody>
                {transactions.map(tx => (
                  <tr key={tx.id}>
                    <td>{new Date(tx.createdAt).toLocaleDateString('vi-VN')}</td>
                    <td>{tx.amount.toLocaleString()}đ</td>
                    <td style={{ color: 'red' }}>{tx.feeAmount ? tx.feeAmount.toLocaleString() : '0'}đ</td>
                    <td><strong style={{ color: 'green' }}>{tx.amount.toLocaleString()}đ</strong></td>
                    <td>{tx.description}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </Layout>
  )
}
