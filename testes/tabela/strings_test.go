package tabela

import (
	"strings"
	"testing"
)

const ErrMessage = "%s (parte: %s) - índices: esperado (%d) encontrado (%d)."

/*
	Por padrão um teste começa com "Test"
*/
func TestIndex(t *testing.T) {
	test := []struct {
		text     string
		parte    string
		esperado int
	}{
		{"Ola", "opa", -1},
		{"Gui", "i", 2},
		{"Ca", "c", 1},
	}

	for _, test := range test {
		t.Logf("Valor: %v", test)
		atual := strings.Index(test.text, test.parte)

		if atual != test.esperado {
			t.Errorf(ErrMessage, test.text, test.parte, test.esperado, atual)
		}
	}
}
