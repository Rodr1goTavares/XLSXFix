package handlers

import (
  "net/http"
  . "xlsxfix/internal/services"
  . "xlsxfix/internal/models"
)

func XSLHandler(w http.ResponseWriter, r *http.Request) {

  // Verifica se a requisição é do tipo multipart/form-data
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse do formulário multipart (o que nos permite acessar o arquivo)
	errSize := r.ParseMultipartForm(10 << 20) // Limite de 10 MB para o arquivo
	if errSize != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

  xlsxFile, _, errFormFile := r.FormFile("file")
	if errFormFile != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer xlsxFile.Close()

  sheetName := r.FormValue("sheetName")
	if sheetName == "" {
		http.Error(w, "Missing message", http.StatusBadRequest)
		return
	}

  var receivedXSLXFileInfo = XLSXFileInfo {
    InputFile: xlsxFile,
    SheetName: sheetName,
  }

  var updatedFile, err = RemoveDuplicates(&receivedXSLXFileInfo)
  if err != nil {
    http.Error(w, "Failed to process file: "+err.Error(), http.StatusBadRequest)
		return
  }

  w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename=updated_file.xlsx")
  w.Write(updatedFile)
}
