package main

import (
	"fmt"
	"time"
)

type car struct {
	baz   float64
	model int
	hacim float64
}

type bus struct {
	baz   float64
	model int
	hacim float64
}

func (ci car) vergi() float64 {
	var z float64
	if ci.model < time.Now().Year()-3 {
		if ci.hacim < 1.6 {
			z = ci.baz * 1
		} else {
			z = ci.baz * 1.1
		}
	}

	if ci.model > time.Now().Year() {
		z = 0
	}

	if ci.model >= time.Now().Year()-3 {
		if ci.hacim < 1.6 {
			z = ci.baz * 2
		} else {
			z = ci.baz * 2.1
		}
	}
	return z
}

func (ci bus) vergi() float64 {
	var z float64
	if ci.model < time.Now().Year()-3 {
		if ci.hacim < 1.6 {
			z = ci.baz * 4
		} else {
			z = ci.baz * 5
		}
	}

	if ci.model > time.Now().Year() {
		z = 0
	}

	if ci.model >= time.Now().Year()-3 {
		if ci.hacim < 1.6 {
			z = ci.baz * 6
		} else {
			z = ci.baz * 7
		}
	}
	return z
}

type hesapla interface {
	vergi() float64
}

// INTERFACE FUNC
func calculate(i hesapla) {
	fmt.Println("tax:", i.vergi())

}

func main() {
	var baz, hacim float64
	baz = 100
	model := 2000
	hacim = 1.3
	ca := car{baz, model, hacim}
	bu := bus{baz, model, hacim}
	fmt.Print("car ")
	calculate(ca)
	fmt.Print("bus ")
	calculate(bu)
}
