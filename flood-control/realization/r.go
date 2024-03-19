package realization

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type FloodControl interface {
	Check(ctx context.Context, userID int64) (bool, error)
}

type UserControl struct {
	N int64
	K int64
}

func (u UserControl) Check(ctx context.Context, userID int64) (bool, error) {
	db, err := sql.Open("sqlite3", "../data/data.db")
	if err != nil {
		return false, err
	}
	defer db.Close()

	currentTime := time.Now()

	stTime := currentTime.Add(-time.Duration(u.N) * time.Second)
	endTime := currentTime
	fmt.Println(stTime, endTime)
	row := db.QueryRowContext(ctx, "SELECT COUNT(userID) as c FROM control WHERE userID=$1 AND request_time BETWEEN $2 AND $3", userID, stTime, endTime)
	var c int64
	row.Scan(&c)
	fmt.Println(c)
	if c >= u.K {
		return false, nil
	}
	return true, nil
}
