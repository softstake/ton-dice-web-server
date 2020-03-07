package webserver

import (
	"fmt"
	"github.com/cloudflare/cfssl/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	api "github.com/tonradar/ton-api/proto"
	"github.com/tonradar/ton-dice-web-server/bets"
	"google.golang.org/grpc"
	"os"
)

var (
	tonApiHost string
	tonApiPort string
)

type Webserver struct {
	router     *gin.Engine
	betService *bets.BetService
	apiClient  api.TonApiClient
}

func NewWebserver(s *bets.BetService) *Webserver {
	tonApiHost = os.Getenv("TON_API_HOST")
	tonApiPort = os.Getenv("TON_API_PORT")

	if tonApiHost == "" || tonApiPort == "" {
		log.Fatal("Some of required ENV vars are empty. The vars are: TON_API_HOST, TON_API_PORT")
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", tonApiHost, tonApiPort), opts...)
	if err != nil {
		log.Fatalf("error dial: %v", err)
	}

	apiClient := api.NewTonApiClient(conn)

	r := gin.Default()

	return &Webserver{
		router:     r,
		betService: s,
		apiClient:  apiClient,
	}
}

func (w *Webserver) Start() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://ton777.io"}
	w.router.Use(cors.New(config))

	//w.router.Use(cors.Default())

	InitializeRoutes(w)

	w.router.Run("0.0.0.0:9999")
}
