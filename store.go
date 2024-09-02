package example

import (
	"context"

	"github.com/deitrix/gqlgen-dataloadgen-example/graph/model"
)

type Store interface {
	CreateUser(ctx context.Context, user model.User) (int, error)
	UpdateUser(ctx context.Context, id int, user model.User) error
	DeleteUser(ctx context.Context, id int) error
	Users(ctx context.Context) ([]model.User, error)
	UserByID(ctx context.Context, id int) (model.User, error)
	UsersByIDs(ctx context.Context, ids []int) (map[int]model.User, error)

	CreatePost(ctx context.Context, userID int, post model.Post) (int, error)
	UpdatePost(ctx context.Context, id int, post model.Post) error
	DeletePost(ctx context.Context, id int) error
	Posts(ctx context.Context) ([]model.Post, error)
	PostByID(ctx context.Context, id int) (model.Post, error)
	PostsByIDs(ctx context.Context, ids []int) (map[int]model.Post, error)
	PostsByUserID(ctx context.Context, userID int) ([]model.Post, error)
	PostsByUserIDs(ctx context.Context, userIDs []int) (map[int][]model.Post, error)
}
