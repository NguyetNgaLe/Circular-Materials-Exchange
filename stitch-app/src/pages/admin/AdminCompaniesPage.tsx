import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge, StarRating } from '../../components/UI'
import { Check, X, Loader2 } from 'lucide-react'

export default function AdminCompaniesPage() {
  const store = useStore()
  const [tab, setTab] = useState<'pending' | 'all'>('pending')
  const [rejectId, setRejectId] = useState<string | null>(null)
  const [rejectReason, setRejectReason] = useState('')
  const [companies, setCompanies] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  const loadCompanies = async () => {
    setLoading(true)
    const data = await store.getCompanies()
    setCompanies(data)
    setLoading(false)
  }

  useEffect(() => {
    loadCompanies()
  }, [])

  const pending = companies.filter(c => c.status === 'pending')
  const all = companies
  const list = tab === 'pending' ? pending : all

  const handleApprove = async (id: string) => {
    const success = await store.approveCompany(id)
    if (success) {
      alert('Đã duyệt doanh nghiệp!')
      loadCompanies()
    }
  }

  const handleReject = async () => {
    if (rejectId && rejectReason) {
      const success = await store.rejectCompany(rejectId, rejectReason)
      if (success) {
        setRejectId(null)
        setRejectReason('')
        alert('Đã từ chối doanh nghiệp!')
        loadCompanies()
      }
    }
  }

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
          <h1>Quản Lý Doanh Nghiệp</h1>
        </div>
      </div>

      <div className="tabs">
        <button className={`tab ${tab === 'pending' ? 'active' : ''}`} onClick={() => setTab('pending')}>Chờ duyệt ({pending.length})</button>
        <button className={`tab ${tab === 'all' ? 'active' : ''}`} onClick={() => setTab('all')}>Tất cả ({all.length})</button>
      </div>

      <div className="table-responsive">
        <table className="table">
          <thead>
            <tr><th>Tên doanh nghiệp</th><th>Mã số thuế</th><th>Thành phố</th><th>Đánh giá</th><th>Trạng thái</th><th>Thao tác</th></tr>
          </thead>
          <tbody>
            {list.map(c => (
              <tr key={c.id}>
                <td>
                  <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                    {c.imageUrl && <img src={c.imageUrl} alt="" style={{ width: 40, height: 40, objectFit: 'cover', borderRadius: 6 }} />}
                    <div>
                      <strong>{c.name}</strong><br />
                      <span className="muted" style={{ fontSize: 12 }}>{c.description?.slice(0, 50)}...</span>
                    </div>
                  </div>
                </td>
                <td>{c.taxCode}</td>
                <td>{c.city}</td>
                <td>{c.rating > 0 ? <><StarRating rating={c.rating} /> ({c.reviewCount})</> : '—'}</td>
                <td><StatusBadge status={c.status} /></td>
                <td>
                  <div className="action-btns">
                    {c.status === 'pending' && (
                      <>
                        <button className="icon-btn-sm success" title="Duyệt" onClick={() => handleApprove(c.id)}><Check size={16} /></button>
                        <button className="icon-btn-sm danger" title="Từ chối" onClick={() => setRejectId(c.id)}><X size={16} /></button>
                      </>
                    )}
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {rejectId && (
        <div className="modal-overlay" onClick={() => setRejectId(null)}>
          <div className="modal" onClick={e => e.stopPropagation()}>
            <h2>Từ chối doanh nghiệp</h2>
            <div className="form-group">
              <label>Lý do từ chối</label>
              <textarea rows={3} value={rejectReason} onChange={e => setRejectReason(e.target.value)} placeholder="Nhập lý do..." />
            </div>
            <div className="form-actions">
              <button className="btn btn-ghost" onClick={() => setRejectId(null)}>Hủy</button>
              <button className="btn btn-danger" onClick={handleReject}>Từ chối</button>
            </div>
          </div>
        </div>
      )}
    </Layout>
  )
}
