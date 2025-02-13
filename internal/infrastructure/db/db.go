package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/koh-yoshimoto/go-testable-project/pkg/config"
)

var DB *sql.DB

func InitDB() error {

	dbConfig := config.DB

	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	DB = db

	return nil
}
