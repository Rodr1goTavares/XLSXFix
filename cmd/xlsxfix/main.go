package main

import (
	"fmt"
  "net/http"
	. "xlsxfix/internal/handlers"
)

func main() {
  http.HandleFunc("/xlsx", XSLHandler)
  port := ":8080"
  fmt.Println("Server running at: ", port)
  http.ListenAndServe(port, nil)
}
