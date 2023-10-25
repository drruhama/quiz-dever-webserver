package main

import "fmt"

func main() {
	var input1, input2 int
	defer catchError()
	fmt.Print("Masukkan angka pertama:")
	fmt.Scan(&input1)
	fmt.Println("")
	fmt.Print("Masukkan angka kedua:")
	fmt.Scan(&input2)
	fmt.Println("")
	err := Validasi(input1, input2)
	if err != nil {
		cetak(err.Error())
	}
}
func Validasi(inp1, inp2 int) error {
	if inp2 == 0 {
		panic("Tidak dapat membagi nol")
	} else {
		hasil := inp1 / inp2
		fmt.Println("Hasil pembagian:", hasil)
	}
	return nil
}

func catchError() {
	err := recover()
	if err != nil {
		cetak("Tidak dapat membagi nol")
	}
}

func cetak(text string) {
	fmt.Println(text)
}
