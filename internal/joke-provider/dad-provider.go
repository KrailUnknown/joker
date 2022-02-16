package jokeprovider

import (
	"io/ioutil"
	"net/http"

	"github.com/spore2102/joker/internal/config"
	"github.com/spore2102/joker/internal/utils"
)

type dadJokeProvider struct {
	apiUrl string
}

type DadJokeProvider interface {
	GetJoke() (string, error)
}

func initDadJokeProvider(cfg config.DadJokesConfig) DadJokeProvider {
	return &dadJokeProvider{
		apiUrl: cfg.Url,
	}
}

func (provider *dadJokeProvider) GetJoke() (string, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", provider.apiUrl, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	joke, err := utils.GetByKeyFromJson(body, "joke")

	if err != nil {
		return "", err
	}

	return joke, nil
}
