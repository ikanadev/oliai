package main

import "oliapi/rest"

func main() {
	server, err := rest.NewRestServer()
	if err != nil {
		panic(err)
	}
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
