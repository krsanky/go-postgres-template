package db

import (
	"database/sql"
	"fmt"

	"github.com/krsanky/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres",
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			config.Get("db.user"), config.Get("db.password"), config.Get("db.name")))
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
}

func Test() {
	//age := 21
	//rows, err := db.Query("SELECT * FROM users WHERE age = $1", age)

	//rows, err := db.Query("SELECT * FROM account")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()

}
