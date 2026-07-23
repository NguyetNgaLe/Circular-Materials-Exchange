package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"order-service/internal/handler"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/pb"
)

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "order_db")
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "")
	natsURL := getEnv("NATS_URL", "nats://localhost:4222")
	grpcPort := getEnv("GRPC_PORT", "50054")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	repo := repository.NewOrderRepository(db)
	var eventPublisher service.NATSConn
	natsConnection, err := nats.Connect(natsURL)
	if err != nil {
		log.Printf("NATS unavailable; order processing continues without async events: %v", err)
	} else {
		defer natsConnection.Close()
		eventPublisher = natsConnection
	}
	svc := service.NewOrderService(repo, eventPublisher)
	h := handler.NewOrderHandler(svc)

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, h)
	reflection.Register(s)

	log.Printf("Order service running on :%s", grpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
