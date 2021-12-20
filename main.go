//go:generate swagger generate spec
package main

import (
	"os"
	"rbarrero/visago/routes"
)

func main() {

	var host string = os.Getenv("HOST")

	r := routes.SetupRoutes()
	r.Run(host)
}
