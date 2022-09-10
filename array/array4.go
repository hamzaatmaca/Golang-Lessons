package array

import "fmt"

func Array4() {
	var sayilar [2][3]int
	sayilar[0][0] = 12
	sayilar[0][1] = 31
	sayilar[0][2] = 98
	sayilar[1][0] = 23
	sayilar[1][1] = 54
	sayilar[1][2] = 13
	fmt.Println(sayilar)
}
