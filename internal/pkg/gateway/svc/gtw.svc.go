package svc

import (
	"app/internal/core/grpc/generated/generic/tenants"
	"context"
	gossiper "github.com/pieceowater-dev/lotof.lib.gossiper/v2"
	"log"
	"sync"
)

type GtwService struct {
	transport gossiper.Transport
	clients   []tenants.AppTenantsServiceClient
	mu        sync.Mutex
}

func NewGtwService(serverAddresses []string) *GtwService {
	factory := gossiper.NewTransportFactory()
	clients := make([]tenants.AppTenantsServiceClient, 0, len(serverAddresses))

	for _, address := range serverAddresses {
		grpcTransport := factory.CreateTransport(
			gossiper.GRPC,
			address,
		)

		clientConstructor := tenants.NewAppTenantsServiceClient
		client, err := grpcTransport.CreateClient(clientConstructor)
		if err != nil {
			log.Printf("Error creating client for %s: %v", address, err)
			continue
		}

		clients = append(clients, client.(tenants.AppTenantsServiceClient))
	}

	return &GtwService{
		clients: clients,
	}
}

func (s *GtwService) AddNamespaceTenant(ctx context.Context, req *tenants.AddNamespaceTenantRequest) (*tenants.AddNamespaceTenantResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var wg sync.WaitGroup
	results := make([]bool, len(s.clients))
	errors := make([]error, len(s.clients))

	for i, client := range s.clients {
		wg.Add(1)
		go func(i int, client tenants.AppTenantsServiceClient) {
			defer wg.Done()
			resp, err := client.NewTenant(ctx, &tenants.AddNamespaceTenantRequest{
				Namespace:   req.Namespace,
				Credentials: req.Credentials,
			})
			if err != nil {
				errors[i] = err
				log.Printf("Error calling NewTenant on client %d: %v", i, err)
				return
			}

			results[i] = resp.Success
		}(i, client)
	}

	wg.Wait()

	for i, success := range results {
		if !success {
			return nil, errors[i]
		}
	}

	return &tenants.AddNamespaceTenantResponse{Success: true}, nil
}
