package start

import (
	"fmt"
	"math"
)

// Definindo uma interface chamada "Forma"
type Forma interface {
	Area() float64
}

// Definindo um tipo "Circulo" que implementa a interface "Forma"
type Circulo struct {
	Raio float64
}

// Implementação do método "Area" para o tipo "Circulo"
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

// Definindo um tipo "Retangulo" que implementa a interface "Forma"
type Retangulo struct {
	Largura, Altura float64
}

// Implementação do método "Area" para o tipo "Retangulo"
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func sampleInterface() {
	// Criando um círculo e um retângulo
	circulo := Circulo{Raio: 5}
	retangulo := Retangulo{Largura: 3, Altura: 4}

	// Criando uma slice de formas (interface)
	formas := []Forma{circulo, retangulo}

	// Iterando sobre as formas e calculando suas áreas
	for _, forma := range formas {
		fmt.Printf("Área da forma: %.2f\n", forma.Area())
	}
}
