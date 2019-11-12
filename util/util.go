package util

import (
	"log"
)

const QTD_EXECUCOES_EXPERIMENTO = 10
const SERVER_HOST = "127.0.0.1"
const SERVER_PORT = 12345
const MIOP_REQUEST = 1

const ID_COMPONENTE_REQUESTOR = "requestor"
const ID_MARSHALLER = "marshaller"

type Request struct {
	MoedaDestino string
	Valor        float64
}

type Reply struct {
	Result float64
}

type Args struct {
	MoedaDestino string
	Valor        float64
}

func ChecaErro(err error, msg string) {

	if err != nil {
		log.Fatalf("%s!!: %s", msg, err)
	}
}
