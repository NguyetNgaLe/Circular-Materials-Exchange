import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { MapPin, Search, SlidersHorizontal } from 'lucide-react'

export default function MarketplacePage() {
  const store = useStore()
  const [search, setSearch] = useState('')
  const [catFilter, setCatFilter] = useState<string[]>([])
  const [locationFilter, setLocationFilter] = useState('')
  const [sortBy, setSortBy] = useState('newest')
  const [listings, setListings] = useState<any[]>([])
  const [categories, setCategories] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true)
      try {
        const [listingsData, categoriesData] = await Promise.all([
          store.getListings(),
          store.getCategories()
        ])
        setListings(listingsData.filter((l: any) => l.status === 'active'))
        setCategories(categoriesData)
      } catch (err) {
        console.error('Failed to fetch data:', err)
      } finally {
        setLoading(false)
      }
    }
    fetchData()
  }, [])

  const filtered = listings.filter(l => {
    if (search && !l.title.toLowerCase().includes(search.toLowerCase())) return false
    if (catFilter.length > 0 && !catFilter.includes(l.categoryId)) return false
    if (locationFilter && l.location !== locationFilter) return false
    return true
  })

  const sorted = [...filtered].sort((a, b) => {
    if (sortBy === 'price_asc') return a.pricePerUnit - b.pricePerUnit
    if (sortBy === 'price_desc') return b.pricePerUnit - a.pricePerUnit
    if (sortBy === 'quantity') return b.quantity - a.quantity
    return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
  })

  const locations = [...new Set(listings.map(l => l.location))]

  const toggleCat = (id: string) => {
    setCatFilter(f => f.includes(id) ? f.filter(c => c !== id) : [...f, id])
  }

  return (
    <Layout showSidebar={!!store.currentUser}>
      <div className="page-header">
        <div>
          <p className="eyebrow">Marketplace</p>
          <h1>Chợ Giao Dịch Vật Liệu</h1>
          <p className="muted">Tìm kiếm vật liệu tái chế từ các doanh nghiệp đã xác minh</p>
        </div>
      </div>

      {loading ? (
        <div className="loading-state">
          <div className="spinner" />
          <p>Đang tải dữ liệu...</p>
        </div>
      ) : (
        <div className="marketplace-layout">
          <aside className="filter-sidebar">
            <h3><SlidersHorizontal size={16} /> Bộ Lọc</h3>
            <div className="filter-section">
              <h4>Danh Mục</h4>
              {categories.map(c => (
                <label key={c.id} className="checkbox-label">
                  <input type="checkbox" checked={catFilter.includes(c.id)} onChange={() => toggleCat(c.id)} />
                  {c.name}
                </label>
              ))}
            </div>
            <div className="filter-section">
              <h4>Địa Điểm</h4>
              <select value={locationFilter} onChange={e => setLocationFilter(e.target.value)}>
                <option value="">Tất cả</option>
                {locations.map(l => <option key={l} value={l}>{l}</option>)}
              </select>
            </div>
            <button className="btn btn-primary btn-full" onClick={() => {}}>Áp Dụng Bộ Lọc</button>
          </aside>

          <div className="marketplace-content">
            <div className="marketplace-toolbar">
              <div className="search-input">
                <Search size={16} />
                <input placeholder="Tìm kiếm vật liệu..." value={search} onChange={e => setSearch(e.target.value)} />
              </div>
              <div className="sort-control">
                <span>Sắp xếp:</span>
                <select value={sortBy} onChange={e => setSortBy(e.target.value)}>
                  <option value="newest">Mới nhất</option>
                  <option value="price_asc">Giá thấp → cao</option>
                  <option value="price_desc">Giá cao → thấp</option>
                  <option value="quantity">Số lượng nhiều</option>
                </select>
              </div>
            </div>

            <div className="material-grid">
              {sorted.map(l => {
                const cat = categories.find(c => c.id === l.categoryId)
                const isOwner = store.currentUser && store.currentUser.id === l.sellerId
                return (
                  <Link key={l.id} to={`/material/${l.id}`} className="material-card" style={isOwner ? { opacity: 0.7 } : {}}>
                    <div className="card-image">
                      {l.imageUrl ? (
                        <img src={l.imageUrl} alt={l.title} style={{ width: '100%', height: '100%', objectFit: 'cover' }} />
                      ) : (
                        <div className="card-image-placeholder">
                          <span style={{ fontSize: 32 }}>📦</span>
                        </div>
                      )}
                      {isOwner ? (
                        <span className="card-badge" style={{ background: '#9ca3af', color: '#fff' }}>Sản phẩm của bạn</span>
                      ) : (
                        <span className="card-badge">Đã Xác Minh</span>
                      )}
                    </div>
                    <div className="card-body">
                      <h3>{l.title}</h3>
                      <div className="card-location">
                        <MapPin size={14} /> {l.location}
                      </div>
                      <div className="card-footer">
                        <div>
                          <span className="card-qty-label">Số Lượng</span>
                          <span className="card-qty">{l.quantity} {l.unit}</span>
                        </div>
                        <div className="card-price">
                          {(l.pricePerUnit || 0).toLocaleString()}đ <span className="card-unit">/ {l.unit}</span>
                        </div>
                      </div>
                    </div>
                  </Link>
                )
              })}
            </div>

            {sorted.length === 0 && (
              <div className="empty-state">
                <span className="empty-icon">🔍</span>
                <h3>Không tìm thấy vật liệu</h3>
                <p>Thử thay đổi bộ lọc hoặc từ khóa tìm kiếm</p>
              </div>
            )}
          </div>
        </div>
      )}
    </Layout>
  )
}
