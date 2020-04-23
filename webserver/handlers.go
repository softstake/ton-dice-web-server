package webserver

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	api "github.com/tonradar/ton-api/proto"
	"github.com/tonradar/ton-dice-web-server/storage"
)

func (w *Webserver) GetAllBets(c *gin.Context) {
	queryLimit := 50

	limit := c.Query("limit")
	if limit != "" {
		queryLimit := strconv.FormatInt(limit, 10)
	}

	req := storage.GetAllBetsReq{Limit: queryLimit}
	resp, err := w.betsService.Store.GetAllBets(context.Background(), req)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, resp)
}

func (w *Webserver) GetPlayerBets(c *gin.Context) {
	queryLimit := 50

	address := c.Param("address")
	if address == "" {
		c.JSON(400, "invalid address")
	}

	limit := c.Query("limit")
	if limit != "" {
		queryLimit := strconv.FormatInt(limit, 10)
	}

	req := storage.GetPlayerBetsReq{
		PlayerAddress: address,
		Limit:         queryLimit,
	}
	resp, err := w.betsService.Store.GetPlayerBets(context.Background(), req)
	fmt.Printf("err: %v", err)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, resp)
}

func (w *Webserver) GetBalance(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		c.JSON(400, "invalid address")
	}
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
