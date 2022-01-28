package main

import "fmt"

type notifier interface{ notify() }

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	u := user{"Bill", "bill@email.com"}
	/*
		cannot use u (variable of type user) as notifier value in argument
		to sendNotification: missing method notify (notify has pointer receiver)

		sendNotification(u)
	*/
	sendNotification(&u)
}
