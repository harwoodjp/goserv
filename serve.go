package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  args := os.Args
  if len(args) < 3 {
    fmt.Println("arg count < 3")
  } else {
    port := args[1]
    dir := args[2]    
    fs := http.FileServer(http.Dir(dir))
    http.Handle("/", fs)
    log.Println("Serving " + dir + " on port " + port)
    http.ListenAndServe(":" + port, nil)
  }
}