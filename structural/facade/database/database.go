package database

import "strings"

var DB db

func init() {
	DB = db{
		"jayson.vibandor@gmail.com": "Jayson Vibandor",
		"jimboy.vibandor@gmail.com": "Jimboy Vibandor",
	}
}

type User struct {
	FirstName, LastName string
}

type db map[string]string

func (d *db) getUserByEmail(email string) *User {
	user := new(User)
	fullname := (*d)[email]
	first, last := parseName(fullname)
	user.FirstName = first
	user.LastName = last
	return user
}

func GetUserByEmail(email string) *User {
	return DB.getUserByEmail(email)
}

func parseName(fullname string) (firstname, lastname string) {
	splitName := strings.Split(fullname, " ")
	firstname = splitName[0]
	lastname = splitName[1]
	return firstname, lastname
}
