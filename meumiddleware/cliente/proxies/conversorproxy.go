package proxies

import (
	"Plataformas-Exercicio03/RobertoMiddleware/aux"
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/clientproxy"
	"Plataformas-Exercicio03/RobertoMiddleware/distribution/requestor"
	"Plataformas-Exercicio03/util"
	"reflect"
)

type ConversorProxy struct {
	Proxy clientproxy.ConversorProxy
}

func NewConversorProxy() ConversorProxy {

	p := new(ConversorProxy)

	p.Proxy.TypeName = reflect.TypeOf(ConversorProxy{}).String()
	p.Proxy.Host = "localhost"
	p.Proxy.Port = util.PORTA

	return *p
}

func (proxy ConversorProxy) Converter(moedaDestino string, valor float64) float64 {

	// prepare invocation
	params := make([]interface{}, 2)
	params[0] = moedaDestino
	params[1] = valor
	request := aux.Request{Op: "Converter", Params: params}
	inv := aux.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	return ter[0].(float64)
}
