package main

import (
  "net/http"
  "log"
  "os"
)

func main() {
  args := os.Args
  fs := http.FileServer(http.Dir(args[1]))
  http.Handle("/", fs)

  log.Println("Listening on 3000...")
  http.ListenAndServe(":3000", nil)
}