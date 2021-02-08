package db

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
)

func Conn(opts *pg.Options) *pg.DB {
	dbConn := pg.Connect(opts)
	dbConn.AddQueryHook(Logger{})
	return dbConn
}

type Logger struct{}

func (d Logger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (d Logger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}
