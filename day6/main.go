package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "hello\n hello")
}

func headers(w http.ResponseWriter, req *http.Request){
	 for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    user := User{
        ID:    1,
        Name:  "Angie",
        Email: "angie@example.com",
    }

    data, err := json.Marshal(user)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(data))

	jsonData := `{
		"id": 1,
		"name": "Angie",
		"email": "angie@example.com"
	}`

	var user2 User

	err2 := json.Unmarshal([]byte(jsonData), &user2)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(user2.Name)
	fmt.Println(user2.Email)

}

// func main() {
//     // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//     //     fmt.Fprintln(w, "Hello, World!")
//     // })
// 	// http.HandleFunc("/hello", hello)
//     // http.HandleFunc("/headers", headers)


//     // http.ListenAndServe(":8080", nil)

	
// }