/*
	Neste exemplo, estou utilziando o pacote custom chamado "area"
	Leia o arquivo "README.md" para ver onde esse pacote foi criado.
*/
package main

import (
	"fmt"
	"github.com/golearning/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
}
