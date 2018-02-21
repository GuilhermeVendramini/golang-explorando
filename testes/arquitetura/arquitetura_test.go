/*
	Teste que verifica a arquiterua da maquina
	Por exemplo:
		* Só roda em um linux,
		* Não funciona com um processador "amd64"
		* etc
*/

package arquitetura

import (
	"runtime"
	"testing"
)

func TestArquitetura(t *testing.T) {
	t.Parallel() // executa o teste em paralelo com outros testes
	if runtime.GOARCH == "amd64" {
		t.Skip("Não fuciona com processadores amd64!")
	}
	t.Fail()
}
