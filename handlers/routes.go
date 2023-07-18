package Artist

import "net/http"

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle)
	mux.HandleFunc("/artists/", GroupHandle)
	mux.Handle("/style/", http.StripPrefix("/style", http.FileServer(http.Dir("./ui/style"))))
	return mux
}
