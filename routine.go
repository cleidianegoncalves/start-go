package start

import (
	"fmt"
	"sync"
	"time"
)

// --------------- go routine
/*
	paralelismo concorrente
	utilizando a palavra go
	Go usa um escalonador de goroutines que distribui o trabalho em um conjunto de threads
	de sistema operacional, permitindo que as goroutines sejam executadas concorrentemente.
*/

// ------------ channel
/*
	Os canais são uma maneira de permitir a comunicação entre goroutines de forma segura e sincronizada.
	podem ser usados para sincronizar goroutines,
	garantindo que uma goroutine não prossiga até que outra tenha enviado um valor relevante.
	usa a função make() com a palavra-chave chan seguida do tipo de dado a ser transportado.
	Por exemplo: canal := make(chan int)
*/

func calcularQuadrado(numero int, resultado chan int) {
	quadrado := numero * numero
	resultado <- quadrado
}

func sampleParallel() {
	// Criando um canal para transportar inteiros
	resultado := make(chan int)

	// Executando uma go routine para calcular o quadrado de 4
	go calcularQuadrado(4, resultado)
	// Executando outra go routine para calcular o quadrado de 9
	go calcularQuadrado(9, resultado)

	// Recebendo e imprimindo o resultado do canal para o quadrado de 4
	quadrado1 := <-resultado
	fmt.Println("O quadrado de 4 é:", quadrado1)

	// Recebendo e imprimindo o resultado do canal para o quadrado de 9
	quadrado2 := <-resultado
	fmt.Println("O quadrado de 9 é:", quadrado2)
}

// ------------ select
/*
	É uma construção que permite aguardar múltiplas operações de comunicação em canais.
*/

func sampleSelect() {
	// Criando canais para transportar inteiros
	canal1 := make(chan int)
	canal2 := make(chan int)

	// Executando uma go routine para calcular o quadrado de 4
	go calcularQuadrado(4, canal1)

	// Executando outra go routine para calcular o quadrado de 9
	go calcularQuadrado(9, canal2)

	// Utilizando select para receber o resultado do canal que estiver pronto primeiro
	select {
	case quadrado1 := <-canal1:
		fmt.Println("O quadrado de 4 é:", quadrado1)
	case quadrado2 := <-canal2:
		fmt.Println("O quadrado de 9 é:", quadrado2)
	}
}

// ------------------ problemas
// panic não tratatado
func samplePanic() {
	go func() {
		panic("Erro inesperado")
	}()
}

// memory leak
func workerVazamento(id int, ch chan bool) {
	fmt.Printf("Worker %d iniciado\n", id)
	time.Sleep(time.Second) // Simula algum trabalho
	fmt.Printf("Worker %d concluído\n", id)
	ch <- true // Sinaliza que o trabalho está concluído
}

func sampleVazamentoMemoria() {
	ch := make(chan bool)

	for i := 0; i < 10; i++ {
		go workerVazamento(i, ch)
	}

	// Espera que todas as goroutines concluam
	for i := 0; i < 10; i++ {
		<-ch
	}

	// O programa termina aqui
	// No entanto, se alguma goroutine ainda estiver em execução,
	// o recurso associado a ela não será liberado
}

// ---------------- wait group de routines
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca a goroutine como concluída quando a função retorna
	fmt.Printf("Worker %d começou\n", id)
	time.Sleep(time.Second) // Simulando algum trabalho
	fmt.Printf("Worker %d concluído\n", id)
}

func sampleWitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Adiciona uma goroutine ao WaitGroup
		go worker(i, &wg)
	}

	wg.Wait() // Espera todas as goroutines adicionadas ao WaitGroup terminarem
	fmt.Println("Todas as goroutines concluídas")
}

/*
Mutexes e Mutexes Rápidos (sync.Mutex e sync.RWMutex): São usados para controlar o acesso concorrente a recursos compartilhados,
permitindo que apenas uma go routine por vez acesse o recurso protegido. O sync.Mutex é usado para bloquear e desbloquear acesso
a um recurso, enquanto o sync.RWMutex permite bloqueio exclusivo (escrita) e bloqueio compartilhado (leitura). */

var (
	counter int
	mutex   sync.Mutex
)

func increment() {

	mutex.Lock()         // Bloqueia o mutex antes de acessar a variável compartilhada
	defer mutex.Unlock() // Garante que o mutex seja desbloqueado mesmo se ocorrer um pânico
	counter++
}

// podemos proteger structs, maps, slices ou funcões
func sampleMutex() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Valor final do contador:", counter)
}

/* Atomic Operations (atomic): O pacote atomic fornece operações atômicas que podem ser usadas para manipular variáveis compartilhadas
de forma segura entre go routines. Isso é útil para operações simples como incremento e decremento sem a necessidade de mutexes.

Pool de Go routines (goroutine pool): Às vezes, é útil limitar o número de go routines concorrentes para evitar problemas de consumo
excessivo de recursos. Isso pode ser implementado manualmente ou usando bibliotecas como golang.org/x/sync/semaphore.
*/
