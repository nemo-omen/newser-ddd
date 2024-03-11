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
	fmt.Printf("u1: %+v", u1)
}
