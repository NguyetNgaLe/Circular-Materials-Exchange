import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import Layout from '../../components/Layout'
import { ArrowDownLeft, ArrowUpRight, Loader2 } from 'lucide-react'

interface SettlementEntry {
  id: string
  listingTitle: string
  direction: 'income' | 'expense'
  counterparty: string
  grossAmount: number
  feeAmount: number
  settledAmount: number
  createdAt: string
}

interface SettlementData {
  totalReceived: number
  totalPaid: number
  entries: SettlementEntry[]
}

const emptySettlement: SettlementData = {
  totalReceived: 0,
  totalPaid: 0,
  entries: [],
}

export default function CompanySettlementPage() {
  const [settlement, setSettlement] = useState<SettlementData>(emptySettlement)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => { loadData() }, [])

  const loadData = async () => {
    setLoading(true)
    setError('')
    try {
      const token = localStorage.getItem('token')
      const response = await fetch('/api/company/settlement', {
        headers: { Authorization: `Bearer ${token}` },
      })
      const payload = await response.json()
      if (!response.ok || !payload.success) {
        throw new Error(payload.message || 'Không thể tải dữ liệu đối soát')
      }
      setSettlement({ ...emptySettlement, ...payload.data, entries: payload.data.entries || [] })
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Không thể tải dữ liệu đối soát')
    } finally {
      setLoading(false)
    }
  }

  if (loading) {
    return <Layout showSidebar><div className="empty-state"><Loader2 className="animate-spin" size={32} /><p>Đang tải dữ liệu đối soát...</p></div></Layout>
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Doanh nghiệp</p>
          <h1>Đối soát thu chi</h1>
          <p className="muted">Tổng hợp các khoản thu và chi từ những giao dịch đã hoàn tất</p>
        </div>
        <div style={{ display: 'flex', gap: 8 }}>
          <button type="button" className="btn btn-outline" onClick={loadData}>Cập nhật dữ liệu</button>
          <Link to="/transactions" className="btn btn-outline">Xem lịch sử giao dịch</Link>
        </div>
      </div>

      {error && (
        <div className="empty-state">
          <h3>Không tải được dữ liệu đối soát</h3>
          <p>{error}</p>
          <button type="button" className="btn btn-primary" onClick={loadData}>Thử lại</button>
        </div>
      )}

      {!error && (
        <>
          <div className="stats-grid">
            <div className="stat-card">
              <div className="stat-icon green"><ArrowDownLeft size={22} /></div>
              <div><strong>{settlement.totalReceived.toLocaleString('vi-VN')}đ</strong><span>Tổng tiền đã nhận</span></div>
            </div>
            <div className="stat-card">
              <div className="stat-icon orange"><ArrowUpRight size={22} /></div>
              <div><strong>{settlement.totalPaid.toLocaleString('vi-VN')}đ</strong><span>Tổng tiền đã trả</span></div>
            </div>
          </div>

          <div className="panel">
            <div className="panel-header"><h2>Chi tiết đối soát</h2></div>
            <div style={{ padding: 12, background: '#eff6ff', borderRadius: 8, marginBottom: 16 }}>
              <p style={{ fontSize: 13, color: '#1e40af' }}>
                Chỉ giao dịch đã hoàn tất mới được tính. Khoản bán được ghi nhận sau khi trừ 2% phí sàn; khoản mua được ghi nhận theo toàn bộ giá trị giao dịch.
              </p>
            </div>
            {settlement.entries.length === 0 ? (
              <div className="empty-state">
                <h3>Chưa có dữ liệu thu chi</h3>
                <p>Dữ liệu sẽ được cập nhật khi doanh nghiệp hoàn tất một giao dịch mua hoặc bán.</p>
              </div>
            ) : (
              <div className="table-responsive">
                <table className="table">
                  <thead>
                    <tr><th>Ngày</th><th>Vật liệu</th><th>Loại</th><th>Đối tác</th><th>Giá trị giao dịch</th><th>Phí sàn</th><th>Ghi nhận</th></tr>
                  </thead>
                  <tbody>
                    {settlement.entries.map(entry => (
                      <tr key={entry.id}>
                        <td>{new Date(entry.createdAt).toLocaleDateString('vi-VN')}</td>
                        <td><Link to={`/transactions/${entry.id}`} className="link-blue">{entry.listingTitle}</Link></td>
                        <td><span className={`tag ${entry.direction === 'income' ? 'tag-green' : 'tag-outline'}`}>{entry.direction === 'income' ? 'Thu' : 'Chi'}</span></td>
                        <td>{entry.counterparty}</td>
                        <td>{entry.grossAmount.toLocaleString('vi-VN')}đ</td>
                        <td>{entry.direction === 'income' ? `${entry.feeAmount.toLocaleString('vi-VN')}đ` : '—'}</td>
                        <td><strong style={{ color: entry.direction === 'income' ? 'green' : '#c2410c' }}>{entry.direction === 'income' ? '+' : '-'}{entry.settledAmount.toLocaleString('vi-VN')}đ</strong></td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}
          </div>
        </>
      )}
    </Layout>
  )
}
