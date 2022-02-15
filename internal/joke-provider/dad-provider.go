package jokeprovider

import (
	"io/ioutil"
	"net/http"

	"github.com/spore2102/joker/internal/config"
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
	resp, err := http.Get(provider.apiUrl)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
