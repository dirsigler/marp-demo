package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/search", handleSearch)

	fmt.Println("Starting insecure web server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	rows, err := db.Query("SELECT * FROM products WHERE name LIKE '%" + query + "%'")
	if err != nil {
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			http.Error(w, "An error occurred", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Product: %s\n", name)
	}
}
