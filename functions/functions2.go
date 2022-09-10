package functions

func Functions2(sayi1 int, sayi2 int) (int, int, int, float32) {
	toplam := sayi1 + sayi2
	carpim := sayi1 * sayi2
	cikarim := sayi1 - sayi2
	bolme := float32(sayi1 / sayi2)
	return toplam, carpim, cikarim, bolme
}
