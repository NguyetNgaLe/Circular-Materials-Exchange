.PHONY: build up down logs clean restart test

build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

logs-%:
	docker-compose logs -f $*

clean:
	docker-compose down -v --rmi all

restart-%:
	docker-compose restart $*

restart: down up

test:
	@echo "Testing API Gateway..."
	@curl -s http://localhost:8080/health | head -c 100
	@echo ""
	@echo "Testing Auth..."
	@curl -s -X POST http://localhost:8080/api/auth/login \
		-H "Content-Type: application/json" \
		-d '{"email":"admin@cme.vn","password":"admin123"}' | head -c 200
	@echo ""

status:
	docker-compose ps
