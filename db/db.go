package db

import (
	"database/sql"
	"fmt"

	"github.com/krsanky/config"
	"github.com/krsanky/lg"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	lg.Log.Printf("init pg db start ...")
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
	Test()
}

func Test() {
	var err error
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	rows, err := DB.Query("SELECT username FROM account")
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			panic(err)
		}
		lg.Log.Printf("username:%s", username)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
