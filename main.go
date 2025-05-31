package main

import (
  "fmt"
  "net/http"
)

func main() {
  fmt.Println("HEllo world")

  mux := http.NewServeMux()

  fileServerHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
  mux.Handle("/app/", fileServerHandler)

  mux.HandleFunc("/healthz", func(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "text/plain; charset=utf-8")
    res.WriteHeader(200)
    res.Write([]byte("OK"))
  })

  server := http.Server{
    Handler: mux,
    Addr: ":8080",
  }
  if err := server.ListenAndServe(); err != nil {
    fmt.Println("Unable start server: %w", err)
  }

}
