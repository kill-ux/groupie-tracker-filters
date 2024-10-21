package groupie

import (
	"net/http"
	"strings"
)

var Data = &Page{}

// Home page contain data about artits
func HandelHome(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
		return
	}
	if req.URL.Path != "/" {
		if strings.ContainsRune(req.URL.Path[1:], '/') {
			http.Redirect(res, req, "/notFound", http.StatusFound)
			return
		}
		Error(res, 404, "Oops!! Page Not Found")
		return
	}
	HandelFilter(res,req)
	RenderPage("index", res)
}
