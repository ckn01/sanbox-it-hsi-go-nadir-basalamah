package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var nilaia,nilaib,nilaic int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&nilaia)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&nilaib)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&nilaic)

	fmt.Printf("Nilai a: %d, Nilai b: %d, Nilai c: %d\n", nilaia, nilaib, nilaic)

	if (nilaia*nilaia) + (nilaib*nilaib) == (nilaic*nilaic) {
		fmt.Println("pytagoras")
	} else {
		fmt.Println("Bukan pytagoras")
	}
}