package main

import (
	"context"
	"database/sql"
	"fmt"
	"task/realization"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func FillDB(id int64, t time.Time) {
	db, err := sql.Open("sqlite3", "../data/data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	res, err := db.Exec("insert into control(userID, request_time) values($1, $2)", id, t)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
}

func main() {
	/*проверка*/
	FillDB(3, time.Now())
	FillDB(3, time.Now())
	FillDB(3, time.Now())
	FillDB(5, time.Now())
	FillDB(5, time.Now())
	u := realization.UserControl{
		N: 30,
		K: 3,
	}
	b, err := u.Check(context.Background(), 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	b, err = u.Check(context.Background(), 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}
