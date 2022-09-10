package slices

import "fmt"

func Slices() {
	isimler := make([]string, 3)
	fmt.Println(isimler)
	isimler[0] = "hamza"
	isimler[1] = "atmaca"
	isimler[2] = "golang"
	isimler = append(isimler, "Test")
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}
