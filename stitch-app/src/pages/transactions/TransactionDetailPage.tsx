import React, { useState, useEffect } from 'react'
import { useParams, Link, useNavigate } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { CheckCircle, Clock } from 'lucide-react'

const statusLabels: Record<string, string> = {
  'offer.accepted': 'Chấp nhận đề nghị',
  'transaction.confirmed': 'Giao dịch xác nhận',
  'transaction.in_progress': 'Đang thực hiện',
  'transaction.buyer_confirmed': 'Bên mua xác nhận hoàn tất',
  'transaction.seller_confirmed': 'Bên bán xác nhận hoàn tất',
  'transaction.completed': 'Hoàn tất',
  'transaction.cancelled': 'Đã hủy',
}

export default function TransactionDetailPage() {
  const { id } = useParams<{ id: string }>()
  const store = useStore()
  const nav = useNavigate()
  const user = store.currentUser

  const [tx, setTx] = useState<any>(null)
  const [loading, setLoading] = useState(true)
  const [updating, setUpdating] = useState(false)

  useEffect(() => {
    async function load() {
      if (!id) return
      setLoading(true)
      const data = await store.getTransaction(id)
      setTx(data)
      setLoading(false)
    }
    load()
  }, [id])

  if (loading) return <Layout showSidebar><div className="empty-state"><h3>Đang tải...</h3></div></Layout>
  if (!tx || !user) return <Layout showSidebar><div className="empty-state"><h3>Không tìm thấy giao dịch</h3></div></Layout>

  const isBuyer = user.id === tx.buyerId
  const myRole = isBuyer ? 'Bên mua' : 'Bên bán'
  const partner = isBuyer ? tx.sellerName : tx.buyerName
  const canConfirmBuyer = isBuyer && tx.status === 'in_progress'
  const canConfirmSeller = !isBuyer && tx.status === 'in_progress'
  const canConfirmComplete = !isBuyer && tx.status === 'buyer_confirmed'

  const handleConfirm = async () => {
    setUpdating(true)
    let newStatus = ''
    let note = ''

    if (isBuyer && tx.status === 'in_progress') {
      newStatus = 'buyer_confirmed'
      note = 'Đã nhận hàng và xác nhận'
    } else if (!isBuyer) {
      if (tx.status === 'buyer_confirmed') {
        newStatus = 'completed'
        note = 'Xác nhận hoàn tất giao dịch'
      } else if (tx.status === 'in_progress') {
        newStatus = 'seller_confirmed'
        note = 'Xác nhận đã giao hàng'
      }
    }

    if (newStatus) {
      const result = await store.updateTransactionStatus(tx.id, newStatus, note)
      if (result) {
        alert('Cập nhật thành công!')
        const updated = await store.getTransaction(tx.id)
        setTx(updated)
      } else {
        alert('Có lỗi xảy ra, vui lòng thử lại.')
      }
    }
    setUpdating(false)
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Chi tiết giao dịch</p>
          <h1>{tx.id.toUpperCase()} — {tx.listingTitle}</h1>
        </div>
        <div className="page-actions">
          <span className={`tag ${isBuyer ? 'tag-outline' : 'tag-green'}`}>{myRole}</span>
          <StatusBadge status={tx.status} />
        </div>
      </div>

      <div className="grid-2">
        <div className="panel">
          <h2>Thông tin giao dịch</h2>
          <div className="info-list">
            <div className="info-row"><span>Vật liệu</span><strong>{tx.listingTitle}</strong></div>
            <div className="info-row"><span>Vai trò của bạn</span><strong>{myRole}</strong></div>
            <div className="info-row"><span>Đối tác</span><strong>{partner}</strong></div>
            <div className="info-row"><span>Số lượng</span><strong>{tx.quantity} {tx.unit}</strong></div>
            <div className="info-row"><span>Giá thỏa thuận</span><strong>{(tx.agreedPrice || 0).toLocaleString()}đ/{tx.unit}</strong></div>
            <div className="info-row total"><span>Tổng giá trị</span><strong>{(tx.quantity * tx.agreedPrice).toLocaleString()}đ</strong></div>
          </div>

          <div className="payment-section">
            <h3>Thanh toán (Demo)</h3>
            <div className="info-list">
              <div className="info-row"><span>Phương thức</span><strong>Thanh toán ngoài hệ thống</strong></div>
              <div className="info-row"><span>Trạng thái</span><strong>Ghi nhận demo</strong></div>
              <div className="info-row"><span>Ghi chú</span><strong>{tx.settlementNote}</strong></div>
            </div>
          </div>

          {canConfirmBuyer && (
            <button className="btn btn-primary btn-full" onClick={handleConfirm} disabled={updating}>
              {updating ? 'Đang xử lý...' : 'Xác Nhận Đã Nhận Hàng'}
            </button>
          )}
          {canConfirmSeller && (
            <button className="btn btn-primary btn-full" onClick={handleConfirm} disabled={updating}>
              {updating ? 'Đang xử lý...' : 'Xác Nhận Đã Giao Hàng'}
            </button>
          )}
          {canConfirmComplete && (
            <button className="btn btn-primary btn-full" onClick={handleConfirm} disabled={updating}>
              {updating ? 'Đang xử lý...' : 'Xác Nhận Hoàn Tất Giao Dịch'}
            </button>
          )}
        </div>

        <div className="panel">
          <h2>Timeline Giao Dịch</h2>
          <div className="timeline">
            {tx.events.map((ev: any, i: number) => (
              <div key={ev.id} className={`timeline-item ${i === tx.events.length - 1 ? 'active' : ''}`}>
                <div className="timeline-dot">
                  {i === tx.events.length - 1 ? <CheckCircle size={16} /> : <Clock size={16} />}
                </div>
                <div className="timeline-content">
                  <div className="timeline-header">
                    <strong>{statusLabels[ev.toStatus] || ev.toStatus}</strong>
                    <span className="timeline-time">{new Date(ev.createdAt).toLocaleString('vi-VN')}</span>
                  </div>
                  <p className="timeline-note">{ev.note}</p>
                  <span className="timeline-actor">{ev.actorName}</span>
                </div>
              </div>
            ))}
          </div>

          {tx.status === 'completed' && (
            <div style={{ marginTop: 16 }}>
              <Link to={`/reviews/new/${tx.id}`} className="btn btn-primary btn-full">Đánh Giá Đối Tác</Link>
            </div>
          )}
        </div>
      </div>
    </Layout>
  )
}


