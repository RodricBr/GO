package main

import "fmt"

const (
  PrintColor = "\033[38;5;%dm%s\033[39;49m\n"
)

func main() {
  for x := 0; x < 256; x++ { // Vai ser executado 255 vezes
    fmt.Printf(PrintColor, j, "Olá, mundo!")
  }
}
