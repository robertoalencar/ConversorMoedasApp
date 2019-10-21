package impl

import (
	"ConversorMoedasApp/util"
)

type ConversorRPC struct{}

func (t *ConversorRPC) Converter(args *util.Args, reply *float64) error {

	switch args.MoedaDestino {

	case "dolar":
		*reply = args.Valor * 4.10
	case "euro":
		*reply = args.Valor * 4.51
	case "libra":
		*reply = args.Valor * 5.02
	}

	return nil
}
