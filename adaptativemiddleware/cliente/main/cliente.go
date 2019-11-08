package main

import (
	"ConversorMoedasApp/adaptativemiddleware/cliente/proxies"
	"ConversorMoedasApp/util"
	"RAMid/services/naming/proxy"
	"fmt"
	"os"
	"strconv"
	"time"
)

func ExecuteExperiment() {

	var i int32

	namingService := proxy.NamingProxy{}
	conversor := namingService.Lookup("Conversor").(proxies.ConversorProxy)

	nomeArquivo := time.Now().Format("2006-01-02 15:04:05") + " - Avaliação de Desempenho.txt"
	arquivo, _ := os.OpenFile(nomeArquivo, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	var moedaDestino string = "dolar"
	var valor float64 = 10.0

	for i = 0; i < util.TAMANHO_AMOSTRA; i++ {

		t1 := time.Now()

		resultadoConversao := conversor.Converter(moedaDestino, valor)
		fmt.Println("Operação de coversão:", i, ", Valor: ", valor, "reais em", moedaDestino, ",", resultadoConversao)

		t2 := time.Now()
		x := t2.Sub(t1)
		arquivo.WriteString(strconv.FormatInt(x.Microseconds(), 10) + "\n")

		time.Sleep(1 * time.Millisecond)
	}

}

func main() {

	go ExecuteExperiment()
	fmt.Scanln()
}
