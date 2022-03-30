# Sumário <br>

```markdown
1- Declaração de variáveis
  + var foo int
  + foo := 42

-------------------------------------------------------------

2- Não podemos redeclarar variáveis,
   mas podemos "sombrea-las" (shadowing)

-------------------------------------------------------------

3- Todas as variáveis com valores setados devem ser usadas

-------------------------------------------------------------

4- Visibilidade <br>
  + Letra minúscula primeira letra no escopo do package
  + Letra maiúscula primeira letra para exportar
  + Sem escopo privado

-------------------------------------------------------------

5- Nomeando convenções (Naming Conventions) <br>
  + Pascal ou camelCase
  + Capitalizar acronomos (siglas), (HTTP, URL, DNS...)

-------------------------------------------------------------

6- Conversões de Tipos <br>
  + DestinationType(variável)
  + Usando `strconv` package(pacote/livraria...) para strings
