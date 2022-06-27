package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-upgradeskill/internal/golang-upgradeskill/handlers"
)

func main() {
	l := log.New(os.Stdout, "user-service", log.LstdFlags)
	p := handlers.NewUsers(l)

	mux := http.NewServeMux()
	mux.Handle("/", p)
	addr := ":8000"
	fmt.Printf("Server run on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
