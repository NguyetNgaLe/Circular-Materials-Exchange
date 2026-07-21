import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { Building2, Package, ArrowLeftRight, TrendingUp, AlertTriangle, Loader2 } from 'lucide-react'

export default function AdminDashboardPage() {
  const store = useStore()
  const [companies, setCompanies] = useState<any[]>([])
  const [transactions, setTransactions] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const loadData = async () => {
      setLoading(true)
      const [companiesData, txData] = await Promise.all([
        store.getCompanies(),
        store.getTransactions()
      ])
      setCompanies(companiesData)
      setTransactions(txData)
      setLoading(false)
    }
    loadData()
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

  const pendingCompanies = companies.filter(c => c.status === 'pending')
  const activeTx = transactions.filter(t => ['confirmed', 'in_progress'].includes(t.status))

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản trị hệ thống</p>
          <h1>Admin Dashboard</h1>
        </div>
      </div>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-icon blue"><Building2 size={22} /></div>
          <div><strong>{companies.length}</strong><span>Doanh nghiệp</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-icon orange"><TrendingUp size={22} /></div>
          <div><strong>{pendingCompanies.length}</strong><span>Chờ duyệt</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-icon blue"><ArrowLeftRight size={22} /></div>
          <div><strong>{activeTx.length}</strong><span>Giao dịch active</span></div>
        </div>
        <div className="stat-card">
          <div className="stat-icon green"><Package size={22} /></div>
          <div><strong>{transactions.length}</strong><span>Tổng giao dịch</span></div>
        </div>
      </div>

      <div className="grid-2">
        <div className="panel">
          <div className="panel-header">
            <h2>Doanh nghiệp chờ duyệt</h2>
            <Link to="/admin/companies" className="link-blue">Xem tất cả</Link>
          </div>
          {pendingCompanies.length === 0 ? (
            <p className="muted">Không có doanh nghiệp chờ duyệt</p>
          ) : (
            <div className="table-responsive">
              <table className="table">
                <thead><tr><th>Tên</th><th>Mã số thuế</th><th>Ngày gửi</th></tr></thead>
                <tbody>
                  {pendingCompanies.map(c => (
                    <tr key={c.id}>
                      <td><Link to="/admin/companies" className="link-blue">{c.name}</Link></td>
                      <td>{c.taxCode}</td>
                      <td>{c.memberSince}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>

        <div className="panel">
          <div className="panel-header">
            <h2>Thống kê giao dịch</h2>
          </div>
          <div className="info-list">
            <div className="info-row"><span>Tổng giao dịch</span><strong>{transactions.length}</strong></div>
            <div className="info-row"><span>Hoàn tất</span><strong>{transactions.filter(t => t.status === 'completed').length}</strong></div>
            <div className="info-row"><span>Đang thực hiện</span><strong>{transactions.filter(t => t.status === 'in_progress').length}</strong></div>
            <div className="info-row"><span>Đã hủy</span><strong>{transactions.filter(t => t.status === 'cancelled').length}</strong></div>
          </div>
          <div style={{ marginTop: 16 }}>
            <Link to="/admin/export" className="btn btn-outline">Xuất CSV</Link>
          </div>
        </div>
      </div>
    </Layout>
  )
}
