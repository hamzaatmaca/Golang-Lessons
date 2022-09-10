package array

import "fmt"

func Array3() {

	sayilar := [5]int{233, 24, 56, 87, 23}
	fmt.Println(sayilar)
	enBüyük := sayilar[0]
	for i := 0; i < len(sayilar); i++ {
		fmt.Println(sayilar[i])
		if sayilar[i] > enBüyük {
			enBüyük = sayilar[i]
		}

	}
	fmt.Println(enBüyük)
}
