package main

import "fmt"

func main() {

	var t int
	fmt.Print("Kaça kadar sayalım:")
	fmt.Scanf("%v", &t)
	for i := 0; i < t; i++ {
		if i == (t - 1) {
			break
		} else if i == 2 {
			var a [3]string
			var f int = 99
			b := [2]int{}
			//var c = [2]int{19, 27}    Aynı ifade aşağıda üç nokta yazılarak oluşturuldu.
			var c = [...]int{19, 27}
			a[0] = "izmir"
			a[1] = "uşak"
			a[2] = "çorum"
			b[0] = 35
			b[1] = 64
			fmt.Print(a, b, c, f)
			fmt.Println("  atlanan değer:i=", 2, "a'nın eleman sayısı:", len(a), "a'nın kapasitesi:", cap(a))
			continue
		}
		fmt.Println("i:", i)
	}
	fmt.Println("İşlem sonu")
}
