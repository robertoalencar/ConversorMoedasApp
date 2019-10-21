package main

import (
	"ConversorMoedasApp/grpc/conversor"
	"ConversorMoedasApp/util"
	"fmt"
	"net"
	"strconv"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type servidorConversor struct{}

func (s *servidorConversor) Converter(ctx context.Context, in *conversor.Request) (*conversor.Reply, error) {

	var resultado float32

	if in.MoedaDestino == "dolar" {
		resultado = in.Valor * 4.10
	} else if in.MoedaDestino == "euro" {
		resultado = in.Valor * 4.51
	} else if in.MoedaDestino == "libra" {
		resultado = in.Valor * 5.02
	}

	return &conversor.Reply{Resultado: resultado}, nil
}

func main() {

	conn, err := net.Listen("tcp", ":"+strconv.Itoa(util.PORTA))
	util.ChecaErro(err, "Não foi possível criar o listener")

	servidor := grpc.NewServer()
	conversor.RegisterGreeterServer(servidor, &servidorConversor{})

	fmt.Println("Servidor pronto ...")

	// Register reflection service on gRPC servidor.
	reflection.Register(servidor)

	err = servidor.Serve(conn)
	util.ChecaErro(err, "Falha ao servir")

}
