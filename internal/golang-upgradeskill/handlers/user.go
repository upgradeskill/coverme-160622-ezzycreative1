package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-upgradeskill/pkg/models"
)

type users struct {
	l *log.Logger
}

func NewUsers(l *log.Logger) *users {
	return &users{l}
}

func (u *users) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// This will get the data as JSON format.
	if r.Method == http.MethodGet {
		u.getUser(rw, r)
		return
	} else if r.Method == http.MethodPost {
		u.addUser(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// Expect the ID in URI.
		// r := regexp.MustCompile(`/([0-9]+)`)
		// g := r.FindAllStringSubmatch(r.URL.Path, -1)
		u.l.Println("The request is ", r.Method)
		t := r.URL.Path

		if len(t) == 1 {
			u.l.Println("Invalid ID")
			return
		}
		id, _ := strconv.Atoi(string(t[1:]))

		u.l.Println("The id is ", id)
		u.updateUser(id, rw, r)
	}

	// Handle UPDATE
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *users) getUser(rw http.ResponseWriter, r *http.Request) {

	lp := models.GetUsers()

	// This converts the lp object into json format and calls he Write method of rw.
	err2 := json.NewEncoder(rw).Encode(lp)

	if err2 != nil {
		http.Error(rw, "Unable to marshall json", http.StatusInternalServerError)
	}
}

func (u *users) addUser(rw http.ResponseWriter, r *http.Request) {

	np := models.NewUser()

	err := json.NewDecoder(r.Body).Decode(np)

	if err != nil {
		fmt.Println("Something wrong")
	}

	u.l.Printf("Prod: %#v", np)

	models.AddUser(np)
}

func (u *users) updateUser(id int, rw http.ResponseWriter, r *http.Request) {

	np := models.NewUser()

	err := json.NewDecoder(r.Body).Decode(np)

	if err != nil {
		fmt.Println("Something wrong")
	}

	models.UpdateUser(id, *np)
}
