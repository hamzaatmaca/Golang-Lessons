package slices

import "fmt"

func Slices2() {
	sehirler := []string{"Ankara", "izmir", "istanbul"}
	sehirler = append(sehirler, "Yozgat")
	sehirlerKopya := make([]string, len(sehirler))

	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	fmt.Println(sehirler[1:3])
}
