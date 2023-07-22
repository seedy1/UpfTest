package packages

// only facebook, youtube, instagram, tiktok and pin contains LIKES
func getTotalLikes(data map[int]Data) int {
	var totalLikes int = 0

	for _, dat := range data {
		if dat.FacebookStatus != nil {
			totalLikes = totalLikes + dat.FacebookStatus.Likes
		} else if dat.YouTubeVideo != nil {
			totalLikes = totalLikes + dat.YouTubeVideo.Likes
		} else if dat.InstagramMedia != nil {
			totalLikes = totalLikes + dat.InstagramMedia.Likes
		} else if dat.TiktokVideo != nil {
			totalLikes = totalLikes + dat.TiktokVideo.Likes
		} else if dat.Pin != nil {
			totalLikes = totalLikes + dat.Pin.Likes
		}
	}

	return totalLikes
}
