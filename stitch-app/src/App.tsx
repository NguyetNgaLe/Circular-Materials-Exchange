import React from 'react'
import { Routes, Route, Navigate } from 'react-router-dom'
import { useStore } from './store'
import LoginPage from './pages/auth/LoginPage'
import DashboardPage from './pages/dashboard/DashboardPage'
import MarketplacePage from './pages/marketplace/MarketplacePage'
import DemandMarketplacePage from './pages/marketplace/DemandMarketplacePage'
import MaterialDetailPage from './pages/material/MaterialDetailPage'
import CompanyProfilePage from './pages/company/CompanyProfilePage'
import PostMaterialPage from './pages/listings/PostMaterialPage'
import MyListingsPage from './pages/listings/MyListingsPage'
import NewOfferPage from './pages/offers/NewOfferPage'
import SentOffersPage from './pages/offers/SentOffersPage'
import ReceivedOffersPage from './pages/offers/ReceivedOffersPage'
import TransactionListPage from './pages/transactions/TransactionListPage'
import TransactionDetailPage from './pages/transactions/TransactionDetailPage'
import NewReviewPage from './pages/reviews/NewReviewPage'
import ReviewsPage from './pages/reviews/ReviewsPage'
import AdminDashboardPage from './pages/admin/AdminDashboardPage'
import AdminCompaniesPage from './pages/admin/AdminCompaniesPage'
import AdminCategoriesPage from './pages/admin/AdminCategoriesPage'
import AdminReportsPage from './pages/admin/AdminReportsPage'
import AdminExportPage from './pages/admin/AdminExportPage'
import AdminListingsPage from './pages/admin/AdminListingsPage'
import AdminTransactionsPage from './pages/admin/AdminTransactionsPage'
import AdminFinancePage from './pages/admin/AdminFinancePage'
import NotificationsPage from './pages/notifications/NotificationsPage'

function ProtectedRoute({ children, role }: { children: React.ReactNode; role?: string }) {
  const store = useStore()
  if (!store.currentUser) return <Navigate to="/login" />
  if (role && store.currentUser.role !== role) return <Navigate to="/dashboard" />
  return <>{children}</>
}

export default function App() {
  return (
    <Routes>
      <Route path="/login" element={<LoginPage />} />
      <Route path="/" element={<Navigate to="/marketplace" />} />
      <Route path="/marketplace" element={<MarketplacePage />} />
      <Route path="/demand" element={<DemandMarketplacePage />} />
      <Route path="/material/:id" element={<MaterialDetailPage />} />

      <Route path="/dashboard" element={<ProtectedRoute><DashboardPage /></ProtectedRoute>} />
      <Route path="/company" element={<ProtectedRoute><CompanyProfilePage /></ProtectedRoute>} />
      <Route path="/listings" element={<ProtectedRoute><MyListingsPage /></ProtectedRoute>} />
      <Route path="/listings/new" element={<ProtectedRoute><PostMaterialPage /></ProtectedRoute>} />
      <Route path="/offers/new/:listingId" element={<ProtectedRoute><NewOfferPage /></ProtectedRoute>} />
      <Route path="/offers/sent" element={<ProtectedRoute><SentOffersPage /></ProtectedRoute>} />
      <Route path="/offers/received" element={<ProtectedRoute><ReceivedOffersPage /></ProtectedRoute>} />
      <Route path="/transactions" element={<ProtectedRoute><TransactionListPage /></ProtectedRoute>} />
      <Route path="/transactions/:id" element={<ProtectedRoute><TransactionDetailPage /></ProtectedRoute>} />
      <Route path="/reviews" element={<ProtectedRoute><ReviewsPage /></ProtectedRoute>} />
      <Route path="/reviews/new/:txId" element={<ProtectedRoute><NewReviewPage /></ProtectedRoute>} />
      <Route path="/notifications" element={<ProtectedRoute><NotificationsPage /></ProtectedRoute>} />

      <Route path="/admin" element={<ProtectedRoute role="admin"><AdminDashboardPage /></ProtectedRoute>} />
      <Route path="/admin/companies" element={<ProtectedRoute role="admin"><AdminCompaniesPage /></ProtectedRoute>} />
      <Route path="/admin/categories" element={<ProtectedRoute role="admin"><AdminCategoriesPage /></ProtectedRoute>} />
      <Route path="/admin/listings" element={<ProtectedRoute role="admin"><AdminListingsPage /></ProtectedRoute>} />
      <Route path="/admin/transactions" element={<ProtectedRoute role="admin"><AdminTransactionsPage /></ProtectedRoute>} />
      <Route path="/admin/finance" element={<ProtectedRoute role="admin"><AdminFinancePage /></ProtectedRoute>} />
      <Route path="/admin/reports" element={<ProtectedRoute role="admin"><AdminReportsPage /></ProtectedRoute>} />
      <Route path="/admin/export" element={<ProtectedRoute role="admin"><AdminExportPage /></ProtectedRoute>} />
    </Routes>
  )
}
