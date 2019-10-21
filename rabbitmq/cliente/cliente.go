package main

import (
	"ConversorMoedasApp/util"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	var i int32

	// conecta ao servidor de mensageria
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	util.ChecaErro(err, "Não foi possível se conectar ao servidor de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	util.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	// declara as filas
	requestQueue, err := ch.QueueDeclare("request", false, false, false, false, nil)
	util.ChecaErro(err, "Não foi possível criar a fila no servidor de mensageria")

	replyQueue, err := ch.QueueDeclare("response", false, false, false, false, nil)
	util.ChecaErro(err, "Não foi possível criar a fila no servidor de mensageria")

	// cria consumidor
	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false, false, false, nil)
	util.ChecaErro(err, "Falha ao registrar o consumidor servidor de mensageria")

	nomeArquivo := time.Now().Format("2006-01-02 15:04:05") + " - Avaliação de Desempenho.txt"
	arquivo, _ := os.OpenFile(nomeArquivo, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	var moedaDestino string = "dolar"
	var valor float64 = 10.0

	for i = 0; i < util.TAMANHO_AMOSTRA; i++ {

		t1 := time.Now()

		// prepara request
		msgRequest := util.Request{MoedaDestino: moedaDestino, Valor: valor}
		msgRequestBytes, err := json.Marshal(msgRequest)
		util.ChecaErro(err, "Falha ao serializar a mensagem")

		// publica request
		err = ch.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})
		util.ChecaErro(err, "Falha ao enviar a mensagem para o servidor de mensageria")

		// recebe resposta
		resultadoConversao := <-msgsFromServer

		fmt.Println("Operação de coversão de", valor, "reais em", moedaDestino, ",", resultadoConversao.Body)

		t2 := time.Now()
		x := t2.Sub(t1)
		arquivo.WriteString(strconv.FormatInt(x.Microseconds(), 10) + "\n")
	}

	arquivo.Close()
}
