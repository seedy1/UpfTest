package packages

type YouTubeVideo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Views       int    `json:"views"`
	Comments    int    `json:"comments"`
	Likes       int    `json:"likes"`
	Dislikes    int    `json:"dislikes"`
	Timestamp   int64  `json:"timestamp"`
	PostID      string `json:"post_id"`
}

type Pin struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Links       string `json:"links"`
	Likes       int    `json:"likes"`
	Comments    int    `json:"comments"`
	Saves       int    `json:"saves"`
	Repins      int    `json:"repins"`
	Timestamp   int64  `json:"timestamp"`
	PostID      string `json:"post_id"`
}

type TiktokVideo struct {
	Comments          int64    `json:"comments"`
	ID                int64    `json:"id"`
	Likes             int      `json:"likes"`
	Link              string   `json:"link"`
	MentionedProfiles []string `json:"mentioned_profiles"`
	Name              string   `json:"name"`
	Plays             int64    `json:"plays"`
	PostID            string   `json:"post_id"`
	Shares            int64    `json:"shares"`
	Tags              []string `json:"tags"`
	ThumbnailURL      string   `json:"thumbnail_url"`
	Timestamp         int64    `json:"timestamp"`
}

type InstagramMedia struct {
	Comments       int64    `json:"comments"`
	HiddenLikes    bool     `json:"hidden_likes"`
	ID             int64    `json:"id"`
	Likes          int      `json:"likes"`
	Link           string   `json:"link"`
	LocationName   string   `json:"location_name"`
	Mtype          int64    `json:"mtype"`
	Plays          int64    `json:"plays"`
	PostID         string   `json:"post_id"`
	TaggedProfiles []string `json:"tagged_profiles"`
	Text           string   `json:"text"`
	ThumbnailURL   string   `json:"thumbnail_url"`
	Timestamp      int64    `json:"timestamp"`
	Type           string   `json:"type"`
	Views          int64    `json:"views"`
}

type Article struct {
	Content     string `json:"content"`
	Description string `json:"description"`
	ID          int64  `json:"id"`
	Timestamp   int64  `json:"timestamp"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

type TwitchStream struct {
	AvgViewers int64 `json:"avg_viewers"`
	Chapters   []struct {
		AvgViewers     int64  `json:"avg_viewers"`
		EndTimestamp   int64  `json:"end_timestamp"`
		Game           string `json:"game"`
		PeakViewers    int64  `json:"peak_viewers"`
		StartTimestamp int64  `json:"start_timestamp"`
		Title          string `json:"title"`
	} `json:"chapters"`
	Duration     int64  `json:"duration"`
	ID           int64  `json:"id"`
	Language     string `json:"language"`
	PeakViewers  int64  `json:"peak_viewers"`
	StreamID     string `json:"stream_id"`
	ThumbnailURL string `json:"thumbnail_url"`
	Timestamp    int64  `json:"timestamp"`
	Title        string `json:"title"`
}

type FacebookStatus struct {
	Comments  int64  `json:"comments"`
	ID        int64  `json:"id"`
	Likes     int    `json:"likes"`
	PostID    string `json:"post_id"`
	PostType  int64  `json:"post_type"`
	Shares    int64  `json:"shares"`
	Text      string `json:"text"`
	Thumbnail string `json:"thumbnail"`
	Timestamp int64  `json:"timestamp"`
	Views     int64  `json:"views"`
}

type Tweet struct {
	Comments  int64  `json:"comments"`
	Content   string `json:"content"`
	Favorites int64  `json:"favorites"`
	ID        int64  `json:"id"`
	IsRetweet bool   `json:"is_retweet"`
	PostID    string `json:"post_id"`
	Retweets  int64  `json:"retweets"`
	Timestamp int64  `json:"timestamp"`
}

type Data struct {
	YouTubeVideo   *YouTubeVideo   `json:"youtube_video,omitempty"`
	Pin            *Pin            `json:"pin,omitempty"`
	Article        *Article        `json:"article,omitempty"`
	InstagramMedia *InstagramMedia `json:"instagram_media,omitempty"`
	TiktokVideo    *TiktokVideo    `json:"tiktok_video,omitempty"`
	TwitchStream   *TwitchStream   `json:"twitch_stream,omitempty"`
	FacebookStatus *FacebookStatus `json:"facebook_status,omitempty"`
	Tweet          *Tweet          `json:"tweet,omitempty"`
}

type AggregationResult struct {
	AverageLikes     *int `json:"average_likes,omitempty"`
	AverageComments  *int `json:"average_comments,omitempty"`
	AverageRetweets  *int `json:"average_retweets,omitempty"`
	AverageFavorites *int `json:"average_favorites,omitempty"`

	TotalPost        int   `json:"total_post"`
	MinimumTimestamp int64 `json:"minimum_timestamp"`
	MaximumTimestamp int64 `json:"maximum_timestamp"`
}

type Config struct {
	Port string `json:"port,omitempty"`
}
