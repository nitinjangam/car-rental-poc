package availabilityserver

import (
	"log"
	"time"

	"github.com/nitinjangam/car-rental-poc/availability-ms/data"
	"github.com/nitinjangam/car-rental-poc/availability-ms/utils"
	"github.com/nitinjangam/car-rental-poc/definitions-store/availabilitypb"
)

type Server struct {
	availabilitypb.UnimplementedAvailabilityServiceServer
}

func (*Server) GetAvailability(req *availabilitypb.AvailabilityRequest, strm availabilitypb.AvailabilityService_GetAvailabilityServer) error {
	usrLoc := req.GetSource()
	lat := usrLoc.Latitude
	lon := usrLoc.Longitude
	for _, c := range data.Carpool {
		var a, b int
		if c.Available {
			a, b = utils.RandomNumberGenerator()
			CarLocation := data.LocationName[a][b]
			m := utils.Distance(a, b, int(lat), int(lon))
			res := &availabilitypb.AvailabilityResponse{
				CarType:  c.Model,
				Location: CarLocation,
				Distance: m,
			}
			strm.Send(res)
			log.Printf("Sent:%v", res)
			time.Sleep(2 * time.Nanosecond)
		}
	}
	return nil
}
