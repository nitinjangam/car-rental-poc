package routingserver

import (
	"context"
	"log"

	"github.com/nitinjangam/car-rental-poc/definitions-store/availabilitypb"
	"github.com/nitinjangam/car-rental-poc/definitions-store/routingpb"
	"google.golang.org/grpc"
)

type Server struct {
	routingpb.UnimplementedRoutingServiceServer
}

func (*Server) GetAvailability(stream routingpb.RoutingService_GetAvailabilityServer) error {
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
	if err != nil {
		log.Fatalf("error while c.GetAvailability: %v", err.Error())
	}
	response := []*routingpb.RoutingAvailabilityResponse{}
	respAvailability := routingpb.ListRoutingAvailabilityResponse{}
	for {
		resp, err := strm.Recv()
		if err != nil {
			break
		}
		response = append(response, &routingpb.RoutingAvailabilityResponse{
			CarType:  resp.GetCarType(),
			Distance: resp.GetDistance(),
			Location: resp.GetLocation(),
		})
	}
	respAvailability.ListRoutingAvailabilityResponse = response
	stream.Send(&respAvailability)

	return nil
}
