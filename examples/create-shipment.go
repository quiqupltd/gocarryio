package examples

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/quiqupltd/gocarryio"
)

const createShipmentBearerToken = "your-token-here"

func CreateShipment(ctx context.Context, shipmentId string) {
	// Create a new client
	client, err := gocarryio.NewClientWithResponses("https://api.carriyo.com")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	createdShipment, err := client.CreateShipmentWithResponse(
		ctx,
		&gocarryio.CreateShipmentParams{},
		gocarryio.CreateShipmentJSONRequestBody{
			References: gocarryio.ReferencesRequest{
				PartnerOrderReference: "1234567890",
			},
		},
		createShipmentRequestAuthFn,
	)

	fmt.Println(createdShipment.JSON200)
}

// function used to add the token to the request, required for authenticated requests
func createShipmentRequestAuthFn(_ context.Context, req *http.Request) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", createShipmentBearerToken))
	return nil
}
