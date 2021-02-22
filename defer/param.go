package main

import (
	"errors"
	"fmt"
)

func f11() {
	var err error
	defer fmt.Println(err)
	err = errors.New("defer error")
	return
}

func f22() {
	var err error

	defer func() {
		fmt.Println(err)
	}()

	err = errors.New("defer error")
	return
}

func f33() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)

	err = errors.New("defer error")
	return
}

func main() {
	f11()
	f22()
	f33()
}
