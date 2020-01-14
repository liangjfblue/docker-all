package server

import (
	"link-gin-db/internal/models"
	"log"
	"net/http"

	"link-gin-db/config"

	"github.com/gin-gonic/gin"
)

type IServer interface {
	Init()
	Start()
	Stop()
}

type Server struct {
	Config *config.Config
	G      *gin.Engine
}

func NewServer(config *config.Config, g *gin.Engine) *Server {
	return &Server{
		Config: config,
		G:      g,
	}
}

func (s *Server) Init() {
	models.Init(s.Config.MysqlConf)
}

func (s *Server) Start() {
	log.Println("server start...")
	if err := http.ListenAndServe(s.Config.HTTPConf.Addr, s.G); err != nil {
		log.Println("server start err :", err.Error())
		return
	}
}

func (s *Server) Stop() {
	log.Println("server stop...")
}
