import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { Loader2 } from 'lucide-react'

export default function AdminListingsPage() {
  const store = useStore()
  const [listings, setListings] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const loadListings = async () => {
      setLoading(true)
      const data = await store.getListings()
      setListings(data)
      setLoading(false)
    }
    loadListings()
  }, [])

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
          <p className="eyebrow">Quản trị</p>
          <h1>Quản Lý Nguồn Cung Vật Liệu</h1>
          <p className="muted">Tất cả nguồn cung vật liệu trên hệ thống</p>
        </div>
      </div>

      <div className="table-responsive">
        <table className="table">
          <thead>
            <tr><th>Vật liệu</th><th>Danh mục</th><th>Doanh nghiệp</th><th>Số lượng</th><th>Giá</th><th>Địa điểm</th><th>Trạng thái</th><th>Thao tác</th></tr>
          </thead>
          <tbody>
            {listings.map(l => (
              <tr key={l.id}>
                <td>
                  <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                    {l.imageUrl && <img src={l.imageUrl} alt="" style={{ width: 40, height: 40, objectFit: 'cover', borderRadius: 6 }} />}
                    <Link to={`/material/${l.id}`} className="link-blue">{l.title}</Link>
                  </div>
                </td>
                <td>{l.categoryName || l.categoryId}</td>
                <td>{l.companyName || l.companyId}</td>
                <td>{l.quantity} {l.unit}</td>
                <td>{l.pricePerUnit?.toLocaleString()}đ</td>
                <td>{l.location}</td>
                <td><StatusBadge status={l.status} /></td>
                <td>
                  <div className="action-btns">
                    <button className="btn btn-sm btn-outline">Duyệt</button>
                    <button className="btn btn-sm btn-ghost">Ẩn</button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Layout>
  )
}
