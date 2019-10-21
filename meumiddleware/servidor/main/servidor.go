package main

import (
	"Plataformas-Exercicio03/ConversorMoedasApp/meumiddleware/cliente/proxies"
	"Plataformas-Exercicio03/ConversorMoedasApp/meumiddleware/servidor/invoker"
	"Plataformas-Exercicio03/RobertoMiddleware/services/naming/proxy"
	"fmt"
)

func main() {

	// create a built-in proxy of naming service
	namingProxy := proxy.NamingProxy{}

	// create a proxy of conversor service
	conversor := proxies.NewConversorProxy()

	// register service in the naming service
	namingProxy.Register("Conversor", conversor)

	// control loop passed to middleware
	fmt.Println("Conversor Server running!!")
	conversorInvoker := invoker.NewConversorInvoker()

	go conversorInvoker.Invoke()

	fmt.Scanln()

}
