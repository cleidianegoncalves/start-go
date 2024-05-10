package start

import "errors"

// ------------------ error err
/*
o tipo error é uma interface embutida que é comumente usada para representar erros em funções.
errors podem ser nil
*/

type error interface {
	Error() string
}

func sampleError(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("não é possível dividir por zero")
	}
	return a / b, nil
}
