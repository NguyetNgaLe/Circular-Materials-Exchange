import type { User, Company, MaterialCategory, SupplyListing, DemandListing, PurchaseOffer, Transaction, Review, Notification, Report } from '../types'

export const users: User[] = [
  { id: 'u1', name: 'Nguyễn Văn An', email: 'an@ecopoly.vn', phone: '0901234567', role: 'business', avatar: '', companyId: 'c1' },
  { id: 'u2', name: 'Trần Thị Bình', email: 'binh@greenpack.vn', phone: '0912345678', role: 'business', avatar: '', companyId: 'c2' },
  { id: 'u3', name: 'Lê Admin', email: 'admin@cme.vn', phone: '0900000000', role: 'admin', avatar: '' },
  { id: 'u4', name: 'Phạm Minh Châu', email: 'chau@recyclehub.vn', phone: '0923456789', role: 'business', avatar: '', companyId: 'c3' },
]

export const companies: Company[] = [
  {
    id: 'c1', name: 'EcoPoly Solutions', taxCode: '0123456789', address: '123 Đường Lê Lợi, Q.1', city: 'TP. Hồ Chí Minh',
    description: 'Chuyên cung cấp nhựa tái chế công nghiệp chất lượng cao. Nhà máy đạt chuẩn ISO 9001.',
    status: 'verified', ownerId: 'u1', rating: 4.7, reviewCount: 23, memberSince: '2023-06-15',
    certifications: ['ISO 9001', 'Chứng nhận GRS'],
  },
  {
    id: 'c2', name: 'GreenPack Việt Nam', taxCode: '9876543210', address: '456 Đường Nguyễn Huệ, Q.3', city: 'TP. Hồ Chí Minh',
    description: 'Doanh nghiệp sản xuất bao bì thân thiện môi trường. Tìm kiếm nguyên liệu tái chế.',
    status: 'verified', ownerId: 'u2', rating: 4.5, reviewCount: 15, memberSince: '2023-09-01',
    certifications: ['ISO 14001'],
  },
  {
    id: 'c3', name: 'RecycleHub VN', taxCode: '1122334455', address: '789 Đường Võ Văn Tần, Q.3', city: 'TP. Hồ Chí Minh',
    description: 'Nền tảng thu gom và xử lý phế liệu công nghiệp.',
    status: 'pending', ownerId: 'u4', rating: 0, reviewCount: 0, memberSince: '2024-01-10',
    certifications: [],
  },
]

export const categories: MaterialCategory[] = [
  { id: 'cat1', name: 'Nhựa', icon: 'recycling' },
  { id: 'cat2', name: 'Kim Loại', icon: 'hardware' },
  { id: 'cat3', name: 'Giấy & Bìa Cứng', icon: 'description' },
  { id: 'cat4', name: 'Gỗ', icon: 'forest' },
  { id: 'cat5', name: 'Dệt May', icon: 'checkroom' },
  { id: 'cat6', name: 'Thủy Tinh', icon: 'local_drink' },
]

export const supplyListings: SupplyListing[] = [
  {
    id: 'sl1', title: 'Nhựa PET tái chế', categoryId: 'cat1', sellerId: 'u1', companyId: 'c1',
    description: 'Nhựa PET tái chế sau công nghiệp, đã rửa sạch và ép hạt. Phù hợp sản xuất chai, hộp.',
    specs: { 'Độ tinh khiết': '> 99%', 'Màu sắc': 'Trong suốt', 'Dạng': 'Hạt' },
    quantity: 15, unit: 'Tấn', pricePerUnit: 12000, currency: 'VND',
    location: 'Bình Dương', minOrderQuantity: 1, packaging: 'Bao 25kg',
    status: 'active', images: [], createdAt: '2024-11-01',
  },
  {
    id: 'sl2', title: 'Sắt vụn công nghiệp', categoryId: 'cat2', sellerId: 'u1', companyId: 'c1',
    description: 'Sắt vụn từ gia công cơ khí, đã phân loại. Phù hợp nấu luyện thép.',
    specs: { 'Loại': 'Sắt carbon', 'Độ sạch': '> 95%', 'Kích thước': '5-50mm' },
    quantity: 42, unit: 'Tấn', pricePerUnit: 8500, currency: 'VND',
    location: 'TP. Hồ Chí Minh', minOrderQuantity: 5, packaging: 'Đống bulk',
    status: 'active', images: [], createdAt: '2024-10-28',
  },
  {
    id: 'sl3', title: 'Thùng Carton cũ', categoryId: 'cat3', sellerId: 'u4', companyId: 'c3',
    description: 'Carton sóng 3 lớp, đã ép gọn. Phù hợp tái chế giấy.',
    specs: { 'Loại': 'Sóng BC', 'Độ ẩm': '< 8%', 'Nén': 'Đã ép' },
    quantity: 5, unit: 'Tấn', pricePerUnit: 3200, currency: 'VND',
    location: 'Đồng Nai', minOrderQuantity: 1, packaging: 'Ép kiện 1 tấn',
    status: 'active', images: [], createdAt: '2024-11-05',
  },
  {
    id: 'sl4', title: 'Nhôm phế liệu loại 1', categoryId: 'cat2', sellerId: 'u4', companyId: 'c3',
    description: 'Nhôm từ khuôn ép, không lẫn tạp chất. Phù hợp nấu nhôm hợp kim.',
    specs: { 'Loại': 'Nhôm 6061', 'Độ sạch': '> 98%', 'Dạng': 'Mảnh' },
    quantity: 8, unit: 'Tấn', pricePerUnit: 35000, currency: 'VND',
    location: 'Long An', minOrderQuantity: 1, packaging: 'Đóng bao 500kg',
    status: 'active', images: [], createdAt: '2024-11-03',
  },
  {
    id: 'sl5', title: 'Pallet gỗ thông cũ', categoryId: 'cat4', sellerId: 'u1', companyId: 'c1',
    description: 'Pallet gỗ thông đã qua sử dụng, còn chắc chắn. Có thể tái sử dụng hoặc nghiền dăm.',
    specs: { 'Kích thước': '1200x1000mm', 'Tình trạng': 'Còn sử dụng được', 'Loại gỗ': 'Thông' },
    quantity: 500, unit: 'Cái', pricePerUnit: 45000, currency: 'VND',
    location: 'Bình Dương', minOrderQuantity: 50, packaging: 'Xếp chồng',
    status: 'active', images: [], createdAt: '2024-10-25',
  },
  {
    id: 'sl6', title: 'Hạt nhựa HDPE tái sinh', categoryId: 'cat1', sellerId: 'u1', companyId: 'c1',
    description: 'HDPE tái chế sau công nghiệp, tạo hạt đồng đều. Thích hợp ép phun và ép đùn.',
    specs: { 'Độ tinh khiết': '> 99.5%', 'MFI': '0.4-0.8 g/10min', 'Mật độ': '0.94-0.96 g/cm³' },
    quantity: 20, unit: 'Tấn', pricePerUnit: 18500, currency: 'VND',
    location: 'TP. Hồ Chí Minh', minOrderQuantity: 5, packaging: 'Bao 25kg',
    status: 'active', images: [], createdAt: '2024-11-08',
  },
]

export const demandListings: DemandListing[] = [
  {
    id: 'dl1', title: 'Cần mua nhựa PP tái chế', categoryId: 'cat1', buyerId: 'u2', companyId: 'c2',
    description: 'Cần nhựa PP tái chế trắng, dùng sản xuất bao bì thực phẩm. Yêu cầu có chứng nhận FDA.',
    quantity: 10, unit: 'Tấn', targetPrice: 15000, location: 'TP. Hồ Chí Minh',
    deadline: '2025-02-01', status: 'open', createdAt: '2024-11-10',
  },
  {
    id: 'dl2', title: 'Tìm carton cuộn tái chế', categoryId: 'cat3', buyerId: 'u2', companyId: 'c2',
    description: 'Carton cuộn sóng, dùng đóng gói sản phẩm. Cần giao hàng tuần.',
    quantity: 3, unit: 'Tấn', targetPrice: 4000, location: 'Bình Dương',
    deadline: '2025-01-15', status: 'open', createdAt: '2024-11-12',
  },
]

export const offers: PurchaseOffer[] = [
  {
    id: 'of1', type: 'buyer_to_seller', listingId: 'sl6', listingTitle: 'Hạt nhựa HDPE tái sinh',
    buyerId: 'u2', buyerName: 'Trần Thị Bình', sellerId: 'u1', sellerName: 'Nguyễn Văn An',
    quantity: 10, unit: 'Tấn', proposedPrice: 17000, currency: 'VND',
    message: 'Chúng tôi cần 10 tấn cho đơn hàng đầu tiên. Mong được hợp tác lâu dài.',
    status: 'pending', createdAt: '2024-11-15',
  },
  {
    id: 'of2', type: 'buyer_to_seller', listingId: 'sl1', listingTitle: 'Nhựa PET tái chế',
    buyerId: 'u2', buyerName: 'Trần Thị Bình', sellerId: 'u1', sellerName: 'Nguyễn Văn An',
    quantity: 5, unit: 'Tấn', proposedPrice: 11000, currency: 'VND',
    message: 'Muốn mua 5 tấn PET, giá có thể thương lượng.',
    status: 'accepted', createdAt: '2024-11-10',
  },
  {
    id: 'of3', type: 'buyer_to_seller', listingId: 'sl4', listingTitle: 'Nhôm phế liệu loại 1',
    buyerId: 'u2', buyerName: 'Trần Thị Bình', sellerId: 'u4', sellerName: 'Phạm Minh Châu',
    quantity: 2, unit: 'Tấn', proposedPrice: 33000, currency: 'VND',
    message: 'Cần nhôm cho sản xuất linh kiện điện tử.',
    status: 'accepted', createdAt: '2024-11-08',
  },
]

export const transactions: Transaction[] = [
  {
    id: 'tx1', offerId: 'of2', listingTitle: 'Nhựa PET tái chế',
    buyerId: 'u2', buyerName: 'Trần Thị Bình', sellerId: 'u1', sellerName: 'Nguyễn Văn An',
    quantity: 5, unit: 'Tấn', agreedPrice: 11000, currency: 'VND',
    paymentStatus: 'bypassed_demo', paymentMethod: 'manual_offline',
    settlementNote: 'Thanh toán được thực hiện ngoài hệ thống trong phạm vi prototype',
    status: 'in_progress', createdAt: '2024-11-11',
    events: [
      { id: 'ev1', actorId: 'u1', actorName: 'Nguyễn Văn An', fromStatus: 'offer.accepted', toStatus: 'transaction.confirmed', note: 'Giao dịch được tạo tự động khi seller chấp nhận offer', createdAt: '2024-11-11T08:00:00Z' },
      { id: 'ev2', actorId: 'u2', actorName: 'Trần Thị Bình', fromStatus: 'transaction.confirmed', toStatus: 'transaction.in_progress', note: 'Xác nhận thỏa thuận giao dịch', createdAt: '2024-11-11T09:30:00Z' },
    ],
  },
  {
    id: 'tx2', offerId: 'of3', listingTitle: 'Nhôm phế liệu loại 1',
    buyerId: 'u2', buyerName: 'Trần Thị Bình', sellerId: 'u4', sellerName: 'Phạm Minh Châu',
    quantity: 2, unit: 'Tấn', agreedPrice: 33000, currency: 'VND',
    paymentStatus: 'bypassed_demo', paymentMethod: 'manual_offline',
    settlementNote: 'Thanh toán được thực hiện ngoài hệ thống trong phạm vi prototype',
    status: 'completed', createdAt: '2024-11-09',
    events: [
      { id: 'ev3', actorId: 'u4', actorName: 'Phạm Minh Châu', fromStatus: 'offer.accepted', toStatus: 'transaction.confirmed', note: 'Giao dịch được tạo tự động', createdAt: '2024-11-09T10:00:00Z' },
      { id: 'ev4', actorId: 'u2', actorName: 'Trần Thị Bình', fromStatus: 'transaction.confirmed', toStatus: 'transaction.in_progress', note: 'Bắt đầu giao dịch', createdAt: '2024-11-09T11:00:00Z' },
      { id: 'ev5', actorId: 'u2', actorName: 'Trần Thị Bình', fromStatus: 'transaction.in_progress', toStatus: 'transaction.buyer_confirmed', note: 'Đã nhận hàng, chất lượng đạt yêu cầu', createdAt: '2024-11-12T14:00:00Z' },
      { id: 'ev6', actorId: 'u4', actorName: 'Phạm Minh Châu', fromStatus: 'transaction.buyer_confirmed', toStatus: 'transaction.completed', note: 'Xác nhận hoàn tất giao dịch', createdAt: '2024-11-12T15:00:00Z' },
    ],
  },
]

export const reviews: Review[] = [
  {
    id: 'rv1', transactionId: 'tx2', reviewerId: 'u2', reviewerName: 'Trần Thị Bình',
    revieweeId: 'u4', revieweeName: 'Phạm Minh Châu', rating: 5,
    comment: 'Giao hàng đúng hẹn, nhôm chất lượng tốt. Sẽ hợp tác lại.',
    createdAt: '2024-11-13',
  },
]

export const notifications: Notification[] = [
  { id: 'n1', userId: 'u1', title: 'Offer mới', message: 'Trần Thị Bình gửi đề nghị mua Hạt nhựa HDPE tái sinh', type: 'offer', read: false, createdAt: '2024-11-15T10:00:00Z' },
  { id: 'n2', userId: 'u2', title: 'Offer được chấp nhận', message: 'Nguyễn Văn An đã chấp nhận đề nghị mua Nhựa PET tái chế', type: 'offer', read: false, createdAt: '2024-11-10T14:00:00Z' },
  { id: 'n3', userId: 'u2', title: 'Giao dịch hoàn tất', message: 'Giao dịch Nhôm phế liệu loại 1 đã hoàn tất', type: 'transaction', read: true, createdAt: '2024-11-12T15:00:00Z' },
  { id: 'n4', userId: 'u1', title: 'Yêu cầu xác minh doanh nghiệp', message: 'RecycleHub VN đã gửi yêu cầu xác minh', type: 'system', read: false, createdAt: '2024-11-10T09:00:00Z' },
]

export const reports: Report[] = [
  {
    id: 'rp1', reporterId: 'u2', reporterName: 'Trần Thị Bình', targetType: 'listing',
    targetId: 'sl3', targetName: 'Thùng Carton cũ', reason: 'Thông tin không chính xác',
    description: 'Khối lượng thực tế ít hơn quảng cáo', status: 'pending', createdAt: '2024-11-14',
  },
]
