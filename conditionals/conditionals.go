package conditionals

import "fmt"

func Conditionals() {
	var hesap float64 = 10000
	var cekilmekistenen float64 = 4000
	if cekilmekistenen > hesap {
		fmt.Println("Geçersiz Tutar")
	}
	if cekilmekistenen <= hesap {
		fmt.Println("Para bitti")
		hesap = hesap - cekilmekistenen
		fmt.Println("Bitti. Hesaptaki Tutar : " + fmt.Sprintf("%f", hesap))
		fmt.Printf("Bitti.Hesaptaki Tutar : %v", hesap)
	}
}
