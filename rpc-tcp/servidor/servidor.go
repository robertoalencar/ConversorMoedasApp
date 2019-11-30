package main

import (
	"ConversorMoedasApp/impl"
	"ConversorMoedasApp/util"
	"fmt"
	"net"
	"net/rpc"
	"strconv"
)

func servidorRPCTCP() {

	// cria instância do conversor
	conversor := new(impl.ConversorRPC)

	// cria um novo servidor rpc e registra o conversor
	server := rpc.NewServer()
	server.RegisterName("Conversor", conversor)

	// cria um listen rpc-sender
	l, err := net.Listen("tcp", ":"+strconv.Itoa(util.SERVER_PORT))
	util.ChecaErro(err, "Servidor não está pronto")

	// aguarda por chamadas
	fmt.Println("Servidor pronto (RPC TCP) ...")
	server.Accept(l)
}

func main() {

	go servidorRPCTCP()

	fmt.Scanln()
}
