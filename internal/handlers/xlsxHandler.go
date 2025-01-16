package handlers

import (
  "log"
  "bytes"
  "io"
  "net/http"
  . "xlsxfix/internal/services"
  . "xlsxfix/internal/models"
)


func XSLHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    return
  }

  err := r.ParseMultipartForm(10 << 20) // 10 MB limit
  if err != nil {
    http.Error(w, "Unable to parse form", http.StatusBadRequest)
    return
  }

  xlsxFile, fileHeader, err := r.FormFile("file")
  if err != nil {
    http.Error(w, "Unable to read file", http.StatusBadRequest)
    return
  }
  defer xlsxFile.Close()

  // Log file information for debugging
  log.Printf("Received file: %s", fileHeader.Filename)
  log.Printf("Content-Type: %s", fileHeader.Header.Get("Content-Type"))

  // Ensure file is XLSX by checking MIME type
  if fileHeader.Header.Get("Content-Type") != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
    http.Error(w, "Invalid file type, expected .xlsx", http.StatusBadRequest)
    return
  }

  // Buffer the file content for Excelize to process
  var buf bytes.Buffer
  _, err = io.Copy(&buf, xlsxFile)
  if err != nil {
    http.Error(w, "Failed to buffer file content", http.StatusInternalServerError)
    return
  }

  sheetName := r.FormValue("sheetName")
  if sheetName == "" {
    http.Error(w, "Missing sheet name", http.StatusBadRequest)
    return
  }

  var receivedXLSXFileInfo = XLSXFileInfo{
    InputFile: &buf,
    SheetName: sheetName,
  }

  updatedFile, err := RemoveDuplicates(&receivedXLSXFileInfo)
  if err != nil {
    log.Printf("Error processing file: %v", err)
    http.Error(w, "Failed to process file: "+err.Error(), http.StatusBadRequest)
    return
  }

  w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
  w.Header().Set("Content-Disposition", "attachment; filename=updated_file.xlsx")
  w.Write(updatedFile)
}


