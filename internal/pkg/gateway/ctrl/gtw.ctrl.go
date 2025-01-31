package ctrl

import (
	"app/internal/core/grpc/generated/generic/tenants"
	"app/internal/pkg/gateway/svc"
	"context"
)

type GtwController struct {
	service *svc.GtwService
	tenants.UnimplementedGatewayServiceServer
}

func NewGtwController(service *svc.GtwService) *GtwController {
	return &GtwController{service: service}
}

func (c *GtwController) AddNamespaceTenant(ctx context.Context, req *tenants.AddNamespaceTenantRequest) (*tenants.AddNamespaceTenantResponse, error) {
	return c.service.AddNamespaceTenant(ctx, req)
}
