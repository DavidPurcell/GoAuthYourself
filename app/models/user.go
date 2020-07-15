package model

import (
	"fmt"
)

type User struct {
	Id       string
	Username string
	Password string
}

func (u User) LeavesRemaining() {
	fmt.Printf("%s %s %s \n", u.Id, u.Username, u.Password)
}
