package server

import (
	"github.com/gin-gonic/gin"
)

type ServerStruct struct {
	R *gin.Engine
}

func (s *ServerStruct) StartServer() {
	s.R.Run("localhost:8080")
}

func NewHTTPServer() *ServerStruct {
	router := gin.Default()

	return &ServerStruct{
		R: router,
	}
}
