package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Com as threads vamos consegui fazer o sistema trabalhar de forma paralela

// green threads
func main() {
	// Criei um canal
	requestId := make(chan int)
	concurrency := 2 // Parti desse valor posso determinar a quantidade de processos
	// quando coloco go antes isso segnifica que gera uma nova thread
	for i := 1; i <= concurrency; i++ {
		go worker(requestId, i)
	}

	for i := 0; i <= 10; i++ {
		// O worker vai rodar duas vezes
		// Ele vai passar valor para o chan da line:17
		requestId <- i
	}
}

func worker(requestId chan int, worker int) {
	for r := range requestId {

		// Dispara requisição para teste de carga na minha api
		res, err := http.Get("http://localhost:3333/user")

		if err != nil {
			log.Fatal("Erro")
		}
		// Pesquisar o que é defer no GO
		defer res.Body.Close() //Encerra todo processo depois que tudo terminar

		content, _ := ioutil.ReadAll(res.Body)

		fmt.Printf("Worker %d. Content: %s, RequestId: %d", worker, string(content), r)

		time.Sleep(time.Second * 2)

	}

}
