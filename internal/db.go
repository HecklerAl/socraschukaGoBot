package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// not done
// idk when will be done
// there are problems related to postgresql

func connect() (*sql.DB, error) {
	connStr := "user=postgres password=password dbname=links sslmode=disable"
	return sql.Open("postgres", connStr)
}

func AddShortLink(actualLink, schortLink string) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Query("CALL \"addShortLink\"('" + actualLink + "', '" + schortLink + "')")
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
