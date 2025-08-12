package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	for {
		var choose int

		fmt.Println("\nCLI Pass Manager")
		fmt.Println("Pilih menu")
		fmt.Println("1 : Buat app password baru")
		fmt.Println("2 : Cari Password")
		fmt.Println("3 : Keluar")
		fmt.Print("Masukkan Pilihan: ")
		_, err := fmt.Scanf("%d", &choose)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			fmt.Scanln() // Clear input buffer
			continue
		}

		switch choose {
		case 1:
			Add()
		case 2:
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter app name: ")
			appName, _ := reader.ReadString('\n')
			appName = strings.TrimSpace(appName)
			Select(appName)
		case 3:
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
		}
	}
}

func Add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter app name: ")
	appName, _ := reader.ReadString('\n')
	appName = strings.TrimSpace(appName)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	f, err := excelize.OpenFile("Pass.xlsx")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	sheet := "Sheet1"

	// Get the last row to append new data
	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Println("Error reading sheet:", err)
		return
	}

	lastRow := len(rows) + 1

	// Write app name and password to the next empty row
	f.SetCellValue(sheet, fmt.Sprintf("A%d", lastRow), appName)
	f.SetCellValue(sheet, fmt.Sprintf("B%d", lastRow), password)

	if err := f.Save(); err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Password saved successfully!")
}

func Select(kode_pass string) {
	if kode_pass == "" {
		fmt.Println("Error: App name cannot be empty.")
		return
	}

	// Open Excel file
	f, err := excelize.OpenFile("Pass.xlsx")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("Error reading sheet:", err)
		return
	}

	// Search for the app name
	found := false
	for _, row := range rows {
		if len(row) > 0 && strings.EqualFold(strings.TrimSpace(row[0]), kode_pass) {
			// Found the app name, get the password from the next cell
			if len(row) > 1 {
				password := row[1]
				fmt.Printf("\nApp: %s\nPassword: %s\n", row[0], password)
			} else {
				fmt.Printf("\nApp: %s\nPassword: [empty]\n", row[0])
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Println("App not found in database.")
	}
}
