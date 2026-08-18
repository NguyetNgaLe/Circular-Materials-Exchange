import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { DollarSign, TrendingUp, ArrowLeftRight, Loader2 } from 'lucide-react'

export default function AdminFinancePage() {
  const store = useStore()
  const [overview, setOverview] = useState<any>(null)
  const [fees, setFees] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const loadData = async () => {
      setLoading(true)
      try {
        const token = localStorage.getItem('token')
        const headers = { 'Authorization': `Bearer ${token}` }
        
        const [overviewRes, feesRes] = await Promise.all([
          fetch('/api/admin/finance/overview', { headers }).then(r => r.json()),
          fetch('/api/admin/finance/fees', { headers }).then(r => r.json()),
        ])
        
        if (overviewRes.success) setOverview(overviewRes.data)
        if (feesRes.success) setFees(feesRes.data.fees || [])
      } catch (err) {
        console.error('Failed to load finance data:', err)
      }
      setLoading(false)
    }
    loadData()
  }, [])

  if (loading) {
    return (
      <Layout showSidebar>
        <div className="empty-state">
          <Loader2 className="animate-spin" size={32} />
          <p>Đang tải du lieu Tài chính...</p>
        </div>
      </Layout>
    )
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản trị Tài chính</p>
          <h1>Dong Tien Thu Chi</h1>
        </div>
      </div>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-icon green"><DollarSign size={22} /></div>
          <div>
            <strong>{(overview?.totalIncome || 0).toLocaleString()}đ</strong>
            <span>Tổng doanh thu</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon blue"><TrendingUp size={22} /></div>
          <div>
            <strong>{(overview?.monthIncome || 0).toLocaleString()}đ</strong>
            <span>Doanh thu tháng nay</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon orange"><ArrowLeftRight size={22} /></div>
          <div>
            <strong>{overview?.totalTransactions || 0}</strong>
            <span>So Giao dịch</span>
          </div>
        </div>
      </div>

      <div className="grid-2">
        <div className="panel">
          <div className="panel-header">
            <h2>Doanh thu theo thang</h2>
          </div>
          {(overview?.monthlyData || []).length === 0 ? (
            <p className="muted" style={{ padding: 20 }}>Chưa có dữ liệu doanh thu</p>
          ) : (
            <div className="table-responsive">
              <table className="table">
                <thead><tr><th>Thang</th><th>Doanh thu</th></tr></thead>
                <tbody>
                  {(overview?.monthlyData || []).map((m: any) => (
                    <tr key={m.month}>
                      <td>{m.month}</td>
                      <td><strong>{m.total.toLocaleString()}đ</strong></td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>

        <div className="panel">
          <div className="panel-header">
            <h2>Thông tin doanh thu</h2>
          </div>
          <div className="info-list">
            <div className="info-row"><span>Tổng doanh thu</span><strong style={{ color: 'green' }}>{(overview?.totalIncome || 0).toLocaleString()}đ</strong></div>
            <div className="info-row"><span>Doanh thu tháng nay</span><strong>{(overview?.monthIncome || 0).toLocaleString()}đ</strong></div>
            <div className="info-row"><span>So Giao dịch Hoàn tất</span><strong>{overview?.totalTransactions || 0}</strong></div>
          </div>
          <div style={{ marginTop: 16, padding: 12, background: '#f0fdf4', borderRadius: 8 }}>
            <p style={{ fontSize: 13, color: '#166534' }}>
              <strong>Mô hình Thanh toán:</strong> Tien tu dong chuyen ve tai khoan Ngân hang da lien ket cua tung Doanh nghiệp khi Giao dịch Hoàn tất. San thu phi 2%.
            </p>
          </div>
        </div>
      </div>

      <div className="panel">
        <div className="panel-header">
          <h2>Lịch sử phi Giao dịch</h2>
        </div>
        {fees.length === 0 ? (
          <p className="muted" style={{ padding: 20 }}>Chưa có phi Giao dịch nao</p>
        ) : (
          <div className="table-responsive">
            <table className="table">
              <thead>
                <tr>
                  <th>Ngay</th>
                  <th>Giao dịch</th>
                  <th>Tien Giao dịch</th>
                  <th>Ty le phi</th>
                  <th>Số tiền phi</th>
                  <th>Trạng thái</th>
                </tr>
              </thead>
              <tbody>
                {fees.map(f => (
                  <tr key={f.id}>
                    <td>{new Date(f.createdAt).toLocaleDateString('vi-VN')}</td>
                    <td>{f.listingTitle || f.transactionId?.slice(0, 8)}</td>
                    <td>{f.transactionAmount.toLocaleString()}đ</td>
                    <td>{(f.feeRate * 100).toFixed(1)}%</td>
                    <td><strong style={{ color: 'green' }}>{f.feeAmount.toLocaleString()}đ</strong></td>
                    <td>
                      <span className={`status-badge ${f.status === 'collected' ? 'status-verified' : 'status-pending'}`}>
                        {f.status === 'collected' ? 'Da thu' : 'Cho thu'}
                      </span>
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
