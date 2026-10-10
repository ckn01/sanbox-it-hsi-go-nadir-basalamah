package main

import (
	"fmt"
)

func main() {
	var inputUkuranBidang int
	var nilaiAkhir int
	fmt.Print("Masukkan ukuran bidang tanah: ")
	fmt.Scan(&inputUkuranBidang)


	if inputUkuranBidang < 0 {
		fmt.Println("hanya menerima bilangan positif")
		return
	}

	for i := 1; i <= inputUkuranBidang; i++ {
		fmt.Println("i:", i)
		for j := 1; j <= inputUkuranBidang; j++ {
			if (i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0) {
				fmt.Println("j:", j)
				nilaiAkhir++
			}
		}
	}


	fmt.Println("Jumlah yang bisa ditanam di bidang tanah", inputUkuranBidang, "adalah", nilaiAkhir)

}
