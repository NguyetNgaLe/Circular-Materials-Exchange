import React, { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { StatusBadge } from '../../components/UI'
import { Package, ArrowLeftRight, TrendingUp, CheckCircle, ShoppingCart, Loader2 } from 'lucide-react'

export default function DashboardPage() {
  const store = useStore()
  const user = store.currentUser
  const [offers, setOffers] = useState<any[]>([])
  const [transactions, setTransactions] = useState<any[]>([])
  const [listings, setListings] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    if (!user) return
    const loadData = async () => {
      setLoading(true)
      const [offersData, txData, listingsData] = await Promise.all([
        store.getOffers(),
        store.getTransactions(),
        store.getListings()
      ])
      setOffers(offersData)
      setTransactions(txData)
      setListings(listingsData)
      setLoading(false)
    }
    loadData()
  }, [user])

  if (!user) return null

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

  const myListings = listings.filter(l => l.sellerId === user.id)
  const sentOffers = offers.filter(o => o.buyerId === user.id)
  const receivedOffers = offers.filter(o => o.sellerId === user.id)
  const pendingReceived = receivedOffers.filter(o => o.status === 'pending')
  const activeTx = transactions.filter(t => ['confirmed', 'in_progress', 'buyer_confirmed', 'seller_confirmed'].includes(t.status))
  const completedTx = transactions.filter(t => t.status === 'completed')

  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Tổng quan</p>
          <h1>Dashboard</h1>
          <p className="muted">Xin chào, {user.name}</p>
        </div>
        <Link to="/listings/new" className="btn btn-primary">+ Đăng vật liệu</Link>
      </div>

      <div className="stats-grid">
        <div className="stat-card">
          <div className="stat-icon green"><Package size={22} /></div>
          <div>
            <strong>{myListings.length}</strong>
            <span>Nguồn cung đang bán</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon blue"><TrendingUp size={22} /></div>
          <div>
            <strong>{pendingReceived.length}</strong>
            <span>Đề nghị chờ xử lý</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon orange"><ShoppingCart size={22} /></div>
          <div>
            <strong>{sentOffers.length}</strong>
            <span>Đề nghị đã gửi</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon blue"><ArrowLeftRight size={22} /></div>
          <div>
            <strong>{activeTx.length}</strong>
            <span>Giao dịch đang thực hiện</span>
          </div>
        </div>
        <div className="stat-card">
          <div className="stat-icon green"><CheckCircle size={22} /></div>
          <div>
            <strong>{completedTx.length}</strong>
            <span>Giao dịch hoàn tất</span>
          </div>
        </div>
      </div>

      <div className="grid-2">
        <div className="panel">
          <div className="panel-header">
            <h2>Đề nghị mua gần đây</h2>
            <Link to="/offers/received" className="link-blue">Xem tất cả</Link>
          </div>
          {receivedOffers.length === 0 ? (
            <p className="muted">Chưa có đề nghị nào</p>
          ) : (
            <div className="table-responsive">
              <table className="table">
                <thead>
                  <tr><th>Vật liệu</th><th>Người mua</th><th>Giá đề xuất</th><th>Trạng thái</th></tr>
                </thead>
                <tbody>
                  {receivedOffers.slice(0, 5).map(o => (
                    <tr key={o.id}>
                      <td><Link to={`/material/${o.listingId}`} className="link-blue">{o.listingTitle}</Link></td>
                      <td>{o.buyerName}</td>
                      <td>{o.proposedPrice?.toLocaleString()}đ/{o.unit}</td>
                      <td><StatusBadge status={o.status} /></td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>

        <div className="panel">
          <div className="panel-header">
            <h2>Giao dịch gần đây</h2>
            <Link to="/transactions" className="link-blue">Xem tất cả</Link>
          </div>
          {transactions.length === 0 ? (
            <p className="muted">Chưa có giao dịch nào</p>
          ) : (
            <div className="table-responsive">
              <table className="table">
                <thead>
                  <tr><th>Vật liệu</th><th>Vai trò</th><th>Đối tác</th><th>Trạng thái</th></tr>
                </thead>
                <tbody>
                  {transactions.slice(0, 5).map(t => {
                    const isBuyer = user.id === t.buyerId
                    return (
                      <tr key={t.id}>
                        <td><Link to={`/transactions/${t.id}`} className="link-blue">{t.listingTitle}</Link></td>
                        <td><span className={`tag ${isBuyer ? 'tag-outline' : 'tag-green'}`}>{isBuyer ? 'Mua' : 'Bán'}</span></td>
                        <td>{isBuyer ? t.sellerName : t.buyerName}</td>
                        <td><StatusBadge status={t.status} /></td>
                      </tr>
                    )
                  })}
                </tbody>
              </table>
            </div>
          )}
        </div>
      </div>
    </Layout>
  )
}
