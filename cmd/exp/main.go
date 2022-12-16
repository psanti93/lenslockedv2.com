package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")

	if err != nil {
		panic(err)
	}

	// //anonymous struct
	// user := struct {
	// 	Name string
	// }{
	// 	Name: "Paul Santiago",
	// }

	user := User{
		Name: "Paul Santiago",
		Age:  29,
		Meta: UserMeta{
			Visits: 4,
		},
	}

	if err = t.Execute(os.Stdout, user); err != nil {
		panic(err)
	}
}
