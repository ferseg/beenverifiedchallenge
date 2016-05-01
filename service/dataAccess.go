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


// Constants for the open process
const (
	DB_ENGINE = "sqlite3"
	BD_PATH = "./service/jrdd.db"
)

// Searches a song by a criteria
func searchSongs(criteria string) (string) {
	db, err := sql.Open(DB_ENGINE, BD_PATH)
	utils.CheckError(err)

	// Gets the sql query from the queryManager.go
	query := createSearchSongQuery(criteria)

	// Excecutes the query
    resultingRows, err := db.Query(query)
    utils.CheckError(err)
    songs := make([]Song, 1)
    
    // Iterate over the resulting rows
    for resultingRows.Next() {
    	var song string
    	var artist string
    	var length int
    	var genre string
    	err = resultingRows.Scan(&song, &artist, &genre, &length)
        utils.CheckError(err)
        currentSong := Song{song, artist, length, genre}
        songs = append(songs, currentSong)

    }
    // Closes 
    defer resultingRows.Close()
    db.Close()

    // converts the map to json
    result := utils.ConvertToJSON(songs)
	return string(result)
}




