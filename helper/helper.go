package helper

import (
	"fmt"
)

func PrintMenu() {
	fmt.Println("\nCLI Pass Manager")
	fmt.Println("Pilih menu")
	fmt.Println("1 : Buat app password baru")
	fmt.Println("2 : Cari Password")
	fmt.Println("3 : Keluar")
	fmt.Print("Masukkan Pilihan: ")
}

func ReadChoice() (int, error) {
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	fmt.Scanln() 
	return choice, err
}
