package database

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var db *pg.DB

func Connect(dbURL string) error {
	opts, optsErr := pg.ParseURL(dbURL)
	if optsErr != nil {
		return optsErr
	}

	instance := pg.Connect(opts)
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return err
	}

	db = instance

	return nil
}

func CreateSchema() error {
	models := []interface{}{
		(*Path)(nil),
	}

	for i := 0; i < len(models); i += 1 {
		if err := db.Model(models[i]).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			return err
		}
	}

	return nil
}
