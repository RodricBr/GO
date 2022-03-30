# Sumário <br>

```markdown
1- Tipos booleanos
 - Valores são verdadeiro ou falso (true / false)
 - Não é um pseudônimo para outros tipos (ex: int)
 - Valor ZERO é false

-------------------------------------------------------------

2- Tipos numérios
 + Integer (int)
   - Tipo int tem um valor variante, mas um mínimo de 32 bits
   - 8 bit (int8) e 64 bit (int64)

 + Unsigned Integer (uint)
   - 8 bit (byte e uint8) e 32 bit (uint32)

 + Operadores aritiméticos
   - Adição, subtração, multiplicação, divisão e
   - resto de divisão

 + Operadores bit-wise
   - And, or, xor, not

 + Valor zero é 0

 + Não se pode misturar tipos da mesma familia (uint16 + uint32 = erro)

 + Floating point (float)
   - Segue o padrão IEEE-754
   - Valor zero é 0
   - Versões de 32 e 64 bit

 + Literal style
   - Decimal (1.33)
   - Exponencial (13e37 OU 1E37)
   - Misturado (13.3e37)

 + Operações aritiméticas
   - Adição, subtração, multiplicação, divisão
 
 + Números complexos
   - Valor zero é (0+0i)
   - Versões de 64 e 128 bit
   + Funções built-in
     - complex --> Faz números complexos de dois floats
     - real    --> Pega a parte real como float
     - imag    --> Pega a parte imaginária como float
   - UTF-8
   - Imutável
   - Pode ser concatenada usando o operador (+)
   - Pode ser convertida para []byte

   + Runa
     - UTF-32
     - Alias para int32
     + Métodos especiais normalmente precisam processar
       - e.g: strings.Reader#ReadRune

-------------------------------------------------------------

## 3 - Tipos de textos 
```

