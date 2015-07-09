package parser

import (
    "fmt"
    "time"
    "net/http"
    "encoding/json"
    "strings"
    "model"
)

const (
    FACEBOOK_PREFIX = "https://www.facebook.com/"

    FACEBOOK_API = "https://graph.facebook.com/v2.3/"

    EVENT_FIELDS = "id,name,description,start_time,end_time,is_date_only,place"

    ACCESS_TOKEN = "CAACEdEose0cBABPr5uft5ktSLPr4mluNShYyAQFZADejEQIAGjYxtuP8c2tpxnfy6LHxynBGaclcVgfysyZCgimZAaI4PFPnZANlCFR48HaapvZCLAhlExBvKVkP6U7VseerorOWkcZAQGF2PKteT4D1ZBn0RUrZA6m83WQ6Swljzc4zNCEmXBB1NX6LDxUZCteXntz2z0u8KmRW2ZCdarDwM4OfmdLjKkXJ4ZD"
)


type fbError struct {
    Message string
    Type string
    Code int
}

type fbPlace struct {
    Id string
    Name string
    Location struct {
        Latitude float64
        Longitude float64
        Street string
        City string
        Zip string
        Country string
    }
}

type fbEvent struct {
    Id string
    Name string
    Description string
    Start_time string
    End_time string
    Is_date_only bool
    Place fbPlace
}

type fbEventList struct {
    Error *fbError
    Data []fbEvent
}


func ParseEvents(url string, events *[]model.Event) (error) {

    if !strings.HasPrefix(url, FACEBOOK_PREFIX) {
        return fmt.Errorf("ParseError: '%s' is not a valid facebook URL", url)
    }

    page := url[len(FACEBOOK_PREFIX):]
    req := FACEBOOK_API + page + "/events" + "?access_token=" + ACCESS_TOKEN + "&fields=" + EVENT_FIELDS

    resp, err := http.Get(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    var msg fbEventList
    if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
        return err
    }
    if msg.Error != nil {
        return fmt.Errorf("%s: %s", msg.Error.Type, msg.Error.Message)
    }

    for _, e := range msg.Data {
        event := model.Event {
            Name: e.Name,
            Desc: e.Description,
            Link: getLink(e.Id),
            Start: parseDate(e.Start_time),
            End: parseDate(e.End_time),
            IsDateOnly: e.Is_date_only,
            Place: model.Place {
                Name: e.Place.Name,
                Lat: e.Place.Location.Latitude,
                Long: e.Place.Location.Longitude,
                Street: e.Place.Location.Street,
                City: e.Place.Location.City,
                Zip: e.Place.Location.Zip,
            },
        }

        if event.IsValid() {
            *events = append(*events, event)
        }
    }

    return nil
}

func parseDate(dt string) time.Time {
    tm, _ := time.Parse("2006-01-02T15:04:05-0700", dt)
    return tm
}

func getLink(id string) string {
    return fmt.Sprintf("https://www.facebook.com/%s", id)
}