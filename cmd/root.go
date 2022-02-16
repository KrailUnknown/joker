/*
Copyright © 2022 spore2102

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spore2102/joker/internal/config"
	jokeprovider "github.com/spore2102/joker/internal/joke-provider"
	"github.com/spore2102/joker/internal/types"
)

var (
	configName = "config"
	configType = "toml"
	configPath = "./config"
)

var (
	// For flags
	isDadJoke   bool
	isChuckJoke bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "joker",
	Short: "Provides you with a random joke",
	Run: func(cmd *cobra.Command, args []string) {
		// Decoding config into a struct
		var cfg config.Config
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("Could not decode config into struct: %v", err)
		}

		// Running the command
		if err := runCmd(&cfg); err != nil {
			log.Fatal(err)
		}
	},
}

func runCmd(cfg *config.Config) error {
	jokeProvider := jokeprovider.InitJokeProvider(cfg.JokesApiConfig)

	var (
		joke string
		err  error
	)

	currentJokeType, err := getCurrentJokeType()

	if err != nil {
		return err
	}

	switch currentJokeType.GetType() {
	case types.DAD_JOKE_TYPE:
		joke, err = jokeProvider.GetDadJoke()
	case types.CHUCK_JOKE_TYPE:
		joke, err = jokeProvider.GetChuckJoke()

	}

	if err != nil {
		return err
	}

	fmt.Println(joke)

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolVarP(&isDadJoke, "dad", "d", false, "Gives you a dad joke")
	rootCmd.PersistentFlags().BoolVarP(&isChuckJoke, "chuck", "c", false, "Gives you a Chuck Norris joke")
	rootCmd.Flags().SortFlags = false

	// viper.Set("currentJokeType", jokeType)

	rootCmd.AddCommand(serveCmd)
}

func initConfig() {
	viper.AddConfigPath(configPath)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Could not read the config file: %v", err)
	}
}

// Returns current joke type from flag
// Returns error if no joke type flag is present
// or if more than one joke type flag is present
func getCurrentJokeType() (*types.JokeType, error) {
	currentJokeType := types.JokeType{}

	// If both flags are present return error
	if isDadJoke && isChuckJoke {
		err := &types.CommandError{Message: "Only one joke type flag is possible (use --help for more information)"}
		return nil, err
	} else if isDadJoke {
		currentJokeType.SetToDadType()
	} else if isChuckJoke {
		currentJokeType.SetToChuckType()
	} else {
		err := &types.CommandError{Message: "Use at least one joke type flag (use --help for more information)"}
		return nil, err
	}

	return &currentJokeType, nil
}
