package main

import (
	"fmt"
	"net/http"
)


func main() {
  mux := http.NewServeMux()

  apiConfig := &apiConfig{}

  fileServerHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
  mux.Handle("/app/", apiConfig.middlewareMetricsInc(fileServerHandler))

  mux.HandleFunc("/metrics", apiConfig.getFileserverHitsHandler)
  mux.HandleFunc("/reset", apiConfig.resetFileserverHitsHandler)
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

  fmt.Println("Server running at port 8080")
}
