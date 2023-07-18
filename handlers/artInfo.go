package Artist

import (
	Structs "groupie/structs"
	"net/http"
	"strconv"
	"text/template"
)

func GroupHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	myid := r.URL.Query().Get("id")
	intId, err := strconv.Atoi(myid)
	if err != nil || intId > 52 || intId < 1 {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	ResponseOneArist, code := Structs.FuncResponseOneArtist(myid)
	ResponseRelation, code2 := Structs.FuncResponseRelation(myid)
	if code != http.StatusOK || code2 != http.StatusOK {
		http.Error(w, "Internal Server Error", code)
		return
	}
	allStrct := Structs.AllStrct{
		ResponseOneArist: ResponseOneArist,
		ResponseRelation: ResponseRelation,
	}
	tmp, err := template.ParseFiles("ui/html/infoArtist.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmp.ExecuteTemplate(w, "infoArtist.html", allStrct)
	if err != nil {
		http.Error(w, "Internal Server Error", code)
		return
	}
}
