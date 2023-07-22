package packages

// only twitter contains RETWEETS
func getTotalRetweets(data map[int]Data) int {

	var totalRetweets int = 0
	for _, dat := range data {
		if dat.Tweet != nil && dat.Tweet.Retweets != 0 {
			totalRetweets = totalRetweets + int(dat.Tweet.Retweets)
		}
	}

	return totalRetweets
}
