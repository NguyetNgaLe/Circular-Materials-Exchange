# Circular Materials Exchange

Deployment guide for the complete Circular Materials Exchange platform. The system connects companies that have surplus materials with companies that need reusable materials, then tracks offers, transactions, reviews, notifications, and company income/expense settlement.

Live application: [http://42.1.71.99:13005](http://42.1.71.99:13005)

## 1. Architecture

```text
Browser
  |
  | HTTP :13005 (or HTTPS :443)
  v
Nginx
  |-- /                 -> React static files
  |-- /api/*            -> API Gateway :8085
  `-- /images/*         -> MinIO bucket cme-images :9000
                              ^
API Gateway :8085             |
  |                           |
  |-- Auth Service :50051     |
  |-- Company Service :50052  |
  |-- Material Service :50053-+
  |-- Order Service :50054
  |-- Review Service :50055
  `-- Notification Service :50056
          |             |
          v             v
   PostgreSQL :5432   NATS :4222
   (six databases)    (events)
```

The backend runs as Docker Compose services. The frontend is built with Vite and served as static files by Nginx. Browser requests use the relative `/api` path, so the frontend does not contain a server-specific backend URL.

## 2. Repository layout

```text
.
|-- circular-materials-exchange/  # Go microservices and Docker Compose
|   |-- api-gateway/
|   |-- auth-service/
|   |-- company-service/
|   |-- material-service/
|   |-- order-service/
|   |-- review-service/
|   |-- notification-service/
|   `-- scripts/                  # Database initialization and demo seed
|-- stitch-app/                   # React + TypeScript + Vite frontend
|-- images/                       # Seed images imported into MinIO
`-- README.md                     # This deployment guide
```

## 3. Requirements

Production host:

- Ubuntu 22.04 or newer
- 4 GB RAM minimum; 8 GB or more recommended
- Git
- Docker Engine 24 or newer with Docker Compose v2
- Node.js 20 LTS and npm
- Nginx
- `curl`, `rsync`, and `openssl`

Install the main packages on Ubuntu:

```bash
sudo apt update
sudo apt install -y git nginx curl rsync openssl
```

Install Docker Engine and Node.js from their official repositories. Confirm the tools before deploying:

```bash
docker --version
docker compose version
node --version
npm --version
nginx -v
```

For local development on Windows, Docker Desktop with WSL 2 and Node.js 20 are sufficient. A machine with 16 GB RAM can run the complete stack comfortably.

## 4. Network ports

| Port | Component | Public exposure |
| ---: | --- | --- |
| 22 | SSH | Allow only from trusted administration addresses |
| 13005 | Nginx web application | Public |
| 8085 | API Gateway | Private; Nginx proxies `/api` |
| 5433 | PostgreSQL host port | Private; use an SSH tunnel for pgAdmin |
| 9000 | MinIO API | Private; Nginx proxies `/images` |
| 9001 | MinIO console | Private |
| 4222, 8222 | NATS client and monitoring | Private |
| 50051-50056 | Internal gRPC services | Private |

The current Compose file publishes backend ports on the host for diagnostics. Keep them blocked by the host firewall. A customer does **not** need direct access to port `9000` to upload an image: the browser sends the file to `/api/upload`, and the API Gateway forwards it internally to the Material Service and MinIO.

Example UFW policy:

```bash
sudo ufw allow OpenSSH
sudo ufw allow 13005/tcp
sudo ufw status verbose
```

Do not enable UFW on a remote host until SSH access has been allowed and verified in a second terminal. Do not add public rules for the private ports in the table above.

## 5. First deployment

The examples use `/home/ubuntu/pttkht_git` as the repository directory and `/home/ubuntu/cme-frontend` as the frontend document root.

### 5.1 Clone the source

```bash
cd /home/ubuntu
git clone https://github.com/NguyetNgaLe/System-Analysis-and-Design.git pttkht_git
cd /home/ubuntu/pttkht_git
```

If the repository already exists:

```bash
cd /home/ubuntu/pttkht_git
git pull --ff-only origin main
```

### 5.2 Configure backend secrets

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
cp .env.example .env
chmod 600 .env
```

Edit `.env` and replace every placeholder. Random values can be generated with:

```bash
openssl rand -base64 36
```

Never commit `.env`, server passwords, private keys, or production database dumps.

Validate the resolved Compose configuration without printing the secret values into a shared log:

```bash
docker compose config --quiet
```

### 5.3 Start the backend

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose pull postgres nats minio
docker compose up -d --build
docker compose ps
```

On an empty PostgreSQL volume, Docker automatically runs:

1. `scripts/init-databases.sql`
2. `scripts/seed-demo-data.sql`

This creates `auth_db`, `company_db`, `material_db`, `order_db`, `review_db`, and `notif_db`. Initialization scripts run only when the `pg-data` volume is empty; editing an initialization SQL file does not migrate an existing database.

### 5.4 Create the MinIO bucket and import seed images

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose exec minio mc mb --ignore-existing local/cme-images
docker compose exec minio mc anonymous set download local/cme-images
docker compose cp ../images/. minio:/tmp/cme-seed-images
docker compose exec minio mc mirror --overwrite /tmp/cme-seed-images local/cme-images
```

The bucket name must remain `cme-images` because the Material Service writes uploaded files to that bucket. Public object download is required only through the Nginx `/images/` route; port `9000` remains blocked from the Internet.

### 5.5 Build the frontend

```bash
cd /home/ubuntu/pttkht_git/stitch-app
npm ci
npm run build
sudo mkdir -p /home/ubuntu/cme-frontend
sudo rsync -a --delete dist/ /home/ubuntu/cme-frontend/
sudo chown -R www-data:www-data /home/ubuntu/cme-frontend
```

`npm run build` must finish successfully before replacing the deployed directory.

### 5.6 Configure Nginx

Create `/etc/nginx/sites-available/circular-materials-exchange`:

```nginx
server {
    listen 13005;
    listen [::]:13005;
    server_name _;

    root /home/ubuntu/cme-frontend;
    index index.html;

    client_max_body_size 6m;

    location /api/ {
        proxy_pass http://127.0.0.1:8085/api/;
        proxy_http_version 1.1;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_connect_timeout 10s;
        proxy_read_timeout 60s;
    }

    location /images/ {
        proxy_pass http://127.0.0.1:9000/cme-images/;
        proxy_set_header Host 127.0.0.1:9000;
        proxy_cache_valid 200 1h;
        expires 1h;
    }

    location /assets/ {
        try_files $uri =404;
        expires 7d;
        add_header Cache-Control "public, immutable";
    }

    location / {
        try_files $uri $uri/ /index.html;
    }
}
```

Enable and validate the site:

```bash
sudo ln -sfn /etc/nginx/sites-available/circular-materials-exchange /etc/nginx/sites-enabled/circular-materials-exchange
sudo nginx -t
sudo systemctl reload nginx
sudo systemctl enable nginx
```

The `try_files ... /index.html` fallback is required for React Router pages such as `/transactions` and listing detail pages. Without it, reloading a frontend route returns a blank page or a 404.

## 6. Verification after deployment

Run all checks from the server:

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose ps
curl --fail http://127.0.0.1:8085/health
curl --fail http://127.0.0.1:8085/ready
curl --fail http://127.0.0.1:13005/
curl --fail http://127.0.0.1:13005/api/categories
```

Expected API responses:

```json
{"status":"ok"}
{"status":"ready"}
```

Check recent logs and restart counts:

```bash
docker compose logs --tail=100 api-gateway
docker compose logs --tail=100 material-service order-service
docker inspect -f '{{.Name}} restart={{.RestartCount}} status={{.State.Status}}' $(docker compose ps -q)
```

Then test the main user flows in a private browser window:

1. Register or log in.
2. Browse supply listings and open a listing detail page.
3. Create, edit, hide, and delete a company-owned listing.
4. Upload an image and confirm the returned `/images/...` URL loads.
5. Send and accept an offer.
6. Complete a transaction and verify company total received/paid values.
7. Check notifications and transaction history.
8. Check admin company, supply, and finance pages.

## 7. Updating an existing deployment

Always inspect the working tree before pulling:

```bash
cd /home/ubuntu/pttkht_git
git status --short
git pull --ff-only origin main
```

### Backend update

Rebuild the entire backend when shared contracts, initialization, or several services changed:

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose up -d --build
docker compose ps
curl --fail http://127.0.0.1:8085/ready
```

For a single service, build and replace only that container:

```bash
docker compose build order-service
docker compose up -d --no-deps order-service

docker compose build api-gateway
docker compose up -d --no-deps api-gateway
```

Use the same pattern for another service. Rebuild the API Gateway when its handlers, routes, or generated gRPC client contracts change.

### Frontend update with a recoverable backup

```bash
cd /home/ubuntu/pttkht_git/stitch-app
npm ci
npm run build

backup_dir="/home/ubuntu/backups/cme-frontend-$(date +%Y%m%d-%H%M%S)"
sudo mkdir -p "$backup_dir"
sudo rsync -a /home/ubuntu/cme-frontend/ "$backup_dir/"
sudo rsync -a --delete dist/ /home/ubuntu/cme-frontend/
sudo chown -R www-data:www-data /home/ubuntu/cme-frontend
sudo nginx -t
sudo systemctl reload nginx
```

The backup is owned by root because these commands use `sudo`. Delete it with `sudo` when it is no longer needed:

```bash
sudo rm -rf -- "/home/ubuntu/backups/cme-frontend-YYYYMMDD-HHMMSS"
```

Replace the placeholder with one exact backup directory after verifying it with `readlink -f`; never run a recursive delete against `/home/ubuntu/backups` itself.

### Image update

When seed files under `images/` change, mirror them into MinIO again:

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose cp ../images/. minio:/tmp/cme-seed-images
docker compose exec minio mc mirror --overwrite /tmp/cme-seed-images local/cme-images
```

## 8. Database administration

### pgAdmin through an SSH tunnel

Do not expose PostgreSQL port `5433` publicly. From an administrator machine, create a tunnel:

```bash
ssh -L 15433:127.0.0.1:5433 ubuntu@SERVER_IP
```

Keep the terminal open and configure pgAdmin with:

| Field | Value |
| --- | --- |
| Host | `127.0.0.1` |
| Port | `15433` |
| Maintenance database | `auth_db` |
| Username | `cme` |
| Password | Value of `DB_PASSWORD` in the server `.env` |

The same connection can access all six application databases.

### PostgreSQL shell

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose exec postgres psql -U cme -d auth_db
```

List databases and tables:

```sql
\l
\c order_db
\dt
```

### Database backup

```bash
mkdir -p /home/ubuntu/backups
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose exec -T postgres pg_dumpall -U cme --clean --if-exists > "/home/ubuntu/backups/cme-all-$(date +%Y%m%d-%H%M%S).sql"
```

Confirm that the resulting file is non-empty and copy it to separate storage.

### Database restore

Restoring can overwrite current data. Stop application traffic, verify the backup file, and take a fresh backup first:

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose stop api-gateway auth-service company-service material-service order-service review-service notification-service
docker compose exec -T postgres psql -U cme -d postgres < /home/ubuntu/backups/VERIFIED-BACKUP.sql
docker compose start auth-service company-service material-service order-service review-service notification-service api-gateway
curl --fail http://127.0.0.1:8085/ready
```

## 9. Local deployment

Start the backend from the repository root:

```bash
cd circular-materials-exchange
cp .env.example .env
# Replace all placeholders in .env
docker compose up -d --build
docker compose exec minio mc mb --ignore-existing local/cme-images
docker compose exec minio mc anonymous set download local/cme-images
docker compose cp ../images/. minio:/tmp/cme-seed-images
docker compose exec minio mc mirror --overwrite /tmp/cme-seed-images local/cme-images
```

Run the frontend development server in another terminal:

```bash
cd stitch-app
npm ci
npm run dev -- --host 0.0.0.0
```

Vite development mode needs an `/api` and `/images` proxy. For a production-like local test, build the frontend and use the Nginx configuration above. The backend health endpoint is available at `http://localhost:8085/health`.

Stop local services without deleting data:

```bash
cd circular-materials-exchange
docker compose down
```

Do not add `-v` unless the PostgreSQL and MinIO volumes are intentionally being deleted.

## 10. Automated build checks

Frontend:

```bash
cd stitch-app
npm ci
npm run build
```

All Go modules:

```bash
cd circular-materials-exchange
for module in api-gateway auth-service company-service material-service order-service review-service notification-service; do
  (cd "$module" && go test ./...) || exit 1
done
```

Compose validation:

```bash
cd circular-materials-exchange
docker compose config --quiet
```

## 11. Troubleshooting

### The frontend loads but API data is missing

```bash
curl -i http://127.0.0.1:8085/ready
curl -i http://127.0.0.1:13005/api/categories
sudo tail -n 100 /var/log/nginx/error.log
```

A `502 Bad Gateway` normally means that Nginx cannot reach the API Gateway or the API Gateway is restarting.

### A frontend route is blank after refresh

Confirm that Nginx uses `try_files $uri $uri/ /index.html;`, then run:

```bash
sudo nginx -t
sudo systemctl reload nginx
```

Also rebuild the frontend and check the browser console for JavaScript errors.

### Upload succeeds but the image does not load

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose exec minio mc ls local/cme-images/listings/
curl -I http://127.0.0.1:9000/cme-images/listings/FILE_NAME
curl -I http://127.0.0.1:13005/images/listings/FILE_NAME
```

Confirm that the bucket exists, anonymous download is enabled, and the Nginx `/images/` proxy includes the trailing slashes shown in the example configuration.

### PostgreSQL initialization changes are not applied

Initialization scripts only run for a new volume. Apply an explicit migration to the running database or create a verified backup before deliberately recreating the volume. Never remove `pg-data` just to retry a deployment.

### Permission denied when deleting a backup

Deployment backups created with `sudo` are root-owned. Inspect the exact directory first, then use `sudo` for that exact path. Avoid changing permissions recursively on all of `/home/ubuntu`.

### Service diagnostics

```bash
cd /home/ubuntu/pttkht_git/circular-materials-exchange
docker compose ps
docker compose logs --tail=200 SERVICE_NAME
docker stats --no-stream
df -h
free -h
```

## 12. Security checklist

- Keep `.env` and credentials outside Git.
- Allow public access only to the web port and tightly restrict SSH.
- Access PostgreSQL and the MinIO console through SSH tunnels or a private network.
- Use HTTPS and a real domain for production traffic.
- Rotate secrets after accidental exposure.
- Back up PostgreSQL and MinIO data before structural changes.
- Pin and regularly update container image versions for long-lived production use.
- Review `docker compose logs` and container restart counts after every deployment.

## 13. Quick command reference

```bash
# Pull source
cd /home/ubuntu/pttkht_git && git pull --ff-only origin main

# Deploy backend
cd circular-materials-exchange && docker compose up -d --build

# Deploy frontend
cd ../stitch-app && npm ci && npm run build
sudo rsync -a --delete dist/ /home/ubuntu/cme-frontend/

# Validate
curl --fail http://127.0.0.1:8085/ready
curl --fail http://127.0.0.1:13005/

# Follow logs
cd ../circular-materials-exchange && docker compose logs -f --tail=100
```
