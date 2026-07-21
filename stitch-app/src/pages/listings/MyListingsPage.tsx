import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { Plus, Edit, Eye, EyeOff, Trash2 } from 'lucide-react'

export default function MyListingsPage() {
  const store = useStore()
  const user = store.currentUser
  const [listings, setListings] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    if (!user) return
    const fetchListings = async () => {
      setLoading(true)
      try {
        const data = await store.getListings()
        setListings(data.filter((l: any) => l.sellerId === user.id))
      } catch (err) {
        console.error('Failed to fetch listings:', err)
      } finally {
        setLoading(false)
      }
    }
    fetchListings()
  }, [user])

  if (!user) return null

  const handleDelete = async (id: string) => {
    if (!confirm('Bạn có chắc chắn muốn xóa vật liệu này?')) return
    try {
      // await store.deleteListing(id) // TODO: Add deleteListing to store
      setListings(prev => prev.filter(l => l.id !== id))
    } catch (err) {
      console.error('Failed to delete listing:', err)
    }
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản lý</p>
          <h1>Nguồn Cung Của Tôi</h1>
        </div>
        <Link to="/listings/new" className="btn btn-primary"><Plus size={16} /> Đăng vật liệu mới</Link>
      </div>

      {loading ? (
        <div className="loading-state">
          <div className="spinner" />
          <p>Đang tải dữ liệu...</p>
        </div>
      ) : listings.length === 0 ? (
        <div className="empty-state">
          <span className="empty-icon">📦</span>
          <h3>Chưa có nguồn cung nào</h3>
          <p>Bắt đầu bằng cách đăng vật liệu đầu tiên của bạn.</p>
          <Link to="/listings/new" className="btn btn-primary">Đăng vật liệu</Link>
        </div>
      ) : (
        <div className="table-responsive">
          <table className="table">
            <thead>
              <tr>
                <th>Vật liệu</th>
                <th>Danh mục</th>
                <th>Số lượng</th>
                <th>Giá</th>
                <th>Địa điểm</th>
                <th>Trạng thái</th>
                <th>Thao tác</th>
              </tr>
            </thead>
            <tbody>
              {listings.map(l => (
                <tr key={l.id}>
                  <td><Link to={`/material/${l.id}`} className="link-blue">{l.title}</Link></td>
                  <td>{l.categoryId}</td>
                  <td>{l.quantity} {l.unit}</td>
                  <td>{(l.pricePerUnit || 0).toLocaleString()}đ</td>
                  <td>{l.location}</td>
                  <td><StatusBadge status={l.status} /></td>
                  <td>
                    <div className="action-btns">
                      <Link to={`/listings/edit/${l.id}`} className="icon-btn-sm" title="Chỉnh sửa"><Edit size={16} /></Link>
                      <button className="icon-btn-sm" title={l.status === 'hidden' ? 'Hiện' : 'Ẩn'}>
                        {l.status === 'hidden' ? <Eye size={16} /> : <EyeOff size={16} />}
                      </button>
                      <button className="icon-btn-sm" title="Xóa" onClick={() => handleDelete(l.id)}>
                        <Trash2 size={16} />
                      </button>
                    </div>
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


