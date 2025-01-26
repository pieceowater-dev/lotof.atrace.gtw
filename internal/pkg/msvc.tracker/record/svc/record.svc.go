package svc

import (
	"app/internal/core/cfg"
	pb "app/internal/core/grpc/generated/lotof.atrace.msvc.tracker/record"
	"github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
)

// RecordService handles the operations related to domain items.
type RecordService struct {
	transport gossiper.Transport
	client    pb.RecordServiceClient // gRPC client for domain item service.
}

// NewRecordService creates a new RecordService with the necessary transport and client.
func NewRecordService() *RecordService {
	factory := gossiper.NewTransportFactory()
	grpcTransport := factory.CreateTransport(
		gossiper.GRPC,
		cfg.Inst().LotofHubMSvcUsersGrpcAddress,
	)

	// Create the client only once and store it as a property.
	clientConstructor := pb.NewRecordServiceClient
	client, err := grpcTransport.CreateClient(clientConstructor)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	return &RecordService{
		transport: grpcTransport,
		client:    client.(pb.RecordServiceClient), // Cast to the correct type.
	}
}
