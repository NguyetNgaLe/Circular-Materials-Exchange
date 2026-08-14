\set ON_ERROR_STOP on

-- Demo marketplace data for a fresh deployment.
-- Fixed UUIDs and ON CONFLICT make this script safe to run more than once.

\connect company_db

BEGIN;

INSERT INTO companies (
    id, name, tax_code, address, city, description, status, owner_id,
    rating, review_count, member_since, certifications, image_url
) VALUES
(
    'c0000000-0000-0000-0000-000000000001',
    'EcoPoly Solutions', '0123456789', '123 Đường Lê Lợi, Quận 1',
    'TP. Hồ Chí Minh',
    'Chuyên thu gom và cung cấp nhựa, kim loại tái chế công nghiệp đã phân loại.',
    'verified', 'a0000000-0000-0000-0000-000000000002',
    4.70, 23, '2023-06-15', 'ISO 9001,Chứng nhận GRS',
    '/images/companies/ecopoly.jpg'
),
(
    'c0000000-0000-0000-0000-000000000002',
    'GreenPack Việt Nam', '9876543210', '456 Đường Nguyễn Huệ, Quận 3',
    'TP. Hồ Chí Minh',
    'Doanh nghiệp sản xuất bao bì tuần hoàn và kết nối tái sử dụng vật liệu sản xuất.',
    'verified', 'a0000000-0000-0000-0000-000000000003',
    4.50, 15, '2023-09-01', 'ISO 14001',
    '/images/companies/greenpack.jpg'
)
ON CONFLICT (id) DO NOTHING;

COMMIT;

\connect material_db

BEGIN;

UPDATE categories
SET image_url = CASE id
    WHEN '00000000-0000-0000-0000-000000000001' THEN '/images/categories/plastic.jpg'
    WHEN '00000000-0000-0000-0000-000000000002' THEN '/images/categories/metal.jpg'
    WHEN '00000000-0000-0000-0000-000000000003' THEN '/images/categories/paper.jpg'
    WHEN '00000000-0000-0000-0000-000000000004' THEN '/images/categories/wood.jpg'
    WHEN '00000000-0000-0000-0000-000000000005' THEN '/images/categories/textile.jpg'
    WHEN '00000000-0000-0000-0000-000000000006' THEN '/images/categories/glass.jpg'
    ELSE image_url
END
WHERE id IN (
    '00000000-0000-0000-0000-000000000001',
    '00000000-0000-0000-0000-000000000002',
    '00000000-0000-0000-0000-000000000003',
    '00000000-0000-0000-0000-000000000004',
    '00000000-0000-0000-0000-000000000005',
    '00000000-0000-0000-0000-000000000006'
);

INSERT INTO categories (id, name, icon, image_url) VALUES
(
    '00000000-0000-0000-0000-000000000007',
    'Phụ Phẩm Hữu Cơ', 'compost', '/images/listings/coffee_grounds.jpg'
)
ON CONFLICT (id) DO NOTHING;

INSERT INTO supply_listings (
    id, title, category_id, seller_id, company_id, description, specs,
    quantity, unit, price_per_unit, currency, location, min_order_quantity,
    packaging, status, images, image_url, created_at, updated_at
) VALUES
(
    '10000000-0000-0000-0000-000000000001', 'Phế liệu nhôm định hình loại 1',
    '00000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Nhôm thu hồi từ dây chuyền gia công, đã loại bỏ tạp chất và phân loại theo mác.',
    '{"Mác nhôm":"6061","Độ sạch":"> 98%","Dạng":"Đoạn và mảnh"}',
    8, 'Tấn', 35000, 'VND', 'Long An', 1, 'Bao jumbo 500kg', 'active',
    '/images/listings/aluminum_scrap.jpg', '/images/listings/aluminum_scrap.jpg',
    '2026-08-01 08:00:00', '2026-08-01 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000002', 'Thùng carton cũ đã ép kiện',
    '00000000-0000-0000-0000-000000000003', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Carton sóng sau sử dụng, khô sạch, đã tháo băng keo và ép kiện để vận chuyển.',
    '{"Loại":"Carton sóng","Độ ẩm":"< 8%","Tình trạng":"Đã ép kiện"}',
    12, 'Tấn', 3200, 'VND', 'Đồng Nai', 1, 'Kiện 500kg', 'active',
    '/images/listings/cardboard_old.jpg', '/images/listings/cardboard_old.jpg',
    '2026-08-02 08:00:00', '2026-08-02 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000003', 'Giấy carton thu hồi dạng tấm',
    '00000000-0000-0000-0000-000000000003', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Giấy carton dư từ công đoạn cắt bao bì, đồng đều và chưa qua sử dụng.',
    '{"Nguồn":"Dư sản xuất","Màu":"Nâu kraft","Tạp chất":"< 2%"}',
    7.5, 'Tấn', 4100, 'VND', 'Bình Dương', 0.5, 'Bó và pallet', 'active',
    '/images/listings/carton_paper.jpg', '/images/listings/carton_paper.jpg',
    '2026-08-03 08:00:00', '2026-08-03 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000004', 'Bã cà phê sau pha đã làm ráo',
    '00000000-0000-0000-0000-000000000007', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Bã cà phê phát sinh hằng ngày từ dây chuyền đồ uống, phù hợp ủ phân hoặc làm giá thể.',
    '{"Độ ẩm":"45-55%","Tần suất":"Hằng tuần","Tạp chất":"Không"}',
    3, 'Tấn', 800, 'VND', 'TP. Hồ Chí Minh', 0.2, 'Thùng kín 50kg', 'active',
    '/images/listings/coffee_grounds.jpg', '/images/listings/coffee_grounds.jpg',
    '2026-08-04 08:00:00', '2026-08-04 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000005', 'Chai thủy tinh màu đã phân loại',
    '00000000-0000-0000-0000-000000000006', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Chai thủy tinh thu hồi, đã bỏ nắp và phân loại theo màu xanh, nâu, trong.',
    '{"Tình trạng":"Nguyên chai","Phân loại":"Theo màu","Tạp chất":"< 3%"}',
    9, 'Tấn', 1500, 'VND', 'Bà Rịa - Vũng Tàu', 1, 'Bao jumbo', 'active',
    '/images/listings/glass_bottles.jpg', '/images/listings/glass_bottles.jpg',
    '2026-08-05 08:00:00', '2026-08-05 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000006', 'Mảnh thủy tinh cullet sạch',
    '00000000-0000-0000-0000-000000000006', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Cullet thủy tinh đã nghiền và sàng, kích thước đồng đều cho lò nấu thủy tinh.',
    '{"Kích thước":"10-40mm","Độ sạch":"> 97%","Kim loại":"Đã tách"}',
    14, 'Tấn', 2200, 'VND', 'Bình Dương', 2, 'Bao jumbo 1 tấn', 'active',
    '/images/listings/glass_cullet.jpg', '/images/listings/glass_cullet.jpg',
    '2026-08-06 08:00:00', '2026-08-06 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000007', 'Hạt nhựa HDPE tái sinh',
    '00000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'HDPE sau công nghiệp được rửa, tạo hạt đồng đều, phù hợp ép phun và ép đùn.',
    '{"MFI":"0.4-0.8 g/10 phút","Mật độ":"0.94-0.96 g/cm3","Màu":"Xám"}',
    20, 'Tấn', 18500, 'VND', 'TP. Hồ Chí Minh', 5, 'Bao 25kg', 'active',
    '/images/listings/hdpe_pellets.jpg', '/images/listings/hdpe_pellets.jpg',
    '2026-08-07 08:00:00', '2026-08-07 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000008', 'Sắt vụn công nghiệp đã phân loại',
    '00000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Sắt vụn từ gia công cơ khí, không lẫn rác sinh hoạt, phù hợp nấu luyện thép.',
    '{"Loại":"Sắt carbon","Độ sạch":"> 95%","Kích thước":"5-50cm"}',
    42, 'Tấn', 8500, 'VND', 'TP. Hồ Chí Minh', 5, 'Hàng rời', 'active',
    '/images/listings/iron_scrap.jpg', '/images/listings/iron_scrap.jpg',
    '2026-08-08 08:00:00', '2026-08-08 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000009', 'Pallet gỗ thông cũ còn sử dụng được',
    '00000000-0000-0000-0000-000000000004', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Pallet gỗ thông đã qua sử dụng, kết cấu chắc chắn, có thể tái dùng hoặc sửa chữa.',
    '{"Kích thước":"1200x1000mm","Loại gỗ":"Thông","Tình trạng":"Loại A/B"}',
    500, 'Cái', 45000, 'VND', 'Bình Dương', 50, 'Xếp chồng', 'active',
    '/images/listings/pallet_wood.jpg', '/images/listings/pallet_wood.jpg',
    '2026-08-09 08:00:00', '2026-08-09 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000010', 'Hạt nhựa PE tái sinh màu tự nhiên',
    '00000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Hạt PE tái sinh từ màng công nghiệp sạch, phù hợp sản xuất túi và màng phủ.',
    '{"Dạng":"Hạt","Màu":"Tự nhiên","Độ ẩm":"< 0.5%"}',
    18, 'Tấn', 16200, 'VND', 'Long An', 3, 'Bao 25kg', 'active',
    '/images/listings/pe_granule.jpg', '/images/listings/pe_granule.jpg',
    '2026-08-10 08:00:00', '2026-08-10 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000011', 'Mảnh nhựa PET tái chế đã rửa',
    '00000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Mảnh PET chai sau tiêu dùng đã rửa nóng, tách nhãn và kim loại.',
    '{"Dạng":"Mảnh","Màu":"Trong/xanh nhạt","PVC":"< 50ppm"}',
    15, 'Tấn', 12000, 'VND', 'Tây Ninh', 2, 'Bao jumbo', 'active',
    '/images/listings/pet_recycled.jpg', '/images/listings/pet_recycled.jpg',
    '2026-08-11 08:00:00', '2026-08-11 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000012', 'Hạt nhựa PP tái sinh',
    '00000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Nguồn PP tái sinh ổn định khoảng 20 tấn mỗi tháng, phù hợp ép phun sản phẩm kỹ thuật.',
    '{"MFI":"8-12 g/10 phút","Màu":"Đen","Nguồn cung":"20 tấn/tháng"}',
    20, 'Tấn', 15800, 'VND', 'Bình Dương', 2, 'Bao 25kg', 'active',
    '/images/listings/pp_granule.jpg', '/images/listings/pp_granule.jpg',
    '2026-08-12 08:00:00', '2026-08-12 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000013', 'Vải vụn cotton đã phân loại màu',
    '00000000-0000-0000-0000-000000000005', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Vải cotton dư từ công đoạn cắt may, đã phân loại theo màu, không lẫn phụ kiện kim loại.',
    '{"Thành phần":"Cotton > 95%","Kích thước":"10-50cm","Phân loại":"Theo màu"}',
    6, 'Tấn', 6500, 'VND', 'Tây Ninh', 0.5, 'Kiện 100kg', 'active',
    '/images/listings/sorted_fabric_scraps.jpg', '/images/listings/sorted_fabric_scraps.jpg',
    '2026-08-13 08:00:00', '2026-08-13 08:00:00'
),
(
    '10000000-0000-0000-0000-000000000014', 'Thép phế liệu sản xuất',
    '00000000-0000-0000-0000-000000000002', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Đầu mẩu và phoi thép từ nhà máy cơ khí, được tập kết riêng và không nhiễm dầu.',
    '{"Vật liệu":"Thép carbon","Dầu":"Không","Tạp chất":"< 3%"}',
    30, 'Tấn', 9200, 'VND', 'Đồng Nai', 5, 'Container/hàng rời', 'active',
    '/images/listings/steel_scrap.jpg', '/images/listings/steel_scrap.jpg',
    '2026-08-13 10:00:00', '2026-08-13 10:00:00'
),
(
    '10000000-0000-0000-0000-000000000015', 'Vải vụn dệt may hỗn hợp',
    '00000000-0000-0000-0000-000000000005', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Vải vụn cotton/polyester từ xưởng may, thích hợp làm giẻ lau hoặc tái chế sợi.',
    '{"Thành phần":"Cotton/Polyester","Độ sạch":"Khô sạch","Phụ kiện":"Đã loại bỏ"}',
    10, 'Tấn', 4200, 'VND', 'TP. Hồ Chí Minh', 1, 'Kiện 100kg', 'active',
    '/images/listings/textile_scraps.jpg', '/images/listings/textile_scraps.jpg',
    '2026-08-13 12:00:00', '2026-08-13 12:00:00'
),
(
    '10000000-0000-0000-0000-000000000016', 'Pallet gỗ dư thừa số lượng lớn',
    '00000000-0000-0000-0000-000000000004', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Lô pallet dư từ kho logistics, số lượng lớn, bán định kỳ theo tháng.',
    '{"Số lượng":"Hơn 1.000 cái","Tình trạng":"Cũ hỗn hợp","Giao hàng":"Theo lô"}',
    1200, 'Cái', 28000, 'VND', 'Đồng Nai', 100, 'Xếp chồng theo lô', 'active',
    '/images/listings/wood_pallet_yard.jpg', '/images/listings/wood_pallet_yard.jpg',
    '2026-08-14 08:00:00', '2026-08-14 08:00:00'
)
ON CONFLICT (id) DO NOTHING;

INSERT INTO demand_listings (
    id, title, category_id, buyer_id, company_id, description, quantity,
    unit, target_price, location, deadline, status, created_at, updated_at
) VALUES
(
    '20000000-0000-0000-0000-000000000001', 'Cần mua nhựa PP tái chế',
    '00000000-0000-0000-0000-000000000001', 'a0000000-0000-0000-0000-000000000003',
    'c0000000-0000-0000-0000-000000000002',
    'Tìm nguồn PP tái chế ổn định để sản xuất bao bì công nghiệp.',
    20, 'Tấn', 15500, 'TP. Hồ Chí Minh', '2026-12-31', 'open', NOW(), NOW()
),
(
    '20000000-0000-0000-0000-000000000002', 'Tìm pallet gỗ tái sử dụng',
    '00000000-0000-0000-0000-000000000004', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Cần pallet 1200x1000mm còn khả năng chịu tải để luân chuyển nội bộ.',
    1000, 'Cái', 35000, 'Bình Dương', '2026-11-30', 'open', NOW(), NOW()
),
(
    '20000000-0000-0000-0000-000000000003', 'Thu mua bã cà phê làm phân hữu cơ',
    '00000000-0000-0000-0000-000000000007', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Tìm nguồn bã cà phê sạch, giao định kỳ để phối trộn compost.',
    5, 'Tấn', 1000, 'TP. Hồ Chí Minh', '2026-12-15', 'open', NOW(), NOW()
),
(
    '20000000-0000-0000-0000-000000000004', 'Cần vải vụn cotton đã phân loại',
    '00000000-0000-0000-0000-000000000005', 'a0000000-0000-0000-0000-000000000002',
    'c0000000-0000-0000-0000-000000000001',
    'Cần vải cotton sạch để tái chế thành sợi và vật liệu chèn.',
    10, 'Tấn', 6000, 'Long An', '2026-12-20', 'open', NOW(), NOW()
)
ON CONFLICT (id) DO NOTHING;

COMMIT;
