package main

import (
	"fmt"

	"github.com/manfered/Go-GettingStarted-Packages/models"
)

func main() {
	// declaring freduser by grabing the user from models package
	fredUser := models.User{
		Id:        1,
		UserName:  "Manfered",
		FirstName: "Fred",
	}

	fmt.Println(fredUser)
}
