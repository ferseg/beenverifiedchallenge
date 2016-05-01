// Author: Fernando Segovia Salgado
// Creation: 04/30/2016
// Last Update: 05/01/2016
// Description: Builds the queries
// file: queryManager.go
// package: service
// dependencies: none
package service

// Builds the query to search all songs depending on the criteria
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

// Builds the query to search songs in a range of length
func createSearchByLengthQuery(min, max string) (query string) {
    query = "SELECT Songs.song, Songs.artist, Genres.name as genre, Songs.length " +
            "FROM Songs " +
            "INNER JOIN Genres " +
                "ON Songs.genre = Genres.id " + 
            "WHERE Songs.length BETWEEN " + min + " AND " + max + " " +
            "ORDER BY Songs.length DESC"
    return
}

// Builds the query to get the info of the genres
// For each genre it shows:
//    Name = the name of the genre
//    Quantity = The amount of songs in that gender
//    TotalLength = The sum of all song length in the genre
func createSearchGenreInfoQuery() (query string) {
    query = "SELECT Genres.name as genre, COUNT(1) as quantity, SUM(Songs.length) as totalLength " +
            "FROM Genres " +
            "INNER JOIN Songs " +
                "ON Songs.genre = Genres.id " +
            "GROUP BY Genres.name"
    return
}
