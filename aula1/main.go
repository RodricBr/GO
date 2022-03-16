// Aual 1: Estrutura de um programa GO

// Golang Tutorial for Beginners | Full Go Course
// v=yyUHQIec83I

// Todo o nosso código deve pertencer a um 'package'
// A primeira instrução no arquivo GO deve ser 'package ...'
package main

import "fmt" // fmt == Format package
// Os programas Go são organizados em pacotes.
// A biblioteca padrão do Go, fornece diferentes
// pacotes principais para usarmos.
// "fmt" é um deles, que você pode usar
// ao importa-la

// Um programa pode usar apenas uma função "main",
// porque você só pode ter 1 entrypoint
func main(){
  fmt.Println("Ola, mundo") // Println == new line no final (exemplo\n)
}

// Criando um modulo: go mod init aprendendo-go
