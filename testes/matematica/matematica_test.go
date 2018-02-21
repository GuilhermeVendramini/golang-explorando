/*
	Para rodar um test execute:

	Função específica:
	go test -run NameOfTest

	Todo o arquivo:
	go test foo_test.go

	Rodar tests da pasta corrente e subdiretórios
	go test -run ./..
*/

package matematica

import (
	"testing"
)

const ErrMessage = "Valo esperado %v, mas o resultado foi %v."

/*
	Por padrão um teste começa com "Test"
*/
func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(ErrMessage, valorEsperado, valor)
	}
}
