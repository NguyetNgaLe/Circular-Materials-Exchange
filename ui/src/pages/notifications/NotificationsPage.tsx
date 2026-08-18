import React, { useState, useEffect } from 'react'
import { useStore } from '../../store'
import Layout from '../../components/Layout'
import { Loader2 } from 'lucide-react'

export default function NotificationsPage() {
  const store = useStore()
  const [notifs, setNotifs] = useState<any[]>([])
  const [loading, setLoading] = useState(true)

  const loadNotifications = async () => {
    setLoading(true)
    const data = await store.getNotifications()
    setNotifs(data)
    setLoading(false)
  }

  useEffect(() => {
    loadNotifications()
  }, [])

  const handleMarkRead = async (id: string) => {
    await store.markNotificationRead(id)
    setNotifs(prev => prev.map(n => n.id === id ? { ...n, read: true } : n))
  }

  const handleMarkAllRead = async () => {
    await store.markAllNotificationsRead()
    setNotifs(prev => prev.map(n => ({ ...n, read: true })))
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
          <p className="eyebrow">Thông báo</p>
          <h1>Thông Báo</h1>
        </div>
        {notifs.some(n => !n.read) && (
          <button className="btn btn-outline" onClick={handleMarkAllRead}>Đánh dấu tất cả đã đọc</button>
        )}
      </div>

      {notifs.length === 0 ? (
        <div className="empty-state">
          <span className="empty-icon">🔔</span>
          <h3>Không có thông báo</h3>
        </div>
      ) : (
        <div className="notifications-list">
          {notifs.map(n => (
            <div key={n.id} className={`notification-item ${n.read ? '' : 'unread'}`} onClick={() => !n.read && handleMarkRead(n.id)}>
              <div className="notif-icon">
                {n.type === 'offer' ? '📄' : n.type === 'transaction' ? '🔄' : n.type === 'review' ? '⭐' : '📢'}
              </div>
              <div className="notif-content">
                <strong>{n.title}</strong>
                <p>{n.message}</p>
                <span className="notif-time">{new Date(n.createdAt).toLocaleString('vi-VN')}</span>
              </div>
              {!n.read && <span className="notif-dot" />}
            </div>
          ))}
        </div>
      )}
    </Layout>
  )
}
