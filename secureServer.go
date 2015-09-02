package main

import (
  "fmt"
  "net/http"
  "flag"
  "strconv"
)

//http://stackoverflow.com/questions/23494082/golang-listenandservetls-returns-data-when-not-using-https-in-the-browser
//https://www.digitalocean.com/community/questions/how-do-i-generate-a-csr-key
const (
  CERT = "../cert.pem"
  KEY = "../key.pem"
)

func main() {
  port := flag.Int("port", 7070, "The port the server will listen on.")
  flag.Parse()

  fmt.Println("Listening on port", *port)

  fs := http.FileServer(http.Dir("./"))
  http.Handle("/", fs)

  //Can't listen on same port for both, because the port is in use?
  //http.ListenAndServe(":7071", nil)
  err := http.ListenAndServeTLS(":" + strconv.Itoa(*port), CERT, KEY, nil)
  if err != nil {
    fmt.Printf("SSL error: %s\n", err)
  }
}
