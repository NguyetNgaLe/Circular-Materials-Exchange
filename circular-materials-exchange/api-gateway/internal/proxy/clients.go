package proxy

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	AuthConn         *grpc.ClientConn
	CompanyConn      *grpc.ClientConn
	MaterialConn     *grpc.ClientConn
	OrderConn        *grpc.ClientConn
	ReviewConn       *grpc.ClientConn
	NotificationConn *grpc.ClientConn
}

func NewGRPCClients() *GRPCClients {
	return &GRPCClients{}
}

func (c *GRPCClients) Connect() {
	c.AuthConn = connect(getEnv("AUTH_SERVICE_ADDR", "localhost:50051"))
	c.CompanyConn = connect(getEnv("COMPANY_SERVICE_ADDR", "localhost:50052"))
	c.MaterialConn = connect(getEnv("MATERIAL_SERVICE_ADDR", "localhost:50053"))
	c.OrderConn = connect(getEnv("ORDER_SERVICE_ADDR", "localhost:50054"))
	c.ReviewConn = connect(getEnv("REVIEW_SERVICE_ADDR", "localhost:50055"))
	c.NotificationConn = connect(getEnv("NOTIFICATION_SERVICE_ADDR", "localhost:50056"))
}

func (c *GRPCClients) Close() {
	if c.AuthConn != nil {
		c.AuthConn.Close()
	}
	if c.CompanyConn != nil {
		c.CompanyConn.Close()
	}
	if c.MaterialConn != nil {
		c.MaterialConn.Close()
	}
	if c.OrderConn != nil {
		c.OrderConn.Close()
	}
	if c.ReviewConn != nil {
		c.ReviewConn.Close()
	}
	if c.NotificationConn != nil {
		c.NotificationConn.Close()
	}
}

func connect(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Warning: Failed to connect to %s: %v", addr, err)
		return nil
	}
	log.Printf("Connected to %s", addr)
	return conn
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
