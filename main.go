package main

import "fmt"

func print_name(name string) {
  fmt.Println(name)

}

func main() {
  fmt.Println("Hello World")
  print_name("Isaac")
  fmt.Print("What happens without ln?")
  //fmt.Println('What happens with single quotes?')
  fmt.Println("Move package main to the bottom is a stupid exercise")

}

