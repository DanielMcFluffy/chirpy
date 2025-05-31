package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("HEllo world")

  mux := http.NewServeMux()
  server := http.Server{
    Handler: mux,
    Addr: ":8080",
  }
  if err := server.ListenAndServe(); err != nil {
    fmt.Errorf("Unable start server: %w", err)
  }

}
