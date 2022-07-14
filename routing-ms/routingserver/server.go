package routingserver

import (
	"context"
	"log"

	"github.com/nitinjangam/car-rental-poc/definitions-store/availabilitypb"
	"github.com/nitinjangam/car-rental-poc/definitions-store/ratepb"
	"github.com/nitinjangam/car-rental-poc/definitions-store/routingpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
		log.Printf("received:%v", resp)
	}
	respAvailability.ListRoutingAvailabilityResponse = response
	stream.Send(&respAvailability)

	return nil
}

func (*Server) GetRate(req *routingpb.RoutingRateRequest, stream routingpb.RoutingService_GetRateServer) error {
	dateRequested := req.GetDate()
	destinationRequested := req.GetDestination()
	sourceRequested := req.GetSource()

	//prepare request for rates ms
	ratesMsRequest := ratepb.RateRequest{
		Source:      sourceRequested,
		Destination: destinationRequested,
		Date:        dateRequested,
	}

	//dial on grpc server where rates ms gRPC server is running
	cc, err := grpc.Dial("0.0.0.0:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while Dial: %v", err)
	}

	c := ratepb.NewRoutingServiceClient(cc)

	strm, err := c.GetRate(context.Background(), &ratesMsRequest)
	if err != nil {
		log.Fatalf("error while c.GetRate() : %v", err.Error())
	}

	for {
		rates, err := strm.Recv()
		if err != nil {
			break
		}
		ratesresp := routingpb.RoutingRateResponse{
			CarType:  rates.GetCarType(),
			Location: rates.GetLocation(),
			Distance: rates.GetDistance(),
			Price:    rates.GetPrice(),
		}
		if err := stream.Send(&ratesresp); err != nil {
			log.Fatalf("error while stream.send() : %v", err.Error())
		}
	}

	return nil
}
