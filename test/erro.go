package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

// Fila global de pedidos
var pedidos []int

func main() {

	for i := 1; i <= 20; i++ {
		pedidos = append(pedidos, i)
	}

	fmt.Printf("Incio dos trabalhos. Total de pedidos na fila: %d \n", len(pedidos))

	// Iniciando com 5 cozinheiros (goroutines) e trabalhar em simultaneo
	for c := 1; c <= 5; c++ {
		go cozinheiro(c)
	}

	time.Sleep(2 * time.Second)

	fmt.Printf("Acabou os trablhos. Total de pedidos restantes em fila: %d\n", len(pedidos))
}

func cozinheiro(id int) {

	for {
		// fila vazia, cozinheiro parado
		if len(pedidos) == 0 {
			break
		}

		// Erro várias goroutines podem ler o mesmo indice
		pedido := pedidos[0]
		fmt.Printf("Cozinheiro %d começou a preparar o pedido %d\n", id, pedido)

		//Simula o tempo de preparação
		time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)

		// Erro remoção do pedido da fila global sem proteção de variavel
		if len(pedidos) > 0 {
			pedidos = pedidos[1:]
			fmt.Printf("Cozinheiro %d terminou e removeu o pedido %d da fila\n", id, pedido)
		}
	}
}
