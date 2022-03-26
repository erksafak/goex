package main

import (
	"fmt"
	"math"
)

//STRUCT
type dikdörtgen struct {
	a, b float64 //a, b birer FIELD dır
}
type daire struct {
	pi, r float64
}
type üçgen struct {
	a, b, ang float64
}

//METHODS
func (x dikdörtgen) alan() float64 {
	fmt.Print("Dikdörtgen Alanı:")
	return x.a * x.b
}
func (x dikdörtgen) çevre() float64 {
	fmt.Print("Dikdörtgen Çevresi:")
	return 2 * (x.a + x.b)
}
func (x daire) alan() float64 {
	fmt.Print("Daire Alanı:")
	return x.pi * x.r * x.r
}
func (x daire) çevre() float64 {
	fmt.Print("Daire Çevresi:")
	return 2 * x.pi * x.r
}
func (x üçgen) alan() float64 {
	fmt.Print("Üçgen Alanı:")
	c := math.Sin(x.ang)
	return (x.a * x.b * c) / 2
}
func (x üçgen) çevre() float64 {
	fmt.Print("Üçgen Çevsesi:")
	c := math.Cos(x.ang)
	ç := x.a*x.a + x.b*x.b - 2*x.a*x.b*c
	s := math.Sqrt(ç)
	return s + x.a + x.b
}

//INTERFACE
type shape interface {
	alan() float64
	çevre() float64
}

// INTERFACE FUNC
func calculate(i shape) {
	fmt.Println(i.alan())
	fmt.Println(i.çevre())
}

//MAİN FUNC
func main() {
	di := new(dikdörtgen)
	di.a = 4
	di.b = 5
	da := daire{math.Pi, 1}
	üç := üçgen{3, 4, math.Pi / 2}

	calculate(di)
	calculate(da)
	calculate(üç)
}
