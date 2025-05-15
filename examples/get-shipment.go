package examples

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/quiqupltd/gocarryio"
)

const bearerToken = "your-token-here"

func GetShipment(ctx context.Context, shipmentId string) {
	// Create a new client
	client, err := gocarryio.NewClientWithResponses("https://api.carriyo.com")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get the shipment, passing the auth functiion to set the correct headers
	shipment, err := client.GetShipmentWithResponse(ctx, shipmentId, &gocarryio.GetShipmentParams{}, requestAuthFn)
	if err != nil {
		log.Fatalf("Failed to get shipment: %v", err)
	}

	fmt.Println(shipment)
}

// function used to add the token to the request, required for authenticated requests
func requestAuthFn(_ context.Context, req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	return nil
}
