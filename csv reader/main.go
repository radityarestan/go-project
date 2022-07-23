package main

import (
	"csv-reader/service"
	"fmt"
)

func main() {
	validationInstance := service.GetInstance()
	fmt.Println(validationInstance.Validate("babi"))
	fmt.Println(validationInstance.Validate("anjing"))
	fmt.Println(validationInstance.Validate("aku suka kamu beb"))
	fmt.Println(validationInstance.Validate("iki Jancuk! iki"))
	fmt.Println(validationInstance.Validate("aku suka kamu babi"))
}
