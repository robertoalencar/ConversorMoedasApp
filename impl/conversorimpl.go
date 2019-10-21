package impl

import "ConversorMoedasApp/util"

type Conversor struct{}

func (Conversor) InvocaConversor(req util.Request) float64 {

	moedaDestino := req.MoedaDestino
	valor := req.Valor
	return Conversor{}.Converter(moedaDestino, valor)
}

func (Conversor) Converter(moedaDestino string, valor float64) float64 {

	var resultado float64

	switch moedaDestino {

	case "dolar":
		resultado = valor * 4.10
	case "euro":
		resultado = valor * 4.51
	case "libra":
		resultado = valor * 5.02
	}

	return resultado
}
