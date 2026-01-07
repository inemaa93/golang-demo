package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var name string = "Dегр"
	var addr string = "127.0.0.1:8080"
	var password string = fmt.Sprintf("%s%s123", name, name)

	// Hello World
	fmt.Println("hello world")

	// Values
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Variables
	var a = "initial"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int
	fmt.Println(e)
	f := "apple"
	fmt.Print(f)
	
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
