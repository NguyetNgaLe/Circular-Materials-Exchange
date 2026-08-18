# Changelog

All notable changes to Circular Materials Exchange are documented in this file.

The project follows [Semantic Versioning](https://semver.org/).

## [1.0.0] - 2026-08-18

### Added

- React and TypeScript user interface for businesses and administrators.
- Authentication, company registration, company approval, and role-based access.
- Supply and demand marketplaces with category and listing details.
- Company-owned supply creation, editing, hiding, image upload, and deletion.
- Offer creation, acceptance, rejection, and transaction generation.
- Transaction status tracking, transaction history, reviews, and notifications.
- Company settlement totals for received income and paid expenses.
- Administrative company, supply, transaction, escrow, withdrawal, and finance views.
- API Gateway and six Go gRPC microservices for authentication, companies, materials, orders, reviews, and notifications.
- PostgreSQL initialization for six logical databases and demo seed data.
- NATS event messaging and MinIO object storage integration.
- Docker Compose and Nginx deployment documentation.
- Material category and listing image set with source and license attribution.

### Changed

- Renamed the frontend source directory from `stitch-app` to `ui`.
- Updated React Router to the patched `6.30.6` release.
- Aligned company settlement with completed purchase and sale transactions.

### Security

- Kept service credentials in environment variables excluded from Git.
- Documented private-port firewall rules and SSH-tunneled database access.
- Updated the frontend routing dependency to address known redirect vulnerabilities in the previous version.

### Known limitations

- Settlement values are internal accounting records; the platform does not connect to an external payment gateway.
- HTTPS termination and automated CI/CD are deployment-environment responsibilities.

[1.0.0]: https://github.com/NguyetNgaLe/Circular-Materials-Exchange/releases/tag/v1.0.0
