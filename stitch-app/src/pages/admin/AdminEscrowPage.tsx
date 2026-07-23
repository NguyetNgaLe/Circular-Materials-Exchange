import React, { useState, useEffect } from 'react'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { Wallet, ArrowDownLeft, Loader2, Check } from 'lucide-react'

export default function AdminEscrowPage() {
  const [escrows, setEscrows] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => { loadData() }, [])

  const loadData = async () => {
    setLoading(true)
    try {
      const token = localStorage.getItem('token')
      const headers = { 'Authorization': `Bearer ${token}` }
      const res = await fetch('/api/admin/escrow', { headers }).then(r => r.json())
      if (res.success) {
        setEscrows(res.data.escrows || [])
      }
    } catch (err) {
      console.error('Failed to load data:', err)
    }
    setLoading(false)
  }

  const handleRelease = async (id: string) => {
    if (!confirm('Xác nhận Giải ngân? Tien se chuyen truc tiep ve tai khoan Ngân hang cua seller.')) return
    try {
      const token = localStorage.getItem('token')
      const res = await fetch(`/api/admin/escrow/${id}/release`, {
        method: 'POST',
        headers: { 'Authorization': `Bearer ${token}` }
      }).then(r => r.json())
      if (res.success) {
        alert('Giải ngân Thành công! Tiền đã chuyển ve tai khoan seller.')
        loadData()
      }
    } catch (err) { alert('Có lỗi xảy ra') }
  }

  if (loading) {
    return <Layout showSidebar><div className="empty-state"><Loader2 className="animate-spin" size={32} /><p>Đang tải...</p></div></Layout>
  }

  const holdingEscrows = escrows.filter(e => e.status === 'holding' && e.amount > 0)
  const releasedEscrows = escrows.filter(e => e.status === 'released')
  const totalHolding = holdingEscrows.reduce((sum, e) => sum + (e.amount || 0), 0)

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản trị Tài chính</p>
          <h1>Quản lý Escrow</h1>
        </div>
      </div>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-icon orange"><Wallet size={22} /></div>
          <div><strong>{totalHolding.toLocaleString()}đ</strong><span>Tiền đang giữ</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-icon blue"><ArrowDownLeft size={22} /></div>
          <div><strong>{holdingEscrows.length}</strong><span>Giao dịch đang giữ</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-icon green"><ArrowDownLeft size={22} /></div>
          <div><strong>{releasedEscrows.length}</strong><span>Giao dịch đã chuyển</span></div>
        </div>
      </div>

      <div style={{ padding: 12, background: '#eff6ff', borderRadius: 8, marginBottom: 16 }}>
        <p style={{ fontSize: 13, color: '#1e40af' }}>
          <strong>Mô hình:</strong> Tiền buyer Thanh toán duoc giu tai escrow. Khi buyer Xác nhận Nhận hàng, tien tu dong chuyen 98% ve tai khoan Ngân hang seller, san giu 2% phi.
        </p>
      </div>

      <div className="panel">
        <div className="panel-header"><h2>Danh sách Escrow</h2></div>
        {escrows.length === 0 ? (
          <p className="muted" style={{ padding: 20 }}>Chưa có escrow nào</p>
        ) : (
          <div className="table-responsive">
            <table className="table">
              <thead><tr><th>Ngày</th><th>Buyer</th><th>Seller</th><th>Tổng tiền</th><th>Phí (2%)</th><th>Seller nhận</th><th>Trạng thái</th><th>Thao tác</th></tr></thead>
              <tbody>
                {escrows.filter(e => e.amount > 0).map(e => (
                  <tr key={e.id}>
                    <td>{e.createdAt ? new Date(e.createdAt).toLocaleDateString('vi-VN') : '—'}</td>
                    <td>{e.buyerName || '—'}</td>
                    <td>{e.sellerName || '—'}</td>
                    <td>{(e.amount || 0).toLocaleString()}đ</td>
                    <td style={{ color: 'green' }}>{(e.feeAmount || 0).toLocaleString()}đ</td>
                    <td><strong>{(e.sellerAmount || 0).toLocaleString()}đ</strong></td>
                    <td><StatusBadge status={e.status} /></td>
                    <td>
                      {e.status === 'holding' && (
                        <button className="icon-btn-sm success" title="Giải ngân" onClick={() => handleRelease(e.id)}>
                          <Check size={16} />
                        </button>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </Layout>
  )
}
