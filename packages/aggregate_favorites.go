package packages

// only twitter contains favorites
func getTotalFavorites(data map[int]Data) int {
	var totalFavorites int = 0

	for _, dat := range data {
		if dat.Tweet != nil && dat.Tweet.Favorites != 0 {
			totalFavorites = totalFavorites + int(dat.Tweet.Favorites)
		}
	}

	return totalFavorites
}
