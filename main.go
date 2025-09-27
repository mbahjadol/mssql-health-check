package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func getConnString() string {
	user := os.Getenv("MSSQL_USER")
	pass := os.Getenv("MSSQL_PASS")
	host := os.Getenv("MSSQL_HOST")
	port := os.Getenv("MSSQL_PORT")
	db := os.Getenv("MSSQL_DB")

	if port == "" {
		port = "1433"
	}
	if db == "" {
		db = "master"
	}

	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, pass, host, port, db,
	)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	connString := getConnString()

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB open error: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var result int
	err = db.QueryRow("SELECT 1").Scan(&result)
	if err != nil || result != 1 {
		http.Error(w, fmt.Sprintf("DB query error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting MSSQL healthcheck service on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
