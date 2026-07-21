import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'

export default function SentOffersPage() {
  const store = useStore()
  const [offers, setOffers] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    async function load() {
      setLoading(true)
      const data = await store.getOffers({ role: 'buyer' })
      setOffers(data)
      setLoading(false)
    }
    load()
  }, [])

  const handleCancel = async (id: string) => {
    if (!confirm('Bạn có chắc muốn hủy đề nghị này?')) return
    // Note: store doesn't have cancelOffer, but we can use rejectOffer as placeholder
    // In real implementation, you'd add a cancelOffer method to the store
    alert('Chức năng hủy đề nghị chưa được hỗ trợ')
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Đề nghị mua</p>
          <h1>Đề Nghị Đã Gửi</h1>
          <p className="muted">Các đề nghị mua bạn đã gửi cho người bán</p>
        </div>
      </div>

      {loading ? (
        <div className="empty-state"><h3>Đang tải...</h3></div>
      ) : offers.length === 0 ? (
        <div className="empty-state">
          <span className="empty-icon">📄</span>
          <h3>Chưa gửi đề nghị nào</h3>
          <p>Tìm vật liệu trên marketplace và gửi đề nghị mua.</p>
          <Link to="/marketplace" className="btn btn-primary">Đến Marketplace</Link>
        </div>
      ) : (
        <div className="table-responsive">
          <table className="table">
            <thead>
              <tr><th>Vật liệu</th><th>Người bán</th><th>Số lượng</th><th>Giá đề xuất</th><th>Ngày gửi</th><th>Trạng thái</th><th>Thao tác</th></tr>
            </thead>
            <tbody>
              {offers.map(o => (
                <tr key={o.id}>
                  <td><Link to={`/material/${o.listingId}`} className="link-blue">{o.listingTitle}</Link></td>
                  <td>{o.sellerName}</td>
                  <td>{o.quantity} {o.unit}</td>
                  <td>{(o.proposedPrice || 0).toLocaleString()}đ/{o.unit}</td>
                  <td>{o.createdAt}</td>
                  <td><StatusBadge status={o.status} /></td>
                  <td>
                    {o.status === 'pending' && (
                      <button className="btn btn-sm btn-ghost" onClick={() => handleCancel(o.id)}>Hủy</button>
                    )}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      )}
    </Layout>
  )
}


