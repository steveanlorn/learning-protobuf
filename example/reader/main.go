package main

import (
	"log"
	"os"

	paymentpb "github.com/steveanlorn/learning-protobuf/gen/go/payment/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	fname := "order.bin"
	b, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("Could not read '%s' file: %v\n", fname, err)
	}

	order := new(paymentpb.Order)
	if err := proto.Unmarshal(b, order); err != nil {
		log.Fatalf("Could not unmarshal order: %v\n", err)
	}

	log.Printf("Order ID: %s\n", order.GetOrderId())
	log.Printf("Recipient ID: %s\n", order.GetRecipientId())
	log.Printf("Payment provider: %s\n", order.GetPaymentProvider().String())
}
