package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "golang"
)

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func main(){
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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