import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { Download, Loader2 } from 'lucide-react'

export default function AdminExportPage() {
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

  const handleExport = () => {
    const headers = ['ID', 'Vật liệu', 'Người mua', 'Người bán', 'Số lượng', 'Giá', 'Trạng thái', 'Ngày tạo']
    const rows = transactions.map(t => [t.id, t.listingTitle, t.buyerName, t.sellerName, `${t.quantity} ${t.unit}`, t.agreedPrice, t.status, t.createdAt])
    const csv = [headers.join(','), ...rows.map(r => r.join(','))].join('\n')
    const blob = new Blob([csv], { type: 'text/csv' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'transactions.csv'
    a.click()
    URL.revokeObjectURL(url)
  }

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
          <h1>Xuất Báo Cáo</h1>
        </div>
      </div>

      <div className="grid-2">
        <div className="panel">
          <h2>Xuất giao dịch CSV</h2>
          <p className="muted">Tải về danh sách tất cả giao dịch dưới dạng file CSV.</p>
          <div className="info-list" style={{ marginTop: 16 }}>
            <div className="info-row"><span>Tổng số giao dịch</span><strong>{transactions.length}</strong></div>
            <div className="info-row"><span>Hoàn tất</span><strong>{transactions.filter(t => t.status === 'completed').length}</strong></div>
            <div className="info-row"><span>Đang thực hiện</span><strong>{transactions.filter(t => t.status === 'in_progress').length}</strong></div>
          </div>
          <button className="btn btn-primary" onClick={handleExport} style={{ marginTop: 16 }}>
            <Download size={16} /> Tải CSV
          </button>
        </div>

        <div className="panel">
          <h2>Xem trước dữ liệu</h2>
          <div className="table-responsive">
            <table className="table">
              <thead>
                <tr><th>ID</th><th>Vật liệu</th><th>Trạng thái</th><th>Giá trị</th></tr>
              </thead>
              <tbody>
                {transactions.map(t => (
                  <tr key={t.id}>
                    <td>{t.id.toUpperCase()}</td>
                    <td>{t.listingTitle}</td>
                    <td>{t.status}</td>
                    <td>{(t.quantity * t.agreedPrice).toLocaleString()}đ</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </Layout>
  )
}
