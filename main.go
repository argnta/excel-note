package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func CreateFile() {
    f := excelize.NewFile()
    if err := f.SaveAs("Pass.xlsx"); err != nil {
        fmt.Println(err)
    }
}

func Create() {
    f := excelize.NewFile()
    defer func() {
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    for idx, row := range [][]interface{}{
        {"App"}, {"Small", 2, 3, 3},
        {"Normal", 5, 2, 4}, {"Large", 6, 7, 8},
    } {
        cell, err := excelize.CoordinatesToCellName(1, idx+1)
        if err != nil {
            fmt.Println(err)
            return
        }
        f.SetSheetRow("Sheet1", cell, &row)
    }
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        fmt.Println(err)
    }
}

func main() {
    fmt.Println("CLI Pass Manager")
    fmt.Println("Pilih menu ")
    f := excelize.NewFile()
    if err := f.SaveAs("Pass.xlsx"); err != nil {
        fmt.Println(err)
    }
}