import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { Loader2 } from 'lucide-react'

export default function NewReviewPage() {
  const { txId } = useParams<{ txId: string }>()
  const store = useStore()
  const nav = useNavigate()
  const user = store.currentUser
  const [tx, setTx] = useState<any>(null)
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)
  const [rating, setRating] = useState(5)
  const [comment, setComment] = useState('')

  useEffect(() => {
    if (!txId) return
    const loadTransaction = async () => {
      setLoading(true)
      const data = await store.getTransaction(txId)
      setTx(data)
      setLoading(false)
    }
    loadTransaction()
  }, [txId])

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

  if (!tx || !user) return <Layout showSidebar><div className="empty-state"><h3>Không tìm thấy giao dịch</h3></div></Layout>

  const revieweeName = user.id === tx.buyerId ? tx.sellerName : tx.buyerName
  const revieweeId = user.id === tx.buyerId ? tx.sellerId : tx.buyerId

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitting(true)
    const result = await store.createReview({
      transactionId: tx.id,
      revieweeId,
      rating,
      comment
    })
    setSubmitting(false)
    if (result) {
      alert('Đã gửi đánh giá thành công!')
      nav('/reviews')
    } else {
      alert('Có lỗi xảy ra, vui lòng thử lại.')
    }
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Đánh giá</p>
          <h1>Đánh Giá Đối Tác</h1>
          <p className="muted">Giao dịch: {tx.id.toUpperCase()} — {tx.listingTitle}</p>
        </div>
      </div>

      <div className="form-card" style={{ maxWidth: 600 }}>
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label>Đối tác</label>
            <input value={revieweeName} disabled />
          </div>
          <div className="form-group">
            <label>Đánh giá</label>
            <div className="rating-input">
              {[1, 2, 3, 4, 5].map(i => (
                <button key={i} type="button" className={`star-btn ${i <= rating ? 'active' : ''}`} onClick={() => setRating(i)}>
                  ★
                </button>
              ))}
            </div>
          </div>
          <div className="form-group">
            <label>Nhận xét</label>
            <textarea rows={4} value={comment} onChange={e => setComment(e.target.value)} placeholder="Chia sẻ trải nghiệm giao dịch..." />
          </div>
          <div className="form-actions">
            <button type="button" className="btn btn-ghost" onClick={() => nav(-1)}>Hủy</button>
            <button type="submit" className="btn btn-primary" disabled={submitting}>
              {submitting ? <><Loader2 className="animate-spin" size={16} /> Đang gửi...</> : 'Gửi Đánh Giá'}
            </button>
          </div>
        </form>
      </div>
    </Layout>
  )
}
