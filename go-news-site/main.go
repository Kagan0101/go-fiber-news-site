package main

import (
	Admin_routers "News/Config"
	database "News/Models"
	"net/http"
)

func main() {
	database.Info{}.Migrate()
	http.ListenAndServe(":8000",Admin_routers.Routers())
}