import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StarRating } from '../../components/UI'
import { Loader2 } from 'lucide-react'

export default function ReviewsPage() {
  const store = useStore()
  const user = store.currentUser
  const [reviews, setReviews] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    if (!user) return
    const loadReviews = async () => {
      setLoading(true)
      const data = await store.getReviews()
      setReviews(data)
      setLoading(false)
    }
    loadReviews()
  }, [user])

  if (!user) return null

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
          <p className="eyebrow">Đánh giá</p>
          <h1>Đánh Giá Của Tôi</h1>
        </div>
      </div>

      {reviews.length === 0 ? (
        <div className="empty-state">
          <span className="empty-icon">⭐</span>
          <h3>Chưa có đánh giá nào</h3>
          <p>Đánh giá sẽ xuất hiện sau khi bạn hoàn tất giao dịch.</p>
        </div>
      ) : (
        <div className="reviews-list">
          {reviews.map(r => (
            <div key={r.id} className="review-card">
              <div className="review-header">
                <div>
                  <strong>{r.reviewerId === user.id ? `Đánh giá ${r.revieweeName}` : `Đánh giá từ ${r.reviewerName}`}</strong>
                  <StarRating rating={r.rating} />
                </div>
                <span className="muted">{r.createdAt}</span>
              </div>
              <p>{r.comment}</p>
            </div>
          ))}
        </div>
      )}
    </Layout>
  )
}
