package main

import (
	"fmt"
	"time"
)

func worker(workerNum int, data chan int) {
	for msg := range data {
		fmt.Printf("Worker %d: %d\n", workerNum, msg)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)

	qtWorkers := 50

	// EU PRECISO INICIAR A GO ROUTINE ANTES DE ENVIAR PARA ELE, SE NAO DA DEADLOCK.
	go worker(1, canal)
	go worker(2, canal)

	for i := 0; i < qtWorkers; i++ {
		canal <- i
	}

	close(canal) // Feche o canal depois de enviar todos os dados.

}
