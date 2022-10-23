package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func migrate(databaseUrl string) error {
	fmt.Println("Started migrations...")

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return err
	}

	var migrations = []string{"./migrations/setup.sql", "./migrations/migration-0.0.1-users.sql"}

	for _, migration := range migrations {
		fmt.Println("Running " + migration)

		data, err := os.ReadFile(migration)
		if err != nil {
			return err
		}

		rows, err := db.Query(string(data))
		if err != nil {
			return err
		}
		defer rows.Close()	
	}

	fmt.Println("Finished migrations...")
	return nil
}