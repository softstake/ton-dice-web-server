package webserver

import (
	"github.com/gin-gonic/gin"
	"github.com/tonradar/ton-dice-web-server/bets"
)

type Webserver struct {
	router     *gin.Engine
	betService *bets.BetService
}

func NewWebserver(s *bets.BetService) *Webserver {
	r := gin.Default()

	return &Webserver{
		router:     r,
		betService: s,
	}
}

func (w *Webserver) Start() {
	InitializeRoutes(w)
	w.router.Run()
}
