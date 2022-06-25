package adapter

import (
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/jmoiron/sqlx"
	"github.com/phamdinhha/go-authorizer/config"
)

func CreateSqlxAdapter(db *sqlx.DB, cfg *config.Config) (*sqlxadapter.Adapter, error) {
	sqlxAdapter, err := sqlxadapter.NewAdapter(db, cfg.Casbin.PostgresTable)
	return sqlxAdapter, err
}
