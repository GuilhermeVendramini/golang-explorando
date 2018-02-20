/*
	Exemplo de conexão em 2 servidores
	O select vai exibir o que conectar mais rápido
	Podendo em alguns casos o processamento ser simultâneo.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func conect(server string) chan string {

	wg.Add(1)
	out := make(chan string)

	go func(server string) {
		//O sleep está simulando o timeout de conexão do servidor
		time.Sleep(time.Duration(rand.Intn(1)) * time.Second)
		fmt.Println("Processou o server:", server)
		out <- server
		wg.Done()
	}(server)

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

var wg sync.WaitGroup

func main() {
	server1 := conect("Server 1")
	server2 := conect("Server 2")

	select {
	case s1 := <-server1:
		fmt.Println("Foi mais rápido:", s1)
	case s2 := <-server2:
		fmt.Println("Foi mais rápido:", s2)
	case <-time.After(1000 * time.Millisecond):
		fmt.Println("Servidores for do ar!")
	}
}
