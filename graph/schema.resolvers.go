package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.52

import (
	"context"

	"github.com/aungmyozaw92/go-graphql/middlewares"
	"github.com/aungmyozaw92/go-graphql/models"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input models.NewUser) (*models.User, error) {
	return models.CreateUser(ctx, &input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*models.LoginInfo, error) {
	return models.Login(ctx, username, password)
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	return models.Logout(ctx)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	return models.CreateUser(ctx, &input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input models.NewUser) (*models.User, error) {
	return models.UpdateUser(ctx, id, &input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID int) (*models.User, error) {
	return models.DeleteUser(ctx, userID)
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, oldPassword string, newPassword string) (*models.User, error) {
	return models.ChangePassword(ctx, oldPassword, newPassword)
}

// CreateModule is the resolver for the createModule field.
func (r *mutationResolver) CreateModule(ctx context.Context, input models.NewModule) (*models.Module, error) {
	return models.CreateModule(ctx, &input)
}

// UpdateModule is the resolver for the updateModule field.
func (r *mutationResolver) UpdateModule(ctx context.Context, id int, input models.NewModule) (*models.Module, error) {
	return models.UpdateModule(ctx, id, &input)
}

// DeleteModule is the resolver for the deleteModule field.
func (r *mutationResolver) DeleteModule(ctx context.Context, id int) (*models.Module, error) {
	return models.DeleteModule(ctx, id)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id int) (*models.User, error) {
	return models.GetUser(ctx, id)
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context, name *string, phone *string, mobile *string, email *string, isActive *bool) ([]*models.User, error) {
	return models.GetUsers(ctx, name, phone, mobile, email, isActive)
}

// PaginateUser is the resolver for the paginateUser field.
func (r *queryResolver) PaginateUser(ctx context.Context, limit *int, after *string, name *string, phone *string, mobile *string, email *string, isActive *bool) (*models.UsersConnection, error) {
	return models.PaginateUser(ctx, limit, after, name, phone, mobile, email, isActive)
}

// GetModule is the resolver for the getModule field.
func (r *queryResolver) GetModule(ctx context.Context, id int) (*models.Module, error) {
	return models.GetModule(ctx, id)
}

// GetModules is the resolver for the getModules field.
func (r *queryResolver) GetModules(ctx context.Context, name *string) ([]*models.Module, error) {
	return models.GetModules(ctx, name)
}

// Role is the resolver for the role field.
func (r *userResolver) Role(ctx context.Context, obj *models.User) (*models.Role, error) {
	return middlewares.GetRole(ctx, obj.RoleId)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
