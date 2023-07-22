package packages

// only facebook, youtube, instagram, tiktok and pin contain COMMENTS
func getTotalComments(data map[int]Data) int {
	var totalComments int = 0

	for _, dat := range data {
		if dat.FacebookStatus != nil {
			totalComments = totalComments + int(dat.FacebookStatus.Comments)
		} else if dat.YouTubeVideo != nil {
			totalComments = totalComments + int(dat.YouTubeVideo.Comments)
		} else if dat.InstagramMedia != nil {
			totalComments = totalComments + int(dat.InstagramMedia.Comments)
		} else if dat.TiktokVideo != nil {
			totalComments = totalComments + int(dat.TiktokVideo.Comments)
		} else if dat.Pin != nil {
			totalComments = totalComments + int(dat.Pin.Comments)
		}
	}

	return totalComments
}
