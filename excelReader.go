package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
)


func readExcelData(filepath string) ([][]string, error) {
    f, err := xlsx.OpenFile(filepath)
    if nil != err {
        return nil, err;
    }
    
    if 0 == len(f.Sheets) {
        return nil, fmt.Errorf("empty excel file no sheets");
    }
    data := make([][]string, len(f.Sheets[0].Rows))
    for i, row := range(f.Sheets[0].Rows) {
        data[i] = make([]string, len(row.Cells))
        for j, cell := range(row.Cells) {
            str, _ := cell.String()
            data[i][j] = str
        }
    }
    return data, nil
}