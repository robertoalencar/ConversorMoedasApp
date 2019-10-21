package main

import (
	"ConversorMoedasApp/util"
	"fmt"
	"os"
	"strconv"
	"time"

	"ConversorMoedasApp/grpc/conversor"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var i int32

	// Estabelece conexão com o servidor
	conn, err := grpc.Dial("localhost"+":"+strconv.Itoa(util.PORTA), grpc.WithInsecure())
	util.ChecaErro(err, "Não foi possível se conectar ao servidor")
	defer conn.Close()

	conv := conversor.NewGreeterClient(conn)

	// contacta o servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	nomeArquivo := time.Now().Format("2006-01-02 15:04:05") + " - Avaliação de Desempenho.txt"
	arquivo, _ := os.OpenFile(nomeArquivo, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	var moedaDestino string = "dolar"
	var valor float32 = 10.0

	for i = 0; i < util.TAMANHO_AMOSTRA; i++ {

		t1 := time.Now()

		// invoca operação remota
		resultadoConversao, _ := conv.Converter(ctx, &conversor.Request{MoedaDestino: moedaDestino, Valor: valor})
		fmt.Println("Operação de coversão de", valor, "reais em", moedaDestino, ",", resultadoConversao)

		t2 := time.Now()
		x := t2.Sub(t1)
		arquivo.WriteString(strconv.FormatInt(x.Microseconds(), 10) + "\n")
	}

	arquivo.Close()
}
