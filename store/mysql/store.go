package mysql

import (
	"context"
	"database/sql"

	"github.com/deitrix/gqlgen-dataloadgen-example/graph/model"
	"github.com/deitrix/gqlgen-dataloadgen-example/store/mysql/db"
	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

var (
	usersTable = goqu.T("users")
	postsTable = goqu.T("posts")
)

func (s Store) CreateUser(ctx context.Context, user model.User) (int, error) {
	return db.Insert(ctx, s.db, db.MySQL.
		Insert(usersTable).
		Rows(goqu.Record{"name": user.Name}))
}

func (s Store) UpdateUser(ctx context.Context, id int, user model.User) error {
	return db.Exec(ctx, s.db, db.MySQL.
		Update(usersTable).
		Set(goqu.Record{"name": user.Name}).
		Where(usersTable.Col("id").Eq(id)))
}

func (s Store) DeleteUser(ctx context.Context, id int) error {
	return db.Exec(ctx, s.db, db.MySQL.
		Delete(usersTable).
		Where(usersTable.Col("id").Eq(id)))
}

func (s Store) Users(ctx context.Context) ([]model.User, error) {
	return db.SelectAll(ctx, s.db, scanUser, usersQuery())
}

func (s Store) UserByID(ctx context.Context, id int) (model.User, error) {
	return db.Select(ctx, s.db, scanUser, usersQuery().
		Where(usersTable.Col("id").Eq(id)))
}

func (s Store) UsersByIDs(ctx context.Context, ids []int) (map[int]model.User, error) {
	users, err := db.SelectAll(ctx, s.db, scanUser, usersQuery().
		Where(usersTable.Col("id").In(ids)))
	if err != nil {
		return nil, err
	}
	byID := make(map[int]model.User)
	for _, user := range users {
		byID[user.ID] = user
	}
	return byID, nil
}

func (s Store) CreatePost(ctx context.Context, userID int, post model.Post) (int, error) {
	return db.Insert(ctx, s.db, db.MySQL.
		Insert(postsTable).
		Rows(goqu.Record{"title": post.Title, "text": post.Text, "user_id": userID}))
}

func (s Store) UpdatePost(ctx context.Context, id int, post model.Post) error {
	return db.Exec(ctx, s.db, db.MySQL.
		Update(postsTable).
		Set(goqu.Record{"title": post.Title, "text": post.Text}).
		Where(postsTable.Col("id").Eq(id)))
}

func (s Store) DeletePost(ctx context.Context, id int) error {
	return db.Exec(ctx, s.db, db.MySQL.
		Delete(postsTable).
		Where(postsTable.Col("id").Eq(id)))
}

func (s Store) Posts(ctx context.Context) ([]model.Post, error) {
	return db.SelectAll(ctx, s.db, scanPost, postsQuery())
}

func (s Store) PostByID(ctx context.Context, id int) (model.Post, error) {
	return db.Select(ctx, s.db, scanPost, postsQuery().
		Where(postsTable.Col("id").Eq(id)))
}

func (s Store) PostsByIDs(ctx context.Context, ids []int) (map[int]model.Post, error) {
	posts, err := db.SelectAll(ctx, s.db, scanPost, postsQuery().
		Where(postsTable.Col("id").In(ids)))
	if err != nil {
		return nil, err
	}
	byID := make(map[int]model.Post)
	for _, post := range posts {
		byID[post.ID] = post
	}
	return byID, nil
}

func (s Store) PostsByUserID(ctx context.Context, userID int) ([]model.Post, error) {
	return db.SelectAll(ctx, s.db, scanPost, postsQuery().
		Where(postsTable.Col("user_id").Eq(userID)))
}

func (s Store) PostsByUserIDs(ctx context.Context, userIDs []int) (map[int][]model.Post, error) {
	posts, err := db.SelectAll(ctx, s.db, scanPost, postsQuery().
		Where(postsTable.Col("user_id").In(userIDs)))
	if err != nil {
		return nil, err
	}
	byUserID := make(map[int][]model.Post)
	for _, post := range posts {
		byUserID[post.UserID] = append(byUserID[post.UserID], post)
	}
	return byUserID, nil
}

func usersQuery() *goqu.SelectDataset {
	return db.MySQL.
		Select("id", "name").
		From("users")
}

func scanUser(row db.Row) (model.User, error) {
	var user model.User
	err := row.Scan(&user.ID, &user.Name)
	return user, err
}

func postsQuery() *goqu.SelectDataset {
	return db.MySQL.
		Select("id", "title", "text", "user_id").
		From("posts")
}

func scanPost(row db.Row) (model.Post, error) {
	var post model.Post
	err := row.Scan(&post.ID, &post.Title, &post.Text, &post.UserID)
	return post, err
}
