package server

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stillsaro/Leenky/internal/database/DBgo"
)

type Server struct {
	ListenAddr string
	DB         *sql.DB
	Queries    *DBgo.Queries
	Auth       *Auth
	Hub        *Hub
}

func New(addr string, db *sql.DB, queries *DBgo.Queries, jwtKey string) *Server {
	return &Server{
		ListenAddr: addr,
		DB:         db,
		Queries:    queries,
		Auth:       NewAuth(jwtKey),
		Hub:        NewHub(),
	}
}

func (s *Server) Run() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = NewTemplate()
	e.Static("/css", "/home/saro/GeekHub/Go/Leenky/viwes/css/")
	go s.Hub.Run()

	e.GET("/", s.RootHandler)

	e.GET("/signup", s.Signup)
	e.POST("/signup", s.SignupHandler)

	e.GET("/login", s.Login)
	e.POST("/login", s.LoginHandler)

	e.GET("/leenkychat", s.LeenkyChat, s.JwtAuthMiddleware())
	e.GET("/ws/leenkychat", s.LeenkyChatWsHandler, s.JwtAuthMiddleware())

	return e.Start(s.ListenAddr)
}
