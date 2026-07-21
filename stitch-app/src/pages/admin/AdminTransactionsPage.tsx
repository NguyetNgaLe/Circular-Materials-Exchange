import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { Loader2 } from 'lucide-react'

export default function AdminTransactionsPage() {
  const store = useStore()
  const [transactions, setTransactions] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const loadTransactions = async () => {
      setLoading(true)
      const data = await store.getTransactions()
      setTransactions(data)
      setLoading(false)
    }
    loadTransactions()
  }, [])

  if (loading) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <Loader2 className="animate-spin" size={32} />
          <p>Đang tải dữ liệu...</p>
        </div>
      </Layout>
    )
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản trị</p>
          <h1>Quản Lý Giao Dịch</h1>
          <p className="muted">Tất cả giao dịch trên hệ thống</p>
        </div>
      </div>

      <div className="stats-grid" style={{ marginBottom: 20 }}>
        <div className="stat-card">
          <div><strong>{transactions.length}</strong><span>Tổng giao dịch</span></div>
        </div>
        <div className="stat-card">
          <div><strong>{transactions.filter(t => t.status === 'completed').length}</strong><span>Hoàn tất</span></div>
        </div>
        <div className="stat-card">
          <div><strong>{transactions.filter(t => ['confirmed', 'in_progress'].includes(t.status)).length}</strong><span>Đang thực hiện</span></div>
        </div>
        <div className="stat-card">
          <div><strong>{transactions.filter(t => t.status === 'cancelled').length}</strong><span>Đã hủy</span></div>
        </div>
      </div>

      <div className="table-responsive">
        <table className="table">
          <thead>
            <tr><th>Mã GD</th><th>Vật liệu</th><th>Người mua</th><th>Người bán</th><th>Số lượng</th><th>Giá trị</th><th>Ngày tạo</th><th>Trạng thái</th></tr>
          </thead>
          <tbody>
            {transactions.map(t => (
              <tr key={t.id}>
                <td><Link to={`/transactions/${t.id}`} className="link-blue">{t.id.toUpperCase()}</Link></td>
                <td>{t.listingTitle}</td>
                <td>{t.buyerName}</td>
                <td>{t.sellerName}</td>
                <td>{t.quantity} {t.unit}</td>
                <td>{(t.quantity * t.agreedPrice).toLocaleString()}đ</td>
                <td>{t.createdAt}</td>
                <td><StatusBadge status={t.status} /></td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  )
}
