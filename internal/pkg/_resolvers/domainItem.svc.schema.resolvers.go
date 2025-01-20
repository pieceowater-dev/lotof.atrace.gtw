package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"app/internal/core/graph"
	"app/internal/core/graph/model"
	"context"
	"fmt"
)

// MutateSomething is the resolver for the mutateSomething field.
func (r *mutationResolver) MutateSomething(ctx context.Context, input model.MutateSomethingDto) (*model.Something, error) {
	panic(fmt.Errorf("not implemented: MutateSomething - mutateSomething"))
}

// Somethings is the resolver for the somethings field.
func (r *queryResolver) Somethings(ctx context.Context) ([]*model.Something, error) {
	return r.DomainItem.API.Somethings(ctx)
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }