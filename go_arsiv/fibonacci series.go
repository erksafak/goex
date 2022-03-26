package main

import "fmt"

func main() {
	t0 := 0
	t1 := 1
	var n int
	fmt.Print("Kaç döngü olacak:")
	fmt.Scanf("%d", &n)
	fmt.Printf("fibonacci series for %d cycle:", n)
	d := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Print(" ", t0)
			d[i] = t0
			continue
		} else if i == 1 {
			fmt.Print(" ", t1)
			d[i] = t1
			continue
		}
		var tt int
		tt = t0 + t1
		fmt.Print(" ", tt)
		t0 = t1
		t1 = tt
		d[i] = tt
	}
	fmt.Print("\nfibonacci array: ", d)
}
