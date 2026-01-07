package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func AddApp() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter app name: ")
	appName, _ := reader.ReadString('\n')
	appName2 := strings.TrimSpace(appName)

	return appName2
}

func AddPass() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password2 := strings.TrimSpace(password)

	return password2
}

func Add() {
	appname := AddApp()
	password := AddPass()

	f, err := excelize.OpenFile("Pass.xlsx")
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	sheet := "Sheet1"

	rows, err := f.GetRows(sheet)
	if err != nil {
		log.Println("Error reading sheet:", err)
		return
	}

	lastRow := len(rows) + 1

	f.SetCellValue(sheet, fmt.Sprintf("A%d", lastRow), appname)
	f.SetCellValue(sheet, fmt.Sprintf("B%d", lastRow), password)

	if err := f.Save(); err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Password saved successfully!")
}
