package pkg

import (
	"app/internal/core/generic/interfaces"
	resolvers "app/internal/pkg/_resolvers"
	postSvc "app/internal/pkg/msvc.tracker/post"
	recordSvc "app/internal/pkg/msvc.tracker/record"
	routeSvc "app/internal/pkg/msvc.tracker/route"
	"google.golang.org/grpc"
	"reflect"
)

// Router manages the modules and initializes the routes for the application.
type Router struct {
	modules map[string]interfaces.IModule // Map of module names to their instances.
}

// NewRouter creates a new Router instance and initializes the domainItem module.
func NewRouter() *Router {
	postModule := postSvc.New()
	recordModule := recordSvc.New()
	routeModule := routeSvc.New()

	return &Router{
		modules: map[string]interfaces.IModule{
			postModule.Name():   postModule,
			recordModule.Name(): recordModule,
			routeModule.Name():  routeModule,
		},
	}
}

// InitializeRouter initializes all modules and returns the GraphQL resolver.
func (r *Router) InitializeRouter() (any, error) {
	// Initialize all modules
	resolver := r.initializeGQLResolvers()
	// r.initializeGRPCRoutes(r.server)
	return resolver, nil
}

// initializeGQLResolvers initializes the GraphQL resolvers for all modules.
func (r *Router) initializeGQLResolvers() *resolvers.Resolver {
	resolver := &resolvers.Resolver{}
	resolverValue := reflect.ValueOf(resolver).Elem()

	for i := 0; i < resolverValue.NumField(); i++ {
		field := resolverValue.Type().Field(i)
		moduleName := field.Name
		if module, ok := r.modules[moduleName]; ok {
			resolverValue.Field(i).Set(reflect.ValueOf(module))
		}
	}

	return resolver
}

// InitializeGRPCRoutes initializes the gRPC routes for all modules.
func (r *Router) InitializeGRPCRoutes(server *grpc.Server) {
	//pb.RegisterGatewayServiceServer(server, r.modules["NamespacesMod"].(ns.Module).API)
}

// GetModules returns the map of modules.
func (r *Router) GetModules() map[string]interfaces.IModule {
	return r.modules
}
