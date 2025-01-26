package svc

import (
	"app/internal/core/cfg"
	pb "app/internal/core/grpc/generated/lotof.atrace.msvc.tracker/route"
	"github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

// RouteService handles the operations related to domain items.
type RouteService struct {
	transport gossiper.Transport
	client    pb.RouteServiceClient // gRPC client for domain item service.
}

// NewRouteService creates a new RouteService with the necessary transport and client.
func NewRouteService() *RouteService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	// Create the client only once and store it as a property.
	clientConstructor := pb.NewRouteServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &RouteService{
		transport: grpcTransport,
		client:    client.(pb.RouteServiceClient), // Cast to the correct type.
	}
}
