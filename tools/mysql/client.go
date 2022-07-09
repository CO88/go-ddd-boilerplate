package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CO88/go-ddd-boilerplate/config"
)

func NewMysqlConn(config config.Configuration) *sql.DB {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Db.UserName,
		config.Db.UserPassword,
		config.Db.Host,
		config.Db.Port,
		config.Db.Name)

	dbConn, err := sql.Open(`mysql`, connection)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return dbConn
}
