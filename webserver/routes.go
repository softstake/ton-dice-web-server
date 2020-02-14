package webserver

func InitializeRoutes(w *Webserver) {
	w.router.GET("/bets", w.GetAllBets)
}
