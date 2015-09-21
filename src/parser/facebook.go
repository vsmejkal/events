package parser

import (
	"encoding/json"
	"fmt"
	"model"
	"net/http"
	"strings"
	"time"
)

const (
	GRAPH_API = "https://graph.facebook.com/v2.4/"
)

type fbError struct {
	Message string
	Type    string
	Code    int
}

type fbPlace struct {
	Id       string
	Name     string
	Location struct {
		Latitude  float64
		Longitude float64
		Street    string
		City      string
		Zip       string
		Country   string
	}
}

type fbEvent struct {
	Id           string
	Name         string
	Description  string
	Start_time   string
	End_time     string
	Place        fbPlace
}

type fbEventList struct {
	Error *fbError
	Data  []fbEvent
}

func ParseEvents(url string, eventChan chan<- model.Event, errChan chan<- error) {
	node, err := getNodeName(url)
	if err != nil {
		errChan <- err
		return
	}
	
	req := GRAPH_API + node + "/events" +
		   "?access_token=" + FACEBOOK_APP_TOKEN +
		   "&fields=id,name,description,start_time,end_time,place" +
		   "&limit=50"

	resp, err := http.Get(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	var msg fbEventList
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		errChan <- err
		return
	}
	if msg.Error != nil {
		errChan <- fmt.Errorf("%s: %s", msg.Error.Type, msg.Error.Message)
		return
	}

	for _, e := range msg.Data {
		eventChan <- model.Event {
			Name:     e.Name,
			Desc:     e.Description,
			Link:     createLink(e.Id),
			Start:    parseDate(e.Start_time),
			End:      parseDate(e.End_time),
			Place: model.Place {
				Name:   e.Place.Name,
				Gps:	model.Gps{e.Place.Location.Latitude, e.Place.Location.Longitude},
				Street: e.Place.Location.Street,
				City:   e.Place.Location.City,
				Zip:    e.Place.Location.Zip,
			},
		}
	}
}

func parseDate(date string) model.Datetime {
	// Date and time
	tm, err := time.Parse("2006-01-02T15:04:05-0700", date)

	// Only date
	if err != nil {
		tm, _ = time.Parse("2006-01-02", date)
	}

	return model.Datetime{tm}
}

func getNodeName(url string) (string, error) {
	if !strings.HasPrefix(url, "https://www.facebook.com/") {
		return "", fmt.Errorf("ParseError: '%s' is not a valid facebook URL", url)
	}

	return url[strings.LastIndex(url, "/") + 1:], nil
}

func createLink(id string) string {
	return fmt.Sprintf("https://www.facebook.com/%s", id)
}
