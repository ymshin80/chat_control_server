package network

import (
	"chat_golang_control/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tower struct {
	server *Server
}

//var once sync.Once

func registerTowerAPI(server *Server) {
	// once.Do(func () {
	// })
	t := &tower{server: server}

	t.server.engine.GET("/server-list", t.serverList)
}

func (t *tower) serverList(c *gin.Context) {


	response(c, http.StatusOK, t.server.service.GetAvailableList())
}



func response(c *gin.Context, s int, res interface{}, data ...string) {
	c.JSON(s, types.NewRes(s, res, data...))
}