package svc

import (
	"app/internal/core/cfg"
	pb "app/internal/core/grpc/generated/lotof.atrace.msvc.tracker/post"
	"github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

// PostService handles the operations related to domain items.
type PostService struct {
	transport gossiper.Transport
	client    pb.PostServiceClient // gRPC client for domain item service.
}

// NewPostService creates a new PostService with the necessary transport and client.
func NewPostService() *PostService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	// Create the client only once and store it as a property.
	clientConstructor := pb.NewPostServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &PostService{
		transport: grpcTransport,
		client:    client.(pb.PostServiceClient), // Cast to the correct type.
	}
}
