package Artist

import (
	Structs "groupie/structs"
	"net/http"
	"text/template"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	} else if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	tmp, err := template.ParseFiles("ui/html/mainPage.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	responseArtists, code := Structs.FuncResponseArtists()
	if code != http.StatusOK {
		http.Error(w, "Internal Server Error", code)
		return
	}
	err = tmp.Execute(w, responseArtists)
	if err != nil {
		http.Error(w, "Internal Server Error", code)
		return
	}
}
