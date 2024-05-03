package main

import (
	"net/http"

	"github.com/oxxi/accel-one/pkg/routers"
)

func main() {

	router := http.NewServeMux()

	routers.RegisterRouter(router)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
