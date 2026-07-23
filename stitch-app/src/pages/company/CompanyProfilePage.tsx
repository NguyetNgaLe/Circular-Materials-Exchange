import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { MapPin, Phone, Mail, Calendar, Loader2 } from 'lucide-react'

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
    if (!formData.name) { alert('Vui lòng nhap ten Doanh nghiệp'); return }
    setSubmitting(true)
    const result = await store.createCompany(formData)
    if (result) {
      alert('Tạo hồ sơ Thành công! Cho admin duyet.')
      // Update user's companyId in store
      const updatedUser = { ...user!, companyId: result.id }
      localStorage.setItem('user', JSON.stringify(updatedUser))
      window.location.reload()
    } else {
      alert('Loi khi Tạo hồ sơ!')
    }
    setSubmitting(false)
  }

  if (!user) return null

  if (loading) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <Loader2 className="animate-spin" size={32} />
          <p>Đang tải du lieu...</p>
        </div>
      </Layout>
    )
  }

  if (!company && !showForm) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <span className="empty-icon">🏢</span>
          <h3>Chưa có hồ sơ Doanh nghiệp</h3>
          <p>Ban can Tạo hồ sơ Doanh nghiệp de dang Vật liệu va thuc hien Giao dịch.</p>
          <button className="btn btn-primary" onClick={() => setShowForm(true)}>Tạo hồ sơ Doanh nghiệp</button>
        </div>
      </Layout>
    )
  }

  if (showForm && !company) {
    return (
      <Layout showSidebar>
        <div className="page-header">
          <div>
            <p className="eyebrow">Đăng ký Doanh nghiệp</p>
            <h1>Tạo hồ sơ Doanh nghiệp</h1>
          </div>
        </div>
        <div className="panel" style={{ maxWidth: 600 }}>
          <div className="form-group">
            <label>Ten Doanh nghiệp *</label>
            <input value={formData.name} onChange={e => setFormData({...formData, name: e.target.value})} placeholder="VD: EcoPoly Solutions" />
          </div>
          <div className="form-group">
            <label>Mã số thuế</label>
            <input value={formData.tax_code} onChange={e => setFormData({...formData, tax_code: e.target.value})} placeholder="VD: 0123456789" />
          </div>
          <div className="form-group">
            <label>Địa chỉ</label>
            <input value={formData.address} onChange={e => setFormData({...formData, address: e.target.value})} placeholder="VD: 123 Le Loi, Q.1" />
          </div>
          <div className="form-group">
            <label>Thành phố</label>
            <input value={formData.city} onChange={e => setFormData({...formData, city: e.target.value})} placeholder="VD: TP.HCM" />
          </div>
          <div className="form-group">
            <label>Mô tả</label>
            <textarea rows={4} value={formData.description} onChange={e => setFormData({...formData, description: e.target.value})} placeholder="Mô tả ve hoat dong kinh doanh..." />
          </div>
          <div className="form-actions">
            <button className="btn btn-ghost" onClick={() => setShowForm(false)}>Huy</button>
            <button className="btn btn-primary" onClick={handleCreate} disabled={submitting}>
              {submitting ? 'Đang tạo...' : 'Tạo hồ sơ'}
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
      </div>
    </Layout>
  )
}
