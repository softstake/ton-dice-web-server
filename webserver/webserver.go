package webserver

import (
	"github.com/cloudflare/cfssl/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	api "github.com/tonradar/ton-api/proto"
	"github.com/tonradar/ton-dice-web-server/bets"
	"google.golang.org/grpc"
)

type Webserver struct {
	router     *gin.Engine
	betService *bets.BetService
	apiClient  api.TonApiClient
}

func NewWebserver(s *bets.BetService) *Webserver {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial("127.0.0.1:5400", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
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
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:8081"}
	//w.router.Use(cors.New(config))

	w.router.Use(cors.Default())

	InitializeRoutes(w)

	w.router.Run()
}
