package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
  resp, err := http.Get("https://www.google.com")
  if err != nil {
    fmt.Println("Error: ", err)
    os.Exit(1)
  }

  // NOTE: We have to create a slice of bytes to pass to the READ method of the interface
  // NOTE: The Slice has a lot of spaces (99999) because READ can not resize the slice
  bs := make([]byte, 99999)
  resp.Body.Read(bs)

  fmt.Println(string(bs))
}
