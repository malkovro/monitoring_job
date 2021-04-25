/*
Copyright © 2021 Leo Figea <figealeo@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/d2r2/go-dht"
	"github.com/spf13/cobra"
)

var BaseUrl string

type Measure struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
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

func postMeasure(measures ...Measure) (bool, error) {
	url := fmt.Sprintf("%s/measures/batch", BaseUrl)
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// climateCmd represents the climate command
var climateCmd = &cobra.Command{
	Use:   "climate",
	Short: "Measure and send Temp & Humidity datapoint",
	Long: `Measure the temperature and humidity sending a pulse to the DHT22.
	Send the datapoint up to the monitoring API based on the baseUrl passed as flag`,
	Run: func(cmd *cobra.Command, args []string) {
		temp, humidity := ReadTempAndHumidity()
		measures := []Measure{Measure{`temperature`, temp}, Measure{`humidity`, humidity}}
		success, err := postMeasure(measures...)
		check(err)
		fmt.Printf("--- Finished the Climate monitoring job (success: %t) ---", success)
	},
}

func init() {
	sendCmd.AddCommand(climateCmd)

	climateCmd.Flags().StringVarP(&BaseUrl, "baseUrl", "u", "", "Base Url of the monitoring Api")
}
