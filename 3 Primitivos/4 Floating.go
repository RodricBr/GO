package main

import "fmt"

func main(){
  // Temos vários tipos diferentes de inteiros, o que nos permite armazenar muitos
  // tipos tamanhos diferentes de números. Mas com tipos inteiros, podemos apenas
  // armazenar inteiros, podendo ser positivo ou negativo, ou zero, mas não
  // podemos armazenar números decimais
  // Para podermos armazenar números decimais em GO, precisamos usar o número
  // float, ou flutuante
  // Os números flutuantes em GO seguem o padrão iEEE 754. E nesse padrão, nós
  // vamos utilizar dois dos tipos. Temos o número Floating point de 32 bits, e
  // 64 bits.
  // float32: armazena números entre ±1.18E-38 até ±3.4E38 (Mais ou menos 1.18
  // elevado a -38)
  // float64: armazena números entre ±2.23E-308 até ±1.8E308

  // Exemplos:
  a := 3.14 // Definindo uma literal num float point (quase sempre)
  a = 13.3e72 // Podemos usar notação exponencial
  a = 2.1E14 // Outro exemplo de notação exponencial usando 'E' maiúsculo
  fmt.Printf("Valor: %v, Tipo: %T\n", a, a)

  // Ou podemos declara-las explicitamente, dessa forma:
  var b float32 = 3.14
  fmt.Printf("Valor: %v, Tipo: %T\n", b, b)
  // Importante:
  // Se formos usar a sintaxe de inicialização num decimal, sempre será um
  // inicializado para float 64, então não podemos fazer operações aritiméticas
  // entre float64's e float32's. Se apenas estivermos usando sintaxe de inicialização, precisamos ter certeza de que tudo estará funcionando em float64. E se esquecermos? Fodase, o compilador vai explicar isso pra gente para poder-mos arrumar!

  // Agora falando de operações aritiméticas, vemos as operações que podemos
  // utilizar com número flutuantes. Não temos o % (Resto da divisão), não
  // podemos usar bitwise ou o operador de bit shifting. (Apenas com inteiros)
  c := 10.2
  d := 3.7
  fmt.Println(c + d)
  fmt.Println(c - d)
  fmt.Println(c * d)
  fmt.Println(c / d)
}
