package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() { fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email) }

type admin struct {
	*user
	level string
}

func main() {
	u := user{name: "john smith", email: "john@yahoo.com"}
	ad := admin{
		// Pointer Semantic Embedding
		user: &u, level: "super",
	}
	ad.user.notify()
	u.name = "william smith"
	ad.notify() // Outer type promotion
}
