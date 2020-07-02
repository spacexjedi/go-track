package main


import "fmt"
import f "fmt"
import "runtime"

func print_name(name string) {
  fmt.Println(name)

}

func main() {
  fmt.Println("Hello World")
  print_name("Isaac")
  fmt.Print("What happens without ln?")
  //fmt.Println('What happens with single quotes?')
  fmt.Println("Move package main to the bottom is a stupid exercise")
  f.Println("Alias to packages")
  f.Print("semicolon");
  f.Print("o");
  // the most stupid thing you can do 
  //"Hello"
  fmt.Println("Hello!" + "!" + "!" + "!" + "?")
  fmt.Println(runtime.Version())
//  fmt.Println("hello")
//	fmt.Println("how")
//	fmt.Println("are")
//	fmt.Println("you")

}

