package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"notification-service/internal/messaging"
	"os"

	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"notification-service/internal/handler"
	"notification-service/internal/repository"
	"notification-service/internal/service"
	"notification-service/pb"
)

func main() {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbName := getEnv("DB_NAME", "notification_db")
	dbUser := getEnv("DB_USER", "cme")
	dbPass := getEnv("DB_PASSWORD", "")
	natsURL := getEnv("NATS_URL", "nats://localhost:4222")
	grpcPort := getEnv("GRPC_PORT", "50056")

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

	repo := repository.NewNotificationRepository(db)
	svc := service.NewNotificationService(repo)
	h := handler.NewNotificationHandler(svc)

	natsConnection, err := nats.Connect(natsURL)
	if err != nil {
		log.Printf("NATS unavailable; synchronous gRPC notifications remain active: %v", err)
	} else {
		defer natsConnection.Close()
		if err := messaging.SubscribeOrderEvents(natsConnection, svc); err != nil {
			log.Printf("Failed to subscribe to order events: %v", err)
		}
	}

	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterNotificationServiceServer(s, h)

	log.Printf("Notification service running on :%s", grpcPort)
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
