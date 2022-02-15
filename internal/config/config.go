package config

type Config struct {
	JokesApiConfig JokesApiConfig `mapstructure:"jokes_api"`
}

type JokesApiConfig struct {
	DadApiConfig   DadJokesConfig   `mapstructure:"dad"`
	ChuckApiConfig ChuckJokesConfig `mapstructure:"chuck"`
}

type ApiConfig struct {
	Url string
}

type DadJokesConfig ApiConfig
type ChuckJokesConfig ApiConfig
