workspace "Circular Materials Exchange" "D11 - Logical Architecture as-built" {
    model {
        guest = person "Guest" "Xem marketplace và nhu cầu."
        businessUser = person "Business User" "Doanh nghiệp đóng vai Buyer hoặc Seller."
        admin = person "Admin" "Duyệt doanh nghiệp và quản trị finance/escrow."

        cme = softwareSystem "Circular Materials Exchange" "Nền tảng trao đổi vật liệu tuần hoàn B2B." {
            web = container "React SPA" "Giao diện marketplace, doanh nghiệp, giao dịch và quản trị." "React + TypeScript + Vite"
            gateway = container "API Gateway" "REST route, bearer-token middleware, HTTP handlers và gRPC clients." "Go + Gin"

            auth = container "Auth Service" "Tài khoản, JWT và xác thực token." "Go + gRPC"
            company = container "Company Service" "Hồ sơ và phê duyệt doanh nghiệp." "Go + gRPC"
            material = container "Material Service" "Danh mục, nguồn cung và nhu cầu." "Go + gRPC"
            order = container "Order Service" "Offer, Transaction và event publish." "Go + gRPC"
            review = container "Review Service" "Đánh giá và điểm uy tín." "Go + gRPC"
            notification = container "Notification Service" "Thông báo trong ứng dụng." "Go + gRPC"

            authDb = container "auth_db" "User." "PostgreSQL" "Database"
            companyDb = container "company_db" "Company." "PostgreSQL" "Database"
            materialDb = container "material_db" "Category, SupplyListing, DemandListing." "PostgreSQL" "Database"
            orderDb = container "order_db" "Offer, Transaction, Event, finance, escrow và wallet." "PostgreSQL" "Database"
            reviewDb = container "review_db" "Review." "PostgreSQL" "Database"
            notifDb = container "notif_db" "Notification." "PostgreSQL" "Database"
            nats = container "NATS" "Kênh event bất đồng bộ của Order." "NATS JetStream"
            minio = container "MinIO" "Lưu ảnh vật liệu." "MinIO"
        }

        guest -> web "Xem dữ liệu công khai" "HTTPS"
        businessUser -> web "Sử dụng nghiệp vụ" "HTTPS"
        admin -> web "Quản trị" "HTTPS"
        web -> gateway "REST API" "HTTP/JSON"

        gateway -> auth "Login/Register/VerifyToken/GetUser" "gRPC"
        gateway -> company "Company CRUD và approval" "gRPC"
        gateway -> material "Category/Listing/Demand/Upload" "gRPC"
        gateway -> order "Offer/Transaction/Finance/Escrow" "gRPC"
        gateway -> review "Review và rating" "gRPC"
        gateway -> notification "Notification đồng bộ" "gRPC"

        auth -> authDb "Đọc/ghi theo thiết kế" "SQL"
        company -> companyDb "Đọc/ghi theo thiết kế" "SQL"
        material -> materialDb "Đọc/ghi theo thiết kế" "SQL"
        order -> orderDb "Đọc/ghi theo thiết kế" "SQL"
        review -> reviewDb "Đọc/ghi theo thiết kế" "SQL"
        notification -> notifDb "Đọc/ghi theo thiết kế" "SQL"
        order -> nats "Publish cme.orders.*" "NATS"
        notification -> nats "Subscribe queue group" "NATS"
        material -> minio "Upload và lưu object ảnh" "HTTP PUT"
    }

    views {
        systemContext cme "SystemContext" {
            include *
            autolayout lr
        }

        container cme "Containers" {
            include *
            autolayout lr
        }

        styles {
            element "Person" {
                background #1565C0
                color #FFFFFF
                shape Person
            }
            element "Software System" {
                background #2E7D32
                color #FFFFFF
            }
            element "Container" {
                background #43A047
                color #FFFFFF
            }
            element "Database" {
                shape Cylinder
                background #6D4C41
                color #FFFFFF
            }
        }
    }
}
