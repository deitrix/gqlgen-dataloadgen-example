package dataloader

import (
	"context"
	"errors"
	"net/http"

	"github.com/deitrix/gqlgen-dataloadgen-example/graph/model"
	"github.com/vikstrous/dataloadgen"

	example "github.com/deitrix/gqlgen-dataloadgen-example"
)

type Loaders struct {
	store example.Store

	postLoader      *dataloadgen.Loader[int, model.Post]
	userLoader      *dataloadgen.Loader[int, model.User]
	userPostsLoader *dataloadgen.Loader[int, []model.Post]
}

func New(store example.Store) *Loaders {
	l := &Loaders{store: store}
	l.postLoader = dataloadgen.NewLoader(l.posts)
	l.userLoader = dataloadgen.NewLoader(l.users)
	l.userPostsLoader = dataloadgen.NewLoader(l.usersPosts)
	return l
}

func (l *Loaders) posts(ctx context.Context, ids []int) ([]model.Post, []error) {
	posts, err := l.store.PostsByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}
	result := make([]model.Post, len(ids))
	errs := make([]error, len(ids))
	for i, id := range ids {
		var ok bool
		result[i], ok = posts[id]
		if !ok {
			errs[i] = errors.New("not found")
		}
	}
	return result, errs
}

func (l *Loaders) users(ctx context.Context, ids []int) ([]model.User, []error) {
	users, err := l.store.UsersByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}
	result := make([]model.User, len(ids))
	errs := make([]error, len(ids))
	for i, id := range ids {
		var ok bool
		result[i], ok = users[id]
		if !ok {
			errs[i] = errors.New("not found")
		}
	}
	return result, errs
}

func (l *Loaders) usersPosts(ctx context.Context, ids []int) ([][]model.Post, []error) {
	posts, err := l.store.PostsByUserIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}
	result := make([][]model.Post, len(ids))
	for i, id := range ids {
		var ok bool
		result[i], ok = posts[id]
		if !ok {
			result[i] = []model.Post{}
		}
	}
	return result, nil
}

type loadersKey struct{}

func Middleware(store example.Store) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), loadersKey{}, New(store))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey{}).(*Loaders)
}

func User(ctx context.Context, id int) (model.User, error) {
	return For(ctx).userLoader.Load(ctx, id)
}

func Post(ctx context.Context, id int) (model.Post, error) {
	return For(ctx).postLoader.Load(ctx, id)
}

func UserPosts(ctx context.Context, id int) ([]model.Post, error) {
	return For(ctx).userPostsLoader.Load(ctx, id)
}
