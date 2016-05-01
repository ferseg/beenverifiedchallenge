// Author: Fernando Segovia Salgado
// Creation: 04/30/2016
// Last Update: 05/01/2016
// Description: Manages the communication with the database
// file: dataAccess.go
// package: service
// dependencies:
//      - "database/sql"
//      - "github.com/ferseg/beenverifiedchallenge/utils"
//      - "github.com/mattn/go-sqlite3"
package service

import (
    "database/sql"
    "github.com/ferseg/beenverifiedchallenge/utils"
    _ "github.com/mattn/go-sqlite3"
)

// Struct used to store the data of the songs
// returned by the query
type Song struct {
    Name string `json:"name"`
    Artist string `json:"artist"`
    Length int `json:"length"`
    Genre string `json:"genre"`
}

type read func(*sql.Rows) string


// Constants for the open process
const (
    DB_ENGINE = "sqlite3"
    BD_PATH = "./service/jrdd.db"
)

// Creates the connection with the database
func createConnection() (*sql.DB) {
    conn, err := sql.Open(DB_ENGINE, BD_PATH)
    utils.CheckError(err)
    return conn
}

// Executes a select query
// Returns the read rows
func executeSelectQuery(query string, conn *sql.DB, fn read) (string) {
    resultingRows, err := conn.Query(query)
    utils.CheckError(err)
    result := fn(resultingRows)
    defer resultingRows.Close()
    return result
}

// Determines the way in which the selected rows will be read
// for the createSearchSongQuery
func readSelectedSongs(resultingRows *sql.Rows) (string) {
    songs := make([]Song, 0)

    // Iterate over the resulting rows
    for resultingRows.Next() {
        var song string
        var artist string
        var length int
        var genre string
        err := resultingRows.Scan(&song, &artist, &genre, &length)
        utils.CheckError(err)
        currentSong := Song{song, artist, length, genre}
        songs = append(songs, currentSong)

    }

    return string(utils.ConvertToJSON(songs))
}

// Determines the way in which the selected rows will be read
// for the createSearchByLengthQuery
func readSelectedByLength (resultingRows *sql.Rows) (string) {
    songs := make([]Song, 0)

    // Iterate over the resulting rows
    for resultingRows.Next() {
        var song string
        var artist string
        var length int
        var genre string
        err := resultingRows.Scan(&song, &artist, &genre, &length)
        utils.CheckError(err)
        currentSong := Song{song, artist, length, genre}
        songs = append(songs, currentSong)

    }

    return string(utils.ConvertToJSON(songs))
}

// Searches a song by a criteria
func searchSongs(criteria string) (string) {
    conn := createConnection()

    // Gets the sql query from the queryManager.go
    query := createSearchSongQuery(criteria)

    // Excecutes the query
    result := executeSelectQuery(query, conn, readSelectedSongs)
    
    // Closes 
    conn.Close()
    
    return result
}


func searchSongByLength(min, max string) (string) {
    conn := createConnection()

    // Gets the sql query from the queryManager.go
    query := createSearchByLengthQuery(min, max)

    // Excecutes the query
    result := executeSelectQuery(query, conn, readSelectedSongs)
    
    // Closes 
    conn.Close()
    
    return result
}
