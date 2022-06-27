package models

import (
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Users []*User

func UpdateUser(id int, u User) {

	for i := 0; i < len(UserList); i++ {
		if UserList[i].ID == id {
			UserList[i] = &u
		}
	}
}

func GetUsers() Users {
	fmt.Println(UserList)
	return UserList
}

func AddUser(u *User) {
	u.ID = GetNextId()
	UserList = append(UserList, u)
}

func GetNextId() int {
	curr := UserList[len(UserList)-1]
	return curr.ID + 1
}

func NewUser() *User {
	return &User{}
}

var UserList = []*User{
	&User{
		ID:   1,
		Name: "Gilang",
	},
	&User{
		ID:   2,
		Name: "Anda",
	},
}
