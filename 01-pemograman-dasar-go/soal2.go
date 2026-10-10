package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var nilaiEH,nilaiEP,nilaiEA float32
	var nilaiAkhir float32
	fmt.Print("Masukkan nilai Evaluasi Harian: ")
	fmt.Scan(&nilaiEH)
	fmt.Print("Masukkan nilai Evaluasi Pekanan: ")
	fmt.Scan(&nilaiEP)
	fmt.Print("Masukkan nilai Evaluasi Akhir: ")
	fmt.Scan(&nilaiEA)

	if (nilaiEH < 0 || nilaiEH > 100) || (nilaiEP < 0 || nilaiEP > 100) || (nilaiEA < 0 || nilaiEA > 100) {
		fmt.Println("Invalid input")
		return
	}

	nilaiAkhir = ((nilaiEH * 0.2) + (nilaiEP * 0.25) + (nilaiEA * 0.55))

	fmt.Println("Nilai Akhir:", nilaiAkhir)


 if nilaiAkhir >= 91 && nilaiAkhir <= 100 {
		fmt.Println("Sangat Baik")
	} else if nilaiAkhir >= 80 && nilaiAkhir <= 90 {
		fmt.Println("Baik")
	} else if nilaiAkhir >= 71 && nilaiAkhir <= 80 {
		fmt.Println("Cukup")
	} else if nilaiAkhir >= 66 && nilaiAkhir <= 79 {
		fmt.Println("Kurang")
	} else if (nilaiAkhir >= 0 && nilaiAkhir <= 65) {
		fmt.Println("Sangat Kurang")
	} else {
		fmt.Println("Nilai akhir tidak valid")
	}

}