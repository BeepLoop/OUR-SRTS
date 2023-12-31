package store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/BeepLoop/registrar-digitized/config"
)

var Db_Conn *sqlx.DB

func Initialize() error {
	db, err := sqlx.Connect("mysql", config.Env.DSN)
	if err != nil {
		return err
	}

	Db_Conn = db

	return nil
}
