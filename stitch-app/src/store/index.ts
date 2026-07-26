import { useState, useCallback } from 'react'
import type { User, Company, SupplyListing, PurchaseOffer, Transaction, Notification } from '../types'
import { api } from '../services/api'

let _currentUser: User | null = null
let _listeners: Array<() => void> = []

function notify() { _listeners.forEach(l => l()) }

const savedUser = localStorage.getItem('user')
const savedToken = localStorage.getItem('token')
if (savedUser) {
  try { _currentUser = JSON.parse(savedUser) } catch {}
}
if (savedToken) api.setToken(savedToken)

export function useStore() {
  const [, setTick] = useState(0)
  const rerender = useCallback(() => setTick((t: number) => t + 1), [])

  if (!_listeners.includes(rerender)) {
    _listeners.push(rerender)
  }

  return {
    get currentUser() { return _currentUser },

    async login(email: string, password: string): Promise<boolean> {
      try {
        const res = await api.login(email, password)
        if (res.success) {
          api.setToken(res.data.token)
          _currentUser = res.data.user
          localStorage.setItem('user', JSON.stringify(_currentUser))
          notify()
          return true
        }
      } catch {}
      return false
    },

    logout() {
      _currentUser = null
      api.clearToken()
      localStorage.removeItem('user')
      localStorage.removeItem('token')
      notify()
    },

    async register(name: string, email: string, password: string): Promise<boolean> {
      try {
        const res = await api.register(name, email, password)
        if (res.success) {
          api.setToken(res.data.token)
          _currentUser = res.data.user
          localStorage.setItem('user', JSON.stringify(_currentUser))
          notify()
          return true
        }
      } catch {}
      return false
    },

    async getListings(params?: { category_id?: string; keyword?: string }) {
      try {
        const res = await api.getListings(params)
        return res.success ? res.data.listings || [] : []
      } catch { return [] }
    },

    async getListing(id: string) {
      try {
        const res = await api.getListing(id)
        return res.success ? res.data : null
      } catch { return null }
    },

    async createListing(data: any) {
      try {
        const res = await api.createListing(data)
        return res.success ? res.data : null
      } catch { return null }
    },

    async deleteListing(id: string): Promise<boolean> {
      try {
        const res = await api.deleteListing(id)
        return res.success === true
      } catch {
        return false
      }
    },

    async getDemands(params?: { category_id?: string; keyword?: string }) {
      try {
        const res = await api.getDemands(params)
        return res.success ? res.data.demands || [] : []
      } catch { return [] }
    },

    async createDemand(data: any) {
      try {
        const res = await api.createDemand(data)
        return res.success ? res.data : null
      } catch { return null }
    },

    async getCategories() {
      try {
        const res = await api.getCategories()
        return res.success ? res.data : []
      } catch { return [] }
    },

    async getOffers(params?: { role?: string; status?: string }) {
      try {
        const res = await api.getOffers(params)
        return res.success ? res.data.offers || [] : []
      } catch { return [] }
    },

    async createOffer(data: any) {
      try {
        const res = await api.createOffer(data)
        return res.success ? res.data : null
      } catch { return null }
    },

    async acceptOffer(id: string) {
      try {
        const res = await api.acceptOffer(id)
        return res.success ? res.data : null
      } catch { return null }
    },

    async rejectOffer(id: string) {
      try {
        const res = await api.rejectOffer(id)
        return res.success ? res.data : null
      } catch { return null }
    },

    async getTransactions(params?: { status?: string }) {
      try {
        const res = await api.getTransactions(params)
        return res.success ? res.data.transactions || [] : []
      } catch { return [] }
    },

    async getTransaction(id: string) {
      try {
        const res = await api.getTransaction(id)
        return res.success ? res.data : null
      } catch { return null }
    },

    async updateTransactionStatus(id: string, status: string, note?: string) {
      try {
        const res = await api.updateTransactionStatus(id, status, note)
        return res.success ? res.data : null
      } catch { return null }
    },

    async getCompanies(params?: { status?: string }) {
      try {
        const res = await api.getCompanies(params)
        return res.success ? res.data.companies || [] : []
      } catch { return [] }
    },

    async getCompany(id: string) {
      try {
        const res = await api.getCompany(id)
        return res.success ? res.data : null
      } catch { return null }
    },

    async createCompany(data: any) {
      try {
        const res = await api.createCompany(data)
        return res.success ? res.data : null
      } catch { return null }
    },

    async approveCompany(id: string) {
      try {
        const res = await api.approveCompany(id)
        return res.success
      } catch { return false }
    },

    async rejectCompany(id: string, reason: string) {
      try {
        const res = await api.rejectCompany(id, reason)
        return res.success
      } catch { return false }
    },

    async getReviews(params?: { reviewee_id?: string }) {
      try {
        const res = await api.getReviews(params)
        return res.success ? res.data.reviews || [] : []
      } catch { return [] }
    },

    async createReview(data: any) {
      try {
        const res = await api.createReview(data)
        return res.success ? res.data : null
      } catch { return null }
    },

    async getNotifications() {
      try {
        const res = await api.getNotifications()
        return res.success ? res.data.notifications || [] : []
      } catch { return [] }
    },

    async getUserNotifications() {
      return this.getNotifications()
    },

    async markNotificationRead(id: string) {
      try {
        await api.markNotificationRead(id)
      } catch {}
    },

    async markAllNotificationsRead() {
      try {
        await api.markAllNotificationsRead()
      } catch {}
    },

    async getUnreadCount() {
      try {
        const res = await api.getUnreadCount()
        return res.success ? res.data.count : 0
      } catch { return 0 }
    },
  }
}
