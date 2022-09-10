package goroutines

import (
	"fmt"
	"time"
)

func Tek() {
	for i := 1; i < 10; i += 2 {

		fmt.Println("tek : ", i)
		time.Sleep(1 * time.Second)
	}
}
func Goroutines() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("çift : ", i)
		time.Sleep(1 * time.Second)
	}

}
