package maps

import "fmt"

func Maps() {
	sozluk := make(map[string]string)
	sozluk["chair"] = "sandalye"
	sozluk["fury"] = "ofke"
	sozluk["car"] = "araba"
	fmt.Println(sozluk["chair"])
	fmt.Println(len(sozluk))
	fmt.Println(sozluk)
	delete(sozluk, "car")
	fmt.Println(sozluk)
	deger, varmi := sozluk["xxx"]
	fmt.Println(deger)
	fmt.Println(varmi)
	sozluk2 := map[string]string{"glass": "bardak", "window": "cam"}
	fmt.Println(sozluk2)
}
