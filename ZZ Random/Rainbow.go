package main

// By: github.com/Yariya

import ("fmt"; "math/rand"; "os"; "time")

// Cores em um array
func rainbow(i string, timeout int){
  colors := []string{
    "\033[;31m",
    "\033[;34m",
    "\033[;32m",
    "\033[;35m",
    // Adicionar mais cores na sequencia...
  }

  for {
    n := rand.Int() % len(colors)
    fmt.Printf("\r%s%s", colors[n], i)
    os.Stdout.Sync()
    time.Sleep(time.Millisecond * time.Duration(timeout))
  }
}

func main(){
  rainbow("Olá, mundo", 300) // Texto & Timeout
}
