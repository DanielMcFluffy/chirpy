package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
  fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc (next http.Handler) http.Handler {
  handler := func(res http.ResponseWriter, req *http.Request) {
    cfg.fileserverHits.Add(1)
    next.ServeHTTP(res, req)
  }
  return http.HandlerFunc(handler)
}

func (cfg *apiConfig) getFileserverHitsHandler(res http.ResponseWriter, req *http.Request) {
  data := fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())
  res.Header().Set("Content-Type", "text/plain; charset=utf-8")
  res.WriteHeader(200)
  res.Write([]byte(data))
}

func (cfg *apiConfig) resetFileserverHitsHandler (res http.ResponseWriter, req *http.Request) {
  cfg.fileserverHits.Store(0)
  data := fmt.Sprintf("Hits: 0")
  res.Header().Set("Content-Type", "text/plain; charset=utf-8")
  res.WriteHeader(200)
  res.Write([]byte(data))
}
