package conditionals

import "fmt"

func Conditionals2() {
	var hesap float64 = 10000
	var cekilmekistenen float64 = 10000
	if cekilmekistenen < hesap {
		fmt.Println("Para bitti")
		hesap = hesap - cekilmekistenen
		fmt.Println("Bitti. Hesaptaki Tutar : " + fmt.Sprintf("%f", hesap))
		fmt.Printf("Bitti.Hesaptaki Tutar : %v", hesap)
	} else if cekilmekistenen == hesap {
		fmt.Printf("Paranızın hepsi çekiliyor")
	} else {
		fmt.Printf("Yetersiz Bakiye")
	}
}
