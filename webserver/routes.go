package webserver

func InitializeRoutes(w *Webserver) {
	w.router.GET("/bets", w.GetAllBets)
	w.router.GET("/bets/:address", w.GetPlayerBets)
	w.router.GET("/balance/:address", w.GetBalance)
}
