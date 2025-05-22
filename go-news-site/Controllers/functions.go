package Controllers

import (
	"fmt"
	"net/http"
	"text/template"
	 database"News/Models"
	"github.com/julienschmidt/httprouter"
	"io"
	"os"
)

func Home(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view,_ := template.ParseFiles("Views/news.html")
	data := make(map[string]interface{})
	data["Posts"] = database.Info{}.GetAll()
	data["Picturea"] = database.Info{}.Get("1")
	view.Execute(w,data)
}

func Add(w http.ResponseWriter,r *http.Request, params httprouter.Params){
	view,_ := template.ParseFiles("Views/add.html")
	view.Execute(w,nil) 
}

func Ekle(w http.ResponseWriter,r *http.Request,params httprouter.Params){
	title := r.FormValue("news-title")
	information := r.FormValue("news-information")
	categories := r.FormValue("news-categories")
	r.ParseMultipartForm(10 << 20)
	file,header,err := r.FormFile("news-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f,err := os.OpenFile("Views/img/"+header.Filename,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_,err = io.Copy(f,file)
	// Upload End
	if err != nil {
		fmt.Println(err)
		return
	}

	database.Info{
		Title: title,
		Information: information,
		Categories: categories,
		Picture_url: "/Views/img/" + header.Filename,
	}.Add()

	http.Redirect(w,r,"/",http.StatusSeeOther)

}