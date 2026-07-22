import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { DollarSign, TrendingUp, ArrowLeftRight, Wallet, Loader2 } from 'lucide-react'

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
          <p>Dang tai du lieu tai chinh...</p>
        </div>
      </Layout>
    )
  }

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quan tri tai chinh</p>
          <h1>Dong Tien Thu Chi</h1>
        </div>
      </div>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-icon green"><DollarSign size={22} /></div>
          <div>
            <strong>{(overview?.totalIncome || 0).toLocaleString()}đ</strong>
            <span>Tong doanh thu</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon blue"><TrendingUp size={22} /></div>
          <div>
            <strong>{(overview?.monthIncome || 0).toLocaleString()}đ</strong>
            <span>Doanh thu thang nay</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon orange"><ArrowLeftRight size={22} /></div>
          <div>
            <strong>{overview?.totalTransactions || 0}</strong>
            <span>So giao dich</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon green"><Wallet size={22} /></div>
          <div>
            <strong>{(overview?.walletBalance || 0).toLocaleString()}đ</strong>
            <span>So du vi</span>
          </div>
        </div>
      </div>

      <div className="grid-2">
        <div className="panel">
          <div className="panel-header">
            <h2>Doanh thu theo thang</h2>
          </div>
          {(overview?.monthlyData || []).length === 0 ? (
            <p className="muted" style={{ padding: 20 }}>Chua co du lieu doanh thu</p>
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
            <h2>Thong tin vi san</h2>
          </div>
          <div className="info-list">
            <div className="info-row"><span>So du hien tai</span><strong>{(overview?.walletBalance || 0).toLocaleString()}đ</strong></div>
            <div className="info-row"><span>Tong thu</span><strong style={{ color: 'green' }}>{(overview?.totalIncome || 0).toLocaleString()}đ</strong></div>
            <div className="info-row"><span>Tong chi</span><strong style={{ color: 'red' }}>{(overview?.totalExpense || 0).toLocaleString()}đ</strong></div>
            <div className="info-row"><span>Loi nhuan</span><strong>{((overview?.totalIncome || 0) - (overview?.totalExpense || 0)).toLocaleString()}đ</strong></div>
          </div>
          <div style={{ marginTop: 16, padding: 12, background: '#f0fdf4', borderRadius: 8 }}>
            <p style={{ fontSize: 13, color: '#166534' }}>
              <strong>Mo hinh phi:</strong> Thu 2% tu seller moi giao dich hoan tat
            </p>
          </div>
        </div>
      </div>

      <div className="panel">
        <div className="panel-header">
          <h2>Lich su phi giao dich</h2>
        </div>
        {fees.length === 0 ? (
          <p className="muted" style={{ padding: 20 }}>Chua co phi giao dich nao</p>
        ) : (
          <div className="table-responsive">
            <table className="table">
              <thead>
                <tr>
                  <th>Ngay</th>
                  <th>Giao dich</th>
                  <th>Tien giao dich</th>
                  <th>Ty le phi</th>
                  <th>So tien phi</th>
                  <th>Trang thai</th>
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
