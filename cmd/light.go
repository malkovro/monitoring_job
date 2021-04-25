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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var On bool

// lightCmd represents the light command
var lightCmd = &cobra.Command{
	Use:   "light",
	Short: "Turn on/off the lights",
	Long:  `Turn on or off the relay plugged onto the PI on GPIO 17`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Turn Light %v called\n", On)
		err := rpio.Open()

		if err != nil {
			panic(err)
		}

		pin := rpio.Pin(17)
		pin.Output()
		if On {
			pin.Low()
		} else {
			pin.High()
		}
	},
}

func init() {
	switchCmd.AddCommand(lightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lightCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	lightCmd.Flags().BoolVarP(&On, "on", "o", true, "Turn the light on")
}
