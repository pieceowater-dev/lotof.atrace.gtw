package graph

import (
	"app/internal/pkg/msvc.tracker/post"
	"app/internal/pkg/msvc.tracker/record"
	"app/internal/pkg/msvc.tracker/route"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Post   post.Module
	Record record.Module
	Route  route.Module
}
