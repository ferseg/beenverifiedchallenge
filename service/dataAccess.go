package service

import (
	"database/sql"
	"github.com/ferseg/beenverifiedchallenge/utils"
	_ "github.com/mattn/go-sqlite3"
)

type Song struct {
	name string
	artist string
	length int
	genre string
}

const (
	DB_ENGINE = "sqlite3"
	BD_PATH = "./service/jrdd.db"
)

func searchSongs(criteria string) (string) {
	db, err := sql.Open(DB_ENGINE, BD_PATH)
	utils.CheckError(err)

	query := songQuery(criteria)

    resultingRows, err := db.Query(query)
    utils.CheckError(err)

    result := ""

    for resultingRows.Next() {
    	var song string
    	var artist string
    	var length int
    	var genre string
    	err = resultingRows.Scan(&song, &artist, &genre, &length)
        utils.CheckError(err)
       	result += song + ", " + genre + "\n"

    }
    defer resultingRows.Close()
    db.Close()

	return result
}
