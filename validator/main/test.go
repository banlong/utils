package main

import (
	"log"
	"utils/validator"
)

func main() {
	nur := NewUserRequest{Username: "som", Name:"A", Age: 20}
	if errs := validator.Validate(nur); errs != nil {
		log.Println(errs.Error())
		return
	}
}

type NewUserRequest struct {
	Username string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Name string     `validate:"nonzero"`
	Age int         `validate:"min=21"`
	Password string `validate:"min=8"`
}


