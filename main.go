package main

import (
	"excel-note/helper"
	"fmt"
)

func main() {
	for {
		helper.PrintMenu()

		choice, err := helper.ReadChoice()
		if err != nil {
			fmt.Println("Input tidak valid. Masukkan angka.")
			continue
		}

		switch choice {
		case 1:
			helper.Add()
		case 2:
			helper.Selector()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Pilihan tidak tersedia. Silakan pilih 1, 2, atau 3.")
		}
	}
}