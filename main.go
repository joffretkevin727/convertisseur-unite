package main

import (
	"fmt"
	"net/http"
	"proj1/router"
)

// FONCTION PRINCIPAL
func main() {
	r := router.New()
	fmt.Println("http://localhost:8080/home")
	http.ListenAndServe(":8080", r)

}
