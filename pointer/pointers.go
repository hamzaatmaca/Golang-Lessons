package pointer

import "fmt"

func Pointers(sayi1 *int) {
	*sayi1 = *sayi1 + 1
	fmt.Println(sayi1)
}
