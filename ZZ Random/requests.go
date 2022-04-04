package main

import (
  "bufio"
  "os"
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
)

// Pega o stdout e joga no stdin(esse programa)
// Ex: echo "vulnweb.com" | ./requests.go
func main(){
  std := bufio.NewScanner(os.Stdin)
  for std.Scan() {
    var URL string = std.Text()
    //fmt.Printf(URL)

    RESP, ERRO := http.Get(URL)
    fmt.Printf("Enviando request GET para: %s\n\n", URL)

    // Quando o ERRO é nil, a response vai conter o body da response e vice verse
    if ERRO != nil {
      log.Fatalln(ERRO)
    }

    // Lendo o body da response:
    BODY, ERRO := ioutil.ReadAll(RESP.Body)
    if ERRO != nil {
      log.Fatalln(ERRO)
    }

    // Convertendo o body para string
    SB := string(BODY)
    log.Printf(SB)

  }
}
