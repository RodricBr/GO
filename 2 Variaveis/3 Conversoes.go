package main

import ("fmt"; "strconv")

// Operador de conversão
// Convertindo um tipo de variável para outro

func main(){
  var i int = 42
  fmt.Printf("Valor: %v, Tipo: %T\n", i, i)

  // Tratar o "i" como um valor do tipo float e assinar esse valor à "j"
  //var j string
  //j = string(i) // Posso fazer isso utilizando esse operador de conversão
  //fmt.Printf("Valor: %v, Tipo: %T", j, j)
  // Somos obrigados a fazer uma conversão explicita quando estamos mudando tipos
  // Dessa forma, é nossa responsabilidade em entender se estamos perdendo informação ou não
  // Outra coisa que é importante saber é se eu decidir trabalhar com strings,
  // que é muito comum de se utilizar, de tentar converter um int para string.
  // Então se eu converter o int 42 para string, o resultado vai sair o UNICODE
  // "*", ao invés do número 42. E para entender como isso funciona, precisamos
  // entender como strings funcionam com GO. Uma string é apenas uma "alias" para
  // uma sequencia de bytes. Então quando pedimos para a aplicação converter 42
  // para uma string, ela procura qual o caractere UNICODE é equivalente a 42, no
  // case, 42 é *.
  // Se quisermos converter entre strings para números, precisamos importar uma
  // package de conversão de string chamada "strconv"
  var j string
  j = strconv.Itoa(i) // Itoa == Integer to ASCII
  fmt.Printf("Valor: %v, Tipo: %T", j, j)
  // Conclusão: se precisamos trabalhar com conversão entre números e strings,
  // então precisamos usar o package strconv, mas se queremos converter entre
  // tipos numéricos, tenha em mente que você não pode fazer essa conversão implicitamente,
  // você tem que fazer explicitamente faze-la. E se não precisa fazer isso, então
  // use o tipo como uma função.
}
