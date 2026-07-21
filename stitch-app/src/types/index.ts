export interface User {
  id: string
  name: string
  email: string
  phone: string
  role: 'business' | 'admin'
  avatar: string
  companyId?: string
}

export interface Company {
  id: string
  name: string
  taxCode: string
  address: string
  city: string
  description: string
  status: 'draft' | 'pending' | 'verified' | 'rejected'
  rejectReason?: string
  ownerId: string
  rating: number
  reviewCount: number
  memberSince: string
  certifications: string[]
}

export interface MaterialCategory {
  id: string
  name: string
  icon: string
}

export interface SupplyListing {
  id: string
  title: string
  categoryId: string
  sellerId: string
  companyId: string
  description: string
  specs: Record<string, string>
  quantity: number
  unit: string
  pricePerUnit: number
  currency: string
  location: string
  minOrderQuantity: number
  packaging: string
  status: 'active' | 'pending_review' | 'sold' | 'hidden'
  images: string[]
  createdAt: string
}

export interface DemandListing {
  id: string
  title: string
  categoryId: string
  buyerId: string
  companyId: string
  description: string
  quantity: number
  unit: string
  targetPrice?: number
  location: string
  deadline: string
  status: 'open' | 'closed' | 'matched'
  createdAt: string
}

export type OfferStatus = 'pending' | 'accepted' | 'rejected' | 'cancelled' | 'expired'

export interface PurchaseOffer {
  id: string
  type: 'buyer_to_seller' | 'seller_to_buyer'
  listingId: string
  listingTitle: string
  buyerId: string
  buyerName: string
  sellerId: string
  sellerName: string
  quantity: number
  unit: string
  proposedPrice: number
  currency: string
  message: string
  status: OfferStatus
  createdAt: string
}

export type TransactionStatus =
  | 'confirmed'
  | 'in_progress'
  | 'buyer_confirmed'
  | 'seller_confirmed'
  | 'completed'
  | 'cancelled'
  | 'disputed'

export interface Transaction {
  id: string
  offerId: string
  listingTitle: string
  buyerId: string
  buyerName: string
  sellerId: string
  sellerName: string
  quantity: number
  unit: string
  agreedPrice: number
  currency: string
  paymentStatus: 'not_required' | 'manual_offline' | 'bypassed_demo'
  paymentMethod: string
  settlementNote: string
  status: TransactionStatus
  createdAt: string
  events: TransactionEvent[]
}

export interface TransactionEvent {
  id: string
  actorId: string
  actorName: string
  fromStatus: string
  toStatus: string
  note: string
  createdAt: string
}

export interface Review {
  id: string
  transactionId: string
  reviewerId: string
  reviewerName: string
  revieweeId: string
  revieweeName: string
  rating: number
  comment: string
  createdAt: string
}

export interface Notification {
  id: string
  userId: string
  title: string
  message: string
  type: 'offer' | 'transaction' | 'system' | 'review'
  read: boolean
  createdAt: string
}

export interface Report {
  id: string
  reporterId: string
  reporterName: string
  targetType: 'listing' | 'company' | 'user' | 'transaction'
  targetId: string
  targetName: string
  reason: string
  description: string
  status: 'pending' | 'reviewed' | 'resolved' | 'dismissed'
  createdAt: string
}
