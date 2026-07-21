import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'

export default function NewOfferPage() {
  const { listingId } = useParams<{ listingId: string }>()
  const store = useStore()
  const nav = useNavigate()
  const user = store.currentUser

  const [listing, setListing] = useState<any>(null)
  const [sellerName, setSellerName] = useState('')
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)
  const [quantity, setQuantity] = useState(0)
  const [price, setPrice] = useState(0)
  const [message, setMessage] = useState('')

  useEffect(() => {
    async function load() {
      if (!listingId) return
      setLoading(true)
      const data = await store.getListing(listingId)
      if (data) {
        setListing(data)
        setQuantity(data.minOrderQuantity)
        setPrice(data.pricePerUnit)
        const company = await store.getCompany(data.companyId)
        if (company) setSellerName(company.name)
      }
      setLoading(false)
    }
    load()
  }, [listingId])

  if (loading) return <Layout showSidebar><div className="empty-state"><h3>Đang tải...</h3></div></Layout>
  if (!listing || !user) return <Layout showSidebar><div className="empty-state"><h3>Không tìm thấy</h3></div></Layout>

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitting(true)
    const result = await store.createOffer({
      type: 'buyer_to_seller',
      listingId: listing.id,
      listingTitle: listing.title,
      buyerId: user.id,
      buyerName: user.name,
      sellerId: listing.sellerId,
      sellerName: sellerName,
      quantity,
      unit: listing.unit,
      proposedPrice: price,
      currency: 'VND',
      message,
    })
    setSubmitting(false)
    if (result) {
      alert('Đã gửi đề nghị mua thành công!')
      nav('/offers/sent')
    } else {
      alert('Có lỗi xảy ra, vui lòng thử lại.')
    }
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Gửi đề nghị</p>
          <h1>Gửi Đề Nghị Mua</h1>
        </div>
      </div>

      <div className="grid-2">
        <form onSubmit={handleSubmit} className="form-card">
          <div className="form-group">
            <label>Vật liệu</label>
            <input value={listing.title} disabled />
          </div>
          <div className="form-group">
            <label>Số lượng ({listing.unit}) *</label>
            <input required type="number" min={listing.minOrderQuantity} value={quantity} onChange={e => setQuantity(Number(e.target.value))} />
            <span className="form-hint">Tối thiểu: {listing.minOrderQuantity} {listing.unit}</span>
          </div>
          <div className="form-group">
            <label>Giá đề xuất (VND/{listing.unit}) *</label>
            <input required type="number" value={price} onChange={e => setPrice(Number(e.target.value))} />
            <span className="form-hint">Giá niêm yết: {(listing.pricePerUnit || 0).toLocaleString()}đ</span>
          </div>
          <div className="form-group">
            <label>Lời nhắn</label>
            <textarea rows={4} value={message} onChange={e => setMessage(e.target.value)} placeholder="Nhập lời nhắn cho người bán..." />
          </div>
          <div className="form-actions">
            <button type="button" className="btn btn-ghost" onClick={() => nav(-1)}>Hủy</button>
            <button type="submit" className="btn btn-primary" disabled={submitting}>{submitting ? 'Đang gửi...' : 'Gửi Đề Nghị'}</button>
          </div>
        </form>

        <div className="panel">
          <h2>Tóm tắt</h2>
          <div className="info-list">
            <div className="info-row"><span>Vật liệu</span><strong>{listing.title}</strong></div>
            <div className="info-row"><span>Người bán</span><strong>{sellerName}</strong></div>
            <div className="info-row"><span>Số lượng</span><strong>{quantity} {listing.unit}</strong></div>
            <div className="info-row"><span>Đơn giá</span><strong>{(price || 0).toLocaleString()}đ/{listing.unit}</strong></div>
            <div className="info-row total"><span>Tổng cộng</span><strong>{(quantity * price).toLocaleString()}đ</strong></div>
          </div>
        </div>
      </div>
    </Layout>
  )
}


