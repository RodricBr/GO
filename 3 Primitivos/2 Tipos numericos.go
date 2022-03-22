package main

import "fmt"

func main(){
  // A primeira coisa que precisamos saber é que o valor zero para todos os
  // tipos numéricos será 0 (ZERO), ou o equivalente à zero para o tipo
  // numérico. 

  // Inteiros (Integer)
  // O primeiro tipo de inteiros que podemos usar, são os "signed integers", e
  // eles tem diversos tipos diferentes de data types. Temos o General Integer,
  // que é um integer com o tamanho não especificado. Não especificado pois
  // cada plataforma pode escolher implementar int com um tamanho diferente.
  // Independente do ambiente, um int vai ser no mínimo 32 bits, mas ele
  // poderia ser 64, ou até mesmo 128 bits, dependendo do tipo de sistema
  // que o código está rodando, e esse será o valor padrão do tipo inteiro.

  // Outros tipos:
  // int8: 8 bits, no qual podem ir de negativo 128 (-128), até 127
  // int16: 16 bits, no qual podem ir de negativo 32,768 (-32,768), até 32,767
  // int32: 32 bits, no qual podem ir aproximadamente negativo para positivo 2
  // BILÕES (-2,147,483,648 até -2,147,483,647)
  // int64: 64 bits, no qual podem ir aproximadamente entre negativo para
  // positivo 9 QUINTILHÕES, se por acaso precisarmos de um número muito grande

  // Relaciona aos "signed integers", temos os "unsigned integers".
  var n uint16 = 20 // Unsigned integer de 16 bits com valor 20
  fmt.Printf("Valor: %v, Tipo: %T\n", n, n)
  // Existe um tipo equivalente para o unsigned integer para cada signed integer,
  // Temos:
  // uint8: unsigned 8 bits integer, que pode ir de 0 até 255
  // uint16: unsigned 16 bits integer, que pode ir de 0 até 65,535
  // uint32: unsigned 32 bits integer, que pode ir de 0 até 4 BILHÕES
  // Apesar de não termos um unsigned integer de 64 bits, mas temos um tipo 
  // "byte". Um byte é uma alias para um unsigned integer de 8 bits, e a razão
  // para isso é que o unsigned integer de 8 bits é muito comum, pois é isso
  // que muitas "data streams" usam para encodar seus dados.
  // Novamente: unsigned e unsigned integer são basicamente o mesmo tipo.

  // Nós também temos vários outros tipos de operadores aritiméticos
  // que podemos usar, Exemplo:
  fmt.Println("\nTipos básicos:")

  a := 10 // Inteiro com valor de 10
  b := 3 // Inteiro com valor de 3

  fmt.Println(a + b) // Soma/Adição
  fmt.Println(a - b) // Subtração
  fmt.Println(a * b) // Multiplicação
  fmt.Println(a / b) // Divisão
  fmt.Println(a % b) // Resto (Resto de uma divisão)

  // Não podemos adicionar dois inteiros de tipos diferentes
  // E para fazer isso funcionar, precisariamos fazer uma conversão de tipos
  // em uma das variáveis para converter para o tipo da outra. Ex:
  var c int = 10
  var d int8 = 3
  //fmt.Println(c + d) // Erro!
  fmt.Println(c + int(d)) // Sucesso!

  // Bit operators (Operadores bit a bit / Lógica binária / Bitwise):
  fmt.Println("\nBitwise:")

  // 10 == 1010
  // 3  == 0011
  e := 10 // Representação binária: 1010 
  // (Apesar de serem 32 ou 64 bits, estou ignorando todos os bits no começo desses números, sobrando só 4 bits alocados)
  f := 3 // 0011 (4 bits alocados)

  fmt.Println(e & f) // AND (Procura por quais bits estão setados para o primeiro número E o segundo).
                     // Se adicionar-mos eles juntos, temos 0010 (1010 + 0011), e em binário, 0010 é igual a 2!
                     // Resultando em: 0010 (2)

  fmt.Println(e | f) // OR (Se um OU o outro está setado). Temos o primero 1,
                     // pois 'e' tem o primeiro bit setado, nenhum dos dois
                     // tem o segundo setado, então é 0, o terceiro os dois são
                     // 1, então fica 1, e o quarto apenas o 'f' tem o bit setado
                     // Resultando em: 1011 (1 + 2 + 8 = 11)

  fmt.Println(e ^ f) // Exclusive OR (OU um tem o bit setado, OU o outro tem,
                     // mas não os dois. A diferença desse para o operador OR
                     // é que o terceiro bit ambos estavam setados para 1 (true),
                     // e por isso não vamos inclui-los)
                     // Resultando em: 1001 (8 + 1 = 9)

  fmt.Println(e &^ f) // AND NOT (Oposto de OR, pois será setado 1 (true) apenas
                      // se nem um dos números tiver um bit setado. Como o primeiro
                      // bit está setado no 'e', não vamos inclui-lo, no segundo
                      // nenhum dos dois tem o bit setado, então vamos incluir-lo,
                      // no terceiro os dois estão setados, então não vamos inclui-los,
                      // e no quarto o 'f' tem um bit setado, então não vamos inclui-lo)
                      // Resultando em: 0100 (8)
}
// 1:07:07
