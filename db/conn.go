package db

import (
	"database/sql"
	"fmt"

	// postgres
	// _ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host     = "qrmemorial.cdmo8e8k81x8.us-east-1.rds.amazonaws.com"
	port     = 3306
	user     = "admin"
	password = ">8!>1%9axOK]H$+>N$SwNzNHOL<y"
	dbname   = "mkt_db"
)

func ConnectDB() (*sql.DB, error) {

	//define a connection string
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	//abre a conexão com o postgres
	// db, err := sql.Open("postgres", psqlInfo)

	//defina a connection string para o mysql
	psqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, dbname)

	//abre a conexão com o mysql
	db, err := sql.Open("mysql", psqlInfo)

	if err != nil {
		panic(err)
	}

	//executa um ping para checar se a conexão responde
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to" + dbname)

	return db, nil
}
