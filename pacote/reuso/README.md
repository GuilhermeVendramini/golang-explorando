# Pacote

Neste exemplo "main.go" estou utilizando um pacote custom,
nomeado "area" que está localizado dentro do seguinte diretório:

```sh
$ /home/ruiner/go/src/github.com/golearning/area
```

Este diretório contém o arquivo "area.go" com o seguinte conteúdo:

```go
package area

import (
	"math"
)

//Pi é...
const Pi = 3.1416

//Cric é...
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é...
func Rect(base, altura float64) float64 {
	return base * altura
}

//Privada
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
```

