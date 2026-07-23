package messaging

import (
	"encoding/json"
	"fmt"
	"log"
	"notification-service/internal/service"

	"github.com/nats-io/nats.go"
)

type orderEvent struct {
	OfferID       string  `json:"offer_id"`
	TransactionID string  `json:"transaction_id"`
	BuyerID       string  `json:"buyer_id"`
	BuyerName     string  `json:"buyer_name"`
	SellerID      string  `json:"seller_id"`
	SellerName    string  `json:"seller_name"`
	ListingTitle  string  `json:"listing_title"`
	Quantity      float64 `json:"quantity"`
	Unit          string  `json:"unit"`
	ProposedPrice float64 `json:"proposed_price"`
}

func SubscribeOrderEvents(connection *nats.Conn, notifications *service.NotificationService) error {
	subscriptions := map[string]func(orderEvent) (string, string, string, string, string){
		"cme.orders.offer.created": func(event orderEvent) (string, string, string, string, string) {
			message := fmt.Sprintf(
				"%s muon mua %s %.0f %s voi gia %.0f VND/%s",
				event.BuyerName, event.ListingTitle, event.Quantity, event.Unit,
				event.ProposedPrice, event.Unit,
			)
			return event.SellerID, "De nghi mua moi", message, "offer", event.OfferID
		},
		"cme.orders.transaction.created": func(event orderEvent) (string, string, string, string, string) {
			message := fmt.Sprintf(
				"%s da chap nhan de nghi mua %s. Giao dich da duoc tao.",
				event.SellerName, event.ListingTitle,
			)
			return event.BuyerID, "De nghi da duoc chap nhan", message, "offer_accepted", event.OfferID
		},
		"cme.orders.transaction.in_progress": func(event orderEvent) (string, string, string, string, string) {
			message := fmt.Sprintf(
				"San pham %s da duoc giao. Vui long xac nhan nhan hang.",
				event.ListingTitle,
			)
			return event.BuyerID, "Seller da giao hang", message, "transaction", event.TransactionID
		},
		"cme.orders.transaction.completed": func(event orderEvent) (string, string, string, string, string) {
			message := fmt.Sprintf(
				"Giao dich %s da hoan tat. Tien da chuyen vao vi cua ban.",
				event.ListingTitle,
			)
			return event.SellerID, "Giao dich hoan tat", message, "transaction", event.TransactionID
		},
	}

	for subject, mapNotification := range subscriptions {
		subject := subject
		mapNotification := mapNotification
		if _, err := connection.QueueSubscribe(subject, "notification-service", func(message *nats.Msg) {
			var event orderEvent
			if err := json.Unmarshal(message.Data, &event); err != nil {
				log.Printf("invalid NATS event on %s: %v", subject, err)
				return
			}
			userID, title, body, notificationType, referenceID := mapNotification(event)
			if userID == "" {
				log.Printf("NATS event on %s has no recipient", subject)
				return
			}
			if _, err := notifications.CreateNotification(
				userID, title, body, notificationType, referenceID,
			); err != nil {
				log.Printf("failed to persist notification from %s: %v", subject, err)
			}
		}); err != nil {
			return fmt.Errorf("subscribe %s: %w", subject, err)
		}
	}

	return connection.Flush()
}
