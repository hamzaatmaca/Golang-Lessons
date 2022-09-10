package pointer

import "fmt"

func Pointers2(sayilar []int) {
	sayilar[0] = 100
	fmt.Println(sayilar[0])
}
