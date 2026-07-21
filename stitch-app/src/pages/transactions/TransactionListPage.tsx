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

  useEffect(() => {
    async function load() {
      setLoading(true)
      const data = await store.getTransactions()
      setTransactions(data)
      setLoading(false)
    }
    load()
  }, [])

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Giao dịch</p>
          <h1>Danh Sách Giao Dịch</h1>
          <p className="muted">Tất cả giao dịch mua và bán của bạn</p>
        </div>
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
              <tr><th>Mã GD</th><th>Vật liệu</th><th>Vai trò</th><th>Đối tác</th><th>Số lượng</th><th>Giá thỏa thuận</th><th>Trạng thái</th></tr>
            </thead>
            <tbody>
              {transactions.map(t => {
                const isBuyer = user?.id === t.buyerId
                return (
                  <tr key={t.id}>
                    <td><Link to={`/transactions/${t.id}`} className="link-blue">{t.id.toUpperCase()}</Link></td>
                    <td>{t.listingTitle}</td>
                    <td><span className={`tag ${isBuyer ? 'tag-outline' : 'tag-green'}`}>{isBuyer ? 'Mua' : 'Bán'}</span></td>
                    <td>{isBuyer ? t.sellerName : t.buyerName}</td>
                    <td>{t.quantity} {t.unit}</td>
                    <td>{(t.agreedPrice || 0).toLocaleString()}đ/{t.unit}</td>
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


