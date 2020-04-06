package response

// OneMeme : Response for a single Meme
type OneMeme struct {
	PostLink  string `json:"postLink"`
	Subreddit string `json:"subreddit"`
	Title     string `json:"title"`
	URL       string `json:"url"`
}
