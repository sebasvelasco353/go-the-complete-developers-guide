package main

import "fmt"

func main() {
  // NOTE: we can create a map using literal syntax
  colors := map[string]string{
    "red": "#FF0000",
    "green": "#00FF00",
    "blue": "#0000FF",
  }
  
  //NOTE: can also create one using var syntax -> var otherColors map[string]string
  //NOTE: the last option is to use make -> anotherColors := make(map[string]string)
  //NOTE: we can add values to a map by using square braces -> colors["white"] = "#FFFFFF"
  //NOTE: To remove something from the map we use the function delete -> delete(colors, "white")

  printMap(colors)
}

func printMap(c map[string]string){
  for key, value := range c {
    fmt.Printf("Hex code for %v is %v\n", key, value)
  }
}
