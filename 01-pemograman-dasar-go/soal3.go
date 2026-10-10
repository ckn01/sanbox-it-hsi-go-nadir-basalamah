package main

import (
	"fmt"
)

func main() {
	var inputBilangan int
	var nilaiAkhir int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&inputBilangan)


	if inputBilangan < 0 {
		fmt.Println("hanya menerima bilangan positif")
		return
	}

	for i := 1; i <= inputBilangan; i++ {
		if inputBilangan%i == 0 {
			fmt.Println(i)
			nilaiAkhir+=i
		}
	}
	fmt.Println("Jumlah faktor dari", inputBilangan, "adalah", nilaiAkhir)

}
