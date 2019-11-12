package proxies

import (
	"ConversorMoedasApp/util"
	"RAMid/aux"
	"RAMid/distribution/clientproxy"
	"RAMid/plugins"
	"plugin"
	"reflect"
)

type ConversorProxy struct {
	Proxy clientproxy.ClientProxy
}

func NewConversorProxy() ConversorProxy {

	p := new(ConversorProxy)

	p.Proxy.TypeName = reflect.TypeOf(ConversorProxy{}).String()
	p.Proxy.Host = util.SERVER_HOST
	p.Proxy.Port = util.SERVER_PORT

	return *p
}

func (proxy ConversorProxy) Converter(moedaDestino string, valor float64) float64 {

	// prepare invocation
	params := make([]interface{}, 2)
	params[0] = moedaDestino
	params[1] = valor
	request := aux.Request{Op: "Converter", Params: params}
	inv := aux.Invocation{Host: proxy.Proxy.Host, Port: proxy.Proxy.Port, Request: request}

	//Carrega o arquivo do componente
	manager := plugins.Manager{}
	componente, err := plugin.Open(manager.ObterComponente("requestor"))
	util.ChecaErro(err, "Falha ao carregar o arquivo do componente")

	funcao, err := componente.Lookup("Invoke")
	util.ChecaErro(err, "Falha ao carregar a função do componente")

	Invoke := funcao.(func(chan interface{}))

	ch := make(chan interface{})
	go Invoke(ch)

	ch <- inv
	retorno := <-ch
	ter := retorno.([]interface{})

	return ter[0].(float64)
}
