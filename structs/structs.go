package structs

import "fmt"

func Structs() {
	fmt.Println(product{name: "Bilgisayar", brand: "Apple", price: 40})
	c := product{"a", "b", 3}
	println(c.brand)
	c.save()
}

func (p product) save() {
	fmt.Println("eklendi")
}

type product struct {
	name  string
	brand string
	price int
}
