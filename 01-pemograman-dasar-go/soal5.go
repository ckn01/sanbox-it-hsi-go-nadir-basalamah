package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var inputKalimat,inputKataKunci,kunciLower string

	fmt.Print("Masukkan kalimat: ")
	if scanner.Scan() {
		inputKalimat = scanner.Text()
	}

	fmt.Print("Masukkan kata kunci: ")
	if scanner.Scan() {
		inputKataKunci = scanner.Text()
	}

	fmt.Println("Kalimat:", inputKalimat)
	fmt.Println("Kata kunci:", inputKataKunci)

	if(inputKalimat == "" || inputKataKunci == "") {
		fmt.Println("Input invalid, harap isi kalimat dan kata kunci")
		return
	}

	stringSplit := strings.Split(strings.ToLower(inputKalimat), " ")
	kunciLower = strings.ToLower(inputKataKunci)

	for i := 0; i < len(stringSplit); i++ {
		if kunciLower == stringSplit[i] {
			fmt.Println("Kata ditemukan")
			return
		}
	}
	fmt.Println("Kata tidak ditemukan")
}
