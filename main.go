package main

import (
	"log"
	"toodlist/routes"
)

func main() {
	r := routes.SetupRouter()

	err := r.Run(":8080")

	if err != nil {
		log.Fatal("Error al iniciar el server")
	}
}
