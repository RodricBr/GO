package main

import "fmt"

func main(){
  // Usado em Data Science, coisa que é raro ver em outras linguagens
  // Existem dois tipos de números complexos:
  // complex64:
  // complex128:
  // Basicamente pegando o float64 + um float64, OU um float32 + um float32
  var n complex64 = 1 + 2i // O parser do GO entende que a literal 'i'
                           // é um número imaginário

  fmt.Printf("Valor: %v, Tipo: %T\n", real(n), real(n))
  fmt.Printf("Valor: %v, Tipo: %T\n", imag(n), imag(n))
  // E se precisarmos pegar o número real do número imaginário?
  // Para isso, o GO tem duas funções imbutidas (built-in), utilizando a função
  // real() e a função imag()
  // E elas vão olhar para o número complexo que providenciamos e irão extrair
  // a parte real ou da parte imaginária
  // O complemento para essas duas funções é o função complex(), e utiliza
  // dois números. O primeiro é a parte real, e o segundo número é a parte
  // imaginária. Exemplo:
  var d complex128 = complex(5, 12)
  fmt.Printf("Valor: %v, Tipo: %T\n", d, d)

  // Operações disponíveis:
  a := 1 + 2i
  b := 2 + 5.2i
  fmt.Println(a + b)
  fmt.Println(a - b)
  fmt.Println(a * b)
  fmt.Println(a / b)
}
