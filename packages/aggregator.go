package packages

import "github.com/seedy1/UpfTest/helpers"

// AggregationResult is the struct that will be returned by the Aggregator function
func AggResult(data map[int]Data, dimension string) AggregationResult {

	result := AggregationResult{}
	totalPost := count(data)
	if dimension == helpers.DIMENSION_LIKES {
		avgLikes := getTotalLikes(data) / totalPost
		result.AverageLikes = &avgLikes
	} else if dimension == helpers.DIMENSION_COMMENTS {
		avgComments := getTotalComments(data) / totalPost
		result.AverageComments = &avgComments
	} else if dimension == helpers.DIMENSION_FAVORITES {
		avgFavorites := getTotalFavorites(data) / totalPost
		result.AverageFavorites = &avgFavorites
	} else if dimension == helpers.DIMENSION_RETWEETS {
		avgRetweets := getTotalRetweets(data) / totalPost
		result.AverageRetweets = &avgRetweets
	}
	result.TotalPost = totalPost
	result.MinimumTimestamp, result.MaximumTimestamp = helpers.GetMinMaxTimeStamp(buildArrayOfTimeStamps(data))
	return result
}

// get total number of items in the data map
func count(data map[int]Data) int {
	var counter int
	counter = len(data)
	return counter
}

// all data types have timestamps present
func buildArrayOfTimeStamps(data map[int]Data) []int64 {

	timeStampArray := make([]int64, len(data))
	for i, dataStreams := range data {
		if dataStreams.FacebookStatus != nil {
			timeStampArray[i] = dataStreams.FacebookStatus.Timestamp
		} else if dataStreams.YouTubeVideo != nil {
			timeStampArray[i] = dataStreams.YouTubeVideo.Timestamp
		} else if dataStreams.InstagramMedia != nil {
			timeStampArray[i] = dataStreams.InstagramMedia.Timestamp
		} else if dataStreams.TiktokVideo != nil {
			timeStampArray[i] = dataStreams.TiktokVideo.Timestamp
		} else if dataStreams.Pin != nil {
			timeStampArray[i] = dataStreams.Pin.Timestamp
		} else if dataStreams.Article != nil {
			timeStampArray[i] = dataStreams.Article.Timestamp
		} else if dataStreams.TwitchStream != nil {
			timeStampArray[i] = dataStreams.TwitchStream.Timestamp
		} else if dataStreams.Tweet != nil {
			timeStampArray[i] = dataStreams.Tweet.Timestamp
		} else {
			timeStampArray[i] = 0 // set to zero if not present/null/unknown
		}
	}
	return timeStampArray
}
