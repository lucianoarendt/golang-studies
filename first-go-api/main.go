package main

import (
	"net/http"

	"github.com/luciano/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
