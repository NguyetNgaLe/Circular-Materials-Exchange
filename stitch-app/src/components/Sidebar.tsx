import React, { useState, useEffect } from 'react'
import { Link, useLocation } from 'react-router-dom'
import { useStore } from '../store'
import {
  LayoutDashboard, Package, ShoppingCart, ArrowLeftRight,
  Star, Building2, FileText, Bell, Settings, Shield, Store, DollarSign
} from 'lucide-react'

interface NavItem { to: string; label: string; icon: React.ReactNode }

const businessNav: NavItem[] = [
  { to: '/dashboard', label: 'Dashboard', icon: <LayoutDashboard size={18} /> },
  { to: '/marketplace', label: 'Chợ Giao Dịch', icon: <Store size={18} /> },
  { to: '/listings', label: 'Nguồn cung của tôi', icon: <Package size={18} /> },
  { to: '/offers/sent', label: 'Đề nghị đã gửi', icon: <FileText size={18} /> },
  { to: '/offers/received', label: 'Đề nghị đã nhận', icon: <ShoppingCart size={18} /> },
  { to: '/transactions', label: 'Giao dịch', icon: <ArrowLeftRight size={18} /> },
  { to: '/reviews', label: 'Đánh giá', icon: <Star size={18} /> },
  { to: '/company', label: 'Doanh nghiệp', icon: <Building2 size={18} /> },
  { to: '/notifications', label: 'Thông báo', icon: <Bell size={18} /> },
]

const adminNav: NavItem[] = [
  { to: '/admin', label: 'Dashboard', icon: <LayoutDashboard size={18} /> },
  { to: '/admin/companies', label: 'Duyệt doanh nghiệp', icon: <Building2 size={18} /> },
  { to: '/admin/categories', label: 'Danh mục vật liệu', icon: <Package size={18} /> },
  { to: '/admin/listings', label: 'Quản lý listing', icon: <FileText size={18} /> },
  { to: '/admin/transactions', label: 'Giao dịch', icon: <ArrowLeftRight size={18} /> },
  { to: '/admin/finance', label: 'Tài chính', icon: <DollarSign size={18} /> },
  { to: '/admin/reports', label: 'Báo cáo vi phạm', icon: <Shield size={18} /> },
  { to: '/admin/export', label: 'Xuất CSV', icon: <Settings size={18} /> },
]

export default function Sidebar() {
  const store = useStore()
  const location = useLocation()
  const role = store.currentUser?.role
  const [company, setCompany] = useState<any>(null)

  const nav = role === 'admin' ? adminNav : businessNav

  useEffect(() => {
    if (store.currentUser?.companyId) {
      store.getCompany(store.currentUser.companyId)
        .then((c: any) => setCompany(c))
        .catch(() => setCompany(null))
    }
  }, [store.currentUser])

  return (
    <aside className="sidebar">
      {company && (
        <div className="company-card">
          <strong>{company.name}</strong>
          <span className={`status-badge status-${company.status}`}>
            {company.status === 'verified' ? 'Đã xác minh' : company.status === 'pending' ? 'Chờ duyệt' : company.status}
          </span>
        </div>
      )}
      {store.currentUser && !company && store.currentUser.role !== 'admin' && (
        <div className="company-card">
          <strong>{store.currentUser.name}</strong>
          <span className="muted">Doanh nghiệp</span>
        </div>
      )}
      <nav className="sidebar-nav">
        {nav.map(item => (
          <Link key={item.to} to={item.to} className={`sidebar-link ${location.pathname === item.to ? 'active' : ''}`}>
            {item.icon}
            <span>{item.label}</span>
          </Link>
        ))}
      </nav>
    </aside>
  )
}
