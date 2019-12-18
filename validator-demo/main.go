package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID   int64  `validate:"gte=0,lte=130"`
	Name string `validate:"required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func main() {
	validate = validator.New()

	user := &User{
		ID: 10000,
	}
	err := validate.Struct(user)
	fmt.Println(err)

	for _, err := range err.(validator.ValidationErrors) {

		fmt.Println(err.Namespace())
		fmt.Println(err.Field())
		fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
		fmt.Println(err.StructField())     // by passing alt name to ReportError like below
		fmt.Println(err.Tag())
		fmt.Println(err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println(err.Param())
		fmt.Println()
	}

}
