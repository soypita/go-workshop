package api

// Client interact with 3-rd party api
type Client interface {
	GetJoke() (*JokeResponse, error)
}
