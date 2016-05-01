// Author: Fernando Segovia Salgado
// Creation: 04/30/2016
// Last Update: 05/01/2016
// Description: Builds the queries
// file: queryManager.go
// package: service
// dependencies: none
package service

// Builds the query to search a song
func createSearchSongQuery(criteria string) (query string) {
    query = "SELECT Songs.song, Songs.artist, Genres.name as genre, Songs.length " +
                "FROM Songs " +
                "INNER JOIN Genres " +
                    "ON Songs.genre = Genres.id " + 
                "WHERE Songs.song LIKE '%" + criteria + "%' " +
                "OR Songs.artist LIKE '%" + criteria + "%' " +
                "OR Genres.name LIKE '%" + criteria + "%'"
    return
}

func createSearchByLengthQuery(min, max string) (query string) {
    query = "SELECT Songs.song, Songs.artist, Genres.name as genre, Songs.length " +
                "FROM Songs " +
                "INNER JOIN Genres " +
                    "ON Songs.genre = Genres.id " + 
                "WHERE Songs.length BETWEEN " + min + " AND " + max
    return
}
