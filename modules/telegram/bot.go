package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	env "os"
	"time"
	"vincentcoreapi/helper"
)

func SendMessage(message *Message) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}
	response, err := http.Post(env.Getenv("TELEGRAM_URL"), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(response.Body)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send successful request. Status was %q", response.Status)
	}
	return nil
}

type InfoIP struct {
	Ip          string `json:"ip"`
	City        string `json:"city"`
	Region      string `json:"region"`
	Query       string `json:"query"`
	Country     string `json:"country"`
	RegionName  string `json:"regionName"`
	CountryCode string `json:"countryCode"`
}

func getStations(body []byte) (*InfoIP, error) {
	var s = new(InfoIP)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func SendMessageTelegram(method string, response helper.Response, request string, device string, ip string, host string) string {
	getLocation := "http://ip-api.com/json/" + ip
	responseLoc, err := http.Get(getLocation)
	if err != nil {
		panic(err)
	}
	defer responseLoc.Body.Close()
	ipLoc, _ := ioutil.ReadAll(responseLoc.Body)

	s, err := getStations([]byte(ipLoc))

	now := time.Now()
	data, _ := json.Marshal(response)
	var message = fmt.Sprintf(env.Getenv("API_TITLE") + "\n" + method + " [" + now.Format("2006-01-02") + "]\n==============================\n Client Req :" + request + "\n==============================\nServer Res :" + string(data) + "\n==============================\n\nReq Info:  \nHost: " + host + " \nIP: " + s.Query + "\nRegion: " + s.RegionName + "  \nCity: " + s.Country + "-" + s.City + " \nPlatform/Device:" + device)
	return message
}

func SendMessageFailureTelegram(method string, response helper.FailureResponse, request string, device string, ip string, host string) string {
	getLocation := "http://ip-api.com/json/" + ip
	responseLoc, err := http.Get(getLocation)
	if err != nil {
		panic(err)
	}
	defer responseLoc.Body.Close()
	ipLoc, _ := ioutil.ReadAll(responseLoc.Body)

	s, err := getStations([]byte(ipLoc))

	now := time.Now()
	data, _ := json.Marshal(response)
	var message = fmt.Sprintf(env.Getenv("API_TITLE") + "\n" + method + " [" + now.Format("2006-01-02") + "]\n==============================\n Client Req :" + request + "\n==============================\n Server Res :" + string(data) + "\n==============================\n\nReq Info: \nHost: " + host + "  \nIP: " + s.Query + "\nRegion: " + s.RegionName + "  \nCity: " + s.Country + "-" + s.City + " \nPlatform/Device:" + device)
	return message
}
