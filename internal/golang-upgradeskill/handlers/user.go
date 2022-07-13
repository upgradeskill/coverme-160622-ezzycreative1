package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	} else {
		// Request has an ID, as in "/task/<id>".
		path := strings.Trim(r.URL.Path, "/")
		pathParts := strings.Split(path, "/")
		if len(pathParts) < 2 {
			http.Error(rw, "expect /<id> in user handler", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathParts[1])
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodPut {
			u.updateUser(int(id), rw, r)
		} else if r.Method == http.MethodGet {
			u.getUser(rw, r)
		} else {
			http.Error(rw, fmt.Sprintf("expect method GET or DELETE at /task/<id>, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}

	// Handle UPDATE
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (u *users) getUser(rw http.ResponseWriter, r *http.Request) {

	lp := models.GetUsers()

	// This converts the lp object into json format andaa calls he Write method of rw.
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
