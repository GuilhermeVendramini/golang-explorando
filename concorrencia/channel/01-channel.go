/*
	Channel - é a forma de comunicação entre goroutines
	Channel é um tipo
*/

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)
}
