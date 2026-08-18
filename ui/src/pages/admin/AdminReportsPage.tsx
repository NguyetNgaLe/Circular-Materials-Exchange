import React from 'react'
import Layout from '../../components/Layout'

export default function AdminReportsPage() {
  return (
    <Layout showSidebar>
      <div className="page-header">
        <div>
          <p className="eyebrow">Quản trị</p>
          <h1>Báo Cáo Vi Phạm</h1>
        </div>
      </div>

      <div className="empty-state">
        <span className="empty-icon">🛡️</span>
        <h3>Chức năng đang phát triển</h3>
        <p className="muted">Tính năng báo cáo vi phạm sẽ sớm được cập nhật.</p>
      </div>
    </Layout>
  )
}
