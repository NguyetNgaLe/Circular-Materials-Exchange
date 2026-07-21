import React, { useState, useEffect } from 'react'
import { Link, useLocation } from 'react-router-dom'
import { useStore } from '../store'
import { Search, Bell, User, Menu, X } from 'lucide-react'

export default function Navbar() {
  const store = useStore()
  const location = useLocation()
  const [menuOpen, setMenuOpen] = useState(false)
  const [unread, setUnread] = useState(0)

  useEffect(() => {
    if (store.currentUser) {
      store.getUserNotifications()
        .then((notifs: any[]) => setUnread(notifs.filter((n: any) => !n.read).length))
        .catch(() => setUnread(0))
    }
  }, [store.currentUser])

  const navLinks = store.currentUser?.role === 'admin'
    ? [
        { to: '/admin', label: 'Dashboard' },
        { to: '/admin/companies', label: 'Doanh nghiệp' },
        { to: '/admin/categories', label: 'Danh mục' },
        { to: '/admin/reports', label: 'Báo cáo' },
      ]
    : [
        { to: '/marketplace', label: 'Chợ Giao Dịch' },
        { to: '/demand', label: 'Nhu Cầu Mua' },
      ]

  return (
    <nav className="navbar">
      <div className="navbar-inner">
        <Link to="/" className="navbar-brand">
          <span className="brand-icon">♻</span>
          Circular Materials
        </Link>

        <div className={`navbar-links ${menuOpen ? 'open' : ''}`}>
          {navLinks.map(l => (
            <Link key={l.to} to={l.to} className={`nav-link ${location.pathname === l.to ? 'active' : ''}`} onClick={() => setMenuOpen(false)}>
              {l.label}
            </Link>
          ))}
        </div>

        <div className="navbar-actions">
          {store.currentUser && (
            <>
              <div className="search-box">
                <Search size={16} />
                <input placeholder="Tìm kiếm vật liệu..." />
              </div>
              <Link to="/notifications" className="icon-btn" style={{ position: 'relative' }}>
                <Bell size={20} />
                {unread > 0 && <span className="badge">{unread}</span>}
              </Link>
              <div className="user-menu">
                <Link to={store.currentUser.role === 'admin' ? '/admin' : '/dashboard'} className="icon-btn">
                  <User size={20} />
                </Link>
                <span className="user-name">{store.currentUser.name}</span>
                <button className="btn-ghost-sm" onClick={() => store.logout()}>Đăng xuất</button>
              </div>
            </>
          )}
          {!store.currentUser && (
            <div style={{ display: 'flex', gap: 8 }}>
              <Link to="/login" className="btn btn-outline btn-sm">Đăng ký</Link>
              <Link to="/login" className="btn btn-primary btn-sm">Đăng nhập</Link>
            </div>
          )}
          <button className="menu-toggle" onClick={() => setMenuOpen(!menuOpen)}>
            {menuOpen ? <X size={24} /> : <Menu size={24} />}
          </button>
        </div>
      </div>
    </nav>
  )
}
