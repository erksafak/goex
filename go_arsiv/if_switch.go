package main

import "fmt"

func main() {
	var score float64
	fmt.Print("Enter score for your last exam:")
	fmt.Scanf("%v", &score)
	if score >= 45 {
		switch {
		case score <= 50:
			fmt.Println("Your grade is F")
		case score <= 60:
			fmt.Println("Your grade is E")
		case score <= 70:
			fmt.Println("Your grade is D")
		case score <= 80:
			fmt.Println("Your grade is C")
		case score <= 90:
			fmt.Println("Your grade is B")
		case score <= 100:
			fmt.Println("Your grade is A")
		default:
			fmt.Println("Wrong data")
		}
	} else {
		fmt.Println("Kicked off")
	}
}
