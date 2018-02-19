package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	/*
		Sem paralelismo
		fale("Gui", "Pq vc não fala comigo?", 3)
		fale("Sophia", "Só posso falar depois de vc!?", 1)
	*/

	// Com paralelismo
	go fale("Gui", "Hey Soh?", 10)
	go fale("Sophia", "Hey Gui!?", 10)

	fmt.Println("Fim!")
}
