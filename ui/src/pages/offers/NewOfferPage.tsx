import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { CheckCircle } from 'lucide-react'

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
  const [buyerCompany, setBuyerCompany] = useState<any>(null)
  const [showConfirm, setShowConfirm] = useState(false)
  const [showSuccess, setShowSuccess] = useState(false)
  const [billData, setBillData] = useState<any>(null)

  useEffect(() => {
    async function load() {
      if (!listingId) return
      setLoading(true)
      const data = await store.getListing(listingId)
      if (data) {
        setListing(data)
        setQuantity(data.minOrderQuantity)
        setPrice(data.pricePerUnit)
        if (data.sellerName) setSellerName(data.sellerName)
      }
      if (user?.companyId) {
        const company = await store.getCompany(user.companyId)
        setBuyerCompany(company)
      }
      setLoading(false)
    }
    load()
  }, [listingId])

  if (loading) return <Layout showSidebar><div className="empty-state"><h3>Đang tải...</h3></div></Layout>
  if (!listing || !user) return <Layout showSidebar><div className="empty-state"><h3>Không tìm thấy</h3></div></Layout>
  if (!user.companyId || !buyerCompany || buyerCompany.status !== 'verified') {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <span className="empty-icon">🏢</span>
          <h3>Doanh nghiệp chưa được duyệt</h3>
          <p>Bạn cần có hồ sơ doanh nghiệp đã được admin duyệt mới có thể mua hàng.</p>
          <button className="btn btn-primary" onClick={() => nav('/company')}>Xem hồ sơ doanh nghiệp</button>
        </div>
      </Layout>
    )
  }

  const totalAmount = quantity * price

  const handleConfirmPayment = async () => {
    setShowConfirm(false)
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
      setBillData({
        id: result.id,
        listingTitle: listing.title,
        sellerName,
        quantity,
        unit: listing.unit,
        price,
        totalAmount,
        buyerName: user.name,
        date: new Date().toLocaleString('vi-VN'),
      })
      setShowSuccess(true)
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
        <form onSubmit={e => { e.preventDefault(); setShowConfirm(true) }} className="form-card">
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
            <button type="submit" className="btn btn-primary" disabled={submitting}>{submitting ? 'Đang xử lý...' : 'Thanh Toán & Gửi Đề Nghị'}</button>
          </div>
        </form>

        <div className="panel">
          <h2>Tóm tắt</h2>
          <div className="info-list">
            <div className="info-row"><span>Vật liệu</span><strong>{listing.title}</strong></div>
            <div className="info-row"><span>Người bán</span><strong>{sellerName}</strong></div>
            <div className="info-row"><span>Số lượng</span><strong>{quantity} {listing.unit}</strong></div>
            <div className="info-row"><span>Đơn giá</span><strong>{(price || 0).toLocaleString()}đ/{listing.unit}</strong></div>
            <div className="info-row total"><span>Tổng cộng</span><strong>{totalAmount.toLocaleString()}đ</strong></div>
          </div>
        </div>
      </div>

      {showConfirm && (
        <div className="modal-overlay" onClick={() => setShowConfirm(false)}>
          <div className="modal" onClick={e => e.stopPropagation()} style={{ maxWidth: 420 }}>
            <h2 style={{ marginBottom: 16 }}>Xác nhận thanh toán</h2>
            <div className="info-list">
              <div className="info-row"><span>Vật liệu</span><strong>{listing.title}</strong></div>
              <div className="info-row"><span>Người bán</span><strong>{sellerName}</strong></div>
              <div className="info-row"><span>Số lượng</span><strong>{quantity} {listing.unit}</strong></div>
              <div className="info-row"><span>Đơn giá</span><strong>{price.toLocaleString()}đ/{listing.unit}</strong></div>
              <div className="info-row total">
                <span>Tổng thanh toán</span>
                <strong style={{ color: 'var(--primary)' }}>{totalAmount.toLocaleString()}đ</strong>
              </div>
            </div>
            <div style={{ padding: 12, background: '#eff6ff', borderRadius: 8, margin: '16px 0', fontSize: 13, color: '#1e40af' }}>
              Số tiền sẽ được giữ tạm trong hệ thống cho đến khi giao dịch hoàn tất.
            </div>
            <div className="form-actions">
              <button className="btn btn-ghost" onClick={() => setShowConfirm(false)}>Hủy</button>
              <button className="btn btn-primary" onClick={handleConfirmPayment} disabled={submitting}>
                {submitting ? 'Đang xử lý...' : `Thanh Toán ${totalAmount.toLocaleString()}đ`}
              </button>
            </div>
          </div>
        </div>
      )}

      {showSuccess && billData && (
        <div className="modal-overlay">
          <div className="modal" onClick={e => e.stopPropagation()} style={{ maxWidth: 480 }}>
            <div style={{ textAlign: 'center', marginBottom: 20 }}>
              <CheckCircle size={48} style={{ color: 'var(--primary)' }} />
              <h2 style={{ marginTop: 12, color: 'var(--primary)' }}>Bill Succeed</h2>
            </div>
            <div style={{ background: 'var(--surface-low)', borderRadius: 8, padding: 20, marginBottom: 16 }}>
              <div className="info-list">
                <div className="info-row"><span>Mã giao dịch</span><strong style={{ fontSize: 12 }}>{billData.id.slice(0, 8).toUpperCase()}</strong></div>
                <div className="info-row"><span>Thời gian</span><strong>{billData.date}</strong></div>
                <div className="info-row"><span>Người mua</span><strong>{billData.buyerName}</strong></div>
                <div className="info-row"><span>Người bán</span><strong>{billData.sellerName}</strong></div>
                <div className="info-row"><span>Vật liệu</span><strong>{billData.listingTitle}</strong></div>
                <div className="info-row"><span>Số lượng</span><strong>{billData.quantity} {billData.unit}</strong></div>
                <div className="info-row"><span>Đơn giá</span><strong>{billData.price.toLocaleString()}đ/{billData.unit}</strong></div>
              </div>
              <div style={{ borderTop: '2px solid var(--outline)', marginTop: 12, paddingTop: 12 }}>
                <div className="info-row total">
                  <span>Tổng thanh toán</span>
                  <strong style={{ fontSize: 20, color: 'var(--primary)' }}>{billData.totalAmount.toLocaleString()}đ</strong>
                </div>
              </div>
            </div>
            <div style={{ padding: 12, background: '#f0fdf4', borderRadius: 8, marginBottom: 16, fontSize: 13, color: '#166534' }}>
              Tiền đang được giữ tạm trong hệ thống. Sẽ chuyển cho người bán khi giao dịch hoàn tất.
            </div>
            <button className="btn btn-primary btn-full" onClick={() => { setShowSuccess(false); nav('/offers/sent') }}>
              Đã hiểu, xem đề nghị đã gửi
            </button>
          </div>
        </div>
      )}
    </Layout>
  )
}


