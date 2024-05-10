package start

import "fmt"

func points() {
	// Declaração de uma variável do tipo inteiro
	var x int = 10

	// Declaração de um ponteiro para um inteiro
	var ponteiro *int

	// Atribuindo o endereço de memória de x ao ponteiro
	ponteiro = &x

	// Impressão do valor de x e do valor apontado pelo ponteiro
	fmt.Println("Valor de x:", x)
	fmt.Println("Valor apontado pelo ponteiro:", *ponteiro)

	// Modificação do valor de x através do ponteiro
	*ponteiro = 20

	// Impressão do novo valor de x
	fmt.Println("Novo valor de x:", x)

	// ponteiros podem ser nil
}
