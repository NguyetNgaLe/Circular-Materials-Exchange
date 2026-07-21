import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { MapPin, Calendar } from 'lucide-react'

export default function DemandMarketplacePage() {
  const store = useStore()
  const [demands, setDemands] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchDemands = async () => {
      setLoading(true)
      try {
        const data = await store.getDemands()
        setDemands(data.filter((d: any) => d.status === 'open'))
      } catch (err) {
        console.error('Failed to fetch demands:', err)
      } finally {
        setLoading(false)
      }
    }
    fetchDemands()
  }, [])

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Sàn nhu cầu</p>
          <h1>Sàn Nhu Cầu Mua Vật Liệu</h1>
          <p className="muted">Xem nhu cầu mua từ các doanh nghiệp và gửi báo giá</p>
        </div>
      </div>

      {loading ? (
        <div className="loading-state">
          <div className="spinner" />
          <p>Đang tải dữ liệu...</p>
        </div>
      ) : (
        <div className="material-grid">
          {demands.map(d => (
            <div key={d.id} className="material-card">
              <div className="card-body">
                <h3>{d.title}</h3>
                <p className="muted" style={{ fontSize: 14, marginBottom: 8 }}>{d.description?.slice(0, 80)}...</p>
                <div className="card-location"><MapPin size={14} /> {d.location}</div>
                <div className="card-footer">
                  <div>
                    <span className="card-qty-label">Cần mua</span>
                    <span className="card-qty">{d.quantity} {d.unit}</span>
                  </div>
                  {d.targetPrice && (
                    <div className="card-price">{(d.targetPrice || 0).toLocaleString()}đ <span className="card-unit">/ {d.unit}</span></div>
                  )}
                </div>
                <div style={{ marginTop: 8, fontSize: 12, color: 'var(--on-muted)' }}>
                  <Calendar size={12} /> Hạn: {d.deadline}
                </div>
              </div>
            </div>
          ))}
          {demands.length === 0 && (
            <div className="empty-state">
              <span className="empty-icon">📋</span>
              <h3>Không có nhu cầu nào</h3>
              <p>Hiện tại chưa có nhu cầu mua vật liệu nào đang mở</p>
            </div>
          )}
        </div>
      )}
    </Layout>
  )
}


