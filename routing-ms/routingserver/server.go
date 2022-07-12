package routingserver

import (
	"context"
	"log"

	"github.com/nitinjangam/car-rental-poc/definitions-store/availabilitypb"
	"github.com/nitinjangam/car-rental-poc/definitions-store/routingpb"
	"google.golang.org/grpc"
)

type server struct {
	routingpb.UnimplementedRoutingServiceServer
}

func (*server) GetAvailability(stream routingpb.RoutingService_GetAvailabilityServer) error {
	req, err := stream.Recv()
	if err != nil {
		log.Fatalf("error while stream.Recv(): %v", err.Error())
	}
	availabilityReq := availabilitypb.AvailabilityRequest{
		Source: &availabilitypb.Location{
			Latitude:  req.Source.Latitude,
			Longitude: req.Source.Longitude,
		},
	}
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}

	c := availabilitypb.NewAvailabilityServiceClient(cc)
	strm, err := c.GetAvailability(context.Background(), &availabilityReq)
	response := []routingpb.RoutingAvailabilityResponse{}
	for {
		resp, err := strm.Recv()
		if err != nil {
			break
		}
		response = append(response, routingpb.RoutingAvailabilityResponse{
			CarType:  resp.GetCarType(),
			Distance: resp.GetDistance(),
			Location: resp.GetLocation(),
		})
	}

	return nil
}
