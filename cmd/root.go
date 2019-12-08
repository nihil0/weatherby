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
		apiKey := os.Getenv("WEATHER_APP_ID")
		url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s", place, country, apiKey)
		fmt.Println(url)
		response, _ := http.Get(url)

		if response.StatusCode != 200 {
			fmt.Println("Sorry, couldn't find that city!")
			os.Exit(1)
		}

		data := models.Main{}
		body, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(body, &data)

		fmt.Printf(
			"\nWeather in %s\n\n"+
				"Temperature: %.2f °C\n"+
				"High: %.2f °C\n"+
				"Low: %.2f °C\n\n"+
				"", place, data.Main.Temperature-273, data.Main.High-273, data.Main.Low-273)

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
	showTemp.Flags().StringP("country", "c", "", "Country in which the city is located. (Optional)")
	showTemp.MarkFlagRequired("place")
	showTemp.MarkFlagRequired("country")

	rootCmd.AddCommand(showTemp)
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
