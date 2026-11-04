package main

import (
	"challenge2/handler"
	"challenge2/repository"

	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sqlx.Connect("postgres",
		"host="+os.Getenv("DB_HOST")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("DB_USER")+
			" password="+os.Getenv("DB_PASSWORD")+
			" dbname="+os.Getenv("DB_NAME")+
			" sslmode="+os.Getenv("DB_SSLMODE"),
	)

	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// Test connection to DB
	if err := db.Ping(); err != nil {
		log.Fatal("cannot connect to DB", err)
	} else {
		log.Println("Successfully Connected")
	}

	h := &handler.Handler{
		Accounts: repository.NewAccountRepo(db),
		Transactions: repository.NewTransactionRepo(db),
	}

	log.Println("server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", h.Routes()))

}
