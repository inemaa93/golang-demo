package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var name string = "негр"
	var addr string = "127.0.0.1:8080"
	var password string = fmt.Sprintf("%s%s123", name, name)

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	// http server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		var htmlDoc string = fmt.Sprintf(`
	hashed password: %s
`, string(hashedPassword))

		_, _ = w.Write([]byte(htmlDoc))
	})

	log.Printf("listening on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
