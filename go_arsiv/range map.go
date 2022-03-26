package main

import "fmt"

func main() {
	plaka := map[string]int{"çorum": 19, "uşak": 64, "izmir": 35}
	for a, b := range plaka {
		if b == 19 {
			fmt.Println(a, "'un plakası", b, "'dur")
		} else if b == 35 {
			fmt.Println(a, "'in plakası", b, "'tir")
		} else if b == 64 {
			fmt.Println(a, "'ın plakası", b, "'tür")
		}
	}
}
