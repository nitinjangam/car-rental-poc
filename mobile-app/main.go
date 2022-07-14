package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitinjangam/car-rental-poc/definitions-store/routingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userRequest struct {
	Location location `json:"source"`
}

type location struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type availabilityResponse struct {
	CarName  string `json:"carName"`
	Location string `json:"location"`
	Distance int    `json:"distance"`
}

type ratesRequest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
}

func main() {
	router := gin.Default()
	rg := router.Group("api/v1/carapp")
	{
		rg.GET("/car", fetchAvailableCarNearby)
		rg.GET("/rate", fetchRatesAvailableCars)
	}

	router.Run()

}

func fetchAvailableCarNearby(c *gin.Context) {

	var usrLoc userRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&usrLoc); err != nil {
		log.Fatalf("error while json unmarshal: %v", err.Error())
	}
	routingRequest := routingpb.RoutingAvailabilityRequest{
		Source: &routingpb.Location{
			Latitude:  int32(usrLoc.Location.Latitude),
			Longitude: int32(usrLoc.Location.Longitude),
		},
	}
	cc, err := grpc.Dial("0.0.0.0:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}

	client := routingpb.NewRoutingServiceClient(cc)
	strm, err := client.GetAvailability(context.Background())
	if err != nil {
		log.Fatalf("error while client.GetAvailability: %v", err.Error())
	}
	strm.Send(&routingRequest)
	resp := &routingpb.ListRoutingAvailabilityResponse{}
	for {
		resp, err = strm.Recv()
		if err != nil {
			break
		}
		c.JSON(http.StatusOK, &resp)
		break
	}
}

func fetchRatesAvailableCars(c *gin.Context) {
	var ratesReq ratesRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&ratesReq); err != nil {
		log.Fatalf("error while json unmarshal: %v", err.Error())
	}

	//prepare rates ms request
	reqRates := routingpb.RoutingRateRequest{
		Source:      ratesReq.Source,
		Destination: ratesReq.Destination,
		Date:        ratesReq.Date,
	}

	cc, err := grpc.Dial("0.0.0.0:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}
	client := routingpb.NewRoutingServiceClient(cc)
	strm, err := client.GetRate(context.Background(), &reqRates)
	if err != nil {
		log.Fatalf("Error while client.GetRate: %v", err)
	}

	resp := []*routingpb.RoutingRateResponse{}
	for {
		res, err := strm.Recv()
		if err != nil {
			break
		}
		resp = append(resp, res)
	}
	log.Printf("rates data received: %v", resp)
	c.JSON(http.StatusOK, &resp)
}
