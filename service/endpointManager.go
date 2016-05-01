package service

import (
	"fmt"
   	"net/http"

	"goji.io"
    "goji.io/pat"
    "golang.org/x/net/context"
)

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
        criteria := pat.Param(ctx, "searchCriteria")
        result := searchSongs(criteria)
        fmt.Fprintf(w, "Hello, %s!", result)
}

/*
 * Initializes the server
 */
func InitServer() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/song/:searchCriteria"), hello)
	http.ListenAndServe("localhost:8000", mux)
}