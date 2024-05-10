package start

import (
	"fmt"
	"time"
)

/*
Ano:
2006 - O ano completo.
06 - O ano em dois dígitos.
Mês:
01 - O número do mês (com zero à esquerda).
1 - O número do mês (sem zero à esquerda).
Jan - O nome do mês abreviado.
January - O nome do mês completo.
Dia:
02 - O dia do mês (com zero à esquerda).
2 - O dia do mês (sem zero à esquerda).
Hora:
15 - A hora (formato 24 horas, com zero à esquerda).
03 - A hora (formato 12 horas, com zero à esquerda).
3 - A hora (formato 12 horas, sem zero à esquerda).
Minutos:
04 - Os minutos (com zero à esquerda).
4 - Os minutos (sem zero à esquerda).
Segundos:
05 - Os segundos (com zero à esquerda).
5 - Os segundos (sem zero à esquerda).
AM/PM:
PM - O período da manhã ou da tarde em letras maiúsculas.
pm - O período da manhã ou da tarde em letras minúsculas.
Fuso horário:
MST - O código do fuso horário.
Zona Horária:
Z07:00 - O deslocamento da zona horária em relação ao UTC.
Z0700 - O deslocamento da zona horária em relação ao UTC sem os dois pontos.

*/

func sampleDate() {
	// Obter a hora atual
	agora := time.Now()
	fmt.Println("Hora atual:", agora)

	// Criar uma data específica
	data := time.Date(2024, time.May, 10, 12, 0, 0, 0, time.UTC)
	fmt.Println("Data específica:", data)

	// Formatar a data
	fmt.Println("Data formatada:", data.Format(time.RFC1123))

	// Adicionar 1 dia à data
	data = data.AddDate(0, 0, 1)
	fmt.Println("Data após adicionar 1 dia:", data)

	// Calcular a diferença entre duas datas
	outraData := time.Date(2024, time.May, 15, 12, 0, 0, 0, time.UTC)
	diferenca := outraData.Sub(data)
	fmt.Println("Diferença entre as datas:", diferenca)
}

const (
	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	DateTime   = "2006-01-02 15:04:05"
	DateOnly   = "2006-01-02"
	TimeOnly   = "15:04:05"
)
