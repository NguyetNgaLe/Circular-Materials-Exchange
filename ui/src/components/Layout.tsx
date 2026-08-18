import React from 'react'
import Navbar from './Navbar'
import Sidebar from './Sidebar'

interface Props { children: React.ReactNode; showSidebar?: boolean }

export default function Layout({ children, showSidebar = false }: Props) {
  return (
    <div className="app-root">
      <Navbar />
      <div className="app-body">
        {showSidebar && <Sidebar />}
        <main className={`main-content ${showSidebar ? 'with-sidebar' : ''}`}>
          {children}
        </main>
      </div>
    </div>
  )
}
