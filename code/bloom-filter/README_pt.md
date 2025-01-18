# Bloom Filter

🌍 *[English](README.md)*

Um **Bloom Filter** é uma estrutura de dados probabilística que mapeia um conjunto de elementos para um vetor de bits. A operação básica de um Bloom Filter é verificar se um elemento pertence ou não a um conjunto, com a possibilidade de **falsos positivos** (afirmar que um elemento está no conjunto quando, na realidade, não está), mas nunca **falsos negativos** (afirmar que um elemento não está no conjunto quando ele realmente está).

O funcionamento básico do Bloom Filter é o seguinte:

- **Inserção**: Quando um item é inserido, várias funções de hash são usadas para calcular vários índices no vetor de bits. Os bits nesses índices são definidos como 1.
- **Consulta**: Quando você consulta a presença de um item, as funções de hash calculam os mesmos índices. Se **todos os bits correspondentes a esses índices estiverem marcados como 1**, o Bloom Filter retorna que o item **pode** estar presente. Caso contrário, ele retorna que o item **não está** no conjunto.

## Descrição Matemática de um Bloom Filter

Podemos definir o comportamento de um Bloom Filter como uma função que mapeia um conjunto de elementos para um vetor de bits, e a operação de verificação de um elemento em relação a esse conjunto.

### Definindo o Bloom Filter

Dado:

- Um conjunto de elementos $S \subseteq \Sigma^*$, onde $\Sigma^*$ é o conjunto de todas as strings possíveis.
- Um **vetor de bits** $B \in \lbrace 0, 1\rbrace^m$ de tamanho $m$, onde $m$ é o número de bits disponíveis no Bloom Filter.
- $k$ funções de hash distintas $h_1, h_2, ..., h_k$, onde cada $h_i: \Sigma^* \rightarrow \lbrace 0, 1, \dots, m-1\rbrace$ é uma função de hash que mapeia um elemento de $\Sigma^*$ para um índice entre 0 e $m-1$.

#### Função de inserção:

A função de inserção do Bloom Filter, dada por $\phi_{\text{add}}: {0,1}^m\times\Sigma^* \rightarrow \lbrace 0, 1\rbrace^m$, atualiza o vetor de bits $B$ quando um elemento $x \in \Sigma^*$ é adicionado. Cada função de hash $h_i(x)$ gera um índice no vetor $B$, e o bit correspondente é marcado como 1. Denotamos $B[i]$ a i-ésima coordenada de $B$. Assim, definimos a função de inserção como:

$$
\phi_{\text{add}}(B_0,x) = B_1 
$$

sendo

$$
B_1[i] = \begin{cases}
1 &\text{ se }i\in\lbrace h_1(x), \dots, h_k(x)\rbrace,\\
B_0[i] &\text{ caso contrário.}
\end{cases}
$$

#### Função de consulta:

A função de consulta $\phi_{\text{query}}: \lbrace 0,1\rbrace^m \times \Sigma^* \rightarrow \lbrace 0, 1\rbrace$ verifica se um elemento $x \in \Sigma^*$ pode estar no conjunto $B$ representado pelo Bloom Filter. Para isso, verificamos os bits correspondentes aos índices $h_1(x), h_2(x), \dots, h_k(x)$. Se algum desses bits for 0, o elemento definitivamente **não** está no conjunto. Se todos os bits forem 1, o elemento **pode** estar no conjunto (com a possibilidade de falsos positivos).

A função de consulta é definida como:

$$
\phi_{\text{query}}(x) = 
\begin{cases} 
1 & \text{ se } B[h_i(x)] = 1 \text{, }\forall i = 1, 2, \dots, k, \\
0 & \text{ caso contrário.}
\end{cases}
$$

Ou seja, $x$ **pode** estar no conjunto se todos os bits $B[h_1(x)], B[h_2(x)], \dots, B[h_k(x)]$ forem iguais a 1, e **definitivamente não** está no conjunto se qualquer um desses bits for 0.


## Probabilidade de Falsos Positivos em um Bloom Filter

A **probabilidade de falsos positivos** é uma das principais características de um **Bloom Filter** e é fundamental para entender o comportamento probabilístico da estrutura. Em um Bloom Filter, um **falso positivo** ocorre quando o filtro retorna **verdadeiro** (ou seja, diz que um item está no conjunto) para um item que **não foi** realmente inserido no filtro.

Lembre-se de que o Bloom Filter nunca retorna um falso negativo. Ou seja, se um item foi inserido no filtro, a consulta sempre indicará corretamente que ele está no conjunto. O que pode acontecer é que o filtro, com base na combinação das funções de hash, possa indicar erroneamente que um item **não inserido** está presente no conjunto.

### Cálculo da Probabilidade de Falsos Positivos

A probabilidade de falsos positivos em um Bloom Filter depende de três fatores principais:

- **$n$**: O número de elementos inseridos no filtro.
- **$m$**: O número de bits no vetor.
- **$k$**: O número de funções de hash utilizadas.

Quando um item é inserido no Bloom Filter, os bits correspondentes aos índices gerados pelas funções de hash são definidos como 1. Ao inserir vários elementos, a probabilidade de um bit ser 1 aumenta. A **probabilidade de falsos positivos** depende da quantidade de bits que são definidos como 1 e da chance de um item não inserido coincidir nos mesmos índices.

### Probabilidade de um bit ser 1

Para calcular a probabilidade de falsos positivos, precisamos primeiro entender a probabilidade de um bit específico ser 1 após a inserção de $n$ elementos. Isso depende do número de funções de hash $k$ e do número de bits $m$.

Cada elemento inserido no filtro afeta $k$ bits (um para cada função de hash), então após a inserção de $n$ elementos, o número total de marcações de bits é $kn$. No entanto, um bit específico não é necessariamente marcado, já que a função de hash distribui os valores uniformemente. A probabilidade de um bit não ser marcado por um único elemento inserido é:

$$
P(\text{bit não marcado por um elemento}) = 1 - \frac{1}{m}
$$

Após inserir $n$ elementos, a probabilidade de um bit não ser marcado por **nenhum** desses $n$ elementos é:

$$
P(\text{bit não marcado após n elementos}) = \left( 1 - \frac{1}{m} \right)^{kn}
$$

Portanto, a probabilidade de um bit ser **marcado** após $n$ inserções é:

$$
P(\text{bit marcado}) = 1 - \left( 1 - \frac{1}{m} \right)^{kn}
$$

### Probabilidade de um falso positivo

Agora, vamos calcular a probabilidade de um falso positivo para um item **não inserido**.

Para um item que não foi inserido, ele será considerado presente no conjunto se, para todas as $k$ funções de hash, os $k$ bits correspondentes forem 1. Como os bits são marcados independentemente, a probabilidade de um **bit específico** estar marcado como 1 (quando o item não foi inserido) é $P(\text{bit marcado})$.

Portanto, a probabilidade de **todos os $k$ bits** correspondentes a um item não inserido estarem marcados como 1 é:

$$
P_{\text{falso positivo}} = \left( 1 - \frac{1}{m} \right)^{kn}
$$

Esta é a probabilidade de que, para um item não inserido, todos os $k$ bits que ele consulta sejam 1. Se todos os bits forem 1, o Bloom Filter retornará "verdadeiro", indicando erroneamente que o item está no conjunto.

### Aproximação da Probabilidade de Falsos Positivos

Para um valor grande de $n$ (número de elementos inseridos) e $m$ (tamanho do vetor de bits), podemos usar uma aproximação mais simples para a probabilidade de falsos positivos. A expressão $\left(1 - \frac{1}{m}\right)^{kn}$ pode ser aproximada por:

$$
P_{\text{falso positivo}} \approx e^{-\frac{kn}{m}}
$$

Isso resulta em uma fórmula muito mais simples e fácil de entender, que nos diz que a probabilidade de um falso positivo diminui exponencialmente com o aumento do número de bits $m$ e do número de funções de hash $k$.

### Otimização do número de funções de hash

Para minimizar a taxa de falsos positivos, é importante escolher o número adequado de funções de hash $k$. Se $k$ for muito pequeno, o filtro pode ser ineficiente, marcando bits em excesso, enquanto se $k$ for muito grande, a probabilidade de um bit ser marcado como 1 também pode ser muito alta, o que aumenta os falsos positivos.

A escolha ótima de $k$ que minimiza a probabilidade de falsos positivos pode ser calculada pela derivada da expressão $P_{\text{falso positivo}}$, resultando na seguinte fórmula:

$$
k^* = \frac{m}{n} \ln(2)
$$

Onde:

- $k^*$ é o número ótimo de funções de hash.
- $m$ é o tamanho do vetor de bits.
- $n$ é o número de elementos inseridos no filtro.

### Resumo da Fórmula para a Probabilidade de Falsos Positivos

Combinando tudo o que discutimos, a **probabilidade de falsos positivos** em um Bloom Filter é dada por:

$$
P_{\text{falso positivo}} \approx e^{-\frac{kn}{m}}
$$

Onde:

- $k$ é o número de funções de hash.
- $n$ é o número de elementos inseridos.
- $m$ é o tamanho do vetor de bits.

Essa fórmula mostra que a probabilidade de falsos positivos depende da relação entre o número de bits, o número de funções de hash e o número de elementos inseridos. 

---

### Implementação Eficiente de um Bloom Filter Usando MurmurHash3

Esta implementação utiliza a função de hash **MurmurHash3** com uma semente, permitindo a geração de múltiplas funções de hash, o que melhora o desempenho e a dispersão dos bits.

#### **Principais Características**
1. **MurmurHash3**: Uma função de hash não criptográfica, projetada para ser rápida e com baixa taxa de colisão.
2. **Manipulação Eficiente de Bits**: Usa um bitset compacto implementado como um array de `uint64` para maior eficiência de memória.
3. **Função de Hash Única com Semente**: Uma única função de hash é reutilizada com sementes diferentes, simulando múltiplas funções de hash, o que simplifica a implementação e melhora a performance.

#### **Visão Geral do Código**

A funcionalidade principal está dividida nos seguintes componentes:

1. **Inicialização**:
   - A função `NewBloomFilter` cria e configura o Bloom Filter com o tamanho desejado e o número de funções de hash.

2. **MurmurHash3 com Semente**:
   - A função `murmurHash3` calcula um valor de hash de 64 bits com base nos dados de entrada e uma semente. Cada semente gera uma função de hash distinta.

3. **Operação de Adição (`Add`)**:
   - O método `Add` calcula múltiplos valores de hash para o item dado e ajusta os bits correspondentes no bitset.

4. **Operação de Verificação (`Contains`)**:
   - O método `Contains` verifica se todos os bits correspondentes aos valores de hash do item estão ajustados. Se algum bit não estiver, o item definitivamente não pertence ao conjunto.

#### **Como Usar**

Veja como usar esta implementação de Bloom Filter em sua aplicação Go:

1. **Importe os Pacotes Necessários**:
   - Certifique-se de importar `encoding/binary`.

2. **Crie um Bloom Filter**:
   ```go
   bf := NewBloomFilter(1000, 3)
   ```
   - `1000`: O tamanho do bitset.
   - `3`: O número de funções de hash.

3. **Adicione Itens ao Bloom Filter**:
   ```go
   bf.Add("exemplo1")
   bf.Add("exemplo2")
   ```

4. **Verifique a Existência de Itens**:
   ```go
   println(bf.Contains("exemplo1")) // Saída: true
   println(bf.Contains("exemplo3")) // Saída: false (ou true no caso de um falso positivo)
   ```

#### **Programa de Exemplo**
```go
package main

func main() {
    // Crie um Bloom Filter com tamanho de 1000 bits e 3 funções de hash
    bf := NewBloomFilter(1000, 3)

    // Adicione elementos ao Bloom Filter
    bf.Add("exemplo1")
    bf.Add("exemplo2")

    // Verifique a existência de elementos
    println(bf.Contains("exemplo1")) // true
    println(bf.Contains("exemplo2")) // true
    println(bf.Contains("exemplo3")) // false (possível falso positivo)
}
```

#### **Considerações de Desempenho**
1. **Eficiência Espacial**: O bitset é implementado como um array de `uint64` para reduzir o uso de memória.
2. **Qualidade do Hash**: O MurmurHash3 oferece excelente dispersão e rapidez, garantindo baixas taxas de colisão.
3. **Trade-Offs**:
   - A probabilidade de falsos positivos aumenta conforme mais itens são adicionados, mas pode ser reduzida ajustando o tamanho do bitset e o número de funções de hash.

Esta implementação oferece um equilíbrio entre simplicidade, desempenho e precisão, tornando-se uma excelente escolha para aplicações que necessitam de verificações rápidas e eficientes de pertencimento a conjuntos.

---

### Considerações Finais

- **Taxa de Falsos Positivos**: O Bloom Filter oferece uma excelente eficiência de espaço, mas com a possibilidade de falsos positivos. A probabilidade de falsos positivos pode ser controlada ajustando o número de funções de hash $k$ e o tamanho do vetor de bits $m$.
- **Impacto da Capacidade**: Quanto mais elementos você inserir no filtro ($n$), maior será a probabilidade de colisões de hash, o que aumenta a chance de falsos positivos.

Veja também [Bloom Filter Calculator](https://hur.st/bloomfilter/?n=30&p=1.0E-7&m=&k=6) 