package main

import (
	"encoding/json"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  serverJSON, _ := json.Marshal(r.Context().Value(http.LocalAddrContextKey).(net.Addr))
  w.Header().Set("content-type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Cache-Control", "no-store")
  w.WriteHeader(http.StatusOK)
  w.Write(serverJSON)
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", handler)
  server := &http.Server{Addr: ":80", Handler: mux}
  server.SetKeepAlivesEnabled(false)
  err := server.ListenAndServe()
  if err != nil {
    panic(err)
  }
}
