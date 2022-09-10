package array

import "fmt"

func Array2() {
	var sehirler [5]string
	sehirler[0] = "Ankara"
	sehirler[1] = "İzmir"
	sehirler[2] = "Bursa"
	sehirler[3] = "İstanbul"
	sehirler[4] = "Mersin"

	fmt.Println(sehirler)
	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
