package data

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func ConnectSQLDb() *sql.DB {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to open db, error: %s", err)
		os.Exit(1)
	}
	return db
}

// type User struct {
// 	ID   int
// 	Name string
// }

// func queryUsers(db *sql.DB) {
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer rows.Close()

// 	var users []User

// 	for rows.Next() {
// 		var user User

// 		if err := rows.Scan(&user.ID, &user.Name); err != nil {
// 			fmt.Println("Error scanning row:", err)
// 			return
// 		}

// 		users = append(users, user)
// 		fmt.Println(user.ID, user.Name)
// 	}

// 	if err := rows.Err(); err != nil {
// 		fmt.Println("Error during rows iteration:", err)
// 	}
// }
