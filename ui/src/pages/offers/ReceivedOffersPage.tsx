import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { Check, X } from 'lucide-react'

export default function ReceivedOffersPage() {
  const store = useStore()
  const [offers, setOffers] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  async function loadOffers() {
    setLoading(true)
    const data = await store.getOffers({ role: 'seller' })
    setOffers(data)
    setLoading(false)
  }

  useEffect(() => {
    loadOffers()
  }, [])

  const handleAccept = async (id: string) => {
    if (!confirm('Bạn có chắc muốn chấp nhận đề nghị này?')) return
    const result = await store.acceptOffer(id)
    if (result) {
      alert('Đã chấp nhận đề nghị!')
      loadOffers()
    } else {
      alert('Có lỗi xảy ra, vui lòng thử lại.')
    }
  }

  const handleReject = async (id: string) => {
    if (!confirm('Bạn có chắc muốn từ chối đề nghị này?')) return
    const result = await store.rejectOffer(id)
    if (result) {
      alert('Đã từ chối đề nghị.')
      loadOffers()
    } else {
      alert('Có lỗi xảy ra, vui lòng thử lại.')
    }
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Đề nghị nhận được</p>
          <h1>Đề Nghị Đã Nhận</h1>
          <p className="muted">Các đề nghị mua từ người mua cho vật liệu của bạn</p>
        </div>
      </div>

      {loading ? (
        <div className="empty-state"><h3>Đang tải...</h3></div>
      ) : offers.length === 0 ? (
        <div className="empty-state">
          <span className="empty-icon">📥</span>
          <h3>Chưa nhận đề nghị nào</h3>
          <p>Khi người mua gửi đề nghị cho vật liệu của bạn, chúng sẽ hiển thị ở đây.</p>
        </div>
      ) : (
        <div className="table-responsive">
          <table className="table">
            <thead>
              <tr><th>Vật liệu</th><th>Người mua</th><th>Số lượng</th><th>Giá đề xuất</th><th>Lời nhắn</th><th>Trạng thái</th><th>Thao tác</th></tr>
            </thead>
            <tbody>
              {offers.map(o => (
                <tr key={o.id}>
                  <td><Link to={`/material/${o.listingId}`} className="link-blue">{o.listingTitle}</Link></td>
                  <td>{o.buyerName}</td>
                  <td>{o.quantity} {o.unit}</td>
                  <td>{(o.proposedPrice || 0).toLocaleString()}đ/{o.unit}</td>
                  <td className="td-ellipsis">{o.message}</td>
                  <td><StatusBadge status={o.status} /></td>
                  <td>
                    {o.status === 'pending' && (
                      <div className="action-btns">
                        <button className="icon-btn-sm success" title="Chấp nhận" onClick={() => handleAccept(o.id)}><Check size={16} /></button>
                        <button className="icon-btn-sm danger" title="Từ chối" onClick={() => handleReject(o.id)}><X size={16} /></button>
                      </div>
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


