package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

func Selector() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter app name: ")
	appName, _ := reader.ReadString('\n')
	appName = strings.TrimSpace(appName)
	Select(appName)
}

func Select(kode_pass string) {
	if kode_pass == "" {
		fmt.Println("Error: App name cannot be empty.")
		return
	}

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

	found := false
	for _, row := range rows {
		if len(row) > 0 && strings.EqualFold(strings.TrimSpace(row[0]), kode_pass) {
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
