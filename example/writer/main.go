package main

import (
	"log"
	"os"

	paymentpb "github.com/steveanlorn/learning-protobuf/gen/go/payment/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	order := paymentpb.Order{
		OrderId:         "1",
		RecipientId:     "2",
		PaymentProvider: paymentpb.PaymentProvider_PAYMENT_PROVIDER_OVO,
	}

	b, err := proto.Marshal(&order)
	if err != nil {
		log.Fatalf("Could not marshal order: %v\n", err)
	}

	if err := os.WriteFile("order.bin", b, 0644); err != nil {
		log.Fatalf("Could not write order binary: %v\n", err)
	}
}
