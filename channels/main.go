package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
  links := []string {
    "https://google.com",
    "https://facebook.com",
    "https://go.dev/",
    "https://amazon.com",
  }
  c := make(chan string)

  for _, link := range links {
    go checkLink(link, c)
  }

  // NOTE: This is an infinite loop
  for l := range c { // NOTE: waits for value on c and sets it as variable l and continues waiting
    go func(link string) {
      time.Sleep(time.Second * 5)
      checkLink(link, c)
    }(l) // NOTE: This is a function literal, the Go equivalent to an anonymous function in js
  }
}

func checkLink(link string, c chan string) {
  _, err := http.Get(link)
  if err != nil {
    fmt.Println(link, "might be down")
    c <- link 
    return
  }
  fmt.Println(link, "is working")
  c <- link
}
