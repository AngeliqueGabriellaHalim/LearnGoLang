package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type User struct{
	Name string `db:"name"`
	Email string `db:"email"`
}
func main(){
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

    // Test the connection to the database
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully Connected")
    }


	place := User{}

	rows,err := db.Queryx("SELECT name, email FROM users")

	for rows.Next(){
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("%#v\n", place)
	}

	var users []User

	err2 := db.Select(&users,
	"SELECT name, email FROM users",)
	
	if err2 != nil {
		log.Fatalln(err)
	}
	
	for  _,user :=  range users{
		fmt.Println(user.Name)
	}
}