package jokeprovider

import (
	"github.com/spore2102/joker/internal/config"
	"github.com/spore2102/joker/internal/types"
)

type jokeProvider struct {
	chuckJokeProvider ChuckJokeProvider
	dadJokeProvider   DadJokeProvider
}

type JokeProvider interface {
	GetDadJoke() (string, error)
	GetChuckJoke() (string, error)
}

func InitJokeProvider(cfg config.JokesApiConfig, jokeType types.JokeType) JokeProvider {
	return &jokeProvider{
		chuckJokeProvider: initChuckJokeProvider(cfg.ChuckApiConfig),
		dadJokeProvider:   initDadJokeProvider(cfg.DadApiConfig),
	}
}

func (provider *jokeProvider) GetChuckJoke() (string, error) {
	joke, err := provider.chuckJokeProvider.GetJoke()

	if err != nil {
		return "", err
	}

	return joke, nil
}

func (provider *jokeProvider) GetDadJoke() (string, error) {
	joke, err := provider.dadJokeProvider.GetJoke()

	if err != nil {
		return "", err
	}

	return joke, nil
}
