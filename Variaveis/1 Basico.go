package main

import "fmt"

// Declarando variável no nível do package
// Nesse nível, não podemos usar a sintaxe :=, temos que
// utilizar a sintaxe de declaração inteira
var x int = 20

// Declarando diversas variáveis
// Agora todas essas variáveis vão ser declaradas pois
// estão dentro dessa bloco var
var (
  meuNome string = "Rodric"
  minhaIdade2 int = 21
  souHomem string = "homem"
)

var (
  counter int = 0
)

func main(){
  // Existem 3 meios de se declararem
  // funções em GO.
  // Primeiro, vamos declarar a função em sí

  var idade int = 20 // Variável chamada idade do tipo inteiro (integer),
  // e assinando um valor numérico à variável idade
  // Ao invés de usarmos esse textão pra declarar uma função com
  // um valor inteiro, pq já que estamos assinando o valor 20 à variável,
  // o compilador pode descobrir qual data type eu preciso ter

  //idade2 := 20 // O compilador já entende que 20 é int e idade recebe esse valor

  // Outros exemplos:
  //var i int // Variável pode ser assinada no decorrer do programa
  //i = 21
  //var j int = 22 // Podemos mudar de int, para float32, ou char... (mais controlável)
  //k := 99. // O . vai transformar a variável que era int em float64
  fmt.Printf("Meu nome é %v\n", meuNome)
  fmt.Printf("Minha idade é %v\n", x)
  fmt.Printf("Daqui alguns meses faço %v anos\n", minhaIdade2)
  fmt.Printf("Eu sou %v\n", souHomem)
  fmt.Printf("Contador: %v\n", counter)
  fmt.Printf("%v, %T\n", idade, idade)
}
