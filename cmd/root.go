package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/nihil0/weatherby/models"
	"github.com/spf13/cobra"
)

var showTemp = &cobra.Command{
	Use:   "show-temp",
	Short: "Show temperature",
	Long:  `Show some information e.g., weatherby show temp or weatherby show forecast`,
	Run: func(cmd *cobra.Command, args []string) {

		place, _ := cmd.Flags().GetString("place")
		country, _ := cmd.Flags().GetString("country")
		response, _ := http.Get("https://samples.openweathermap.org/data/2.5/weather?q=Brighton,uk&appid=9d1983b0c5fbd4fc34206018740f4c88")
		foo, _ := ioutil.ReadAll(response.Body)
		data := models.Main{}
		json.Unmarshal(foo, &data)
		fmt.Printf("I am showing you something: %v", data.Main.Temperature)
		fmt.Printf("These are my args: %v, %v", place, country)
	},
}

var rootCmd = &cobra.Command{
	Use:   "weatherby",
	Short: "Weatherby tells you the weather",
	Long:  `Weatherby is a CLI tool written in Go to get weather information from OpenWeatherMap.com`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	showTemp.Flags().StringP("place", "p", "helsinki", "City for showing temperature")
	showTemp.Flags().StringP("country", "c", "finland", "Country in which the city is located. (Optional)")

	rootCmd.AddCommand(showTemp)
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
