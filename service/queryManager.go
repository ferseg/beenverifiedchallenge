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
