package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

var MySQL = goqu.Dialect("mysql")

func SelectAll[T any](ctx context.Context, db *sql.DB, scan ScanFunc[T], ds *goqu.SelectDataset) ([]T, error) {
	sql, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}
	fmt.Println(sql)
	rows, err := db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanRows(rows, scan)
}

func Select[T any](ctx context.Context, db *sql.DB, scan ScanFunc[T], ds *goqu.SelectDataset) (T, error) {
	sql, args, err := ds.ToSQL()
	if err != nil {
		var zero T
		return zero, err
	}
	fmt.Println(sql)
	row := db.QueryRowContext(ctx, sql, args...)
	return scan(row)
}

type dataset interface {
	ToSQL() (string, []interface{}, error)
}

func Insert(ctx context.Context, db *sql.DB, ds *goqu.InsertDataset) (int, error) {
	result, err := ExecResult(ctx, db, ds)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func Exec[T dataset](ctx context.Context, db *sql.DB, ds T) error {
	_, err := ExecResult(ctx, db, ds)
	return err
}

func ExecResult[T dataset](ctx context.Context, db *sql.DB, ds T) (sql.Result, error) {
	sql, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}
	fmt.Println(sql)
	return db.ExecContext(ctx, sql, args...)
}

type Row interface {
	Scan(dest ...interface{}) error
}

type ScanFunc[T any] func(Row) (T, error)

func scanRows[T any](rows *sql.Rows, f ScanFunc[T]) ([]T, error) {
	var results []T
	for rows.Next() {
		result, err := f(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
