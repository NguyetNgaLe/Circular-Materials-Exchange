const state = {
  role: "buyer",
  lang: "vi",
  view: "dashboard",
  selectedListing: "SUP-1001",
  selectedDemand: "DEM-3001",
  selectedTransaction: "TX-9001",
  selectedReport: "RPT-7001",
};

const roles = {
  guest: {
    label: "Guest",
    name: "Khách truy cập",
    company: "Chưa đăng nhập",
    subtitle: "Public marketplace",
  },
  buyer: {
    label: "Buyer",
    name: "Nguyễn Minh Anh",
    company: "EcoPack Vietnam",
    subtitle: "Buyer · verified",
  },
  seller: {
    label: "Seller",
    name: "Trần Quốc Bảo",
    company: "GreenTech Manufacturing",
    subtitle: "Seller · verified",
  },
  admin: {
    label: "Admin",
    name: "System Admin",
    company: "Circular Materials Exchange",
    subtitle: "Superuser",
  },
};

const navItems = [
  ["dashboard", "Dashboard"],
  ["auth", "Auth & bảo mật"],
  ["company", "Hồ sơ doanh nghiệp"],
  ["catalog", "Danh mục vật liệu"],
  ["marketplace", "Marketplace"],
  ["supply", "Nguồn cung"],
  ["demand", "Nhu cầu mua"],
  ["offers", "Offer / RFQ"],
  ["transactions", "Giao dịch"],
  ["reviews", "Đánh giá"],
  ["reports", "Báo cáo vi phạm"],
  ["notifications", "Thông báo"],
  ["admin", "Admin"],
  ["exports", "Báo cáo & export"],
  ["usecases", "Use case matrix"],
];

const roleDefaultView = {
  guest: "auth",
  buyer: "dashboard",
  seller: "dashboard",
  admin: "admin",
};

const staticI18n = {
  en: {
    "Circular Materials": "Circular Materials",
    "Circular Materials Exchange": "Circular Materials Exchange",
    "Guest": "Guest",
    "Buyer": "Buyer",
    "Seller": "Seller",
    "Admin": "Admin",
    "Khách truy cập": "Guest user",
    "Chưa đăng nhập": "Not signed in",
    "Dashboard": "Dashboard",
    "Auth & bảo mật": "Auth & security",
    "Hồ sơ doanh nghiệp": "Company profile",
    "Danh mục vật liệu": "Material catalog",
    "Marketplace": "Marketplace",
    "Nguồn cung": "Supply listings",
    "Nhu cầu mua": "Demand listings",
    "Offer / RFQ": "Offers / RFQ",
    "Giao dịch": "Transactions",
    "Đánh giá": "Reviews",
    "Báo cáo vi phạm": "Violation reports",
    "Thông báo": "Notifications",
    "Báo cáo & export": "Reports & export",
    "Use case matrix": "Use case matrix",
    "UC054-UC055 · Dashboard": "UC054-UC055 · Dashboard",
    "Dashboard doanh nghiệp": "Business dashboard",
    "Dashboard quản trị hệ thống": "Admin dashboard",
    "Số liệu mock thể hiện nguồn cung, nhu cầu, offer, giao dịch và uy tín doanh nghiệp.": "Mock metrics show supply, demand, offers, transactions, and company reputation.",
    "Đi marketplace": "Go to marketplace",
    "Xuất báo cáo": "Export report",
    "Nguồn cung active": "Active supply",
    "Nhu cầu active": "Active demand",
    "Offer pending": "Pending offers",
    "Giao dịch hoàn tất": "Completed transactions",
    "Hoạt động gần đây": "Recent activity",
    "Hồ sơ hiện tại": "Current profile",
    "Actor": "Actor",
    "Người dùng": "User",
    "Doanh nghiệp": "Company",
    "Điểm uy tín": "Reputation score",
    "UC001-UC006 · Auth": "UC001-UC006 · Auth",
    "Đăng nhập, đăng ký và bảo mật": "Sign in, registration, and security",
    "Mô phỏng đầy đủ đăng ký, OTP, đăng nhập, quên mật khẩu, đổi mật khẩu và đăng xuất.": "Full mock flow for registration, OTP, sign-in, password reset, password change, and sign-out.",
    "Đăng xuất demo": "Demo sign out",
    "Đăng nhập buyer": "Sign in as buyer",
    "Đăng nhập": "Sign in",
    "Email hoặc số điện thoại": "Email or phone",
    "Mật khẩu": "Password",
    "Đăng ký tài khoản": "Create account",
    "Họ tên": "Full name",
    "Số điện thoại": "Phone number",
    "Email": "Email",
    "Loại tài khoản": "Account type",
    "Đại diện doanh nghiệp": "Company representative",
    "OTP demo:": "Demo OTP:",
    "Tài khoản mới ở trạng thái pending_verification.": "New accounts stay in pending_verification.",
    "Quên / đổi mật khẩu": "Forgot / change password",
    "Email nhận mã": "Recovery email",
    "Mật khẩu mới": "New password",
    "Xác nhận OTP": "Confirm OTP",
    "Cập nhật mật khẩu": "Update password",
    "UC007-UC013 · Company": "UC007-UC013 · Company",
    "Hồ sơ cá nhân và doanh nghiệp": "Personal and company profile",
    "Bao gồm cập nhật hồ sơ cá nhân, tạo/cập nhật doanh nghiệp, gửi xác minh, thành viên và phân quyền.": "Includes personal profile updates, company create/update, verification submission, members, and permissions.",
    "Gửi xác minh": "Submit verification",
    "Admin duyệt": "Admin approval",
    "Tên pháp lý": "Legal name",
    "Mã số thuế": "Tax code",
    "Lĩnh vực": "Industry",
    "Tỉnh/thành": "Province / city",
    "Email liên hệ": "Contact email",
    "Địa chỉ": "Address",
    "Mô tả": "Description",
    "Vai trò": "Role",
    "Bảo mật": "Security",
    "OTP email demo": "Demo email OTP",
    "Thành viên & phân quyền": "Members & permissions",
    "Mời thành viên": "Invite member",
    "Thành viên": "Member",
    "Trạng thái": "Status",
    "UC014-UC017 · Material catalog": "UC014-UC017 · Material catalog",
    "Danh mục vật liệu": "Material catalog",
    "Admin chuẩn hóa vật liệu, buyer/seller dùng để lọc marketplace và tạo listing/demand.": "Admin standardizes materials; buyers/sellers use them to filter marketplace and create listings/demands.",
    "Thêm vật liệu": "Add material",
    "Danh mục": "Category",
    "Vật liệu": "Material",
    "Đơn vị": "Unit",
    "Mã": "Code",
    "Tạo/cập nhật vật liệu": "Create/update material",
    "Tên vật liệu": "Material name",
    "Mức nguy hại": "Hazard level",
    "Lưu danh mục": "Save catalog item",
    "UC022, UC029-UC033 · Marketplace": "UC022, UC029-UC033 · Marketplace",
    "Sàn giao dịch nguồn cung vật liệu": "Supply marketplace",
    "Buyer tìm kiếm, lọc, sắp xếp, xem seller và gửi đề nghị mua từ listing.": "Buyers search, filter, sort, inspect sellers, and submit purchase offers from listings.",
    "Gửi đề nghị mua": "Send purchase offer",
    "Tất cả danh mục": "All categories",
    "Tất cả địa điểm": "All locations",
    "Sắp xếp: mới nhất": "Sort: newest",
    "Giá thấp đến cao": "Price: low to high",
    "Khối lượng lớn nhất": "Largest volume",
    "Khối lượng còn lại": "Available quantity",
    "Giá niêm yết": "Listed price",
    "Xem chi tiết": "View details",
    "Chi tiết nguồn cung": "Supply detail",
    "Mã listing": "Listing ID",
    "Chất lượng": "Quality",
    "Hạn khả dụng": "Available until",
    "Uy tín seller": "Seller reputation",
    "Lưu quan tâm": "Save listing",
    "UC018-UC023 · Supply listing": "UC018-UC023 · Supply listing",
    "Quản lý nguồn cung vật liệu": "Supply listing management",
    "Seller đăng, sửa, ẩn/đóng listing; admin có thể kiểm duyệt pending_review.": "Sellers create, edit, hide/close listings; admins can moderate pending_review listings.",
    "Đăng nguồn cung": "Post supply",
    "Khối lượng": "Quantity",
    "Giá": "Price",
    "Hành động": "Actions",
    "Sửa": "Edit",
    "Đóng": "Close",
    "Form đăng nguồn cung": "Supply form",
    "Số lượng": "Quantity",
    "Giá/kg": "Price/kg",
    "Tài liệu đính kèm": "Attachments",
    "Lưu draft / đăng bán": "Save draft / publish",
    "UC024-UC028, UC058-UC059 · Demand": "UC024-UC028, UC058-UC059 · Demand",
    "Sàn nhu cầu mua vật liệu": "Material demand marketplace",
    "Buyer đăng nhu cầu; seller gửi báo giá seller_to_buyer để khép kín luồng demand.": "Buyers post demands; sellers send seller_to_buyer quotes to complete the demand flow.",
    "Đăng nhu cầu mua": "Post demand",
    "Số lượng cần": "Required quantity",
    "Ngân sách": "Budget",
    "Báo giá nhận được": "Quotes received",
    "Xem nhu cầu": "View demand",
    "Seller gửi báo giá": "Seller quote",
    "Cần mua": "Needs",
    "Yêu cầu": "Requirements",
    "Khối lượng cung cấp": "Offered quantity",
    "Đơn giá báo": "Quoted unit price",
    "Ghi chú": "Note",
    "Gửi báo giá": "Send quote",
    "UC034-UC038, UC058-UC059 · Offer/RFQ": "UC034-UC038, UC058-UC059 · Offers/RFQ",
    "Quản lý đề nghị mua và báo giá": "Offer and quote management",
    "Bao gồm buyer gửi offer, seller xử lý offer, buyer hủy pending offer và seller gửi báo giá cho demand.": "Includes buyer offers, seller processing, buyer cancellation of pending offers, and seller quotes for demands.",
    "Accept OFF-5001 demo": "Accept OFF-5001 demo",
    "Form gửi đề nghị mua": "Purchase offer form",
    "Ngày nhận": "Expected receipt date",
    "Lời nhắn": "Message",
    "Lịch sử đàm phán": "Negotiation history",
    "Buyer gửi offer": "Buyer sent offer",
    "Seller chưa xử lý": "Seller has not processed",
    "Chi tiết": "Details",
    "Hủy/Từ chối": "Cancel/Reject",
    "UC039-UC043 · Transactions": "UC039-UC043 · Transactions",
    "Giao dịch được tạo tự động sau khi seller chấp nhận offer; thanh toán được bypass/manual offline trong MVP.": "A transaction is created automatically after seller accepts an offer; payment is bypassed/manual offline in the MVP.",
    "Yêu cầu hủy": "Request cancellation",
    "Xác nhận hoàn tất": "Confirm completion",
    "Giá cuối": "Final price",
    "Timeline trạng thái": "Status timeline",
    "Chi tiết giao dịch": "Transaction details",
    "Không gọi payment gateway. payment_status chỉ ghi nhận manual_offline hoặc bypassed_demo.": "No payment gateway is called. payment_status only records manual_offline or bypassed_demo.",
    "UC044-UC046 · Reviews": "UC044-UC046 · Reviews",
    "Đánh giá và điểm uy tín": "Reviews and reputation",
    "Buyer và seller đánh giá nhau sau khi giao dịch completed; điểm uy tín hiển thị ở hồ sơ doanh nghiệp.": "Buyer and seller review each other after a completed transaction; reputation is shown on company profiles.",
    "Gửi đánh giá": "Submit review",
    "Form đánh giá đối tác": "Partner review form",
    "Đối tác": "Partner",
    "Điểm vật liệu đúng mô tả": "Material matched description score",
    "Thời gian phản hồi": "Response time",
    "Nhận xét": "Comment",
    "Uy tín doanh nghiệp": "Company reputation",
    "UC047-UC048 · Violation reports": "UC047-UC048 · Violation reports",
    "Báo cáo vi phạm và moderation": "Violation reports and moderation",
    "Business user tạo report; admin xử lý bằng cảnh cáo, ẩn listing, khóa user/company hoặc reject report.": "Business users create reports; admins handle them with warning, hiding listing, locking user/company, or rejecting reports.",
    "Gửi report": "Submit report",
    "Form báo cáo vi phạm": "Violation report form",
    "Đối tượng": "Target",
    "Mã đối tượng": "Target ID",
    "Loại vi phạm": "Violation type",
    "Bằng chứng": "Evidence",
    "Admin xử lý": "Admin handling",
    "Lý do": "Reason",
    "Bỏ qua": "Ignore",
    "Ẩn listing": "Hide listing",
    "Khóa công ty": "Lock company",
    "UC049-UC053 · Notifications": "UC049-UC053 · Notifications",
    "Thông báo in-app": "In-app notifications",
    "Thông báo offer mới, kết quả offer, trạng thái giao dịch, report và đánh dấu đã đọc.": "Notifications for new offers, offer results, transaction status, reports, and mark-as-read.",
    "Đánh dấu đã đọc": "Mark as read",
    "Admin · UC011, UC015-UC017, UC023, UC048, UC055": "Admin · UC011, UC015-UC017, UC023, UC048, UC055",
    "Trang quản trị hệ thống": "System admin console",
    "Quản trị user, doanh nghiệp, danh mục vật liệu, listing, demand, offer, transaction, report và review.": "Manage users, companies, material catalog, listings, demands, offers, transactions, reports, and reviews.",
    "Reports": "Reports",
    "Tin chờ duyệt": "Listings pending review",
    "Report pending": "Pending reports",
    "Hồ sơ doanh nghiệp chờ duyệt": "Company verification queue",
    "Request": "Request",
    "Company": "Company",
    "Submitted": "Submitted",
    "Docs": "Docs",
    "Action": "Action",
    "Duyệt": "Approve",
    "Từ chối": "Reject",
    "Moderation nhanh": "Quick moderation",
    "Listing chờ duyệt": "Listings pending review",
    "Material restricted": "Restricted material",
    "Report xử lý": "Report in handling",
    "Export": "Export",
    "Last login": "Last login",
    "UC056-UC057 · Reports/export": "UC056-UC057 · Reports/export",
    "Báo cáo giao dịch và xuất dữ liệu": "Transaction reports and data export",
    "Bộ lọc báo cáo, preview dữ liệu và xuất CSV/PDF demo cho admin hoặc business user.": "Report filters, data preview, and CSV/PDF demo export for admins or business users.",
    "Xuất CSV": "Export CSV",
    "Xuất PDF demo": "Export demo PDF",
    "Bộ lọc báo cáo": "Report filters",
    "Khoảng thời gian": "Time range",
    "30 ngày gần nhất": "Last 30 days",
    "Quý này": "This quarter",
    "Năm nay": "This year",
    "Tất cả": "All",
    "Loại vật liệu": "Material type",
    "Định dạng": "Format",
    "UC001-UC059": "UC001-UC059",
    "Coverage · UC001-UC059": "Coverage · UC001-UC059",
    "Ma trận này chứng minh prototype có mock data và màn hình cho toàn bộ use case, bao gồm hai UC bổ sung cho luồng demand.": "This matrix proves the prototype has mock data and screens for all use cases, including two added UCs for the demand flow.",
    "Bắt đầu demo": "Start demo",
    "Màn:": "Screen:",
    "use cases": "use cases"
  }
};

function translateStaticText(text) {
  if (state.lang === "vi") return text;
  const dictionary = staticI18n[state.lang] || {};
  const trimmed = text.trim();
  if (!trimmed) return text;
  const translated = dictionary[trimmed] || translatePrefix(trimmed, dictionary);
  if (!translated) return text;
  return text.replace(trimmed, translated);
}

function translatePrefix(text, dictionary) {
  const prefixes = ["Admin xử lý "];
  for (const prefix of prefixes) {
    if (text.startsWith(prefix) && dictionary[prefix.trim()]) {
      return text.replace(prefix, `${dictionary[prefix.trim()]} `);
    }
  }
  return "";
}

function applyStaticI18n(root) {
  if (state.lang === "vi") return;
  const walker = document.createTreeWalker(root, NodeFilter.SHOW_TEXT, {
    acceptNode(node) {
      const parent = node.parentElement;
      if (!parent) return NodeFilter.FILTER_REJECT;
      if (["SCRIPT", "STYLE", "INPUT", "TEXTAREA"].includes(parent.tagName)) return NodeFilter.FILTER_REJECT;
      if (parent.closest("[data-dynamic]")) return NodeFilter.FILTER_REJECT;
      return NodeFilter.FILTER_ACCEPT;
    },
  });
  const nodes = [];
  while (walker.nextNode()) nodes.push(walker.currentNode);
  nodes.forEach((node) => {
    node.nodeValue = translateStaticText(node.nodeValue);
  });
}

const mock = {
  users: [
    { id: "USR-001", name: "Nguyễn Minh Anh", email: "minhanh@ecopack.vn", phone: "+84 909 234 881", role: "Business Owner", companyId: "COM-002", status: "active", lastLogin: "05/07/2026 14:18" },
    { id: "USR-002", name: "Trần Quốc Bảo", email: "bao@greentech.vn", phone: "+84 908 778 100", role: "Sales Manager", companyId: "COM-001", status: "active", lastLogin: "05/07/2026 13:42" },
    { id: "USR-003", name: "Lê Thanh Huy", email: "huy@metal-loop.vn", phone: "+84 913 040 882", role: "Operations", companyId: "COM-003", status: "pending_verification", lastLogin: "04/07/2026 09:20" },
    { id: "USR-004", name: "System Admin", email: "admin@cme.local", phone: "+84 000 000 000", role: "Admin", companyId: "SYS", status: "active", lastLogin: "05/07/2026 15:00" },
  ],
  companies: [
    { id: "COM-001", name: "GreenTech Manufacturing", taxCode: "0318-456-982", type: "Sản xuất nhựa tái chế", province: "Bình Dương", address: "KCN VSIP II, Thủ Dầu Một", email: "contact@greentech.vn", phone: "+84 274 889 010", status: "verified", rating: 4.8, completed: 36, successRate: "94%" },
    { id: "COM-002", name: "EcoPack Vietnam", taxCode: "0319-288-551", type: "Bao bì công nghiệp", province: "Long An", address: "KCN Tân Đức, Đức Hòa", email: "procurement@ecopack.vn", phone: "+84 272 552 019", status: "verified", rating: 4.6, completed: 28, successRate: "91%" },
    { id: "COM-003", name: "MetalLoop Co.", taxCode: "3602-889-100", type: "Thu gom kim loại", province: "Đồng Nai", address: "KCN Biên Hòa 2", email: "sales@metalloop.vn", phone: "+84 251 883 220", status: "pending_verification", rating: 4.2, completed: 12, successRate: "87%" },
    { id: "COM-004", name: "Saigon Paper Hub", taxCode: "0317-998-442", type: "Giấy và bao bì", province: "TP. Hồ Chí Minh", address: "Quận 7", email: "hello@paperhub.vn", phone: "+84 28 777 1200", status: "rejected", rating: 3.9, completed: 8, successRate: "75%" },
  ],
  companyMembers: [
    { id: "MBR-001", companyId: "COM-002", user: "Nguyễn Minh Anh", role: "Owner", status: "active" },
    { id: "MBR-002", companyId: "COM-002", user: "Phạm Quỳnh Mai", role: "Buyer", status: "active" },
    { id: "MBR-003", companyId: "COM-001", user: "Trần Quốc Bảo", role: "Seller", status: "active" },
    { id: "MBR-004", companyId: "COM-001", user: "Đặng Hoài Nam", role: "Inventory Manager", status: "invited" },
  ],
  verificationRequests: [
    { id: "VRF-4101", companyId: "COM-003", company: "MetalLoop Co.", submittedAt: "05/07/2026 12:04", status: "pending_verification", docs: ["business-license.pdf", "tax-registration.pdf", "factory-address.jpg"] },
    { id: "VRF-4098", companyId: "COM-004", company: "Saigon Paper Hub", submittedAt: "04/07/2026 10:12", status: "rejected", reason: "Thiếu giấy phép kinh doanh hợp lệ", docs: ["tax-code.pdf"] },
  ],
  categories: [
    { id: "CAT-PLASTIC", name: "Nhựa", description: "PET, HDPE, PP, rPET flakes, nhựa hỗn hợp", status: "active" },
    { id: "CAT-METAL", name: "Kim loại", description: "Nhôm, sắt, đồng, thép phế liệu", status: "active" },
    { id: "CAT-PAPER", name: "Giấy & carton", description: "OCC, bìa carton, giấy văn phòng", status: "active" },
    { id: "CAT-WOOD", name: "Gỗ", description: "Pallet, ván, mùn cưa", status: "active" },
    { id: "CAT-TEXTILE", name: "Dệt may", description: "Vải vụn, cotton offcuts, sợi", status: "active" },
  ],
  materials: [
    { id: "MAT-HDPE", category: "Nhựa", name: "Hạt nhựa HDPE tái sinh", unit: "kg", recyclable: true, hazardous: "none", status: "active" },
    { id: "MAT-RPET", category: "Nhựa", name: "rPET flakes trong/blue", unit: "kg", recyclable: true, hazardous: "none", status: "active" },
    { id: "MAT-ALU", category: "Kim loại", name: "Nhôm phế liệu loại 1", unit: "kg", recyclable: true, hazardous: "none", status: "active" },
    { id: "MAT-OCC", category: "Giấy & carton", name: "Carton OCC ép kiện", unit: "kg", recyclable: true, hazardous: "none", status: "active" },
    { id: "MAT-PCB", category: "Điện tử", name: "Bo mạch điện tử thải", unit: "kg", recyclable: "có điều kiện", hazardous: "medium", status: "restricted" },
  ],
  supplyListings: [
    { id: "SUP-1001", companyId: "COM-001", company: "GreenTech Manufacturing", materialId: "MAT-HDPE", title: "Hạt nhựa HDPE tái sinh - industrial grade", category: "Nhựa", quantity: 24, available: 18, unit: "tấn", price: 11800, location: "Bình Dương", quality: "Grade A", purity: "97%", transactionType: "Bán", status: "active", availableUntil: "31/07/2026", offers: 3 },
    { id: "SUP-1002", companyId: "COM-003", company: "MetalLoop Co.", materialId: "MAT-ALU", title: "Nhôm phế liệu loại 1 đã phân loại", category: "Kim loại", quantity: 8, available: 5, unit: "tấn", price: 37500, location: "Đồng Nai", quality: "Clean", purity: "95%", transactionType: "Bán", status: "pending_review", availableUntil: "20/07/2026", offers: 2 },
    { id: "SUP-1003", companyId: "COM-004", company: "Saigon Paper Hub", materialId: "MAT-OCC", title: "Thùng carton OCC ép kiện", category: "Giấy & carton", quantity: 12, available: 12, unit: "tấn", price: 2200, location: "TP. Hồ Chí Minh", quality: "Mixed", purity: "90%", transactionType: "Bán", status: "active", availableUntil: "18/07/2026", offers: 1 },
    { id: "SUP-1004", companyId: "COM-001", company: "GreenTech Manufacturing", materialId: "MAT-RPET", title: "rPET flakes màu xanh nhạt", category: "Nhựa", quantity: 16, available: 0, unit: "tấn", price: 10400, location: "Bình Dương", quality: "Grade B", purity: "92%", transactionType: "Bán", status: "sold_out", availableUntil: "10/07/2026", offers: 5 },
  ],
  demandListings: [
    { id: "DEM-3001", companyId: "COM-002", company: "EcoPack Vietnam", materialId: "MAT-RPET", title: "Cần mua rPET flakes trong/blue", quantity: 30, unit: "tấn/tháng", targetPrice: 10500, location: "Long An", quality: "Clear/Blue, ít tạp chất", frequency: "Hàng tháng", deadline: "20/07/2026", status: "active", quotes: 5 },
    { id: "DEM-3002", companyId: "COM-002", company: "EcoPack Vietnam", materialId: "MAT-OCC", title: "Thu mua thùng carton OCC ép kiện", quantity: 18, unit: "tấn", targetPrice: 2000, location: "TP. Hồ Chí Minh", quality: "Khô, ép kiện", frequency: "Một lần", deadline: "18/07/2026", status: "active", quotes: 2 },
    { id: "DEM-3003", companyId: "COM-001", company: "GreenTech Manufacturing", materialId: "MAT-ALU", title: "Tìm nguồn cung nhôm phế liệu sạch", quantity: 5, unit: "tấn", targetPrice: 37000, location: "Bình Dương", quality: "Không lẫn sắt", frequency: "Một lần", deadline: "15/07/2026", status: "fulfilled", quotes: 3 },
  ],
  offers: [
    { id: "OFF-5001", type: "buyer_to_seller", buyer: "EcoPack Vietnam", seller: "GreenTech Manufacturing", listingId: "SUP-1001", demandId: null, material: "Hạt nhựa HDPE tái sinh", quantity: 8, unit: "tấn", offerPrice: 11200, status: "pending", message: "Cần COA nếu có, buyer tự bố trí vận chuyển.", createdAt: "05/07/2026 10:30" },
    { id: "OFF-5002", type: "buyer_to_seller", buyer: "EcoPack Vietnam", seller: "MetalLoop Co.", listingId: "SUP-1002", demandId: null, material: "Nhôm phế liệu loại 1", quantity: 5, unit: "tấn", offerPrice: 37500, status: "accepted", message: "Chấp nhận giá niêm yết, nhận tại Đồng Nai.", createdAt: "05/07/2026 09:05" },
    { id: "OFF-5003", type: "buyer_to_seller", buyer: "EcoPack Vietnam", seller: "Saigon Paper Hub", listingId: "SUP-1003", demandId: null, material: "Carton OCC ép kiện", quantity: 10, unit: "tấn", offerPrice: 2000, status: "rejected", message: "Seller từ chối do chưa đủ tồn kho.", createdAt: "04/07/2026 16:22" },
    { id: "OFF-5004", type: "seller_to_buyer", buyer: "EcoPack Vietnam", seller: "GreenTech Manufacturing", listingId: "SUP-1001", demandId: "DEM-3001", material: "rPET/HDPE tái sinh", quantity: 12, unit: "tấn", offerPrice: 10800, status: "pending", message: "Có thể giao 12 tấn trước, phần còn lại sau 7 ngày.", createdAt: "05/07/2026 14:10" },
  ],
  transactions: [
    { id: "TX-9001", offerId: "OFF-5002", buyer: "EcoPack Vietnam", seller: "MetalLoop Co.", listingId: "SUP-1002", material: "Nhôm phế liệu loại 1", quantity: 5, unit: "tấn", finalPrice: 37500, status: "in_progress", paymentStatus: "manual_offline", buyerConfirmed: false, sellerConfirmed: true, createdAt: "05/07/2026 09:20" },
    { id: "TX-9002", offerId: "OFF-4998", buyer: "VinaPack", seller: "Saigon Paper Hub", listingId: "SUP-1003", material: "Carton OCC ép kiện", quantity: 10, unit: "tấn", finalPrice: 2100, status: "completed", paymentStatus: "bypassed_demo", buyerConfirmed: true, sellerConfirmed: true, createdAt: "02/07/2026 15:30" },
    { id: "TX-9003", offerId: "OFF-4980", buyer: "Mekong Reuse", seller: "GreenTech Manufacturing", listingId: "SUP-1004", material: "rPET flakes", quantity: 3, unit: "tấn", finalPrice: 10400, status: "cancelled", paymentStatus: "not_required", buyerConfirmed: false, sellerConfirmed: false, createdAt: "01/07/2026 11:05" },
  ],
  transactionEvents: [
    { tx: "TX-9001", title: "Seller chấp nhận offer", note: "OFF-5002 accepted", time: "05/07/2026 09:20", done: true },
    { tx: "TX-9001", title: "Hệ thống tạo giao dịch", note: "Transaction liên kết offer và listing", time: "05/07/2026 09:20", done: true },
    { tx: "TX-9001", title: "Ghi nhận thanh toán demo", note: "payment_status = manual_offline", time: "05/07/2026 09:22", done: true },
    { tx: "TX-9001", title: "Seller xác nhận đã bàn giao", note: "seller_confirmed = true", time: "05/07/2026 13:00", done: true },
    { tx: "TX-9001", title: "Chờ buyer xác nhận hoàn tất", note: "Sau khi nhận hàng và đối soát chứng từ", time: "Đang chờ", done: false },
  ],
  reviews: [
    { id: "REV-6001", transactionId: "TX-9002", reviewer: "VinaPack", reviewed: "Saigon Paper Hub", rating: 5, type: "buyer_to_seller", comment: "Carton đúng mô tả, giao đúng hẹn." },
    { id: "REV-6002", transactionId: "TX-9002", reviewer: "Saigon Paper Hub", reviewed: "VinaPack", rating: 4, type: "seller_to_buyer", comment: "Buyer phản hồi nhanh, xác nhận đầy đủ." },
  ],
  reports: [
    { id: "RPT-7001", reporter: "EcoPack Vietnam", targetType: "SupplyListing", targetId: "SUP-1001", reason: "Listing sai thông tin", description: "Độ tinh khiết HDPE thực tế thấp hơn mô tả, thiếu chứng từ COA.", status: "processing", evidence: ["sample-photo.jpg", "buyer-note.pdf"] },
    { id: "RPT-7002", reporter: "GreenTech Manufacturing", targetType: "Company", targetId: "COM-004", reason: "Không thực hiện giao dịch", description: "Buyer hủy nhiều lần sau khi seller giữ hàng.", status: "pending", evidence: ["chat-log.pdf"] },
  ],
  notifications: [
    { id: "NOT-8001", user: "EcoPack Vietnam", title: "Seller đã chấp nhận OFF-5002", content: "Hệ thống đã tạo giao dịch TX-9001.", type: "offer_accepted", read: false, time: "05/07/2026 09:21" },
    { id: "NOT-8002", user: "GreenTech Manufacturing", title: "Bạn có offer mới OFF-5001", content: "EcoPack đề nghị mua 8 tấn HDPE.", type: "new_offer", read: false, time: "05/07/2026 10:30" },
    { id: "NOT-8003", user: "EcoPack Vietnam", title: "Có thể đánh giá seller", content: "Giao dịch TX-9002 đã hoàn tất.", type: "review_available", read: true, time: "03/07/2026 08:00" },
  ],
};

const useCases = [
  ["UC001", "Đăng ký tài khoản", "Auth & bảo mật", "Mock form đăng ký + rule email/password"],
  ["UC002", "Xác thực tài khoản", "Auth & bảo mật", "OTP demo 123456"],
  ["UC003", "Đăng nhập", "Auth & bảo mật", "Role switcher mô phỏng session"],
  ["UC004", "Đăng xuất", "Auth & bảo mật", "Chuyển role về Guest"],
  ["UC005", "Quên mật khẩu", "Auth & bảo mật", "Form gửi link/mã reset"],
  ["UC006", "Đổi mật khẩu", "Auth & bảo mật", "Security panel"],
  ["UC007", "Cập nhật hồ sơ cá nhân", "Hồ sơ doanh nghiệp", "User profile mock"],
  ["UC008", "Tạo hồ sơ doanh nghiệp", "Hồ sơ doanh nghiệp", "Company form"],
  ["UC009", "Cập nhật hồ sơ doanh nghiệp", "Hồ sơ doanh nghiệp", "Company edit state"],
  ["UC010", "Gửi yêu cầu xác minh doanh nghiệp", "Hồ sơ doanh nghiệp", "Upload docs mock"],
  ["UC011", "Duyệt doanh nghiệp", "Admin", "Verification queue"],
  ["UC012", "Quản lý thành viên doanh nghiệp", "Hồ sơ doanh nghiệp", "Member table"],
  ["UC013", "Phân quyền thành viên", "Hồ sơ doanh nghiệp", "Role chips Owner/Buyer/Seller"],
  ["UC014", "Xem danh mục vật liệu", "Danh mục vật liệu", "Catalog table"],
  ["UC015", "Tạo danh mục vật liệu", "Danh mục vật liệu", "Admin form"],
  ["UC016", "Cập nhật danh mục vật liệu", "Danh mục vật liệu", "Admin edit form"],
  ["UC017", "Khóa/ẩn danh mục vật liệu", "Danh mục vật liệu", "Status active/restricted/inactive"],
  ["UC018", "Đăng nguồn cung vật liệu", "Nguồn cung", "Create listing form"],
  ["UC019", "Cập nhật nguồn cung", "Nguồn cung", "Seller inventory table"],
  ["UC020", "Ẩn/đóng nguồn cung", "Nguồn cung", "Close action"],
  ["UC021", "Xem nguồn cung của doanh nghiệp", "Nguồn cung", "My supply list"],
  ["UC022", "Xem chi tiết nguồn cung", "Marketplace", "Listing detail panel"],
  ["UC023", "Admin kiểm duyệt nguồn cung", "Admin", "Pending listing review"],
  ["UC024", "Đăng nhu cầu mua", "Nhu cầu mua", "Demand create form"],
  ["UC025", "Cập nhật nhu cầu mua", "Nhu cầu mua", "My demand table"],
  ["UC026", "Đóng nhu cầu mua", "Nhu cầu mua", "Close demand action"],
  ["UC027", "Xem danh sách nhu cầu mua", "Nhu cầu mua", "Demand marketplace"],
  ["UC028", "Xem chi tiết nhu cầu mua", "Nhu cầu mua", "Demand detail panel"],
  ["UC029", "Tìm kiếm vật liệu", "Marketplace", "Search/filter mock"],
  ["UC030", "Lọc vật liệu", "Marketplace", "Category/location/status filters"],
  ["UC031", "Sắp xếp danh sách vật liệu", "Marketplace", "Sort select"],
  ["UC032", "Lưu nguồn cung quan tâm", "Marketplace", "Saved marker mock"],
  ["UC033", "Xem thông tin seller", "Marketplace", "Supplier profile detail"],
  ["UC034", "Gửi đề nghị mua", "Offer / RFQ", "Buyer-to-seller offer form"],
  ["UC035", "Seller xem đề nghị mua", "Offer / RFQ", "Received offer kanban"],
  ["UC036", "Seller chấp nhận đề nghị", "Offer / RFQ", "Accept action tạo transaction"],
  ["UC037", "Seller từ chối đề nghị", "Offer / RFQ", "Reject action/status"],
  ["UC038", "Buyer hủy đề nghị mua", "Offer / RFQ", "Cancel pending offer"],
  ["UC039", "Tạo giao dịch", "Giao dịch", "Auto transaction record"],
  ["UC040", "Theo dõi trạng thái giao dịch", "Giao dịch", "Timeline"],
  ["UC041", "Buyer xác nhận hoàn tất", "Giao dịch", "Confirm button"],
  ["UC042", "Seller xác nhận hoàn tất", "Giao dịch", "Seller confirmed state"],
  ["UC043", "Hủy giao dịch", "Giao dịch", "Cancel reason mock"],
  ["UC044", "Buyer đánh giá seller", "Đánh giá", "Review form"],
  ["UC045", "Seller đánh giá buyer", "Đánh giá", "Review form"],
  ["UC046", "Xem điểm uy tín doanh nghiệp", "Đánh giá", "Reputation summary"],
  ["UC047", "Báo cáo vi phạm", "Báo cáo vi phạm", "Report form"],
  ["UC048", "Admin xử lý báo cáo vi phạm", "Báo cáo vi phạm", "Moderation panel"],
  ["UC049", "Nhận thông báo đề nghị mua", "Thông báo", "New offer notification"],
  ["UC050", "Nhận thông báo kết quả đề nghị", "Thông báo", "Offer result notification"],
  ["UC051", "Nhận thông báo trạng thái giao dịch", "Thông báo", "Transaction notification"],
  ["UC052", "Xem danh sách thông báo", "Thông báo", "Notification list"],
  ["UC053", "Đánh dấu thông báo đã đọc", "Thông báo", "Mark read action"],
  ["UC054", "Xem dashboard doanh nghiệp", "Dashboard", "Business metrics"],
  ["UC055", "Xem dashboard admin", "Admin", "System metrics"],
  ["UC056", "Xem báo cáo giao dịch", "Báo cáo & export", "Transaction report table"],
  ["UC057", "Xuất báo cáo", "Báo cáo & export", "CSV/PDF mock"],
  ["UC058", "Seller gửi báo giá cho nhu cầu mua", "Nhu cầu mua", "Seller-to-buyer offer form"],
  ["UC059", "Buyer xử lý báo giá từ seller", "Offer / RFQ", "Quote accept/reject mock"],
];

function formatMoney(value) {
  return `${Number(value).toLocaleString("vi-VN")}đ`;
}

function classForStatus(status) {
  const value = String(status).toLowerCase();
  if (["active", "accepted", "completed", "verified", "resolved"].some((x) => value.includes(x))) return "green";
  if (["pending", "processing", "in_progress", "manual"].some((x) => value.includes(x))) return "amber";
  if (["draft", "review", "bypassed", "sent"].some((x) => value.includes(x))) return "blue";
  if (["rejected", "cancelled", "locked", "inactive", "suspended"].some((x) => value.includes(x))) return "red";
  return "orange";
}

function status(status) {
  return `<span class="status ${classForStatus(status)}"><span class="dot-icon"></span>${status}</span>`;
}

function actionButton(label, view, variant = "secondary") {
  return `<button class="btn ${variant}" data-view="${view}"><span class="dot-icon"></span>${label}</button>`;
}

function metric(label, value, tone = "primary") {
  return `<div class="card metric"><div class="icon" style="color:${tone === "blue" ? "var(--blue)" : tone === "orange" ? "var(--orange)" : "var(--primary)"}"></div><div><strong>${value}</strong><span>${label}</span></div></div>`;
}

function rowTitle(title, sub) {
  return `<div class="row-title"><div class="thumb"></div><div><strong>${title}</strong><br><span class="muted">${sub}</span></div></div>`;
}

function currentCompany() {
  if (state.role === "seller") return mock.companies[0];
  if (state.role === "buyer") return mock.companies[1];
  if (state.role === "admin") return { name: "System Admin", status: "superuser", rating: "-", completed: mock.transactions.length };
  return { name: "Guest", status: "public", rating: "-", completed: 0 };
}

function render() {
  const role = roles[state.role];
  document.getElementById("app").innerHTML = `
    <div class="mobile-top">
      <strong>Circular Materials</strong>
      <div class="mobile-selects">
        <select data-mobile-lang>
          <option value="vi" ${state.lang === "vi" ? "selected" : ""}>Tiếng Việt</option>
          <option value="en" ${state.lang === "en" ? "selected" : ""}>English</option>
        </select>
        <select data-mobile-role>${Object.entries(roles).map(([key, item]) => `<option value="${key}" ${key === state.role ? "selected" : ""}>${item.label}</option>`).join("")}</select>
        <select data-mobile-view>${navItems.map(([key, label]) => `<option value="${key}" ${key === state.view ? "selected" : ""}>${label}</option>`).join("")}</select>
      </div>
    </div>
    <div class="shell">
      <aside class="sidebar">
        <div class="brand"><span class="mark"></span><span>Circular Materials<br>Exchange</span></div>
        <div class="context-card">
          <strong>${role.company}</strong>
          <span>${role.subtitle}</span>
        </div>
        <div class="role-switcher">
          ${Object.entries(roles).map(([key, item]) => `<button class="role-btn ${key === state.role ? "active" : ""}" data-role="${key}">${item.label}</button>`).join("")}
        </div>
        <div class="lang-switcher">
          <button class="role-btn ${state.lang === "vi" ? "active" : ""}" data-lang="vi">Tiếng Việt</button>
          <button class="role-btn ${state.lang === "en" ? "active" : ""}" data-lang="en">English</button>
        </div>
        <nav class="nav">
          ${navItems.map(([key, label]) => `<button class="nav-btn ${key === state.view ? "active" : ""}" data-view="${key}"><span class="dot-icon"></span>${label}</button>`).join("")}
        </nav>
      </aside>
      <main class="main">${renderView()}</main>
    </div>
  `;
  applyStaticI18n(document.getElementById("app"));
}

function header(eyebrow, title, copy, actions = "") {
  return `
    <header class="topbar">
      <div>
        <p class="eyebrow">${eyebrow}</p>
        <h1>${title}</h1>
        <p class="muted">${copy}</p>
      </div>
      <div class="actions">${actions}</div>
    </header>
  `;
}

function renderView() {
  const views = {
    dashboard: renderDashboard,
    auth: renderAuth,
    company: renderCompany,
    catalog: renderCatalog,
    marketplace: renderMarketplace,
    supply: renderSupply,
    demand: renderDemand,
    offers: renderOffers,
    transactions: renderTransactions,
    reviews: renderReviews,
    reports: renderReports,
    notifications: renderNotifications,
    admin: renderAdmin,
    exports: renderExports,
    usecases: renderUseCases,
  };
  return (views[state.view] || renderDashboard)();
}

function renderDashboard() {
  const company = currentCompany();
  const activeSupply = mock.supplyListings.filter((x) => x.status === "active").length;
  const activeDemand = mock.demandListings.filter((x) => x.status === "active").length;
  const completedTx = mock.transactions.filter((x) => x.status === "completed").length;
  const pendingOffers = mock.offers.filter((x) => x.status === "pending").length;
  return `
    ${header("UC054-UC055 · Dashboard", state.role === "admin" ? "Dashboard quản trị hệ thống" : `Dashboard doanh nghiệp`, "Số liệu mock thể hiện nguồn cung, nhu cầu, offer, giao dịch và uy tín doanh nghiệp.", `${actionButton("Đi marketplace", "marketplace")} ${actionButton("Xuất báo cáo", "exports", "primary")}`)}
    <section class="grid four">
      ${metric("Nguồn cung active", activeSupply)}
      ${metric("Nhu cầu active", activeDemand, "blue")}
      ${metric("Offer pending", pendingOffers, "orange")}
      ${metric("Giao dịch hoàn tất", completedTx)}
    </section>
    <section class="split" style="margin-top:18px">
      <div class="panel">
        <div class="section-title">
          <h2>Hoạt động gần đây</h2>
          ${status(company.status)}
        </div>
        <div class="timeline">
          <div class="timeline-step"><div class="timeline-dot"></div><div class="timeline-copy"><strong>Offer OFF-5002 được chấp nhận</strong><span>Hệ thống tạo TX-9001 và ghi nhận payment_status = manual_offline</span></div></div>
          <div class="timeline-step"><div class="timeline-dot"></div><div class="timeline-copy"><strong>Listing SUP-1001 nhận 3 offer</strong><span>Seller cần xử lý offer pending trong màn Offer / RFQ</span></div></div>
          <div class="timeline-step"><div class="timeline-dot"></div><div class="timeline-copy"><strong>Demand DEM-3001 nhận 5 báo giá</strong><span>Luồng UC058/UC059 được mock bằng seller_to_buyer offer</span></div></div>
          <div class="timeline-step"><div class="timeline-dot pending"></div><div class="timeline-copy"><strong>Report RPT-7001 đang xử lý</strong><span>Admin có thể cảnh cáo, ẩn listing hoặc khóa doanh nghiệp</span></div></div>
        </div>
      </div>
      <aside class="panel">
        <h2>Hồ sơ hiện tại</h2>
        <div class="detail-list">
          <div class="detail-item"><span>Actor</span><strong>${roles[state.role].label}</strong></div>
          <div class="detail-item"><span>Người dùng</span><strong>${roles[state.role].name}</strong></div>
          <div class="detail-item"><span>Doanh nghiệp</span><strong>${company.name}</strong></div>
          <div class="detail-item"><span>Điểm uy tín</span><strong>${company.rating}</strong></div>
          <div class="detail-item"><span>Giao dịch</span><strong>${company.completed}</strong></div>
        </div>
      </aside>
    </section>
  `;
}

function renderAuth() {
  return `
    ${header("UC001-UC006 · Auth", "Đăng nhập, đăng ký và bảo mật", "Mô phỏng đầy đủ đăng ký, OTP, đăng nhập, quên mật khẩu, đổi mật khẩu và đăng xuất.", `<button class="btn ghost" data-role="guest">Đăng xuất demo</button><button class="btn primary" data-role="buyer">Đăng nhập buyer</button>`)}
    <section class="grid three">
      <div class="panel">
        <h2>Đăng nhập</h2>
        <div class="field"><label>Email hoặc số điện thoại</label><input value="minhanh@ecopack.vn"></div>
        <div class="field"><label>Mật khẩu</label><input type="password" value="password123"></div>
        <button class="btn primary" style="width:100%; margin-top:14px" data-role="buyer">Đăng nhập</button>
      </div>
      <div class="panel">
        <h2>Đăng ký tài khoản</h2>
        <div class="field-grid">
          <div class="field"><label>Họ tên</label><input value="Lê Thanh Huy"></div>
          <div class="field"><label>Số điện thoại</label><input value="+84 913 040 882"></div>
          <div class="field"><label>Email</label><input value="huy@metalloop.vn"></div>
          <div class="field"><label>Loại tài khoản</label><select><option>Đại diện doanh nghiệp</option></select></div>
        </div>
        <div class="notice" style="margin-top:14px">OTP demo: <strong>123456</strong>. Tài khoản mới ở trạng thái pending_verification.</div>
      </div>
      <div class="panel">
        <h2>Quên / đổi mật khẩu</h2>
        <div class="field"><label>Email nhận mã</label><input value="minhanh@ecopack.vn"></div>
        <div class="field"><label>Mật khẩu mới</label><input type="password" value="newPassword123"></div>
        <div class="field"><label>Xác nhận OTP</label><input value="123456"></div>
        <button class="btn secondary" style="width:100%; margin-top:14px">Cập nhật mật khẩu</button>
      </div>
    </section>
  `;
}

function renderCompany() {
  const company = currentCompany();
  const members = mock.companyMembers.filter((x) => x.companyId === "COM-002" || x.companyId === "COM-001");
  return `
    ${header("UC007-UC013 · Company", "Hồ sơ cá nhân và doanh nghiệp", "Bao gồm cập nhật hồ sơ cá nhân, tạo/cập nhật doanh nghiệp, gửi xác minh, thành viên và phân quyền.", `${actionButton("Gửi xác minh", "company", "primary")} ${actionButton("Admin duyệt", "admin")}`)}
    <section class="split">
      <div class="panel">
        <div class="section-title"><h2>${company.name}</h2>${status(company.status)}</div>
        <div class="field-grid">
          <div class="field"><label>Tên pháp lý</label><input value="${company.name} JSC"></div>
          <div class="field"><label>Mã số thuế</label><input value="${company.taxCode || "0318-456-982"}"></div>
          <div class="field"><label>Lĩnh vực</label><input value="${company.type || "Marketplace operator"}"></div>
          <div class="field"><label>Tỉnh/thành</label><input value="${company.province || "TP. Hồ Chí Minh"}"></div>
          <div class="field"><label>Email liên hệ</label><input value="${company.email || "admin@cme.local"}"></div>
          <div class="field"><label>Số điện thoại</label><input value="${company.phone || "+84 000 000 000"}"></div>
          <div class="field" style="grid-column:1/-1"><label>Địa chỉ</label><input value="${company.address || "Cloud workspace"}"></div>
          <div class="field" style="grid-column:1/-1"><label>Mô tả</label><textarea>Nền tảng kết nối nguồn cung và nhu cầu vật liệu tuần hoàn B2B. Bản MVP dùng xác minh thủ công và thanh toán bypass.</textarea></div>
        </div>
      </div>
      <aside class="panel">
        <h2>Hồ sơ cá nhân</h2>
        <div class="detail-list">
          <div class="detail-item"><span>Họ tên</span><strong>${roles[state.role].name}</strong></div>
          <div class="detail-item"><span>Vai trò</span><strong>${roles[state.role].label}</strong></div>
          <div class="detail-item"><span>Email</span><strong>${state.role === "buyer" ? "minhanh@ecopack.vn" : "bao@greentech.vn"}</strong></div>
          <div class="detail-item"><span>Bảo mật</span><strong>OTP email demo</strong></div>
        </div>
      </aside>
    </section>
    <section class="panel" style="margin-top:18px">
      <div class="section-title"><h2>Thành viên & phân quyền</h2><button class="btn secondary">Mời thành viên</button></div>
      ${table(["Thành viên", "Doanh nghiệp", "Vai trò", "Trạng thái"], members.map((m) => [m.user, m.companyId, m.role, status(m.status)]))}
    </section>
  `;
}

function renderCatalog() {
  return `
    ${header("UC014-UC017 · Material catalog", "Danh mục vật liệu", "Admin chuẩn hóa vật liệu, buyer/seller dùng để lọc marketplace và tạo listing/demand.", `<button class="btn primary">Thêm vật liệu</button>`)}
    <section class="grid four">
      ${metric("Danh mục", mock.categories.length)}
      ${metric("Vật liệu", mock.materials.length, "blue")}
      ${metric("Hazardous", mock.materials.filter((x) => x.hazardous !== "none").length, "orange")}
      ${metric("Active", mock.materials.filter((x) => x.status === "active").length)}
    </section>
    <section class="split" style="margin-top:18px">
      <div class="table-wrap">
        ${table(["Mã", "Vật liệu", "Danh mục", "Đơn vị", "Recyclable", "Hazardous", "Trạng thái"], mock.materials.map((m) => [m.id, m.name, m.category, m.unit, String(m.recyclable), m.hazardous, status(m.status)]))}
      </div>
      <aside class="panel">
        <h2>Tạo/cập nhật vật liệu</h2>
        <div class="field"><label>Tên vật liệu</label><input value="Hạt nhựa HDPE tái sinh"></div>
        <div class="field"><label>Danh mục</label><select>${mock.categories.map((c) => `<option>${c.name}</option>`).join("")}</select></div>
        <div class="field"><label>Đơn vị</label><select><option>kg</option><option>tấn</option><option>m3</option></select></div>
        <div class="field"><label>Mức nguy hại</label><select><option>none</option><option>low</option><option>medium</option><option>high</option></select></div>
        <button class="btn primary" style="width:100%; margin-top:14px">Lưu danh mục</button>
      </aside>
    </section>
  `;
}

function renderMarketplace() {
  const selected = mock.supplyListings.find((x) => x.id === state.selectedListing) || mock.supplyListings[0];
  return `
    ${header("UC022, UC029-UC033 · Marketplace", "Sàn giao dịch nguồn cung vật liệu", "Buyer tìm kiếm, lọc, sắp xếp, xem seller và gửi đề nghị mua từ listing.", `<button class="btn primary" data-view="offers">Gửi đề nghị mua</button>`)}
    <section class="toolbar">
      <div class="actions">
        <select class="btn ghost"><option>Tất cả danh mục</option><option>Nhựa</option><option>Kim loại</option><option>Giấy & carton</option></select>
        <select class="btn ghost"><option>Tất cả địa điểm</option><option>Bình Dương</option><option>Đồng Nai</option><option>TP. Hồ Chí Minh</option></select>
        <select class="btn ghost"><option>Sắp xếp: mới nhất</option><option>Giá thấp đến cao</option><option>Khối lượng lớn nhất</option></select>
      </div>
      <input style="min-height:40px; min-width:260px; border:1px solid var(--outline); border-radius:8px; padding:8px 12px" value="HDPE, rPET, nhôm">
    </section>
    <section class="split">
      <div class="grid two">
        ${mock.supplyListings.map((l) => `
          <article class="card">
            <div class="section-title"><h3>${l.title}</h3>${status(l.status)}</div>
            <p class="muted">${l.company} · ${l.location} · ${l.quality} · ${l.purity}</p>
            <div class="detail-list">
              <div class="detail-item"><span>Khối lượng còn lại</span><strong>${l.available}/${l.quantity} ${l.unit}</strong></div>
              <div class="detail-item"><span>Giá niêm yết</span><strong>${formatMoney(l.price)}/kg</strong></div>
              <div class="detail-item"><span>Offer</span><strong>${l.offers}</strong></div>
            </div>
            <button class="btn secondary" style="width:100%; margin-top:14px" data-listing="${l.id}">Xem chi tiết</button>
          </article>
        `).join("")}
      </div>
      <aside class="panel">
        <div class="section-title"><h2>Chi tiết nguồn cung</h2>${status(selected.status)}</div>
        <div class="detail-list">
          <div class="detail-item"><span>Mã listing</span><strong>${selected.id}</strong></div>
          <div class="detail-item"><span>Seller</span><strong>${selected.company}</strong></div>
          <div class="detail-item"><span>Vật liệu</span><strong>${selected.title}</strong></div>
          <div class="detail-item"><span>Chất lượng</span><strong>${selected.quality} · ${selected.purity}</strong></div>
          <div class="detail-item"><span>Hạn khả dụng</span><strong>${selected.availableUntil}</strong></div>
          <div class="detail-item"><span>Uy tín seller</span><strong>4.8 · 36 giao dịch</strong></div>
        </div>
        <button class="btn primary" style="width:100%; margin-top:16px" data-view="offers">Gửi đề nghị mua</button>
        <button class="btn ghost" style="width:100%; margin-top:10px">Lưu quan tâm</button>
      </aside>
    </section>
  `;
}

function renderSupply() {
  return `
    ${header("UC018-UC023 · Supply listing", "Quản lý nguồn cung vật liệu", "Seller đăng, sửa, ẩn/đóng listing; admin có thể kiểm duyệt pending_review.", `<button class="btn primary">Đăng nguồn cung</button>`)}
    <section class="split">
      <div class="table-wrap">
        ${table(["Nguồn cung", "Khối lượng", "Giá", "Offer", "Trạng thái", "Hành động"], mock.supplyListings.map((l) => [
          rowTitle(l.title, `${l.id} · ${l.location}`),
          `${l.available}/${l.quantity} ${l.unit}`,
          `${formatMoney(l.price)}/kg`,
          l.offers,
          status(l.status),
          `<button class="btn ghost">Sửa</button> <button class="btn warn">Đóng</button>`,
        ]))}
      </div>
      <aside class="panel">
        <h2>Form đăng nguồn cung</h2>
        <div class="field"><label>Tên nguồn cung</label><input value="Hạt nhựa HDPE tái sinh - industrial grade"></div>
        <div class="field"><label>Vật liệu</label><select>${mock.materials.map((m) => `<option>${m.name}</option>`).join("")}</select></div>
        <div class="field-grid">
          <div class="field"><label>Số lượng</label><input value="24"></div>
          <div class="field"><label>Giá/kg</label><input value="11800"></div>
        </div>
        <div class="field"><label>Tài liệu đính kèm</label><input value="COA-HDPE-2026.pdf, warehouse-photo.jpg"></div>
        <button class="btn primary" style="width:100%; margin-top:14px">Lưu draft / đăng bán</button>
      </aside>
    </section>
  `;
}

function renderDemand() {
  const selected = mock.demandListings.find((x) => x.id === state.selectedDemand) || mock.demandListings[0];
  return `
    ${header("UC024-UC028, UC058-UC059 · Demand", "Sàn nhu cầu mua vật liệu", "Buyer đăng nhu cầu; seller gửi báo giá seller_to_buyer để khép kín luồng demand.", `<button class="btn primary">Đăng nhu cầu mua</button>`)}
    <section class="split">
      <div class="grid">
        ${mock.demandListings.map((d) => `
          <article class="card">
            <div class="section-title"><h3>${d.title}</h3>${status(d.status)}</div>
            <p class="muted">${d.company} · ${d.location} · deadline ${d.deadline}</p>
            <div class="detail-list">
              <div class="detail-item"><span>Số lượng cần</span><strong>${d.quantity} ${d.unit}</strong></div>
              <div class="detail-item"><span>Ngân sách</span><strong>${formatMoney(d.targetPrice)}/kg</strong></div>
              <div class="detail-item"><span>Báo giá nhận được</span><strong>${d.quotes}</strong></div>
            </div>
            <button class="btn secondary" style="width:100%; margin-top:14px" data-demand="${d.id}">Xem nhu cầu</button>
          </article>
        `).join("")}
      </div>
      <aside class="panel">
        <div class="section-title"><h2>Seller gửi báo giá</h2>${status("seller_to_buyer")}</div>
        <div class="detail-list">
          <div class="detail-item"><span>Nhu cầu</span><strong>${selected.id}</strong></div>
          <div class="detail-item"><span>Buyer</span><strong>${selected.company}</strong></div>
          <div class="detail-item"><span>Cần mua</span><strong>${selected.quantity} ${selected.unit}</strong></div>
          <div class="detail-item"><span>Yêu cầu</span><strong>${selected.quality}</strong></div>
        </div>
        <div class="field" style="margin-top:14px"><label>Khối lượng cung cấp</label><input value="12 tấn"></div>
        <div class="field"><label>Đơn giá báo</label><input value="10.800đ/kg"></div>
        <div class="field"><label>Ghi chú</label><textarea>Có thể giao trước 12 tấn, phần còn lại sau 7 ngày.</textarea></div>
        <button class="btn primary" style="width:100%">Gửi báo giá</button>
      </aside>
    </section>
  `;
}

function renderOffers() {
  const pending = mock.offers.filter((x) => x.status === "pending");
  const accepted = mock.offers.filter((x) => x.status === "accepted");
  const rejected = mock.offers.filter((x) => x.status === "rejected" || x.status === "cancelled");
  return `
    ${header("UC034-UC038, UC058-UC059 · Offer/RFQ", "Quản lý đề nghị mua và báo giá", "Bao gồm buyer gửi offer, seller xử lý offer, buyer hủy pending offer và seller gửi báo giá cho demand.", `<button class="btn primary" data-action="accept-offer">Accept OFF-5001 demo</button>`)}
    <section class="split">
      <div class="panel">
        <div class="section-title"><h2>Form gửi đề nghị mua</h2>${status("pending draft")}</div>
        <div class="field-grid">
          <div class="field"><label>Listing</label><select>${mock.supplyListings.map((l) => `<option>${l.id} · ${l.title}</option>`).join("")}</select></div>
          <div class="field"><label>Số lượng</label><input value="8 tấn"></div>
          <div class="field"><label>Giá đề nghị</label><input value="11.200đ/kg"></div>
          <div class="field"><label>Ngày nhận</label><input value="18/07/2026"></div>
          <div class="field" style="grid-column:1/-1"><label>Lời nhắn</label><textarea>Cần mẫu COA nếu có. Buyer tự bố trí vận chuyển sau khi hai bên xác nhận.</textarea></div>
        </div>
      </div>
      <aside class="panel">
        <h2>Lịch sử đàm phán</h2>
        <div class="timeline">
          <div class="timeline-step"><div class="timeline-dot"></div><div class="timeline-copy"><strong>Buyer gửi offer</strong><span>8 tấn · 11.200đ/kg · pending</span></div></div>
          <div class="timeline-step"><div class="timeline-dot pending"></div><div class="timeline-copy"><strong>Seller chưa xử lý</strong><span>Có thể chấp nhận hoặc từ chối</span></div></div>
        </div>
      </aside>
    </section>
    <section class="kanban" style="margin-top:18px">
      ${offerColumn("Pending", pending)}
      ${offerColumn("Accepted", accepted)}
      ${offerColumn("Rejected / Cancelled", rejected)}
    </section>
  `;
}

function offerColumn(title, rows) {
  return `<div class="kanban-col"><h3>${title}</h3>${rows.map((o) => `
    <article class="card">
      <div class="section-title"><strong>${o.id}</strong>${status(o.status)}</div>
      <p><strong>${o.material}</strong></p>
      <p class="muted">${o.buyer} → ${o.seller}</p>
      <div class="detail-list">
        <div class="detail-item"><span>Số lượng</span><strong>${o.quantity} ${o.unit}</strong></div>
        <div class="detail-item"><span>Giá</span><strong>${formatMoney(o.offerPrice)}/kg</strong></div>
        <div class="detail-item"><span>Type</span><strong>${o.type}</strong></div>
      </div>
      <div class="actions" style="margin-top:12px"><button class="btn secondary">Chi tiết</button><button class="btn ghost">Hủy/Từ chối</button></div>
    </article>
  `).join("")}</div>`;
}

function renderTransactions() {
  const tx = mock.transactions.find((x) => x.id === state.selectedTransaction) || mock.transactions[0];
  const events = mock.transactionEvents.filter((x) => x.tx === tx.id);
  return `
    ${header("UC039-UC043 · Transactions", `Giao dịch ${tx.id}`, "Giao dịch được tạo tự động sau khi seller chấp nhận offer; thanh toán được bypass/manual offline trong MVP.", `<button class="btn ghost">Yêu cầu hủy</button><button class="btn primary" data-action="complete-tx">Xác nhận hoàn tất</button>`)}
    <section class="grid four">
      ${metric("Khối lượng", `${tx.quantity} ${tx.unit}`)}
      ${metric("Giá cuối", `${formatMoney(tx.finalPrice)}/kg`, "blue")}
      ${metric("Payment", tx.paymentStatus, "orange")}
      ${metric("Trạng thái", tx.status)}
    </section>
    <section class="split" style="margin-top:18px">
      <div class="panel">
        <div class="section-title"><h2>Timeline trạng thái</h2>${status(tx.status)}</div>
        <div class="timeline">
          ${events.map((e) => `<div class="timeline-step"><div class="timeline-dot ${e.done ? "" : "pending"}"></div><div class="timeline-copy"><strong>${e.title}</strong><span>${e.note} · ${e.time}</span></div></div>`).join("")}
        </div>
      </div>
      <aside class="panel">
        <h2>Chi tiết giao dịch</h2>
        <div class="detail-list">
          <div class="detail-item"><span>Offer</span><strong>${tx.offerId}</strong></div>
          <div class="detail-item"><span>Buyer</span><strong>${tx.buyer}</strong></div>
          <div class="detail-item"><span>Seller</span><strong>${tx.seller}</strong></div>
          <div class="detail-item"><span>Vật liệu</span><strong>${tx.material}</strong></div>
          <div class="detail-item"><span>Buyer confirmed</span><strong>${tx.buyerConfirmed}</strong></div>
          <div class="detail-item"><span>Seller confirmed</span><strong>${tx.sellerConfirmed}</strong></div>
        </div>
        <div class="notice" style="margin-top:14px">Không gọi payment gateway. payment_status chỉ ghi nhận manual_offline hoặc bypassed_demo.</div>
      </aside>
    </section>
    <section class="table-wrap" style="margin-top:18px">
      ${table(["Transaction", "Buyer", "Seller", "Vật liệu", "Khối lượng", "Status", "Payment"], mock.transactions.map((t) => [t.id, t.buyer, t.seller, t.material, `${t.quantity} ${t.unit}`, status(t.status), t.paymentStatus]))}
    </section>
  `;
}

function renderReviews() {
  return `
    ${header("UC044-UC046 · Reviews", "Đánh giá và điểm uy tín", "Buyer và seller đánh giá nhau sau khi giao dịch completed; điểm uy tín hiển thị ở hồ sơ doanh nghiệp.", `<button class="btn primary">Gửi đánh giá</button>`)}
    <section class="split">
      <div class="panel">
        <div class="section-title"><h2>Form đánh giá đối tác</h2>${status("available after completed")}</div>
        <div class="field-grid">
          <div class="field"><label>Transaction</label><select>${mock.transactions.filter((t) => t.status === "completed").map((t) => `<option>${t.id}</option>`).join("")}</select></div>
          <div class="field"><label>Đối tác</label><input value="Saigon Paper Hub"></div>
          <div class="field"><label>Điểm vật liệu đúng mô tả</label><select><option>5</option><option>4</option><option>3</option></select></div>
          <div class="field"><label>Thời gian phản hồi</label><select><option>4</option><option>5</option><option>3</option></select></div>
          <div class="field" style="grid-column:1/-1"><label>Nhận xét</label><textarea>Vật liệu đúng mô tả, giao nhận rõ ràng, chứng từ đầy đủ.</textarea></div>
        </div>
      </div>
      <aside class="panel">
        <h2>Uy tín doanh nghiệp</h2>
        <div class="detail-list">
          ${mock.companies.map((c) => `<div class="detail-item"><span>${c.name}</span><strong>${c.rating} · ${c.completed} giao dịch</strong></div>`).join("")}
        </div>
      </aside>
    </section>
    <section class="table-wrap" style="margin-top:18px">
      ${table(["Review", "Transaction", "Reviewer", "Reviewed", "Rating", "Comment"], mock.reviews.map((r) => [r.id, r.transactionId, r.reviewer, r.reviewed, r.rating, r.comment]))}
    </section>
  `;
}

function renderReports() {
  const report = mock.reports.find((x) => x.id === state.selectedReport) || mock.reports[0];
  return `
    ${header("UC047-UC048 · Violation reports", "Báo cáo vi phạm và moderation", "Business user tạo report; admin xử lý bằng cảnh cáo, ẩn listing, khóa user/company hoặc reject report.", `<button class="btn primary">Gửi report</button>`)}
    <section class="split">
      <div class="panel">
        <div class="section-title"><h2>Form báo cáo vi phạm</h2>${status("pending")}</div>
        <div class="field-grid">
          <div class="field"><label>Đối tượng</label><select><option>SupplyListing</option><option>Transaction</option><option>Company</option><option>User</option></select></div>
          <div class="field"><label>Mã đối tượng</label><input value="SUP-1001"></div>
          <div class="field"><label>Loại vi phạm</label><select><option>Listing sai thông tin</option><option>Vật liệu nguy hại không khai báo</option><option>Không thực hiện giao dịch</option></select></div>
          <div class="field"><label>Bằng chứng</label><input value="sample-photo.jpg, buyer-note.pdf"></div>
          <div class="field" style="grid-column:1/-1"><label>Mô tả</label><textarea>Độ tinh khiết HDPE thực tế thấp hơn mô tả, thiếu chứng từ COA.</textarea></div>
        </div>
      </div>
      <aside class="panel">
        <div class="section-title"><h2>Admin xử lý ${report.id}</h2>${status(report.status)}</div>
        <div class="detail-list">
          <div class="detail-item"><span>Reporter</span><strong>${report.reporter}</strong></div>
          <div class="detail-item"><span>Target</span><strong>${report.targetType} ${report.targetId}</strong></div>
          <div class="detail-item"><span>Lý do</span><strong>${report.reason}</strong></div>
          <div class="detail-item"><span>Evidence</span><strong>${report.evidence.join(", ")}</strong></div>
        </div>
        <div class="actions" style="margin-top:14px">
          <button class="btn ghost">Bỏ qua</button>
          <button class="btn warn">Ẩn listing</button>
          <button class="btn danger">Khóa công ty</button>
        </div>
      </aside>
    </section>
    <section class="table-wrap" style="margin-top:18px">
      ${table(["Report", "Reporter", "Target", "Reason", "Status"], mock.reports.map((r) => [r.id, r.reporter, `${r.targetType} ${r.targetId}`, r.reason, status(r.status)]))}
    </section>
  `;
}

function renderNotifications() {
  return `
    ${header("UC049-UC053 · Notifications", "Thông báo in-app", "Thông báo offer mới, kết quả offer, trạng thái giao dịch, report và đánh dấu đã đọc.", `<button class="btn secondary" data-action="mark-read">Đánh dấu đã đọc</button>`)}
    <section class="grid">
      ${mock.notifications.map((n) => `
        <article class="notification ${n.read ? "" : "unread"}">
          <div class="icon" style="color:var(--blue)"></div>
          <div><strong>${n.title}</strong><p class="muted" style="margin:4px 0 0">${n.content}</p><span class="muted">${n.time}</span></div>
          ${status(n.read ? "read" : "unread")}
        </article>
      `).join("")}
    </section>
  `;
}

function renderAdmin() {
  const pendingCompanies = mock.verificationRequests.filter((x) => x.status === "pending_verification");
  const pendingListings = mock.supplyListings.filter((x) => x.status === "pending_review");
  return `
    ${header("Admin · UC011, UC015-UC017, UC023, UC048, UC055", "Trang quản trị hệ thống", "Quản trị user, doanh nghiệp, danh mục vật liệu, listing, demand, offer, transaction, report và review.", `${actionButton("Danh mục", "catalog")} ${actionButton("Reports", "reports", "primary")}`)}
    <section class="grid four">
      ${metric("Người dùng", mock.users.length)}
      ${metric("Doanh nghiệp", mock.companies.length, "blue")}
      ${metric("Tin chờ duyệt", pendingListings.length, "orange")}
      ${metric("Report pending", mock.reports.filter((x) => x.status !== "resolved").length)}
    </section>
    <section class="split" style="margin-top:18px">
      <div class="panel">
        <div class="section-title"><h2>Hồ sơ doanh nghiệp chờ duyệt</h2>${status("pending_verification")}</div>
        ${table(["Request", "Company", "Submitted", "Docs", "Status", "Action"], pendingCompanies.map((v) => [v.id, v.company, v.submittedAt, v.docs.join(", "), status(v.status), `<button class="btn primary">Duyệt</button> <button class="btn danger">Từ chối</button>`]))}
      </div>
      <aside class="panel">
        <h2>Moderation nhanh</h2>
        <div class="detail-list">
          <div class="detail-item"><span>Listing chờ duyệt</span><strong>${pendingListings.map((x) => x.id).join(", ")}</strong></div>
          <div class="detail-item"><span>Material restricted</span><strong>MAT-PCB</strong></div>
          <div class="detail-item"><span>Report xử lý</span><strong>RPT-7001</strong></div>
          <div class="detail-item"><span>Export</span><strong>CSV/PDF demo</strong></div>
        </div>
      </aside>
    </section>
    <section class="table-wrap" style="margin-top:18px">
      ${table(["User", "Email", "Company", "Role", "Status", "Last login"], mock.users.map((u) => [u.name, u.email, u.companyId, u.role, status(u.status), u.lastLogin]))}
    </section>
  `;
}

function renderExports() {
  return `
    ${header("UC056-UC057 · Reports/export", "Báo cáo giao dịch và xuất dữ liệu", "Bộ lọc báo cáo, preview dữ liệu và xuất CSV/PDF demo cho admin hoặc business user.", `<button class="btn primary">Xuất CSV</button><button class="btn secondary">Xuất PDF demo</button>`)}
    <section class="panel">
      <div class="section-title"><h2>Bộ lọc báo cáo</h2>${status("preview")}</div>
      <div class="field-grid">
        <div class="field"><label>Khoảng thời gian</label><select><option>30 ngày gần nhất</option><option>Quý này</option><option>Năm nay</option></select></div>
        <div class="field"><label>Trạng thái</label><select><option>Tất cả</option><option>completed</option><option>in_progress</option><option>cancelled</option></select></div>
        <div class="field"><label>Loại vật liệu</label><select><option>Tất cả</option><option>Nhựa</option><option>Kim loại</option></select></div>
        <div class="field"><label>Định dạng</label><select><option>CSV</option><option>PDF demo</option></select></div>
      </div>
    </section>
    <section class="grid three" style="margin-top:18px">
      ${metric("Giao dịch", mock.transactions.length)}
      ${metric("Khối lượng", `${mock.transactions.reduce((sum, t) => sum + t.quantity, 0)} tấn`, "blue")}
      ${metric("Thanh toán online", "0", "orange")}
    </section>
    <section class="table-wrap" style="margin-top:18px">
      ${table(["Transaction", "Buyer", "Seller", "Material", "Quantity", "Status", "Payment"], mock.transactions.map((t) => [t.id, t.buyer, t.seller, t.material, `${t.quantity} ${t.unit}`, status(t.status), t.paymentStatus]))}
    </section>
  `;
}

function renderUseCases() {
  const grouped = useCases.reduce((acc, item) => {
    acc[item[2]] = acc[item[2]] || [];
    acc[item[2]].push(item);
    return acc;
  }, {});
  return `
    ${header("Coverage · UC001-UC059", "Use case matrix", "Ma trận này chứng minh prototype có mock data và màn hình cho toàn bộ use case, bao gồm hai UC bổ sung cho luồng demand.", `<button class="btn primary" data-view="dashboard">Bắt đầu demo</button>`)}
    ${Object.entries(grouped).map(([group, items]) => `
      <section class="panel" style="margin-bottom:18px">
        <div class="section-title"><h2>${group}</h2><span class="status blue">${items.length} use cases</span></div>
        <div class="usecase-list">
          ${items.map(([id, name, screen, coverage]) => `<article class="usecase-card"><strong>${id} · ${name}</strong><span class="muted">Màn: ${screen}</span><span>${coverage}</span></article>`).join("")}
        </div>
      </section>
    `).join("")}
  `;
}

function table(headers, rows) {
  return `<table><thead><tr>${headers.map((h) => `<th>${h}</th>`).join("")}</tr></thead><tbody>${rows.map((row) => `<tr>${row.map((cell) => `<td>${cell}</td>`).join("")}</tr>`).join("")}</tbody></table>`;
}

document.addEventListener("click", (event) => {
  const langButton = event.target.closest("[data-lang]");
  if (langButton) {
    state.lang = langButton.dataset.lang;
    render();
    return;
  }

  const roleButton = event.target.closest("[data-role]");
  if (roleButton) {
    state.role = roleButton.dataset.role;
    state.view = roleDefaultView[state.role] || "dashboard";
    render();
    return;
  }

  const viewButton = event.target.closest("[data-view]");
  if (viewButton) {
    state.view = viewButton.dataset.view;
    render();
    return;
  }

  const listingButton = event.target.closest("[data-listing]");
  if (listingButton) {
    state.selectedListing = listingButton.dataset.listing;
    render();
    return;
  }

  const demandButton = event.target.closest("[data-demand]");
  if (demandButton) {
    state.selectedDemand = demandButton.dataset.demand;
    render();
    return;
  }

  const action = event.target.closest("[data-action]")?.dataset.action;
  if (action === "mark-read") {
    mock.notifications.forEach((n) => { n.read = true; });
    render();
  }
  if (action === "accept-offer") {
    const offer = mock.offers.find((o) => o.id === "OFF-5001");
    if (offer) offer.status = "accepted";
    mock.notifications.unshift({ id: "NOT-DEMO", user: "EcoPack Vietnam", title: "OFF-5001 đã được chấp nhận", content: "Demo action tạo thông báo kết quả offer.", type: "offer_accepted", read: false, time: "Vừa xong" });
    render();
  }
  if (action === "complete-tx") {
    const tx = mock.transactions.find((t) => t.id === state.selectedTransaction);
    if (tx) {
      tx.buyerConfirmed = true;
      tx.sellerConfirmed = true;
      tx.status = "completed";
      mock.transactionEvents.push({ tx: tx.id, title: "Hai bên xác nhận hoàn tất", note: "transaction_status = completed", time: "Vừa xong", done: true });
    }
    render();
  }
});

document.addEventListener("change", (event) => {
  if (event.target.matches("[data-mobile-lang]")) {
    state.lang = event.target.value;
    render();
  }
  if (event.target.matches("[data-mobile-role]")) {
    state.role = event.target.value;
    state.view = roleDefaultView[state.role] || "dashboard";
    render();
  }
  if (event.target.matches("[data-mobile-view]")) {
    state.view = event.target.value;
    render();
  }
});

render();
