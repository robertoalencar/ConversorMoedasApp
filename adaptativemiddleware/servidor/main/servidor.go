package main

import (
	"ConversorMoedasApp/adaptativemiddleware/cliente/proxies"
	"ConversorMoedasApp/adaptativemiddleware/servidor/invoker"
	"RAMid/services/naming/proxy"
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
