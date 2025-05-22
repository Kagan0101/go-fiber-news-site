package Config

import (
	Functions "News/Controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routers() *httprouter.Router{
	r := httprouter.New()
	r.GET("/",Functions.Home)
	r.GET("/add",Functions.Add)
	r.POST("/ekle",Functions.Ekle)

	r.ServeFiles("/Views/*filepath",http.Dir("Views"))

	return r
}