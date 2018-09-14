package main

import "fmt"

type Man interface {
	name() string
	age() int
}

// woman
type Woman struct {
}

func (woman Woman) name() string {
	return "katherine"
}

func (woman Woman) age() int {
	return 23
}

// man
type Men struct {
}

func (men Men) name() string {
	return "damon"
}

func (men Men) age() int {
	return 24
}

// main
func main() {
	var man Man

	man = new(Woman)
	fmt.Println(man.name())
	fmt.Println(man.age())
	man = new(Men)
	fmt.Println(man.name())
	fmt.Println(man.age())
}
