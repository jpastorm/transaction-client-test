package main

import (
	"database/sql"
	"fmt"

	"github.com/jpastorm/transaction-client-test/model"

	_ "github.com/lib/pq"
)

func newDBConnection(dbConf model.Database) (*sql.DB, error) {
	if dbConf.SSLMode == "" {
		dbConf.SSLMode = "disable"
	}

	dns := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConf.User,
		dbConf.Password,
		dbConf.Server,
		dbConf.Port,
		dbConf.Name,
		dbConf.SSLMode,
	)

	return sql.Open("postgres", dns)
}
