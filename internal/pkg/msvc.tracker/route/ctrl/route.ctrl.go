package ctrl

import (
	"app/internal/pkg/msvc.tracker/route/svc"
)

// RouteController handles the operations related to domain items.
type RouteController struct {
	trackerService *svc.RouteService // Service for fetching somethings.
}

// NewRouteController creates a new RouteController with the provided service.
func NewRouteController(service *svc.RouteService) *RouteController {
	return &RouteController{trackerService: service}
}
