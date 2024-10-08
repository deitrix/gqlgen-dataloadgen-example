package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"

	"github.com/deitrix/gqlgen-dataloadgen-example/dataloader"
	"github.com/deitrix/gqlgen-dataloadgen-example/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, name string) (model.User, error) {
	id, err := r.store.CreateUser(ctx, model.User{Name: name})
	if err != nil {
		return model.User{}, err
	}
	return r.Query().User(ctx, id)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, name string) (model.User, error) {
	if err := r.store.UpdateUser(ctx, id, model.User{Name: name}); err != nil {
		return model.User{}, err
	}
	return r.Query().User(ctx, id)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	if err := r.store.DeleteUser(ctx, id); err != nil {
		return false, err
	}
	return true, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, userID int, title string, text string) (model.Post, error) {
	id, err := r.store.CreatePost(ctx, userID, model.Post{Title: title, Text: text})
	if err != nil {
		return model.Post{}, err
	}
	return r.Query().Post(ctx, id)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, id int, title string, text string) (model.Post, error) {
	if err := r.store.UpdatePost(ctx, id, model.Post{Title: title, Text: text}); err != nil {
		return model.Post{}, err
	}
	return r.Query().Post(ctx, id)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id int) (bool, error) {
	if err := r.store.DeletePost(ctx, id); err != nil {
		return false, err
	}
	return true, nil
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (model.User, error) {
	return dataloader.User(ctx, obj.UserID)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]model.User, error) {
	return r.store.Users(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (model.User, error) {
	return dataloader.User(ctx, id)
}

// UserPosts is the resolver for the userPosts field.
func (r *queryResolver) UserPosts(ctx context.Context, id int) ([]model.Post, error) {
	return dataloader.UserPosts(ctx, id)
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]model.Post, error) {
	return r.store.Posts(ctx)
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id int) (model.Post, error) {
	return dataloader.Post(ctx, id)
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *model.User) ([]model.Post, error) {
	return dataloader.UserPosts(ctx, obj.ID)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
