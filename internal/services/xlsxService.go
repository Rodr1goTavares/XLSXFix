package services

import (
	"fmt"
  "bytes"
	. "xlsxfix/internal/models"
	"github.com/xuri/excelize/v2"
)


func RemoveDuplicates(xlsxFile *XLSXFileInfo) ([]byte, error) {
	// Abre o arquivo de entrada
	f, err := excelize.OpenReader(xlsxFile.InputFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %w", err)
	}
	defer f.Close()

	// Obtém as linhas da planilha
	rows, err := f.GetRows(xlsxFile.SheetName)
	if err != nil {
		return nil, fmt.Errorf("Failed to read lines: %w", err)
	}

	// Filtra as linhas duplicadas
	unique := make(map[string]bool)
	var filteredRows [][]string
	for _, row := range rows {
		if len(row) > 0 {
			if !unique[row[0]] {
				unique[row[0]] = true
				filteredRows = append(filteredRows, row)
			}
		}
	}

	// Create a new XSL file
	newFile := excelize.NewFile()
	index, err := newFile.NewSheet(xlsxFile.SheetName)
	if err != nil {
		return nil, fmt.Errorf("Error to create new sheet: %w", err)
	}

	// Add filtered lines in new file
	for i, row := range filteredRows {
		for j, cell := range row {
			cellAddress, _ := excelize.CoordinatesToCellName(j+1, i+1)
			newFile.SetCellValue(xlsxFile.SheetName, cellAddress, cell)
		}
	}

	// Define active sheet and return bytes[] 
	newFile.SetActiveSheet(index)
	var buf bytes.Buffer
	if err := newFile.Write(&buf); err != nil {
		return nil, fmt.Errorf("Failed to write file to buffer: %w", err)
	}
	return buf.Bytes(), nil
}
