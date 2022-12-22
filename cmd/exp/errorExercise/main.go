package main

import (
	"errors"
	"fmt"
)

func main() {

	err := A()

	if err != nil {

		if errors.Is(err, ErrorInternal) {
			fmt.Println("Error is of type ErrorNotFound: ", err)
			return
		}
		fmt.Println("Error is not of type ErrorNotFound ")
	}

}

var ErrorNotFound = errors.New("Not found")

var ErrorInternal = errors.New("InternalError")

func A() error {
	return ErrorNotFound
}
