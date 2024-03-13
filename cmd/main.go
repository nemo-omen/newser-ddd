package main

import (
	"fmt"
	"log"
	"newser/internal/service"
)

func main() {
	userService, err := service.NewUserService(service.WithMemoryUserRepository())
	if err != nil {
		log.Println(err.Error())
	}

	user1, err := userService.Signup("jeff@jeff.jeff", "Jeff Caldwell", "123456")
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(user1)

}
