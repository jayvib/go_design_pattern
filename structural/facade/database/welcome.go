package database

import "fmt"

type greet struct {}

func (greet) welcome(name string) string { // this is a form of facade design pattern
	return fmt.Sprintf("Welcome %s!", name)
}

func WelcomePage(email string) string {
	user := DB.getUserByEmail(email)
	g := greet{}
	return g.welcome(user.FirstName)
}