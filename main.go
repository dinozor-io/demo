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

	all := &server.Group{}

	all.Cond(func(c interfaces.Controller) bool {
		return true
	})

	router.AddRoute(methods.GET, "/say-hello", func(c interfaces.Controller) {
		c.Res().JSON(http.StatusOK, map[string]any{
			"msg": "Hello Guys",
		})
	}, all)

	server := server.New(engine, router)

	server.Serve()
}
