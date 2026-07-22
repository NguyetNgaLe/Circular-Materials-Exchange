workspace "Circular Materials Exchange" "D11 - Logical Architecture" {
    model {
        businessUser = person "Business User" "Doanh nghiệp đóng vai Buyer hoặc Seller."
        admin = person "Admin" "Quản trị và kiểm duyệt hệ thống."

        cme = softwareSystem "Circular Materials Exchange" "Nền tảng trao đổi vật liệu tuần hoàn B2B." {
            web = container "React SPA" "Giao diện marketplace, doanh nghiệp và quản trị." "React + TypeScript + Vite"
            gateway = container "API Gateway" "REST API, JWT middleware và gRPC clients." "Go + Gin"

            auth = container "Auth Service" "Tài khoản, đăng nhập và JWT." "Go + gRPC"
            company = container "Company Service" "Hồ sơ và phê duyệt doanh nghiệp." "Go + gRPC"
            material = container "Material Service" "Danh mục, nguồn cung và nhu cầu." "Go + gRPC"
            order = container "Order Service" "Offer, Transaction và timeline." "Go + gRPC"
            review = container "Review Service" "Đánh giá và điểm uy tín." "Go + gRPC"
            notification = container "Notification Service" "Thông báo trong ứng dụng." "Go + gRPC"

            authDb = container "auth_db" "Lưu User." "PostgreSQL" "Database"
            companyDb = container "company_db" "Lưu Company." "PostgreSQL" "Database"
            materialDb = container "material_db" "Lưu Category, SupplyListing, DemandListing." "PostgreSQL" "Database"
            orderDb = container "order_db" "Lưu Offer, Transaction, TransactionEvent." "PostgreSQL" "Database"
            reviewDb = container "review_db" "Lưu Review." "PostgreSQL" "Database"
            notifDb = container "notif_db" "Lưu Notification." "PostgreSQL" "Database"
            nats = container "NATS" "Event bus; Order Service đang publish sự kiện giao dịch." "NATS"
            minio = container "MinIO" "Object storage được khai báo trong Compose." "MinIO"
        }

        businessUser -> web "Sử dụng" "HTTPS"
        admin -> web "Quản trị" "HTTPS"
        web -> gateway "Gọi REST API" "HTTP/JSON"

        gateway -> auth "Gọi" "gRPC"
        gateway -> company "Gọi" "gRPC"
        gateway -> material "Gọi" "gRPC"
        gateway -> order "Gọi" "gRPC"
        gateway -> review "Gọi" "gRPC"
        gateway -> notification "Gọi" "gRPC"

        auth -> authDb "Đọc/ghi" "SQL"
        company -> companyDb "Đọc/ghi" "SQL"
        material -> materialDb "Đọc/ghi" "SQL"
        order -> orderDb "Đọc/ghi" "SQL"
        review -> reviewDb "Đọc/ghi" "SQL"
        notification -> notifDb "Đọc/ghi" "SQL"
        order -> nats "Publish transaction events"
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

