package main

import "fmt"

func main(){
  // O último exemplo com inteiros é com o chamado "Bit Shifting"
  a := 8
  // 8 na verdade é 2^3, e quando fazemos bit shifting, estamos basicamente
  // adicionando à esse exponente, enquanto estivermos lidando com 2 elevado

  fmt.Println(a << 3) // Vai mudar 3 lugares para a esquerda (Multiplicar)
                      // 2^3 * 2^3 = 2^6, que é igual a 64

  fmt.Println(a >> 3) // Vai mudar 3 lugares para a direita (Dividir)
                      // 2^3 / 2^3 = 2^0, e qualquer numero maior que zero elevado a zero é 1
}
