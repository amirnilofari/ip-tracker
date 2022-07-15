package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

// traceCmd represents the trace command
var traceCmd = &cobra.Command{
	Use:   "trace",
	Short: "Trace IP",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				showData(ip)
			}
		} else {
			fmt.Println("Please provide IP to trace")
		}
	},
}

func init() {
	rootCmd.AddCommand(traceCmd)
}

type Ip struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Timezone string `json:"timezone"`
	Postal   string `json:"postal"`
}

func showData(ip string) {
	url := "https://ipinfo.io/" + ip + "/geo"
	responseByte := getData(url)

	data := Ip{}

	err := json.Unmarshal(responseByte, &data)
	if err != nil {
		log.Println("Something wrong to Unmarshal json the response")
	}
	c := color.New(color.FgCyan).Add(color.Underline).Add(color.Italic).Add(color.Bold)
	c.Println("Data found:")
	fmt.Printf("IP: %s\nCity: %s\nRegion: %s\nCountry: %s\nLocation: %s\nTimezone: %s\nPostal: %s\n",
		data.IP, data.City, data.Region, data.Country, data.Location, data.Timezone, data.Postal)
}

func getData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Println("Something wrong to get the response")
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Something wrong to read the response")
	}

	return responseByte
}
