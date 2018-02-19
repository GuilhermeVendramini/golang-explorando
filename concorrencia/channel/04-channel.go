/*
	O buffer não trava a execução da operação. Ex.:
	Se eu definir o buffer 3, o canal vai continuar rodando até atingir 3 itens
	Após atingir os 3 itens, o duffer precisa tirar 1 item para que outro item entra em seu lugar
*/

package main

import (
	"fmt"
)

func rotina(ch chan int) {
	ch <- 1
	fmt.Println("Executou 1") // Executa
	ch <- 2
	fmt.Println("Executou 2") // Executa
	ch <- 3
	fmt.Println("Executou 3") // Executa

	/*
		Pode executar, pq um item vai ser tirado, dando espaço para o 4 item
		O 4 item será printado dependendo do time de execução
	*/
	ch <- 4
	fmt.Println("Executou 4")
	ch <- 5
	fmt.Println("Executou 5")
	ch <- 6
	fmt.Println("Executou 6")
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	fmt.Println(<-ch)
}
