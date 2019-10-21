package main

import (
	"ConversorMoedasApp/util"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
	"time"
)

func clienteRPCTCP() {

	var i int32
	var reply int

	// conectar ao servidor
	client, err := rpc.Dial("tcp", ":"+strconv.Itoa(util.PORTA))
	util.ChecaErro(err, "Servidor não está pronto")

	defer client.Close()

	nomeArquivo := time.Now().Format("2006-01-02 15:04:05") + " - Avaliação de Desempenho.txt"
	arquivo, _ := os.OpenFile(nomeArquivo, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	var moedaDestino string = "dolar"
	var valor float64 = 10.0

	for i = 0; i < util.TAMANHO_AMOSTRA; i++ {

		t1 := time.Now()

		// invoca operação remota
		args := util.Args{MoedaDestino: moedaDestino, Valor: valor}

		client.Call("Conversor.Converter", args, &reply)
		fmt.Println("Coversão n : ", i)

		t2 := time.Now()
		x := t2.Sub(t1)
		arquivo.WriteString(strconv.FormatInt(x.Microseconds(), 10) + "\n")
	}

	arquivo.Close()
}

func main() {

	go clienteRPCTCP()

	fmt.Scanln()
}
