package main

import "fmt"

func main(){
  // Texto em GO se extende em duas categorias básicas

  // O 1º text type que temos disponível é a string, e uma string pode armazenar
  // quaquer caractere UTF-8, o que significa que nem todo tipo pode ser
  // armazenado nela, e pra isso, precisados do outros text type.
  // Podemos tratar uma string como um array (coleção de letras), utilizando
  // o nome da variável e a posição do caractere, por exemplo:
  a := "Primeiro(0), Segundo(1), Terceiro(2)"
  fmt.Printf("Valor: %v, Tipo: %T\n", a[2], a[1])
  // Mas, se executarmos isso, o resultado vai ser: Valor: 105, Tipo: uint8
  // E o que está acontecendo é que uma string em GO é um pseudônimo (alias)
  // para bytes, para printar como string, utilizamos a função: string()
  fmt.Printf("Valor: %v, Tipo: %T\n", string(a[2]), string(a[1]))

  // Concatenação de string
  a2 := " Outras strings"
  fmt.Printf("Valor: %v, Tipo: %T\n", a + a2, a + a2)
  // Outra coisa que podemos fazer com strings é converte-las para coleções de
  // bytes, o que em GO é chamada de "slice of bytes", da seguinte forma:
  a3 := []byte(a)
  fmt.Printf("Valor: %v, Tipo: %T\n", a3, a3)
  // O output é uma coleção de valores ASCII/UTF para cada caractere da string a3,
  // e podemos ver que o Tipo é "uint8", que é um alias de tipo para bytes.
  // E qual é o uso disso? Muito mais flexível de trabalhar do que com strings
  // hard coded. Exemplo: Se quisermos mandar uma resposta para um serviço web,
  // se podemos mandar a string de volta, podemos facilmente converte-las em uma
  // coleção de bytes. Mas se quisermos mandar um arquivo de volta, um arquivo
  // no nosso disco rígido é só uma coleção de bytes também. Então podemos trabalhar
  // com os dados sem nos preocupar com final de linhas, ou coisas do tipo. Conclusão:
  // Quando quisermos enviar ou receber esses dados para aplicações ou serviços
  // externos, iremos precisar tirar proveito dessa habilidade para converter
  // para um btye slice

  // O último tipo de texto que vamos trabalhar é uma runa (rune)
  // Enquanto uma string representa caracteres UTF-8, uma runa representa qualquer
  // tipo de caractere UTF-32. Enquanto qualquer caractere no UTF-32 PODE ser
  // de 32 bits, ela não NECESSITA ser de 32 bits. Por exemplo: Qualquer caractere
  // UTF-8, que tem 8 bits de comprimento, é um caractere válido para UTF-32.
  // Para saber mais sobre runas, você precisa olhar pela API de pacotes em GO
}
