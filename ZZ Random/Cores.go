package main

import "fmt"

// Cores
var (
 deuCERTO = "\033[32m%s\033[0m"
 deuERRADO = "\033[31m%s\033[0m"
 deuMAISouMENOS = "\033[33m%s\033[0m"
)

func main(){

  VERDE := "\033[32m"
  FIM := "\033[0m"
  // Método 1
  fmt.Println(string(VERDE), "VERDE\n", string(FIM))

  // Método 2
  fmt.Printf(deuCERTO, "Deu certo!\n")

  fmt.Printf(deuERRADO, "Deu errado!\n")

  fmt.Printf(deuMAISouMENOS, "Deu + ou -\n")
}
