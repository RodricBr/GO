package main

import "fmt"

// Declarando pela primeira vez teste como inteiro valendo 20
var teste int = 20

func main(){
  // Como variáveis funcionam quando queremos redeclara-las
  fmt.Printf("Antes: %v\n",teste)
  var teste int = 40 // Declarando dentro da função main
  //teste := 1337 // Redeclarando impondo um novo valor à teste

  // Podemos reassinar usando: teste = 1337,
  // mas não podemos declarar uma nova variável teste, pois
  // ela já está sendo declarada na linha 10 (var teste int = 40)
  // Portando, não podemos declarar a variável duas vezes no mesmo escopo, e
  // se removemos a linha 11 (teste := 1337), o valor teste será 40, da segunda
  // declaração. A variável que está mais íntima com o escopo é priorizada,
  // isso se chama "shadowing". A variável de nível do package ainda está
  // disponível, mas a variável que recebeu o novo valor vai ser "escondido"
  // dentro da função main
  fmt.Printf("Depois: %v\n", teste)
  
  // Uma outra coisa sobre variáveis, é que não podemos deixar uma variável
  // com um valor setado sendo inutilizada

  // Naming convention:
  // Variáveis em letra minúscula estão no escopo do package, o que significa
  // que qualquer coisa que consome o package, não pode ve-la, e não pode
  // trabalhar com ela. Mas se mudarmos para maiúcsulo, então isso vai ativar
  // o compilador do GO para expor essa variável ao mundo externo. Só existem 3
  // níveis de visibilidade para variáveis no GO: 1: Se você tem uma variável no
  // nível do package, e se a letra minúscula é escopo do package, então qualquer
  // arquivo no mesmo package pode acessar essa variável. 2: Se está em letra maiúscula
  // no nível do package, então é exportado na frente do package, e se torna
  // globalmente visível. 3: Block scope. Quando nós temos uma função main, na
  // verdade ela está estabelecendo um bloco de escopo. Então quando declaro uma
  // variável var x int = 42, essa variável está no escopo do bloco main, e nunca
  // estará visível fora do bloco em sí
}
