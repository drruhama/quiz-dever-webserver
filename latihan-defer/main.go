package main

import "fmt"

func main() {
	var input1, input2 int
	defer cetak()
	fmt.Print("Masukkan angka pertama:")
	fmt.Scan(&input1)
	fmt.Println("")
	fmt.Print("Masukkan angka kedua:")
	fmt.Scan(&input2)
	fmt.Println("")
	hasil := input1 + input2
	fmt.Println("Hasil penjumlahan:", hasil)
}

func cetak() {
	fmt.Println("Program selesai")
}
