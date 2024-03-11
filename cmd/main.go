package main

import (
	"fmt"
	"log"
	"newser/internal/domain/user"
)

func main() {
	u1, err := user.NewUser("jeff@hello.whatever", "Jeff Caldwell")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("u1: %+v\n", u1)

	u2, err := user.NewUser("eml@nothing", "Dookiepants")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("u2: %+v\n", u2)
}
