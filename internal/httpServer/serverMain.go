package httpServer

import (
	"github.com/GoSeoTaxi/yandex_go_05_p2/internal/config"
	"github.com/GoSeoTaxi/yandex_go_05_p2/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func MainServer() {

	r := chi.NewRouter()
	r.Use(middleware.Compress(2, "gzip"))

	//r.With(handlers.SetCookies).Post("/api/user/register", handlers.APIJSON)

	r.Post("/api/user/register", handlers.ApiUserRegister)

	http.ListenAndServe(":"+config.PortServer, r)

}
