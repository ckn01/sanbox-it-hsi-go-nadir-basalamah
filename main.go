package main

import (
	"fmt"
)

func main() {
	// Percabangan IF
	// if kondisi {
	// 	// aksi jika kondisi benar
	// } else {
	// 	// aksi jika kondisi salah
	// }

	// Contoh penggunaan IF
	// Program : megecekan bilangan genap atau ganjil
	// var bilangan int
	// fmt.Print("Masukkan bilangan: ")
	// fmt.Scan(&bilangan)

	num := 23

	// bilangan genap -> bilangan yang habis dibagi 2  
	// sisa hasil pembagiannya adalah 0
	if num%2 == 0 {
		fmt.Println(num, "adalah bilangan genap")
	} else {
		fmt.Println(num, "adalah bilangan ganjil")
	}

	course := "Golang"
	fmt.Println("Saya sedang belajar", course)

	var role string = "admin"

	switch role {
	case "admin":
		fmt.Println("Selamat datang admin")
	case "user":
		fmt.Println("Selamat datang user")
	default:
		fmt.Println("Selamat datang tamu")
	}


	// case study : perhitungan BMI (Body Mass Index) / IMT (Indeks Massa Tubuh)

	var beratBadan float64 // dalam kilogram
	var tinggiBadan float64  // dalam meter

	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scan(&beratBadan)

	fmt.Print("Masukkan tinggi badan (cm): ")
	fmt.Scan(&tinggiBadan)

	tinggiBadan = tinggiBadan / 100 // konversi dari cm ke m

	bmi := beratBadan / (tinggiBadan * tinggiBadan)
	// Contoh: 86,5 kg dan 170 cm menghasilkan BMI sekitar 29,93.
	// Contoh: 72,1 kg dan 170 cm menghasilkan BMI sekitar 24,95.

	fmt.Printf("BMI Anda adalah: %.2f\n", bmi)

	// if bmi < 18.5 {
	// 	fmt.Println("Kategori: Kekurangan berat badan", bmi)
	// } else if bmi >= 18.5 && bmi < 24.9 {
	// 	fmt.Println("Kategori: Normal (ideal)", bmi)
	// } else if bmi >= 25 && bmi < 29.9 {
	// 	fmt.Println("Kategori: Kelebihan berat badan", bmi)
	// } else {
	// 	fmt.Println("Kategori: Obesitas", bmi)
	// }

	if bmi < 18.5 {
		fmt.Println("Kategori: Kekurangan berat badan", bmi)
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Kategori: Normal (ideal)", bmi)
	} else if bmi >= 25 && bmi < 29.9 {
		fmt.Println("Kategori: Kelebihan berat badan", bmi)
	} else {
		fmt.Println("Kategori: Obesitas", bmi)
	}

	number := 1

	for {
		fmt.Println(number)
		number++
		if number == 10 {
			break
		}

	}
}
