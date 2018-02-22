package main

import (
  "net/http"
  "log"
  "os"
  "fmt"
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

  // fs := http.FileServer(http.Dir(args[1]))
  // http.Handle("/", fs)

  // log.Println("Listening on 3000...")
  // http.ListenAndServe(":3000", nil)
}