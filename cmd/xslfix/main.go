package main

import (
	"fmt"
  "net/http"
	. "xslfix/internal/handlers"
)

func main() {
  http.HandleFunc("/xsl", XSLHandler)
  port := ":8080"
  fmt.Println("Server running at: ", port)
  http.ListenAndServe(port, nil)
}
