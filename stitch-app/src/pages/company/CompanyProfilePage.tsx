import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge, StarRating } from '../../components/UI'
import { MapPin, Phone, Mail, Shield, Calendar, Loader2 } from 'lucide-react'

export default function CompanyProfilePage() {
  const store = useStore()
  const user = store.currentUser
  const [company, setCompany] = useState<any>(null)
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [formData, setFormData] = useState({ name: '', tax_code: '', address: '', city: '', description: '' })
  const [submitting, setSubmitting] = useState(false)

  useEffect(() => {
    if (!user?.companyId) {
      setLoading(false)
      return
    }
    const loadCompany = async () => {
      setLoading(true)
      const data = await store.getCompany(user.companyId!)
      setCompany(data)
      setLoading(false)
    }
    loadCompany()
  }, [user])

  const handleCreate = async () => {
    if (!formData.name) { alert('Vui long nhap ten doanh nghiep'); return }
    setSubmitting(true)
    const result = await store.createCompany(formData)
    if (result) {
      alert('Tao ho do thanh cong! Cho admin duyet.')
      // Update user's companyId in store
      const updatedUser = { ...user!, companyId: result.id }
      localStorage.setItem('user', JSON.stringify(updatedUser))
      window.location.reload()
    } else {
      alert('Loi khi tao ho do!')
    }
    setSubmitting(false)
  }

  if (!user) return null

  if (loading) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <Loader2 className="animate-spin" size={32} />
          <p>Dang tai du lieu...</p>
        </div>
      </Layout>
    )
  }

  if (!company && !showForm) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <span className="empty-icon">🏢</span>
          <h3>Chua co ho do doanh nghiep</h3>
          <p>Ban can tao ho do doanh nghiep de dang vat lieu va thuc hien giao dich.</p>
          <button className="btn btn-primary" onClick={() => setShowForm(true)}>Tao ho do doanh nghiep</button>
        </div>
      </Layout>
    )
  }

  if (showForm && !company) {
    return (
      <Layout showSidebar>
        <div className="page-header">
          <div>
            <p className="eyebrow">Dang ky doanh nghiep</p>
            <h1>Tao Ho Do Doanh Nghiep</h1>
          </div>
        </div>
        <div className="panel" style={{ maxWidth: 600 }}>
          <div className="form-group">
            <label>Ten doanh nghiep *</label>
            <input value={formData.name} onChange={e => setFormData({...formData, name: e.target.value})} placeholder="VD: EcoPoly Solutions" />
          </div>
          <div className="form-group">
            <label>Ma so thue</label>
            <input value={formData.tax_code} onChange={e => setFormData({...formData, tax_code: e.target.value})} placeholder="VD: 0123456789" />
          </div>
          <div className="form-group">
            <label>Dia chi</label>
            <input value={formData.address} onChange={e => setFormData({...formData, address: e.target.value})} placeholder="VD: 123 Le Loi, Q.1" />
          </div>
          <div className="form-group">
            <label>Thanh pho</label>
            <input value={formData.city} onChange={e => setFormData({...formData, city: e.target.value})} placeholder="VD: TP.HCM" />
          </div>
          <div className="form-group">
            <label>Mo ta</label>
            <textarea rows={4} value={formData.description} onChange={e => setFormData({...formData, description: e.target.value})} placeholder="Mo ta ve hoat dong kinh doanh..." />
          </div>
          <div className="form-actions">
            <button className="btn btn-ghost" onClick={() => setShowForm(false)}>Huy</button>
            <button className="btn btn-primary" onClick={handleCreate} disabled={submitting}>
              {submitting ? 'Dang tao...' : 'Tao ho do'}
            </button>
          </div>
        </div>
      </Layout>
    )
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Hồ sơ doanh nghiệp</p>
          <h1>{company.name}</h1>
        </div>
        <div className="page-actions">
          <StatusBadge status={company.status} />
          <button className="btn btn-outline">Chỉnh sửa</button>
        </div>
      </div>

      {company.status === 'rejected' && company.rejectReason && (
        <div className="alert alert-error">Lý do từ chối: {company.rejectReason}</div>
      )}

      <div className="grid-2">
        <div className="panel">
          <h2>Thông tin doanh nghiệp</h2>
          <div className="info-list">
            <div className="info-row"><span>Mã số thuế</span><strong>{company.taxCode}</strong></div>
            <div className="info-row"><span>Địa chỉ</span><strong><MapPin size={14} /> {company.address}, {company.city}</strong></div>
            <div className="info-row"><span>Email</span><strong><Mail size={14} /> {user.email}</strong></div>
            <div className="info-row"><span>Điện thoại</span><strong><Phone size={14} /> {user.phone}</strong></div>
            <div className="info-row"><span>Ngày tham gia</span><strong><Calendar size={14} /> {company.memberSince}</strong></div>
          </div>
          <div className="info-desc">
            <h3>Mô tả</h3>
            <p>{company.description}</p>
          </div>
        </div>

        <div className="panel">
          <h2>Đánh giá & Uy tín</h2>
          <div className="rating-display">
            <span className="rating-big">{company.rating?.toFixed(1) || '0.0'}</span>
            <StarRating rating={company.rating || 0} size={20} />
            <span className="muted">{company.reviewCount || 0} đánh giá</span>
          </div>
          <div className="cert-section">
            <h3>Chứng nhận</h3>
            {company.certifications?.length > 0 ? (
              <div className="cert-tags">
                {company.certifications.map((c: string) => (
                  <span key={c} className="cert-tag"><Shield size={14} /> {c}</span>
                ))}
              </div>
            ) : (
              <p className="muted">Chưa có chứng nhận</p>
            )}
          </div>
        </div>
      </div>
    </Layout>
  )
}
