package server

import (
	"sinarmas-test/internal/domain/auth"
	"sinarmas-test/internal/domain/items"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	authH  auth.Handler
	itemH  items.Handler
}

func New(authH auth.Handler, itemH items.Handler) Server {
	s := Server{
		authH: authH,
		itemH: itemH,
	}

	s.router = gin.Default()

	s.RegisterHandler()

	return s
}

func (s *Server) RegisterHandler() {
	s.router.POST("/login", s.authH.Login)

	s.router.POST("/items", s.itemH.CreateItem)
	s.router.GET("/items", s.itemH.GetAllItems)
}

func (s *Server) Start() error {
	return s.router.Run(":8080")
}
