package api


import (
	"os"
	"net/http"
	//"fmt"
	"io/ioutil"
	"log"
	"github.com/joho/godotenv"
)




func InsightServiceAPI(coin_name string) string {


	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	resp, err := http.Get(os.Getenv("INSIGHT_SERVICE_API_ENDPOINT") + "/ingest/" + coin_name)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: %v", err)
	}

	//fmt.Printf("Response: %s\n", body)



	return string(body)

}


func SentimentServiceAPI(coin_name string) {

	return
}
