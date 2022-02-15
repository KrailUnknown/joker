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
	jokeType := types.JokeType{
		Category: "dad",
	}

	jokeProvider := jokeprovider.InitJokeProvider(cfg.JokesApi, jokeType)
	joke, err := jokeProvider.GetJoke()

	if err != nil {
		log.Fatal(err)
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
	rootCmd.PersistentFlags().BoolP("dad", "d", false, "Gives you a dad joke")
	rootCmd.PersistentFlags().BoolP("chuck", "c", false, "Gives you a Chuck Norris joke")
	rootCmd.Flags().SortFlags = false

	// jokeType := rootCmd.Flags().Lookup("dad")

	// viper.BindPFlag("j", )

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
