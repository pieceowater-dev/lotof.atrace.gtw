package gateway

import (
	"app/internal/core/cfg"
	"app/internal/pkg/gateway/ctrl"
	"app/internal/pkg/gateway/svc"
)

// Module represents the domain item module, including its name, version, and API controller.
type Module struct {
	name    string
	version string
	API     *ctrl.GtwController
}

// New creates a new instance of the Module, initializing the service and controller.
func New() Module {
	service := svc.NewGtwService(
		[]string{
			cfg.Inst().LotofHubMSvcUsersGrpcAddress,
		},
	)
	controller := ctrl.NewGtwController(service)

	return Module{
		name:    "gateway",
		version: "v1",
		API:     controller,
	}
}

// Initialize initializes the module. Currently not implemented.
func (m Module) Initialize() error {
	panic("Not implemented")
}

// Version returns the version of the module.
func (m Module) Version() string {
	return m.version
}

// Name returns the name of the module.
func (m Module) Name() string {
	return m.name
}
