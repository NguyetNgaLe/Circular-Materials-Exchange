import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { ArrowDownLeft, DollarSign, Loader2 } from 'lucide-react'

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
          <p className="eyebrow">Tài khoản người bán</p>
          <h1>Ví người bán</h1>
          <p className="muted">Theo dõi số dư và các khoản tiền bán hàng đã được ghi nhận</p>
        </div>
        <Link to="/transactions" className="btn btn-outline">Xem lịch sử giao dịch</Link>
      </div>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-icon blue"><DollarSign size={22} /></div>
          <div><strong>{(wallet?.balance || 0).toLocaleString()}đ</strong><span>Số dư khả dụng</span></div>
        </div>
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
          <h2>Lịch sử tiền về ví</h2>
        </div>
        <div style={{ padding: 12, background: '#eff6ff', borderRadius: 8, marginBottom: 16 }}>
          <p style={{ fontSize: 13, color: '#1e40af' }}>
            Tiền được ghi nhận vào ví người bán khi giao dịch hoàn tất. Phí giao dịch 2% đã được khấu trừ trước khi tiền về ví.
          </p>
        </div>
        {transactions.length === 0 ? (
          <div className="empty-state">
            <h3>Chưa có khoản tiền nào được ghi nhận</h3>
            <p>Trang này chỉ hiển thị dòng tiền bán hàng. Lịch sử mua và bán đầy đủ nằm trong mục Lịch sử giao dịch.</p>
            <Link to="/transactions" className="link-blue">Xem lịch sử giao dịch</Link>
          </div>
        ) : (
          <div className="table-responsive">
            <table className="table">
              <thead><tr><th>Ngày</th><th>Loại</th><th>Số tiền thực nhận</th><th>Mô tả</th></tr></thead>
              <tbody>
                {transactions.map(tx => (
                  <tr key={tx.id}>
                    <td>{new Date(tx.createdAt).toLocaleDateString('vi-VN')}</td>
                    <td>{tx.type === 'credit' ? 'Tiền vào' : 'Tiền ra'}</td>
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
