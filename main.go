package main

import (
	"flag"
	"fmt"
	Artist "groupie/handlers"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	srv := http.Server{
		Addr:    *addr,
		Handler: Artist.Routes(),
	}
	fmt.Printf("http://localhost%s", *addr)
	fmt.Println()
	srv.ListenAndServe()
}
