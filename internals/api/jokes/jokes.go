package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/soypita/go-workshop/internals/api"
)

const getJokePath = "/api?format=json"

// Joke client implementation
type JokeClient struct {
	url string
}

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

func (j *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := j.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Api response status: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
