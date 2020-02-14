package webserver

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/tonradar/ton-dice-web-server/storage"
)

func (w *Webserver) GetAllBets(c *gin.Context) {
	resp, err := w.betService.Store.GetAllBets(context.Background(), storage.GetAllBetsReq{})
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, resp)
}
