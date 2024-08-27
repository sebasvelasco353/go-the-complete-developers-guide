package main

import (
	"fmt"
	"net/http"
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

  // for { // NOTE: This is an infinite loop
  for link := range c {
    // NOTE: this is other type of infinite loop, basically it waits for a value on c and sets it in variable link,
    // then continues waiting
    go checkLink(link, c)
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
