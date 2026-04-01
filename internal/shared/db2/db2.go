package db2

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ibmdb/go_ibm_db"
)

type DB2Client struct {
	DB *sql.DB
}

func NewDB2Client() (*DB2Client, error) {
	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatal("error loading .env file!")
	// }

	connStr := buildConnString()
	db, err := sql.Open("go_ibm_db", connStr)
	if err != nil {
		return nil, err
	}

	return &DB2Client{DB: db}, nil
}

func buildConnString() string {
	host := os.Getenv("DB2_HOST")
	user := os.Getenv("DB2_USER")
	pwd := os.Getenv("DB2_PWD")

	connStr := fmt.Sprintf("HOSTNAME=%s;DATABASE=BLUDB;PORT=50001;UID=%s;PWD=%s;SECURITY=SSL",
		host, user, pwd)
	log.Println("Connection string: " + connStr)

	return connStr
}
