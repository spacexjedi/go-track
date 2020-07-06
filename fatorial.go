package main


func fatorial(x int8) int8 {
  if x > 1 {
  return x * fatorial(x - 1)
  } else {
    return 1
  }
}
