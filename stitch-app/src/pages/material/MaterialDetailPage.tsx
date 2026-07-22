import React, { useState, useEffect } from 'react'
import { useParams, Link, useNavigate } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StarRating } from '../../components/UI'
import { MapPin, Package, Shield, Truck, ChevronRight } from 'lucide-react'

export default function MaterialDetailPage() {
  const { id } = useParams<{ id: string }>()
  const store = useStore()
  const nav = useNavigate()
  const [listing, setListing] = useState<any>(null)
  const [company, setCompany] = useState<any>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      if (!id) return
      setLoading(true)
      try {
        const listingData = await store.getListing(id)
        setListing(listingData)
        if (listingData?.companyId) {
          const companyData = await store.getCompany(listingData.companyId)
          setCompany(companyData)
        }
      } catch (err) {
        console.error('Failed to fetch listing:', err)
      } finally {
        setLoading(false)
      }
    }
    fetchData()
  }, [id])

  if (loading) {
    return (
      <Layout>
        <div className="loading-state">
          <div className="spinner" />
          <p>Đang tải dữ liệu...</p>
        </div>
      </Layout>
    )
  }

  if (!listing) {
    return <Layout><div className="empty-state"><h3>Không tìm thấy vật liệu</h3></div></Layout>
  }

  return (
    <Layout showSidebar={!!store.currentUser}>
      <nav className="breadcrumbs">
        <Link to="/marketplace">Chợ</Link>
        <ChevronRight size={14} />
        <span>{listing.title}</span>
      </nav>

      <div className="detail-layout">
        <div className="detail-main">
          <div className="detail-tags">
            <span className="tag tag-green">Đã xác minh</span>
            <span className="tag tag-outline">{listing.categoryId}</span>
          </div>

          <h1>{listing.title}</h1>
          <p className="detail-desc">{listing.description}</p>

          <div className="detail-gallery">
            {listing.imageUrl ? (
              <img src={listing.imageUrl} alt={listing.title} style={{ width: '100%', maxHeight: 400, objectFit: 'cover', borderRadius: 12 }} />
            ) : (
              <div className="gallery-placeholder">
                <span style={{ fontSize: 48 }}>📦</span>
                <p>Ảnh vật liệu</p>
              </div>
            )}
          </div>

          <div className="detail-tabs">
            <button className="tab active">Thông số kỹ thuật</button>
            <button className="tab">Tác động môi trường</button>
            <button className="tab">Thông tin giao nhận</button>
          </div>

          <div className="specs-grid">
            {Object.entries(listing.specs || {}).map(([key, val]) => (
              <div key={key} className="spec-row">
                <span className="spec-label">{key}</span>
                <span className="spec-value">{val}</span>
              </div>
            ))}
          </div>

          <div className="detail-section">
            <h3>Mô tả chi tiết</h3>
            <p>{listing.description}</p>
          </div>

          <div className="detail-section">
            <h3>🌿 Tác động Môi trường (Ước tính)</h3>
            <div className="eco-stats">
              <div className="eco-stat">
                <span className="eco-value">-68%</span>
                <span className="eco-label">Carbon so với nguyên sinh</span>
              </div>
              <div className="eco-stat">
                <span className="eco-value">-45%</span>
                <span className="eco-label">Sử dụng nước</span>
              </div>
              <div className="eco-stat">
                <span className="eco-value">100%</span>
                <span className="eco-label">Năng lượng tái tạo</span>
              </div>
            </div>
          </div>
        </div>

        <div className="detail-sidebar">
          <div className="panel sticky">
            <div className="price-section">
              <span className="price-label">Ước tính giá</span>
              <span className="price-value">{(listing.pricePerUnit || 0).toLocaleString()}đ <span className="price-unit">/ {listing.unit}</span></span>
            </div>

            <div className="detail-info">
              <div className="info-row"><span>Số lượng có sẵn</span><strong>{listing.quantity} {listing.unit}</strong></div>
              <div className="info-row"><span>Đơn hàng tối thiểu</span><strong>{listing.minOrderQuantity} {listing.unit}</strong></div>
              <div className="info-row"><span>Đóng gói</span><strong>{listing.packaging}</strong></div>
              <div className="info-row"><span>Địa điểm</span><strong><MapPin size={14} /> {listing.location}</strong></div>
            </div>

            {store.currentUser && store.currentUser.id !== listing.sellerId && (
              <button className="btn btn-primary btn-full" onClick={() => nav(`/offers/new/${listing.id}`)}>
                Gửi Đề Nghị Mua
              </button>
            )}
            {store.currentUser && store.currentUser.id === listing.sellerId && (
              <div className="alert alert-info" style={{ textAlign: 'center', padding: 12 }}>
                Đây là sản phẩm của bạn
              </div>
            )}
            {!store.currentUser && (
              <Link to="/login" className="btn btn-primary btn-full" style={{ textAlign: 'center' }}>Đăng nhập để mua</Link>
            )}

            <p className="disclaimer">Giá chưa bao gồm phí vận chuyển & thuế. Chốt lại khi ký hợp đồng.</p>
          </div>

          {company && (
            <div className="panel">
              <h3 className="panel-title">Hồ sơ nhà cung cấp</h3>
              <div className="seller-card">
                <div className="seller-avatar"><Package size={24} /></div>
                <div>
                  <strong>{company.name}</strong>
                  <StarRating rating={company.rating} />
                </div>
              </div>
              <div className="seller-info">
                <span><Shield size={14} /> Thành viên từ {company.memberSince}</span>
                <span><Truck size={14} /> Thời gian điều phối: 3-5 ngày</span>
              </div>
              <div className="cert-tags">
                {company.certifications.map(c => <span key={c} className="tag tag-outline">{c}</span>)}
              </div>
            </div>
          )}
        </div>
      </div>
    </Layout>
  )
}


