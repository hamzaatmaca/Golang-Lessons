package loops

import "fmt"

func Loops() {
	sayi := 80
	tahmin := 0
	// for i <= 5 {
	// 	fmt.Printf("test")
	// 	i = i + 1
	// }
	fmt.Scanln(&tahmin)
	for tahmin != sayi {
		if tahmin < sayi {
			fmt.Printf("küçük")
			fmt.Scanln(&tahmin)
		} else if tahmin > sayi {
			fmt.Printf("Büyük")
			fmt.Scanln(&tahmin)

		}

	}
	fmt.Printf("Bildiniz")

}
