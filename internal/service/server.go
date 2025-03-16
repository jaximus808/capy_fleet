package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jaximus808/capy_websocket/internal/service/multiplayer"
	"github.com/jaximus808/capy_websocket/internal/service/routes"
)

func CreateServer() {
	r := gin.Default()

	routes.CreateRoutes(r)

	r.LoadHTMLGlob("client/pages/*")
	r.Static("/static", "./client/public")

	//run multi server
	multiplayer.SpinUpMultiPlayerGame()

	r.Run(":3000")
}
