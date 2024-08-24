package main

import (
	"errors"
	"fmt"
)

var error1 = fmt.Errorf("This is error 1")
var error2 = fmt.Errorf("This is error 2")

func findOddNum(number int32) (int32, error) {
	if number%2 == 0 {
		return 0, error1
	} else if number < 0 {
		return 0, errors.New("Number is negative")
	}
	return number, nil
}

func main() {

	if _, error := findOddNum(2); error != nil {
		fmt.Println(error)
	} else if _, error := findOddNum(3); error == nil {
		fmt.Println(error)
	} else {
		fmt.Println("success")

	}
}
