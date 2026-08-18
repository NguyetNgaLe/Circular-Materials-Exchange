import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'

export default function TransactionListPage() {
  const store = useStore()
  const user = store.currentUser
  const [transactions, setTransactions] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  const loadTransactions = async () => {
    setLoading(true)
    const data = await store.getTransactions()
    setTransactions(data)
    setLoading(false)
  }

  useEffect(() => { loadTransactions() }, [user?.id])

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Doanh nghiệp</p>
          <h1>Lịch sử giao dịch</h1>
          <p className="muted">Dữ liệu mua và bán thực tế của doanh nghiệp được cập nhật từ hệ thống</p>
        </div>
        <button type="button" className="btn btn-outline" onClick={loadTransactions} disabled={loading}>
          {loading ? 'Đang cập nhật...' : 'Cập nhật dữ liệu'}
        </button>
      </div>

      {loading ? (
        <div className="empty-state"><h3>Đang tải...</h3></div>
      ) : transactions.length === 0 ? (
        <div className="empty-state">
          <span className="empty-icon">🔄</span>
          <h3>Chưa có giao dịch nào</h3>
          <p>Giao dịch sẽ được tạo tự động khi bạn chấp nhận hoặc được chấp nhận đề nghị mua.</p>
        </div>
      ) : (
        <div className="table-responsive">
          <table className="table">
            <thead>
              <tr><th>Mã GD</th><th>Ngày tạo</th><th>Vật liệu</th><th>Vai trò</th><th>Đối tác</th><th>Số lượng</th><th>Giá thỏa thuận</th><th>Thành tiền</th><th>Trạng thái</th></tr>
            </thead>
            <tbody>
              {transactions.map(t => {
                const isBuyer = user?.id === t.buyerId
                return (
                  <tr key={t.id}>
                    <td><Link to={`/transactions/${t.id}`} className="link-blue">{t.id.toUpperCase()}</Link></td>
                    <td>{new Date(t.createdAt).toLocaleDateString('vi-VN')}</td>
                    <td>{t.listingTitle}</td>
                    <td><span className={`tag ${isBuyer ? 'tag-outline' : 'tag-green'}`}>{isBuyer ? 'Mua' : 'Bán'}</span></td>
                    <td>{isBuyer ? t.sellerName : t.buyerName}</td>
                    <td>{t.quantity} {t.unit}</td>
                    <td>{(t.agreedPrice || 0).toLocaleString()}đ/{t.unit}</td>
                    <td><strong>{((t.agreedPrice || 0) * (t.quantity || 0)).toLocaleString()}đ</strong></td>
                    <td><StatusBadge status={t.status} /></td>
                  </tr>
                )
              })}
            </tbody>
          </table>
        </div>
      )}
    </Layout>
  )
}


