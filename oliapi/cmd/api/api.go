package main

import "oliapi/rest"

func main() {
	server := rest.NewRestServer()
	// server.Migrate()
	server.Populate()
	server.Start()
}
