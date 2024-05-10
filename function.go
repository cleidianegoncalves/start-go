package start

import "fmt"

// ----------- funções normais
func somar(a, b int) int {
	return a + b
}

// ----------- funções variádicas
func media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}

// -------------- funções recursivas
func fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * fatorial(n-1)
}

// ------------- funções anonimas (clojures)
func clojure() {
	// Função anônima para dobrar um número
	dobrar := func(x int) int {
		return x * 2
	}

	resultado := dobrar(5) // resultado será 10
	fmt.Print(resultado)
}

// --------------- método receptor ou apenas metodo
// Definindo um tipo struct chamado "Ponto"
type Ponto struct {
	x, y int
}

// Método receptor para o tipo Ponto
func (p Ponto) Imprimir() {
	fmt.Printf("Coordenadas: (%d, %d)\n", p.x, p.y)
}

func sampleMetodo() {
	// Criando uma instância da struct Ponto
	ponto := Ponto{3, 5}

	// Chamando o método Imprimir da instância ponto
	ponto.Imprimir()
}
