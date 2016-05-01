package service

import (
    "fmt"
    "net/http"

    "goji.io"
    "goji.io/pat"
    "golang.org/x/net/context"
)

// Searches a song in the database and returns the information in JSON format
func searchSong(ctx context.Context, w http.ResponseWriter, r *http.Request) {
        criteria := pat.Param(ctx, "searchCriteria")
        result := searchSongs(criteria)
        fmt.Fprintf(w, result)
}

/*
 * Initializes the server
 */
func InitServer() {
    mux := goji.NewMux()
    mux.HandleFuncC(pat.Get("/song/:searchCriteria"), searchSong)
    http.ListenAndServe("localhost:8000", mux)
}