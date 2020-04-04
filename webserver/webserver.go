package webserver

import (
	"fmt"

	"github.com/cloudflare/cfssl/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	api "github.com/tonradar/ton-api/proto"
	"github.com/tonradar/ton-dice-web-server/bets"
	"github.com/tonradar/ton-dice-web-server/config"
	"google.golang.org/grpc"
)

var (
	tonApiHost string
	tonApiPort string
)

type Webserver struct {
	router      *gin.Engine
	betsService *bets.BetsService
	apiClient   api.TonApiClient
	config      *config.TonWebServerConfig
}

func NewWebserver(s *bets.BetsService, cfg *config.TonWebServerConfig) *Webserver {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", cfg.TonAPIHost, cfg.TonAPIPort), opts...)
	if err != nil {
		log.Fatalf("error dial: %v", err)
	}

	apiClient := api.NewTonApiClient(conn)

	r := gin.Default()

	return &Webserver{
		router:      r,
		betsService: s,
		apiClient:   apiClient,
		config:      cfg,
	}
}

func (w *Webserver) Start() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{w.config.WebDomain}
	w.router.Use(cors.New(config))
	//w.router.Use(cors.Default())

	InitializeRoutes(w)
	listenAddr := fmt.Sprintf("0.0.0.0:%d", w.config.WebListenPort)
	w.router.Run(listenAddr)
}
