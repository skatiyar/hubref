package database

import (
	"context"
	"errors"
	"regexp"

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
	if err := instance.Ping(ctx); err != nil {
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
		tableOpts := &orm.CreateTableOptions{
			IfNotExists: true,
		}
		if err := db.Model(models[i]).CreateTable(tableOpts); err != nil {
			return err
		}
	}

	return nil
}

func GetDB() (*pg.DB, error) {
	if db == nil {
		return nil, errors.New("DB not initialized")
	}
	return db, nil
}

func NilRowError(err error) bool {
	return err.Error() == "pg: no rows in result set"
}

var uErrReg = regexp.MustCompile("violates unique constraint")

func UniqueConstraintError(err error) bool {
	return uErrReg.MatchString(err.Error())
}
