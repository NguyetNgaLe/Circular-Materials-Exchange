package proxy

import (
	authpb "api-gateway/internal/pb/auth"
	companypb "api-gateway/internal/pb/company"
	materialpb "api-gateway/internal/pb/material"
	notificationpb "api-gateway/internal/pb/notification"
	orderpb "api-gateway/internal/pb/order"
	reviewpb "api-gateway/internal/pb/review"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClients struct {
	Auth         authpb.AuthServiceClient
	Company      companypb.CompanyServiceClient
	Material     materialpb.MaterialServiceClient
	Order        orderpb.OrderServiceClient
	Review       reviewpb.ReviewServiceClient
	Notification notificationpb.NotificationServiceClient

	connections map[string]*grpc.ClientConn
}

func NewGRPCClients() *GRPCClients {
	return &GRPCClients{connections: make(map[string]*grpc.ClientConn)}
}

func (c *GRPCClients) Connect() error {
	var err error
	if c.connections["auth"], err = connect(getEnv("AUTH_SERVICE_ADDR", "auth-service:50051")); err != nil {
		return fmt.Errorf("auth service: %w", err)
	}
	if c.connections["company"], err = connect(getEnv("COMPANY_SERVICE_ADDR", "company-service:50052")); err != nil {
		return fmt.Errorf("company service: %w", err)
	}
	if c.connections["material"], err = connect(getEnv("MATERIAL_SERVICE_ADDR", "material-service:50053")); err != nil {
		return fmt.Errorf("material service: %w", err)
	}
	if c.connections["order"], err = connect(getEnv("ORDER_SERVICE_ADDR", "order-service:50054")); err != nil {
		return fmt.Errorf("order service: %w", err)
	}
	if c.connections["review"], err = connect(getEnv("REVIEW_SERVICE_ADDR", "review-service:50055")); err != nil {
		return fmt.Errorf("review service: %w", err)
	}
	if c.connections["notification"], err = connect(getEnv("NOTIFICATION_SERVICE_ADDR", "notification-service:50056")); err != nil {
		return fmt.Errorf("notification service: %w", err)
	}

	c.Auth = authpb.NewAuthServiceClient(c.connections["auth"])
	c.Company = companypb.NewCompanyServiceClient(c.connections["company"])
	c.Material = materialpb.NewMaterialServiceClient(c.connections["material"])
	c.Order = orderpb.NewOrderServiceClient(c.connections["order"])
	c.Review = reviewpb.NewReviewServiceClient(c.connections["review"])
	c.Notification = notificationpb.NewNotificationServiceClient(c.connections["notification"])
	return nil
}

func (c *GRPCClients) Close() {
	for _, conn := range c.connections {
		_ = conn.Close()
	}
}

func (c *GRPCClients) Ready() bool {
	if len(c.connections) != 6 {
		return false
	}
	for _, conn := range c.connections {
		state := conn.GetState()
		if state == connectivity.Shutdown || state == connectivity.TransientFailure {
			return false
		}
	}
	return true
}

func connect(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{MinConnectTimeout: 5 * time.Second}),
	)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
