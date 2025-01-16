package services

import (
	"fmt"
  "bytes"
	. "xlsxfix/internal/models"
	"github.com/xuri/excelize/v2"
)


func RemoveDuplicates(xlsxFile *XLSXFileInfo) ([]byte, error) {
  // Open the XLSX file using the buffered input
  f, err := excelize.OpenReader(xlsxFile.InputFile)
  if err != nil {
    return nil, fmt.Errorf("failed to open file: %w", err)
  }
  defer f.Close()

  // Validate if the specified sheet exists
  sheets := f.GetSheetList()
  sheetExists := false
  for _, sheet := range sheets {
    if sheet == xlsxFile.SheetName {
      sheetExists = true
        break
    }
  }
  if !sheetExists {
    return nil, fmt.Errorf("sheet %s not found in the file", xlsxFile.SheetName)
  }

  // Read rows from the specified sheet
  rows, err := f.GetRows(xlsxFile.SheetName)
  if err != nil {
    return nil, fmt.Errorf("failed to read rows from sheet: %w", err)
  }

  // Filter out duplicate rows based on the first column
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

  // Create a new XLSX file
  newFile := excelize.NewFile()
  index, err := newFile.NewSheet(xlsxFile.SheetName)
  if err != nil {
    return nil, fmt.Errorf("error creating new sheet: %w", err)
  }

  // Write filtered rows to the new file
  for i, row := range filteredRows {
    for j, cell := range row {
      cellAddress, _ := excelize.CoordinatesToCellName(j+1, i+1)
      newFile.SetCellValue(xlsxFile.SheetName, cellAddress, cell)
    }
  }

  // Set the active sheet and write to buffer
  newFile.SetActiveSheet(index)
  var buf bytes.Buffer
  if err := newFile.Write(&buf); err != nil {
    return nil, fmt.Errorf("failed to write file to buffer: %w", err)
  }
  return buf.Bytes(), nil
}

