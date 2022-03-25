package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random joke",
	Long:  `This command fetches a random joke from icanhazdadjoke.com API`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Could not unmarshal response - %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseAPI,
		nil,
	)
	if err != nil {
		log.Printf("Could not request a joke - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "github.com/paravkaushal/random-jokes")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Could not make a request - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body - %v", err)
	}
	return responseBytes
}
