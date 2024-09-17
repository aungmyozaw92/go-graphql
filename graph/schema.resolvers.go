package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.52

import (
	"context"

	"github.com/aungmyozaw92/go-graphql/models"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input models.NewUser) (*models.User, error) {
	return models.CreateUser(ctx, &input)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	return models.GetUser(ctx, id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
