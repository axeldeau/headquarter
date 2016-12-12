package web

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	srv *echo.Echo
}

// NewConfig load the default and minimal configurations
func NewServer() (s *Server) {
	s = &Server{srv: echo.New()}
	s.srv.Use(middleware.Logger())
	s.srv.Use(middleware.Recover())
	return s
}

func (s Server) Start() error {
	port := os.Getenv("PORT")
	ip := os.Getenv("IP")
	return s.srv.Start(ip + ":" + port)
}

func (s *Server) AddRoute(path string, h echo.HandlerFunc, method string) {
	switch strings.ToUpper(method) {
	case "POST":
		s.srv.POST(path, h)
	case "GET":
		s.srv.GET(path, h)
	case "PUT":
		s.srv.PUT(path, h)
	case "DELETE":
		s.srv.DELETE(path, h)
	default:
		s.srv.GET(path, h)
	}
}
