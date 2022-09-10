package ranges

import "fmt"

func Rangess() {
	sehirler := []string{"Anakra", "izmir", "istanbul"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])
	}
	for _, sehir := range sehirler {
		//fmt.Println(i)
		fmt.Println(sehir)
	}
}
