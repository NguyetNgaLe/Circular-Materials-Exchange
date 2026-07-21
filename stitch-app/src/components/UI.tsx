import React from 'react'

const statusMap: Record<string, { label: string; cls: string }> = {
  active: { label: 'Đang bán', cls: 'status-active' },
  pending_review: { label: 'Chờ duyệt', cls: 'status-pending' },
  sold: { label: 'Đã bán', cls: 'status-sold' },
  hidden: { label: 'Đã ẩn', cls: 'status-hidden' },
  open: { label: 'Đang mở', cls: 'status-active' },
  closed: { label: 'Đã đóng', cls: 'status-hidden' },
  matched: { label: 'Đã khớp', cls: 'status-verified' },
  draft: { label: 'Nháp', cls: 'status-draft' },
  pending: { label: 'Chờ duyệt', cls: 'status-pending' },
  verified: { label: 'Đã xác minh', cls: 'status-verified' },
  rejected: { label: 'Từ chối', cls: 'status-rejected' },
  accepted: { label: 'Chấp nhận', cls: 'status-verified' },
  cancelled: { label: 'Đã hủy', cls: 'status-rejected' },
  expired: { label: 'Hết hạn', cls: 'status-hidden' },
  confirmed: { label: 'Đã xác nhận', cls: 'status-pending' },
  in_progress: { label: 'Đang thực hiện', cls: 'status-active' },
  buyer_confirmed: { label: 'Buyer xác nhận', cls: 'status-pending' },
  seller_confirmed: { label: 'Seller xác nhận', cls: 'status-pending' },
  completed: { label: 'Hoàn tất', cls: 'status-verified' },
  disputed: { label: 'Tranh chấp', cls: 'status-rejected' },
  reviewed: { label: 'Đã xem', cls: 'status-pending' },
  resolved: { label: 'Đã xử lý', cls: 'status-verified' },
  dismissed: { label: 'Bỏ qua', cls: 'status-hidden' },
}

export function StatusBadge({ status }: { status: string }) {
  const info = statusMap[status] || { label: status, cls: '' }
  return <span className={`status-badge ${info.cls}`}>{info.label}</span>
}

export function StarRating({ rating, size = 16 }: { rating: number; size?: number }) {
  return (
    <span className="star-rating">
      {[1, 2, 3, 4, 5].map(i => (
        <span key={i} className={`star ${i <= rating ? 'filled' : i - 0.5 <= rating ? 'half' : ''}`} style={{ fontSize: size }}>★</span>
      ))}
    </span>
  )
}

export function EmptyState({ icon, title, description }: { icon: string; title: string; description: string }) {
  return (
    <div className="empty-state">
      <span className="empty-icon">{icon}</span>
      <h3>{title}</h3>
      <p>{description}</p>
    </div>
  )
}
