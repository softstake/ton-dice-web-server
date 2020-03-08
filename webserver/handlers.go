package webserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	api "github.com/tonradar/ton-api/proto"
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

func (w *Webserver) GetPlayerBets(c *gin.Context) {
	address := c.Param("address")
	req := storage.GetPlayerBetsReq{
		PlayerAddress: address,
	}
	resp, err := w.betService.Store.GetPlayerBets(context.Background(), req)
	fmt.Printf("err: %v", err)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, resp)
}

func (w *Webserver) GetBalance(c *gin.Context) {
	address := c.Param("address")
	getAccountStateRequest := &api.GetAccountStateRequest{
		AccountAddress: address,
	}
	getAccountStateResponse, err := w.apiClient.GetAccountState(c, getAccountStateRequest)
	if err != nil {
		c.JSON(500, err)
		return
	}

	balance := getAccountStateResponse.Balance
	response := map[string]interface{}{
		"balance": balance,
	}

	c.JSON(200, response)
}
