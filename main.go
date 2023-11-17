package main

import (
	"net/http"

	"github.com/dinozor-io/consts/methods"
	"github.com/dinozor-io/gineng"
	"github.com/dinozor-io/interfaces"
	"github.com/dinozor-io/server"
)

func main() {

	engine := &gineng.Engine{}

	router := &gineng.Router{}

	router.AddRoute(methods.GET, "/say-hello", func(c interfaces.Controller) {
		c.Res().JSON(http.StatusOK, map[string]any{
			"msg": "Hello Guys",
		})
	})

	server := server.New(engine, router)

	server.Serve()
}
