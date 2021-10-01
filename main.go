package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("DATABASE_HOST")
    port     = os.Getenv("DATABASE_PORT")
    user     = os.Getenv("DATABASE_USER")
    password = os.Getenv("DATABASE_PASSWORD")
    dbname   = os.Getenv("DATABASE_NAME")
)

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func main(){
	portint, _ := strconv.Atoi(port)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, portint, user, password, dbname)
	fmt.Println(psqlconn)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	createUsers := `CREATE TABLE IF NOT EXISTS users (name VARCHAR (255), email VARCHAR (255))`
	_, e := db.Exec(createUsers)
    CheckError(e)

	insertToUsers := `INSERT INTO users (name, email) VALUES('Lazuardy Khatulistiwa', 'lazuardy@nodeflux.io')`
	_, e = db.Exec(insertToUsers)
    CheckError(e)
}