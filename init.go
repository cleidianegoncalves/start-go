package start

// ---------if else
if idade >= 18 {
    fmt.Println("Você é maior de idade.")
} else {
    fmt.Println("Você é menor de idade.")
}

// ----------for

for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// --------- func
func soma(a, b, c, d int) (int, int) {
    return 1, 2
}

resultado1, _ := soma(3, 4)
fmt.Println(resultado1) // Saída: 7


// --------pacotes
import "fmt"

func main() {
    fmt.Println("Olá, mundo!")
}

// --------- array
func arraySample() {
    // Declaração e inicialização de um array de inteiros com tamanho 3
    var numeros [3]int
    numeros[0] = 1
    numeros[1] = 2
    numeros[2] = 3

    // Acessando elementos do array e imprimindo-os
    fmt.Println("Primeiro número:", numeros[0])
    fmt.Println("Segundo número:", numeros[1])
    fmt.Println("Terceiro número:", numeros[2])
}

// -------- slice
func sliceSample() {
    // Criando um slice vazio com capacidade inicial 3
    numeros := make([]int, 0, 3)

    // Adicionando elementos ao slice
    numeros = append(numeros, 1)
    numeros = append(numeros, 2)
    numeros = append(numeros, 3)
	numeros = append(numeros, 4)

    // Imprimindo o slice
    fmt.Println("Slice de números:", numeros)
}

// ----------- maps
package main

import "fmt"

func mapsSample() {
    // Criando um map de string para int
    idadePorNome := make(map[string]int)

    // Adicionando elementos ao map
    idadePorNome["Soares"] = 30
    idadePorNome["Cleidiane"] = 31
    idadePorNome["Rodolfo"] = 35

    // Acessando elementos do map e imprimindo-os
    fmt.Println("Idade de Soares:", idadePorNome["Soares"])
    fmt.Println("Idade de Cleidiane:", idadePorNome["Cleidiane"])
    fmt.Println("Idade de Rodolfo:", idadePorNome["Rodolfo"])
}

// --------- structs
type Pessoa struct {
    Nome string `bson:"name"`
    Idade int
	Cpf string
	Endereco Endereco
}

Endereco struct {
	Rua string
	Numero int
}

func structSample() {
    // Criando uma instância da struct Pessoa
    pessoa1 := Pessoa{"Soares", 30}

    // Acessando e imprimindo os campos da struct
    fmt.Println("Nome:", pessoa1.Endereco)
    fmt.Println("Idade:", pessoa1.Idade)
}

// --------- funções variádicas
// Função soma que aceita um número variável de argumentos inteiros
func soma(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func varidicSample() {
    // Chamando a função soma com diferentes números de argumentos
    fmt.Println("Soma de 1, 2, 3:", soma(1, 2, 3))
    fmt.Println("Soma de 1, 2, 3, 4, 5:", soma(1, 2, 3, 4, 5))
}

// ------------ defer
// adia a execução de uma função
func funcaoComDefer() {
    defer fmt.Println("Este é o último a ser executado")
    defer fmt.Println("Este é o segundo a ser executado")
    fmt.Println("Este é o primeiro a ser executado")
}

func sampleDefer() {
    funcaoComDefer()
}

// impressão
/*
Este é o primeiro a ser executado
Este é o segundo a ser executado
Este é o último a ser executado
*/