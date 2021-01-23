package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/d2r2/go-dht"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Measure struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

func postMeasure(measures ...Measure) (bool, error) {
	url := fmt.Sprintf("%s/measures/batch", LoadEnv().MonitoringBaseUrl)
	params, err := json.Marshal(measures)
	if err != nil {
		return false, err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(params))
	if err != nil {
		return false, err
	}
	return true, nil
}

func ReadTempAndHumidity() (float32, float32) {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
	return temperature, humidity
}

func main() {
	fmt.Println("--- Starting the monitoring job ---")
	temp, humidity := ReadTempAndHumidity()
	measures := []Measure{Measure{`temperature`, temp}, Measure{`humidity`, humidity}}
	success, err := postMeasure(measures...)
	check(err)
	fmt.Printf("--- Finished the monitoring job (success: %t) ---", success)
}
