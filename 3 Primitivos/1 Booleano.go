// https://youtu.be/YS4e4q9oBaU?t=3426
package main

import "fmt"

func main(){
  n := 1 == 1 // Operador de sinal duplo de igual para testar se 1 é igual a 1
  m := 1 == 2 // "Equals operator"
  // É importante saber sobre os primitivos, que toda vez que inicializamos
  // umas variável no GO, ele começa com valor 0 (ZERO). Então nós estamos
  // definindo um valor para 'n' e 'm' no exemplo. Em GO, se criarmos uma
  // variável: (var n bool), sem setar nenhum valor à ela, quando usarmos o
  // Printf, o valor padrão será false, que é um valor zero, e o valor zero
  // para booleano é FALSE!

  fmt.Printf("Valor: %v, Tipo: %T\n", n, n)
  fmt.Printf("Valor: %v, Tipo: %T\n", m, m)
}
