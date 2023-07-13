package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	r := random(1, 20)
	err := a(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = b(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn((max-min)+1) + min
}

func a(i int) error {
	if i < 10 {
		fmt.Println("error is at a()")
		return errors.New("incorrect value")
	}
	return nil
}

func b(i int) error {
	if i >= 10 {
		fmt.Println("error is at b()")
		return errors.New("incorrect value")
	}
	return nil
}
