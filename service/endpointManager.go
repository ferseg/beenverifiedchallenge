// Author: Fernando Segovia Salgado
// Creation: 04/30/2016
// Last Update: 05/01/2016
// Description: Manages the endpoint requests
// file: endpointManager.go
// package: service
// dependencies:
//      - "fmt"
//      - "net/http"
//      - "goji.io"
//      - "goji.io/pat"
//      - "golang.org/x/net/context"
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

func searchSongsByLength(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    min := pat.Param(ctx, "min")
    max := pat.Param(ctx, "max")
    result := searchSongByLength(min, max)
    fmt.Fprintf(w, result)
}

/*
 * Initializes the server
 */
func InitServer() {
    mux := goji.NewMux()
    mux.HandleFuncC(pat.Get("/song/:searchCriteria"), searchSong)
    mux.HandleFuncC(pat.Get("/song/:min/:max"), searchSongsByLength)
    http.ListenAndServe("localhost:8000", mux)
}
