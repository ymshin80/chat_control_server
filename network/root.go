package network

import (
	"chat_golang_control/service"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type Server struct{
	engine *gin.Engine

	service *service.Service
	
	port string
	

}

// type Network struct{
// 	engine *gin.Engine
// 	service *service.Service
// 	repository *repository.Repository
	
// 	port string
// 	ip string
// }

func NewNetwork(service *service.Service,  port string ) *Server {
	s := &Server{engine: gin.New(), service: service, port: port}

	///////////////////////middleware 설정///////////////////////////////
	//default -- log4J
	s.engine.Use(gin.Logger())
	s.engine.Use(gin.Recovery())
	//cross site 설정
	s.engine.Use(cors.New(cors.Config{
		AllowWebSockets: true,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET","POST","PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"ORIGIN","Content-Type", "Content-Length", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		ExposeHeaders: []string{"ORIGIN","Content-Type", "Content-Length", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "expires"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	registerTowerAPI(s)
	return s
}


func (s *Server) Start() error {
	return s.engine.Run(s.port)
}